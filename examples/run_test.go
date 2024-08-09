package examples

import (
	"fmt"
	"github.com/elizabevil/ffmpegx/paramx"
	"github.com/elizabevil/ffmpegx/paramx/formatx/muxerx"
	"github.com/elizabevil/ffmpegx/paramx/optionx"
	"github.com/elizabevil/ffmpegx/paramx/typex"
	"github.com/elizabevil/ffmpegx/transcoderx"
	"github.com/elizabevil/ffmpegx/transcoderx/interfacex"
	"log"
	"os"
	"path/filepath"
	"testing"
	"time"
)

var pwd, _ = os.Getwd()

func makeTranscoder() transcoderx.Transcoder {
	transcode := transcoderx.NewTranscoder(transcoderx.NewConfig())
	transcoderx.SetMode(true)
	return transcode
}

func makeParams() paramx.Param {
	param := paramx.Param{}
	param.GlobalHandle(func(input *optionx.Global) {
		input.Overwrite = true
	}).InputHandle(func(input *optionx.Input) {
		input.Re = true
		input.Inputs = []string{filepath.Join(pwd, InputVedio)}
	}).OutputHandle(func(output *optionx.Output) {
		output.Outputs = []string{OutputVideo}
		output.Vcodec = "libx264"
		output.Aq = "aac"
	}).CommonHandle(paramx.PositionOutput, optionx.Common{C: "copy"})
	fmt.Println("args", param.Args())
	return param
}
func makeParamsToHls() interfacex.IArgs {
	getwd, _ := os.Getwd()
	argInterface := paramx.BuildIArgInterface(
		optionx.Expert{NoStdin: true},
		optionx.Generic{HideBanner: true},
		optionx.Global{Overwrite: true, StatsPeriod: 0.5},
		optionx.Input{
			Re:    false,
			Input: filepath.Join(getwd, InputVedio),
		},
		optionx.Common{
			Acodec: "copy",
			Vcodec: "copy",
			Scodec: "copy",
		},
		muxerx.HLS{
			HlsTime:            typex.TimeDurationSecondI(1000 * time.Millisecond),
			HlsDeleteThreshold: 1,
			Strftime:           1,
			HlsSegmentFilename: "%s.ts",
		},
		optionx.Output{Output: filepath.Base(OutputVideoHls)})
	return argInterface
}

func process(tran transcoderx.Transcoder) {
	tran.Args = makeParams()
	process, err := tran.StartProcess()
	if err != nil {
		fmt.Println(err)
		return
	}
	wait, err := process.Wait()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(wait, err)
}

func TestProcess(t *testing.T) {
	tran := makeTranscoder()
	tran.Args = CreateCutTimeParam()
	process(tran)
}

// ============

func TestCmdx(t *testing.T) {
	cmd, _ := transcoderx.Cmd(FfmpegBin, CreateCutTimeParam())
	cmd.Stderr = os.Stdout
	cmd.Stdout = os.Stdout
	err := cmd.Start()
	if err != nil {
		log.Panicln(err)
	}
	err = cmd.Wait()
	if err != nil {
		log.Panicln(err)
	}
	fmt.Println(err)
}
func TestCmd(t *testing.T) {
	tran := makeTranscoder()
	var cmd = func(tran transcoderx.Transcoder) {
		cmd, _ := tran.Cmd()
		cmd.Stderr = os.Stdout
		cmd.Stdout = os.Stdout
		err := cmd.Start()
		if err != nil {
			fmt.Errorf("%w cmd", err)
			return
		}
		err = cmd.Wait()
		if err != nil {
			fmt.Errorf("%w Wait", err)
			return
		}
		fmt.Println(err)
	}
	tran.Args = CreateCutTimeParam()
	cmd(tran)
}
