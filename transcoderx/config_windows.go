//go:build windows

package transcoderx

import (
	"github.com/elizabevil/ffmpegx/paramx"
	"os"
	"path/filepath"
)

func NewConfig(handles ...func(*Config)) Config {
	co := Config{}
	for _, handle := range handles {
		handle(&co)
	}
	return co
}

func NewConfigWithDir(dir string) (Config, error) {
	config := Config{
		filepath.Join(dir, "ffmpeg.exe"),
		filepath.Join(dir, "ffprobe.exe"),
		""}
	stat, err := os.Stat(config.FFmpegBin)
	if err != nil || stat.IsDir() {
		return config, paramx.ErrNotFountBin
	}
	stat, err = os.Stat(config.FFprobeBin)
	if err != nil || stat.IsDir() {
		return config, paramx.ErrNotFountBin
	}
	return config, nil
}
