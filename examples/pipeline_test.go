package examples

import (
	"context"
	"fmt"
	"github.com/elizabevil/ffmpegx/metadatax"
	"github.com/elizabevil/ffmpegx/transcoderx"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"testing"
	"time"
)

// ============

func TestTranscoderPipeline(t *testing.T) {
	transcoder := makeTranscoder()
	transcoder.Args = hlsParams()
	pipline, err := transcoder.Pipeline(func(cmd *exec.Cmd) {
		//getwd, _ := os.Getwd()
		//cmd.Dir = filepath.Join(getwd, "file")
	})
	if err != nil {
		log.Panicln(err)
	}
	for progress := range pipline {
		fmt.Println("progress", progress)
	}
}
func TestPipeline(t *testing.T) {
	transcoderx.SetMode(true)
	pipline, err := transcoderx.Pipeline(FfmpegBin, transcoderx.NewProgressMaker(), makeParams2(), func(cmd *exec.Cmd) {
		getwd, _ := os.Getwd()
		cmd.Dir = filepath.Join(getwd, "file")
	}, func(process *os.Process) {
	})
	if err != nil {
		log.Panicln(err)
	}
	for progress := range pipline {
		fmt.Println("progress", progress)
	}
	fmt.Println("end！！ Done")
}
func TestPipelineX(t *testing.T) {
	transcoderx.SetMode(true)
	runFunc, err := transcoderx.PipelineX(FfmpegBin, transcoderx.NewProgressMaker(), makeParams2(), func(cmd *exec.Cmd) {
		cmd.Dir = "/tmp"
	})
	if err != nil {
		log.Panic("PipelineX", err)
	}
	now := time.Now()
	defer func() {
		fmt.Println("time.Since", time.Since(now))
	}()
	timeout, cancelFunc := context.WithTimeout(context.Background(), 10*1000*time.Millisecond)
	defer cancelFunc()
	err = runFunc(timeout, func(progress metadatax.Progress) {
		fmt.Println("progress:", progress)
	}, func(process *os.Process) {
	})
	fmt.Println("Err:", err)
}
