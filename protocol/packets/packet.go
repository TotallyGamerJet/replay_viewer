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

/*
type Direction int

const (
	Serverbound Direction = iota
	Clientbound
)*/
/*
type State uint8

const (
	Handshake State = iota
	Status
	Login
	Play
)*/

var PacketList = map[int]Holder{
	//0x00: PlayKeepAlive{},
	//0x01: PlayJoinGame{},
	0x02: ChatMessage{},
	//0x05: PlaySpawnPosition{},
	//0x07: PlayPositionAndLook{},
	0x0D: SpawnObject{},
	0x18: EntityHeadLook{},
	0x25: BlockBreakAnimation{},
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

func (p *Packet) readString() (str string, err error) {
	str, err = utils.ReadString(&p.Data)
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
	var uuid codecs.UUID
	data, err := uuid.Decode(&p.Data)
	if err != nil {
		return
	}
	val, ok := data.(codecs.UUID)
	if !ok {
		err = fmt.Errorf("failed to decode uuid")
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
