package packets

import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
	"io"
	"replay_viewer/protocol"
	"replay_viewer/protocol/codecs"
	"replay_viewer/utils"
)

type Packet struct {
	ID int
	//Direction Direction
	Data bytes.Buffer
}

var (
	UnknownPacketType   = errors.New("unknown packet type")
	InvalidPacketLength = errors.New("received packet is below zero or above maximum size")
	UnExpectedPacket    = errors.New("recieved an unexpected packet")
)

var PacketList = map[int]Holder{
	0x02: ChatMessage{}, // Not Actaully Correct should be Spawn Global Entity
	0x08: BlockBreakAnimation{},
	0x0D: ServerDifficulty{},
	0x0E: ChatMessage{},
	0x11: DeclareCommands{}, //TODO: Finish
	0x1B: EntityStatus{},
	0x18: PluginMessage{}, //Ignored
	0x20: KeepAlive{},
	0x21: ChunkData{},
	0x24: UpdateLight{},
	0x25: JoinGame{},
	0x31: PlayerAbilities{},
	0x33: PlayerInfo{},            //TODO: Finish error checking
	0x35: PlayerPositionAndLook{}, //HAS Data example
	0x36: UnlockRecipes{},
	0x3F: PlayerChangeHeldItem{},
	0x40: UpdateViewPosition{},
	0x4D: SpawnPosition{},
	0x5A: DeclareRecipe{},
	0x5B: DeclareTags{}, //TODO: Finish
}

func ReadInt(reader io.Reader) (val int, err error) {
	data := make([]byte, 4)
	n, err := reader.Read(data)
	if err != nil {
		return
	}
	if n != 4 {
		err = fmt.Errorf("reading file failed with %d bytes", n)
		return
	}
	val = int(binary.BigEndian.Uint32(data))
	return
}

func ReadPacket(reader io.Reader, length int) (pkt *Packet, err error) {
	/*data = make([]byte, length)
	n, err := reader.Read(data)
	if err != nil {
		return
	}
	if n != length {
		err = fmt.Errorf("reading file failed with %d bytes", n)
	}
	fmt.Println(data)*/
	payload := make([]byte, length)       //make var to store data
	_, err = io.ReadFull(reader, payload) //read data

	if err != nil {
		return nil, fmt.Errorf("reading payload failed: %s", err)
	}
	buffer := bytes.NewBuffer(payload)
	id, err := utils.ReadVarInt(buffer)

	if err != nil {
		return nil, fmt.Errorf("id lookup failed: %s", err)
	}
	return &Packet{
		ID: id,
		//Direction: protocol.Clientbound,
		Data: *buffer,
	}, nil
}

func (p *Packet) readString() (str codecs.String, err error) {
	s, err := utils.ReadString(&p.Data)
	str = codecs.String(s)
	return
}

func (p *Packet) readByte() (b codecs.Byte, err error) {
	data, err := utils.ReadInt8(&p.Data)
	if err != nil {
		return
	}
	b = codecs.Byte(data)
	return
}

func (p *Packet) readVarInt() (val codecs.VarInt, err error) {
	i, err := utils.ReadVarInt(&p.Data)
	if err != nil {
		return
	}
	val = codecs.VarInt(i)
	return
}

func (p *Packet) readBlockPos() (val protocol.BlockPos, err error) {
	long, err := utils.ReadUint64(&p.Data)
	if err != nil {
		return
	}
	val = protocol.FromLong(int64(long))
	return
}

func (p *Packet) readUUID() (val codecs.UUID, err error) {
	i, err := utils.ReadUUID(&p.Data)
	if err != nil {
		return
	}
	val, ok := i.(codecs.UUID)
	if !ok {
		err = fmt.Errorf("failed to decode UUID")
	}
	return
}

func (p *Packet) readDouble() (val codecs.Double, err error) {
	i, err := utils.ReadFloat64(&p.Data)
	if err != nil {
		return
	}
	val = codecs.Double(i)
	return
}

func (p *Packet) readInt() (val codecs.Int, err error) {
	i, err := utils.ReadInt32(&p.Data)
	if err != nil {
		return
	}
	val = codecs.Int(i)
	return
}

func (p *Packet) readShort() (val codecs.Short, err error) {
	i, err := utils.ReadInt16(&p.Data)
	if err != nil {
		return
	}
	val = codecs.Short(i)
	return
}

func (p *Packet) readUByte() (val codecs.UnsignedByte, err error) {
	i, err := utils.ReadUint8(&p.Data)
	if err != nil {
		return
	}
	val = codecs.UnsignedByte(i)
	return
}

func (p *Packet) readBool() (val codecs.Boolean, err error) {
	i, err := utils.ReadBool(&p.Data)
	if err != nil {
		return
	}
	val = codecs.Boolean(i)
	return
}

func (p *Packet) readFloat() (val codecs.Float, err error) {
	i, err := utils.ReadFloat32(&p.Data)
	if err != nil {
		return
	}
	val = codecs.Float(i)
	return
}

func (p *Packet) readIdentifier() (val codecs.Identifier, err error) {
	str, err := p.readString()
	if err != nil {
		return
	}
	val = codecs.Identifier(str)
	return
}

func (p *Packet) readByteArray() (val codecs.ByteArray, err error) {
	i, err := utils.ReadByteArray(&p.Data)
	if err != nil {
		return
	}
	val = codecs.ByteArray(i)
	return
}
