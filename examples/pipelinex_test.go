package examples

import (
	"context"
	"fmt"
	"github.com/elizabevil/ffmpegx/metadatax"
	"github.com/elizabevil/ffmpegx/transcoderx"
	"os"
	"os/exec"
	"testing"
	"time"
)

func TestPipelineXx(t *testing.T) {
	transcoderx.Debug = true
	timeout, cancelFunc := context.WithTimeout(context.Background(), 2*time.Second) // run time dur
	defer cancelFunc()
	count := 0
	ctx, c := context.WithCancel(timeout)
	go func() {
		for {
			time.Sleep(2 * time.Second)
			fmt.Println("========---------=====")

			child(ctx)
			count++
			if count > 1 {
				fmt.Println("=========cc======")
				c()
			}
		}
	}()
	select {
	case <-timeout.Done():
		fmt.Println("Exit")
		return
	case <-ctx.Done():
		fmt.Println("ctx Exit")
	}
}
func child(ctx context.Context) {
	fmt.Println("Start", time.Now().Format(time.DateTime))
	defer fmt.Println("C child end")
	runFunc, err := transcoderx.PipelineX(FfmpegBin, transcoderx.NewProgressMaker(), makeParamsToHls(), func(cmd *exec.Cmd) {
		cmd.Dir = "/tmp"
	})
	if err != nil {
		return
	}
	err = runFunc(ctx, func(process *os.Process) {
		go func() {
			time.Sleep(2 * time.Second)
			process.Kill()
		}()
	}, func(progress metadatax.Progress) {
		fmt.Println("progress::", progress)
	})
	if err != nil {
		fmt.Println("transcoderx.PipelineX", err)
		return
	}
}
