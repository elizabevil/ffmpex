package transcoderx

import (
	"context"
	"fmt"
	"github.com/elizabevil/ffmpegx/metadatax"
	"github.com/elizabevil/ffmpegx/paramx"
	"github.com/elizabevil/ffmpegx/transcoderx/interfacex"
	"os"
	"os/exec"
)

type CmdHandle = func(cmd *exec.Cmd)
type ProgressHandle = func(progress metadatax.Progress)
type ProcessCtxHandle = func(ctx context.Context, progressHandle ProgressHandle, handles ...func(process *os.Process)) error

func verifyArgs(args interfacex.IArg) error {
	if args == nil {
		return paramx.ErrNotArgs
	}
	return nil
}
func mustVerifyBin(bin string) {
	stat, err := os.Stat(bin)
	if err != nil || stat.IsDir() {
		panic(fmt.Errorf("bin:%s:%w", bin, paramx.ErrNotFountBin))
	}
}
func verifyBin(bin string) error {
	stat, err := os.Stat(bin)
	if err != nil || stat.IsDir() {
		return paramx.ErrNotFountBin
	}
	return nil
}
