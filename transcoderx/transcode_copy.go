package transcoderx

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/elizabevil/ffmpegx/metadatax"
	"github.com/elizabevil/ffmpegx/paramx"
	"github.com/elizabevil/ffmpegx/transcoderx/interfacex"
	"io"
	"os"
	"os/exec"
)

// Metadata File `s Metadata
func Metadata(ffprobeBin string, input string) (metadatax.Metadata, error) {
	var metadata metadatax.Metadata
	err := verifyBin(ffprobeBin)
	if err != nil {
		return metadata, err
	}
	if len(input) == 0 {
		return metadata, paramx.ErrNotInputs
	}
	args := []string{"-i", input, "-print_format", "json", "-show_format", "-show_streams", "-show_error"}
	cmd := exec.Command(ffprobeBin, args...)
	DebugPrint("%s", cmd.String())
	var outb, errb bytes.Buffer
	cmd.Stdout = &outb
	cmd.Stderr = &errb
	err = cmd.Run()
	if err != nil {
		return metadata, fmt.Errorf("executing (%s)| error: %s | message: %s %s", cmd.String(), err, outb.String(), errb.String())
	}
	if err := json.Unmarshal([]byte(outb.String()), &metadata); err != nil {
		return metadata, err
	}
	return metadata, nil
}

// CommandLine Make command
func CommandLine(ffmpegBin string, args interfacex.IArgs) string {
	err := verifyArgs(args)
	if err != nil {
		return ""
	}
	command := exec.Command(ffmpegBin, args.Args()...)
	DebugPrint("%s", command.String())
	return command.String()
}

// Cmd Use command
func Cmd(ffmpegBin string, args interfacex.IArg, handles ...func(command *exec.Cmd)) (*exec.Cmd, error) {
	err := verifyBin(ffmpegBin)
	if err != nil {
		return nil, err
	}
	err = verifyArgs(args)
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
	err := verifyBin(ffmpegBin)
	if err != nil {
		return nil, err
	}
	err = verifyArgs(args)
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
func Pipeline(ffmpegBin string, ph interfacex.ProgressHandle, args interfacex.IArg, cmdHandle CmdHandle, handles ...func(process *os.Process)) (<-chan metadatax.Progress, error) {
	err := verifyBin(ffmpegBin)
	if err != nil {
		return nil, err
	}
	err = verifyArgs(args)
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
		return nil, err
	}
	err = cmd.Start()
	if err != nil {
		return nil, err
	}
	for _, handle := range handles {
		go handle(cmd.Process)
	}
	out := make(chan metadatax.Progress)
	go func() {
		ph.MakeProgress(stderrIn, out)
	}()
	go func() {
		defer close(out)
		err = cmd.Wait()
	}()
	return out, nil
}

// PipelineX Overload Realtime: Chan ProgressHandle
func PipelineX(ffmpegBin string, ph interfacex.ProgressHandle, args interfacex.IArg, handles ...CmdHandle) (ProcessCtxHandle, error) {
	err := verifyBin(ffmpegBin)
	if err != nil {
		return nil, err
	}
	err = verifyArgs(args)
	if err != nil {
		return nil, err
	}
	cmd := exec.Command(ffmpegBin, args.Args()...)
	for _, handle := range handles {
		handle(cmd)
	}
	DebugPrint("%s", cmd.String())
	return func(ctx context.Context, progressHandle ProgressHandle, handles ...func(process *os.Process)) error {
		var stderrIn io.ReadCloser
		stderrIn, err := cmd.StderrPipe()
		if err != nil {
			return fmt.Errorf("pipe %w", err)
		}
		err = cmd.Start()
		if err != nil {
			return fmt.Errorf("start %w", err)
		}
		for _, handle := range handles {
			handle(cmd.Process)
		}
		hasFunc := progressHandle != nil
		out := make(chan metadatax.Progress)
		go func() {
			ph.MakeProgress(stderrIn, out)
		}()
		//done := make(chan error, 1)
		go func() {
			defer close(out)
			err = cmd.Wait()
		}()
		for {
			select {
			case <-ctx.Done():
				return ctx.Err()
			case data, ok := <-out:
				if ok && hasFunc {
					progressHandle(data)
				}
				if !ok {
					return err
				}
			}
		}
	}, nil
}
