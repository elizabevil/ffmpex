package transcoderx

import (
	"context"
	"github.com/elizabevil/ffmpegx/metadatax"
	"github.com/elizabevil/ffmpegx/transcoderx/interfacex"
	"os"
	"os/exec"
)

type Transcoder struct {
	Args       interfacex.IArg
	Ph         interfacex.ProgressHandle
	Config     Config   `json:"config"`
	MiddleWire []func() `json:"middleware"` // may be use
}

func NewTranscoder(config Config) Transcoder {
	config.MustVerify()
	return Transcoder{Config: config, Ph: NewProgressMaker()}
}

func (r Transcoder) Metadata(input string) (metadatax.Metadata, error) {
	return Metadata(r.Config.FFprobeBin, input)
}

func (r Transcoder) MetadataWithArgs(args interfacex.IArg, unmarshal interfacex.Unmarshal) (metadatax.Metadata, error) {
	return MetadataWithArgs(r.Config.FFprobeBin, args, unmarshal)
}

func (r Transcoder) CommandLine(args interfacex.IArg) string {
	command := exec.Command(r.Config.FFmpegBin, args.Args()...)
	return command.String()
}

func (r Transcoder) Cmd(handles ...metadatax.CmdHandle) (*exec.Cmd, error) {
	return Cmd(r.Config.FFmpegBin, r.Args, handles...)

}
func (r Transcoder) StartProcess(handles ...func(command *os.ProcAttr)) (*os.Process, error) {
	return StartProcess(r.Config.FFmpegBin, r.Args, handles...)
}

func (r Transcoder) Pipeline(ctx context.Context, cmdHandle metadatax.CmdHandle, handles ...func(process *os.Process)) (<-chan metadatax.Progress, error) {
	return Pipeline(ctx, r.Config.FFmpegBin, r.Ph, r.Args, cmdHandle, handles...)
}

func (r Transcoder) PipelineX(middles ...metadatax.CmdHandle) (metadatax.ProcessCtxHandle, error) {
	return PipelineX(r.Config.FFmpegBin, r.Ph, r.Args, middles...)
}
