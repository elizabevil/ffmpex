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

// ============

func TestPipeline(t *testing.T) {
	transcoderx.SetMode(true)
	duration := 100 * 1 * time.Millisecond //
	killDuration := duration * 10
	//=========
	background, cc := context.WithTimeout(context.Background(), duration)
	defer cc()
	now := time.Now()
	defer func() {
		fmt.Println("Since", time.Since(now))
	}()
	pipline, err := transcoderx.Pipeline(background, FfmpegBin, transcoderx.NewProgressMaker(), makeParamsToHls(), func(cmd *exec.Cmd) {
		cmd.Dir = "/tmp"
	}, func(process *os.Process) {
		go func() {
			time.Sleep(killDuration)
			process.Kill()
			fmt.Println("kill")
		}()
	})
	fmt.Println("Pipeline", err)
	for progress := range pipline {
		fmt.Println("progress", progress)
	}
	fmt.Println("end！！ Done")
}
func TestPipelineX(t *testing.T) {
	transcoderx.SetMode(true)
	duration := 300 * time.Millisecond
	killDuration := duration * 100
	background, cc := context.WithTimeout(context.Background(), duration)
	defer cc()
	//=========
	fmt.Println("Start", time.Now().Format(time.DateTime))
	now := time.Now()
	defer func() {
		fmt.Println("Since", time.Since(now))
	}()
	runFunc, err := transcoderx.PipelineX(FfmpegBin, transcoderx.NewProgressMaker(), makeParamsToHls(), func(cmd *exec.Cmd) {
		cmd.Dir = "/tmp"
	})
	if err != nil {
		return
	}
	err = runFunc(background, func(process *os.Process) {
		go func() {
			time.Sleep(killDuration)
			process.Kill()
		}()
	}, func(progress metadatax.Progress) {
		fmt.Println("progress::", progress)
	})
	fmt.Println("transcoderx.PipelineX", err)
}
