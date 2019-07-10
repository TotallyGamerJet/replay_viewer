package packets

import "replay_viewer/protocol/codecs"

type PlayerPositionAndLook struct {
	X     codecs.Double
	Y     codecs.Double
	Z     codecs.Double
	Yaw   codecs.Float
	Pitch codecs.Float
	Flags codecs.Byte
	Data  codecs.VarInt
}

//Name returns the name of the packet as a string
func (_ PlayerPositionAndLook) Name() string { return "ServerPlayerPositionRotationPacket" }

//ID returns the id in hex of the packet
func (_ PlayerPositionAndLook) ID() int { return 0x35 }

func (p PlayerPositionAndLook) ReadPacketData(pkt *Packet) (holder Holder, err error) {
	p.X, err = pkt.readDouble()
	if err != nil {
		return
	}
	p.Y, err = pkt.readDouble()
	if err != nil {
		return
	}
	p.Z, err = pkt.readDouble()
	if err != nil {
		return
	}
	p.Yaw, err = pkt.readFloat()
	if err != nil {
		return
	}
	p.Pitch, err = pkt.readFloat()
	if err != nil {
		return
	}
	p.Flags, err = pkt.readByte()
	if err != nil {
		return
	}
	p.Data, err = pkt.readVarInt()
	if err != nil {
		return
	}
	return p, nil
}
