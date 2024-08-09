package transcoderx

import (
	"fmt"
	"github.com/elizabevil/ffmpegx/paramx"
	"github.com/elizabevil/ffmpegx/transcoderx/interfacex"
	"os"
)

func verifyProgressHandle(args interfacex.ProgressHandle) error {
	if args == nil {
		return paramx.ErrNotProgressHandle
	}
	return nil
}
func verifyArgs(args interfacex.IArg) error {
	if args == nil {
		return paramx.ErrNotArgs
	}
	return nil
}
func verifyBin(bin string) error {
	stat, err := os.Stat(bin)
	if err != nil || stat.IsDir() {
		return paramx.ErrNotFountBin
	}
	return nil
}

func mustVerifyBin(bin string) {
	stat, err := os.Stat(bin)
	if err != nil || stat.IsDir() {
		panic(fmt.Errorf("bin:%s:%w", bin, paramx.ErrNotFountBin))
	}
}
func verify(bin string, args interfacex.IArg) error {
	err := verifyBin(bin)
	if err != nil {
		return err
	}
	err = verifyArgs(args)
	if err != nil {
		return err
	}
	return nil
}
func verifyPipeline(bin string, args interfacex.IArg, ph interfacex.ProgressHandle) error {
	err := verifyBin(bin)
	if err != nil {
		return err
	}
	err = verifyArgs(args)
	if err != nil {
		return err
	}
	err = verifyProgressHandle(ph)
	if err != nil {
		return err
	}
	return nil
}
