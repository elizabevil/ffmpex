package transcoderx

import (
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

// SetInputs overwrite input.Inputs

func (r Transcoder) GetMetadata(input string) (metadatax.Metadata, error) {
	return Metadata(r.Config.FFprobeBin, input)
}

func (r Transcoder) CommandLine(args interfacex.IArg) string {
	command := exec.Command(r.Config.FFmpegBin, args.Args()...)
	return command.String()
}

// run
func (r Transcoder) Cmd(handles ...CmdHandle) (*exec.Cmd, error) {
	return Cmd(r.Config.FFmpegBin, r.Args, handles...)

}
func (r Transcoder) StartProcess(handles ...func(command *os.ProcAttr)) (*os.Process, error) {
	return StartProcess(r.Config.FFmpegBin, r.Args, handles...)
}

// Pipeline totalDuration -> progress%(can be empty)
func (r Transcoder) Pipeline(cmdHandle CmdHandle, handles ...func(process *os.Process)) (<-chan metadatax.Progress, error) {
	return Pipeline(r.Config.FFmpegBin, r.Ph, r.Args, cmdHandle, handles...)
}

// PipelineX totalDuration -> progress%(can be empty)
func (r Transcoder) PipelineX(middles ...CmdHandle) (ProcessCtxHandle, error) {
	return PipelineX(r.Config.FFmpegBin, r.Ph, r.Args, middles...)
}
