package codecs

import (
	"errors"
)

var UnknownCodecType = errors.New("unknown codec type")

/*type Codec interface {
	Decode(r io.Reader) (interface{}, error)
	Encode(w io.Writer) error
}*/
