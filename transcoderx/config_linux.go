//go:build linux

package transcoderx

import (
	"github.com/elizabevil/ffmpegx/paramx"
	"os"
	"path/filepath"
)

func NewConfig(handles ...func(*Config)) Config {
	co := Config{ // default
		FFmpegBin:  "/usr/bin/ffmpeg",
		FFprobeBin: "/usr/bin/ffprobe",
	}
	for _, handle := range handles {
		handle(&co)
	}
	return co
}

func NewConfigWithDir(dir string) (Config, error) {
	config := Config{
		filepath.Join(dir, "ffmpeg"),
		filepath.Join(dir, "ffprobe"),
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
