package transcoderx

import (
	"fmt"
	"github.com/elizabevil/ffmpegx/paramx"
	"github.com/elizabevil/ffmpegx/transcoderx/interfacex"
	"os"
)

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
