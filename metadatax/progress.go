package metadatax

import (
	"bufio"
	"bytes"
	"context"
	"fmt"
	"github.com/elizabevil/ffmpegx/transcoderx/utilx"
	"io"
	"regexp"
	"strconv"
	"strings"
)

// Progress  fftools\ffmpeg.c
type Progress struct {
	Frame string `json:"frame"`
	Fps   string `json:"fps"`
	Q     string `json:"q"`
	Size  string `json:"size"`
	Time  string `json:"time"`

	Bitrate   string `json:"bitrate"`
	TotalSize string `json:"total_size"`

	Speed      string `json:"speed"`
	OutTimeUs  string `json:"out_time_us"`
	OutTimeMs  string `json:"out_time_ms"`
	OutTime    string `json:"out_time"`
	Dup        string `json:"dup"`
	DropFrames string `json:"drop_frames"`

	Progress float64 `json:"progress"`
}

type DefaultProgress struct {
	TotalDuration string
	Filter        func(str string) bool
}

func (r DefaultProgress) MakeProgress(ctx context.Context, stream io.ReadCloser, out chan Progress) {
	if ctx == nil {
		return
	}
	defer stream.Close()
	split := func(data []byte, atEOF bool) (advance int, token []byte, spliterror error) {
		if atEOF && len(data) == 0 {
			return 0, nil, nil
		}
		if i := bytes.IndexByte(data, '\n'); i >= 0 {
			// We have a full newline-terminated line.
			return i + 1, data[0:i], nil
		}
		if i := bytes.IndexByte(data, '\r'); i >= 0 {
			// We have a cr terminated line
			return i + 1, data[0:i], nil
		}
		if atEOF {
			return len(data), data, nil
		}
		return 0, nil, nil
	}
	scanner := bufio.NewScanner(stream)
	scanner.Split(split)
	buf := make([]byte, 2)
	scanner.Buffer(buf, bufio.MaxScanTokenSize)
	var re = regexp.MustCompile(`=\s+`)
	pp := Progress{}
	dursec, _ := strconv.ParseFloat(r.TotalDuration, 64)
	for scanner.Scan() {
		line := scanner.Text()
		next := false
		if r.Filter != nil {
			next = r.Filter(line)
		} else {
			next = strings.Contains(line, "time=") && strings.Contains(line, "size=") && strings.Contains(line, "bitrate=") && strings.Contains(line, "speed=")
		}
		if next {
			st := re.ReplaceAllString(line, `=`)
			f := strings.Fields(st)
			for j := 0; j < len(f); j++ {
				field := f[j]
				fieldSplit := strings.Split(field, "=")
				if len(fieldSplit) > 1 {
					fieldname := strings.Split(field, "=")[0]
					fieldvalue := strings.Split(field, "=")[1]
					switch fieldname {
					case "frame":
						pp.Frame = fieldvalue
					case "fps":
						pp.Fps = fieldvalue
					case "q":
						pp.Q = fieldvalue
					case "size":
						pp.Size = fieldvalue
					case "time":
						pp.Time = fieldvalue
					case "bitrate":
						pp.Bitrate = fieldvalue
					case "total_size":
						pp.TotalSize = fieldvalue
					case "speed":
						pp.Speed = fieldvalue
					case "out_time_us":
						pp.OutTimeUs = fieldvalue
					case "out_time_ms":
						pp.OutTimeMs = fieldvalue
					case "out_time":
						pp.OutTime = fieldvalue
					case "dup":
						pp.Dup = fieldvalue
					case "drop_frames":
						pp.DropFrames = fieldvalue
					}
				}
			}
			pp.Progress = (utilx.DurToSec(pp.Time) * 100) / dursec
			select {
			case <-ctx.Done():
				fmt.Println("make progress end")
				close(out)
				return
			default:
				out <- pp
			}
		}
	}
}
func (r DefaultProgress) MakeProgressX(ctx context.Context, stream io.ReadCloser, handles ...ProgressHandle) {
	out := make(chan Progress)
	defer stream.Close()
	split := func(data []byte, atEOF bool) (advance int, token []byte, spliterror error) {
		if atEOF && len(data) == 0 {
			return 0, nil, nil
		}
		if i := bytes.IndexByte(data, '\n'); i >= 0 {
			// We have a full newline-terminated line.
			return i + 1, data[0:i], nil
		}
		if i := bytes.IndexByte(data, '\r'); i >= 0 {
			// We have a cr terminated line
			return i + 1, data[0:i], nil
		}
		if atEOF {
			return len(data), data, nil
		}
		return 0, nil, nil
	}
	scanner := bufio.NewScanner(stream)
	scanner.Split(split)
	buf := make([]byte, 2)
	scanner.Buffer(buf, bufio.MaxScanTokenSize)
	var re = regexp.MustCompile(`=\s+`)
	pp := Progress{}
	dursec, _ := strconv.ParseFloat(r.TotalDuration, 64)
	for scanner.Scan() {
		line := scanner.Text()
		next := false
		if r.Filter != nil {
			next = r.Filter(line)
		} else {
			next = strings.Contains(line, "time=") && strings.Contains(line, "size=") && strings.Contains(line, "bitrate=") && strings.Contains(line, "speed=")
		}
		if next {
			st := re.ReplaceAllString(line, `=`)
			f := strings.Fields(st)
			for j := 0; j < len(f); j++ {
				field := f[j]
				fieldSplit := strings.Split(field, "=")
				if len(fieldSplit) > 1 {
					fieldname := strings.Split(field, "=")[0]
					fieldvalue := strings.Split(field, "=")[1]
					switch fieldname {
					case "frame":
						pp.Frame = fieldvalue
					case "fps":
						pp.Fps = fieldvalue
					case "q":
						pp.Q = fieldvalue
					case "size":
						pp.Size = fieldvalue
					case "time":
						pp.Time = fieldvalue
					case "bitrate":
						pp.Bitrate = fieldvalue
					case "total_size":
						pp.TotalSize = fieldvalue
					case "speed":
						pp.Speed = fieldvalue
					case "out_time_us":
						pp.OutTimeUs = fieldvalue
					case "out_time_ms":
						pp.OutTimeMs = fieldvalue
					case "out_time":
						pp.OutTime = fieldvalue
					case "dup":
						pp.Dup = fieldvalue
					case "drop_frames":
						pp.DropFrames = fieldvalue
					}
				}
			}
			pp.Progress = (utilx.DurToSec(pp.Time) * 100) / dursec
			select {
			case <-ctx.Done():
				fmt.Println("make progress end")
				close(out)
			default:
				out <- pp
			}
		}
	}
}
