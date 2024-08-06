package metadatax

type codecType = string

// libavutil\utils.c
const (
	AVMEDIA_TYPE_UNKNOWN              = -1 ///< Usually treated as AVMEDIA_TYPE_DATA
	AVMEDIA_TYPE_VIDEO      codecType = "video"
	AVMEDIA_TYPE_AUDIO      codecType = "audio"
	AVMEDIA_TYPE_DATA       codecType = "data" ///< Opaque data information usually continuous
	AVMEDIA_TYPE_SUBTITLE   codecType = "subtitle"
	AVMEDIA_TYPE_ATTACHMENT codecType = "attachment" ///< Opaque data information usually sparse
	AVMEDIA_TYPE_NB         codecType = "nb"
)

type Tags = map[string]any

// Metadata ...
type Metadata struct {
	Format  Format   `json:"format"`
	Streams []Stream `json:"streams"`
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
		if v.CodecType == AVMEDIA_TYPE_AUDIO {
			data = append(data, StreamAudio{v.BaseStream, v.StreamAudioOnly})
		}
	}
	return data, nil
}
func (m Metadata) VideoStream() ([]StreamVideo, error) {
	var data []StreamVideo
	for _, v := range m.Streams {
		if v.CodecType == AVMEDIA_TYPE_VIDEO {
			data = append(data, StreamVideo{v.BaseStream, v.StreamVideoOnly})
		}
	}
	return data, nil
}
func (m Metadata) SubtitleStream() ([]StreamSubtitle, error) {
	var data []StreamSubtitle
	for _, v := range m.Streams {
		if v.CodecType == AVMEDIA_TYPE_SUBTITLE {
			data = append(data, StreamSubtitle{v.BaseStream, v.StreamSubtitleOnly})
		}
	}
	return data, nil
}

func (m Metadata) ParseTypes() ([]StreamAudio, []StreamVideo, []StreamSubtitle) {
	var videos []StreamVideo
	var audios []StreamAudio
	var subtitles []StreamSubtitle
	for _, v := range m.Streams {
		switch v.CodecType {
		case AVMEDIA_TYPE_VIDEO:
			videos = append(videos, StreamVideo{v.BaseStream, v.StreamVideoOnly})
		case AVMEDIA_TYPE_AUDIO:
			audios = append(audios, StreamAudio{v.BaseStream, v.StreamAudioOnly})
		case AVMEDIA_TYPE_SUBTITLE:
			subtitles = append(subtitles, StreamSubtitle{v.BaseStream, v.StreamSubtitleOnly})
		}
	}
	return audios, videos, subtitles
}

type Stream struct {
	BaseStream
	StreamAudioOnly
	StreamVideoOnly
	StreamSubtitleOnly
}
type StreamAudio struct {
	BaseStream
	StreamAudioOnly
}
type StreamVideo struct {
	BaseStream
	StreamVideoOnly
}
type StreamSubtitle struct {
	BaseStream
	StreamSubtitleOnly
}
