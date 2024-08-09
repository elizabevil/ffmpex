package transcoderx

import (
	"encoding/json"
	"github.com/elizabevil/ffmpegx/metadatax"
)

var NewProgressMaker = metadatax.NewDefaultProgress

type Unmarshal struct {
}

func (u Unmarshal) Unmarshal(data []byte, some any) error {
	return json.Unmarshal(data, &some)
}

var JsonUnmarshal = Unmarshal{}
