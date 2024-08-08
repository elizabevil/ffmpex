package examples

import (
	"fmt"
	"github.com/elizabevil/ffmpegx/transcoderx"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"testing"
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
