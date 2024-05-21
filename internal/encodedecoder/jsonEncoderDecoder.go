package encodedecoder

import "encoding/json"

type JSONEncodeDecoder struct {
	enc json.Encoder
	dec json.Decoder
}
