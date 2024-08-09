package metadatax

import "github.com/elizabevil/ffmpegx/paramx/typex"

// Metadata ...
type Metadata struct {
	ProgramVersion   ProgramVersion   `json:"program_version,omitempty"`
	LibraryVersions  []LibraryVersion `json:"library_versions,omitempty"`
	PixelFormats     []PixelFormat    `json:"pixel_formats,omitempty"`
	PacketsAndFrames []PacketAndFrame `json:"packets_and_frames,omitempty"`
	Packets          []Packet         `json:"packets,omitempty"`
	Frames           []Frame          `json:"frames,omitempty"`
	Programs         []Program        `json:"programs,omitempty"`
	Format           Format           `json:"format,omitempty"`
	Chapters         []Chapter        `json:"chapters,omitempty"`
	Streams          []Stream         `json:"streams,omitempty"`
}

func (m Metadata) GetFormat() Format {
	return m.Format
}
func (m Metadata) GetStreams() []Stream {
	return m.Streams

}
func (m Metadata) AudioStream() ([]StreamAudio, error) {
	var data []StreamAudio
	for _, v := range m.Streams {
		if typex.CodecType(v.CodecType) == typex.AVMEDIA_TYPE_AUDIO {
			data = append(data, StreamAudio{v.BaseStream, v.StreamAudioOnly})
		}
	}
	return data, nil
}
func (m Metadata) VideoStream() ([]StreamVideo, error) {
	var data []StreamVideo
	for _, v := range m.Streams {
		if typex.CodecType(v.CodecType) == typex.AVMEDIA_TYPE_VIDEO {
			data = append(data, StreamVideo{v.BaseStream, v.StreamVideoOnly})
		}
	}
	return data, nil
}
func (m Metadata) StreamType(ty typex.CodecType) ([]StreamSubtitle, error) {
	var data []StreamSubtitle
	for _, v := range m.Streams {
		if typex.CodecType(v.CodecType) == ty {
			data = append(data, StreamSubtitle{v.BaseStream, v.StreamSubtitleOnly})
		}
	}
	return data, nil
}
