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

func MapStructs(dest, src interface{}, tag ...string) error {
	d := decoder
	if len(tag) > 0 {
		config := jsoniter.Config{
			EscapeHTML: true,
			TagKey:     tag[0],
		}
		d = godecoder.NewDecoder(config)
	}
	var b bytes.Buffer

	err := d.Encode(&b, src)
	if err != nil {
		return err
	}

	err = d.Decode(&b, dest)
	if err != nil {
		return err
	}

	return nil
}
