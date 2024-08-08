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

func TestPipelineCtx(t *testing.T) {
	//s := "-y -hide_banner -fflags genpts -async 1 -analyzeduration 512000 -probesize 512000 -re -i http://107.167.16.226:8880/live/64h0lf458wue/f80d9e99f93e47a3/478426.ts -vcodec copy -scodec copy -acodec copy -hls_flags delete_segments -hls_delete_threshold 1 -hls_list_size 5 -hls_time 60 /home/debi/Music/stream/live/1/1_.m3u8 -vcodec copy -scodec copy -acodec copy -strftime 1 -hls_segment_filename %Y%m%d%H%M%S.ts -hls_delete_threshold 1 -hls_list_size 5 -hls_time 60 /home/debi/Music/stream/record/1/1_.m3u8"
	transcoderx.Debug = true
	//handle, err := transcoderx.PipelineX(FfmpegBin, transcoderx.NewProgressMaker(), typex.Args(strings.Split(s, " ")), func(cmd *exec.Cmd) {
	//	cmd.Dir = "/home/debi/Desktop/AA/"
	//})
	//if err != nil {
	//	panic(err)
	//}
	outCtx, cancelFunc := context.WithTimeout(context.Background(), time.Second*4)
	defer func() {
		cancelFunc()
		fmt.Println("F cancelFunc")
	}()
	//================
	out := make(chan int)
	ctx, c := context.WithCancel(outCtx)
	var err error
	go write(ctx, out)
	go func() {
		defer close(out)
		time.Sleep(3 * time.Second)
		err = os.ErrNotExist
		c()
	}()
	//go child(outCtx, handle, 1*time.Second)
	for {
		select {
		case as, ok := <-out:
			if ok {
				fmt.Println(as)
			}
		case <-outCtx.Done():
			fmt.Println("F outCtx.Done", err) //F outCtx.Done context deadline exceeded
			return
		default:

		}
	}
}

func write(ctx context.Context, ints chan int) {
	ticker := time.Tick(1 * time.Millisecond)
	for t := range ticker {
		select {
		case <-ctx.Done():
			fmt.Println("ctx.Done")
			return
		default:
			ints <- 1
			fmt.Println("ch send", t)
		}

	}

}

func TestPipelineXx(t *testing.T) {
	transcoderx.Debug = true
	timeout, cancelFunc := context.WithTimeout(context.Background(), 10*time.Second) // run time dur
	defer cancelFunc()
	count := 0
	ctx, c := context.WithCancel(timeout)
	go func() {
		for {
			time.Sleep(1 * time.Second)
			child(ctx)
			count++
			if count > 2 {
				c()
				fmt.Println("=========cc======")
			}
		}
	}()
	select {
	case <-timeout.Done():
		fmt.Println("Exit")
		return
	case <-ctx.Done():
		fmt.Println("ctx Exit")
		return
	}
}
func child(ctx context.Context) {
	fmt.Println("Start", time.Now().Format(time.DateTime))
	defer fmt.Println("C child end")
	runFunc, err := transcoderx.PipelineX(FfmpegBin, transcoderx.NewProgressMaker(), makeParams2(), func(cmd *exec.Cmd) {
		cmd.Dir = "/tmp"
	})
	if err != nil {
		return
	}
	err = runFunc(ctx, func(process *os.Process) {
	}, func(progress metadatax.Progress) {
		fmt.Println("progress:", progress)
	})
	if err != nil {
		fmt.Println("transcoderx.PipelineX", err)
		return
	}
}
