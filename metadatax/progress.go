package metadatax

import (
	"bufio"
	"bytes"
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

	Speed string `json:"speed"`

	Progress float64 `json:"progress"`
}

// GetFramesProcessed ...
func (p Progress) GetFramesProcessed() string {
	return p.Frame
}

// GetCurrentTime ...
func (p Progress) GetCurrentTime() string {
	return p.Time
}

// GetCurrentBitrate ...
func (p Progress) GetCurrentBitrate() string {
	return p.Bitrate
}

// GetProgress ...
func (p Progress) GetProgress() float64 {
	return p.Progress
}

// GetSpeed ...
func (p Progress) GetSpeed() string {
	return p.Speed
}

type DefaultProgress struct {
	TotalDuration string
}

func (r DefaultProgress) MakeProgress(stream io.ReadCloser, out chan Progress) {
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
	var framesProcessed string
	var currentTime string
	var currentBitrate string
	var currentSpeed string
	var q string
	var size string
	var total_size string
	var fps string
	dursec, _ := strconv.ParseFloat(r.TotalDuration, 64)
	for scanner.Scan() {
		pp := Progress{}
		line := scanner.Text()
		if strings.Contains(line, "frame=") && strings.Contains(line, "time=") && strings.Contains(line, "bitrate=") {
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
						framesProcessed = fieldvalue
					case "fps":
						fps = fieldvalue
					case "q":
						q = fieldvalue

					case "size":
						size = fieldvalue
					case "time":
						currentTime = fieldvalue
					case "bitrate":
						currentBitrate = fieldvalue
					case "total_size":
						total_size = fieldvalue
					case "speed":
						currentSpeed = fieldvalue
					}
				}
			}

			pp.Frame = framesProcessed
			pp.Fps = fps
			pp.Q = q
			pp.Size = size
			pp.Time = currentTime
			pp.Bitrate = currentBitrate
			pp.TotalSize = total_size
			pp.Speed = currentSpeed

			pp.Progress = (utilx.DurToSec(currentTime) * 100) / dursec

			out <- pp
		}
	}
}
