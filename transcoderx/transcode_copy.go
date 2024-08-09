package transcoderx

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/elizabevil/ffmpegx/metadatax"
	"github.com/elizabevil/ffmpegx/paramx"
	"github.com/elizabevil/ffmpegx/transcoderx/interfacex"
	"os"
	"os/exec"
)

// Metadata File `s Metadata With Default args
func Metadata(ffprobeBin string, input string) (metadatax.Metadata, error) {
	var metadata metadatax.Metadata
	err := verifyBin(ffprobeBin)
	if err != nil {
		return metadata, err
	}
	if len(input) == 0 {
		return metadata, paramx.ErrNotInputs
	}
	args := []string{"-hide_banner", "-i", input, "-print_format", "json", "-show_format", "-show_streams", "-show_error"}
	cmd := exec.Command(ffprobeBin, args...)
	DebugPrint("%s", cmd.String())
	var outb, errb bytes.Buffer
	cmd.Stdout = &outb
	cmd.Stderr = &errb
	err = cmd.Run()
	if err != nil {
		return metadata, fmt.Errorf("executing (%s)| error: %s | message: %s %s", cmd.String(), err, outb.String(), errb.String())
	}
	if err := json.Unmarshal(outb.Bytes(), &metadata); err != nil {
		return metadata, err
	}
	return metadata, nil
}

// MetadataWithArgs File `s Metadata With  args
func MetadataWithArgs(ffprobeBin string, args interfacex.IArg, unmarshal interfacex.Unmarshal) (metadatax.Metadata, error) {
	var metadata metadatax.Metadata
	err := verifyBin(ffprobeBin)
	if err != nil {
		return metadata, err
	}
	err = verifyArgs(args)
	if err != nil {
		return metadata, err
	}
	cmd := exec.Command(ffprobeBin, args.Args()...)
	DebugPrint("%s", cmd.String())
	var outb, errb bytes.Buffer
	cmd.Stdout = &outb
	cmd.Stderr = &errb
	err = cmd.Run()
	if err != nil {
		return metadata, fmt.Errorf("executing (%s)| error: %s | message: %s %s", cmd.String(), err, outb.String(), errb.String())
	}
	if err := unmarshal.Unmarshal(outb.Bytes(), &metadata); err != nil {
		return metadata, err
	}
	return metadata, nil
}

// CommandLine Make command
func CommandLine(ffmpegBin string, args interfacex.IArg) string {
	if verifyArgs(args) != nil {
		return ""
	}
	command := exec.Command(ffmpegBin, args.Args()...)
	DebugPrint("%s", command.String())
	return command.String()
}

// Cmd Use command
func Cmd(ffmpegBin string, args interfacex.IArg, handles ...func(command *exec.Cmd)) (*exec.Cmd, error) {
	err := verify(ffmpegBin, args)
	if err != nil {
		return nil, err
	}
	command := exec.Command(ffmpegBin, args.Args()...)
	for _, handle := range handles {
		handle(command)
	}
	DebugPrint("%s", command.String())
	return command, nil
}

// StartProcess Use os.StartProcess
func StartProcess(ffmpegBin string, args interfacex.IArg, handles ...func(command *os.ProcAttr)) (*os.Process, error) {
	err := verify(ffmpegBin, args)
	if err != nil {
		return nil, err
	}
	pa := os.ProcAttr{}
	for _, handle := range handles {
		handle(&pa)
	}
	process, err := os.StartProcess(ffmpegBin, args.Args(), &pa)
	if err != nil {
		return nil, err
	}
	return process, nil
}

// Pipeline Realtime: Chan ProgressHandle
func Pipeline(ctx context.Context, ffmpegBin string, ph interfacex.ProgressHandle, args interfacex.IArg, cmdHandle metadatax.CmdHandle, handles ...func(process *os.Process)) (<-chan metadatax.Progress, error) {
	err := verifyPipeline(ffmpegBin, args, ph)
	if err != nil {
		return nil, err
	}
	cmd := exec.Command(ffmpegBin, args.Args()...)
	if cmdHandle != nil {
		cmdHandle(cmd)
	}
	DebugPrint("%s", cmd.String())
	stderrIn, err := cmd.StderrPipe()
	if err != nil {
		return nil, fmt.Errorf("pipe %w", err)
	}
	err = cmd.Start()
	if err != nil {
		return nil, fmt.Errorf("start %w", err)
	}
	for _, handle := range handles {
		handle(cmd.Process)
	}
	out := make(chan metadatax.Progress)
	if ctx == nil {
		ctx = context.Background()
	}
	ctxCancel, cancelFunc := context.WithCancel(ctx)
	go func() {
		ph.MakeProgressPipe(ctxCancel, stderrIn, out)
	}()
	go func() {
		defer func() {
			cancelFunc()
			close(out)
			out = nil
		}()
		err = cmd.Wait()
	}()
	return out, err
}

// PipelineX Overload Realtime: Chan ProgressHandle
func PipelineX(ffmpegBin string, ph interfacex.ProgressHandle, args interfacex.IArg, handles ...metadatax.CmdHandle) (metadatax.ProcessCtxHandle, error) {
	err := verifyPipeline(ffmpegBin, args, ph)
	if err != nil {
		return nil, err
	}
	cmd := exec.Command(ffmpegBin, args.Args()...)
	for _, handle := range handles {
		handle(cmd)
	}
	DebugPrint("%s", cmd.String())
	return func(ctx context.Context, handle func(process *os.Process), progressHandle ...metadatax.ProgressHandle) error {
		stderrIn, err := cmd.StderrPipe()
		if err != nil {
			return fmt.Errorf("pipe %w", err)
		}
		err = cmd.Start()
		if err != nil {
			return fmt.Errorf("start %w", err)
		}
		if handle != nil {
			handle(cmd.Process)
		}
		if ctx == nil {
			ctx = context.Background()
		}
		cancelCtx, cancelFunc := context.WithCancel(ctx)
		out := make(chan metadatax.Progress)
		go func() {
			ph.MakeProgress(cancelCtx, stderrIn, out)
		}()
		go func() {
			err = cmd.Wait()
			cancelFunc()
		}()
		hasFunc := progressHandle != nil
		for {
			select {
			case <-cancelCtx.Done():
				return err
			case data, ok := <-out:
				if ok && hasFunc {
					for _, item := range progressHandle {
						item(data)
					}
				}
			}
		}
	}, nil
}
