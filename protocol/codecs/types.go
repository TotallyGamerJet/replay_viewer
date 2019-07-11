package codecs

type (
	String string
	JSON   struct {
		V interface{}
	}
	VarInt        int
	Boolean       bool
	Byte          byte
	UnsignedByte  uint8
	ByteArray     []byte
	Short         int16
	UnsignedShort uint16
	Int           int32
	UnsignedInt   int32
	Long          int64
	UnsignedLong  uint64
	Float         float32
	Double        float64
	UUID          struct {
		most, least int64
	}
	Identifier String
)
