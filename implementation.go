package gomapper

import (
	"bytes"
	jsoniter "github.com/json-iterator/go"
	"github.com/ptflp/godecoder"
)

var decoder godecoder.Decoder

func init() {
	config := jsoniter.Config{
		EscapeHTML: true,
		TagKey:     "mapper",
	}

	decoder = godecoder.NewDecoder(config)
}

func MapStructs(dest, src interface{}) error {
	var b bytes.Buffer

	err := decoder.Encode(&b, src)
	if err != nil {
		return err
	}

	err = decoder.Decode(&b, dest)
	if err != nil {
		return err
	}

	return nil
}
