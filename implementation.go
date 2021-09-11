package gomapper

import (
	"bytes"
	jsoniter "github.com/json-iterator/go"
	"github.com/ptflp/godecoder"
)

var decode decoder.Decoder

func init() {
	config := jsoniter.Config{
		EscapeHTML: true,
		TagKey: "mapper",
	}

	decode = decoder.NewDecoder()
}

func MapStructs(dest, src interface{}) error {
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