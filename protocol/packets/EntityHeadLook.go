package packets

import "replay_viewer/protocol/codecs"

type EntityHeadLook struct {
	entityID codecs.VarInt
	headYaw  codecs.Byte
}

func (_ EntityHeadLook) Name() string { return "*EntityHeadLook" }

func (_ EntityHeadLook) ID() int { return 0x39 }

func (e EntityHeadLook) ReadPacketData(pkt *Packet) (holder Holder, err error) {
	e.entityID, err = pkt.readVarInt()
	if err != nil {
		return
	}
	e.headYaw, err = pkt.readByte()
	if err != nil {
		return
	}
	return e, nil
}
