package metadatax

import (
	"context"
	"os"
	"os/exec"
)

type CmdHandle = func(cmd *exec.Cmd)
type ProgressHandle = func(progress Progress)
type ProcessCtxHandle = func(ctx context.Context, handle func(process *os.Process), progressHandles ...ProgressHandle) error

func NewDefaultProgress() DefaultProgress { return DefaultProgress{} }
