package packets

import (
	"replay_viewer/protocol"
	"replay_viewer/protocol/codecs"
)

type BlockBreakAnimation struct {
	entityID     codecs.VarInt
	location     protocol.BlockPos
	destroyStage codecs.Byte
}

func (_ BlockBreakAnimation) ID() int { return 0x08 }

func (p BlockBreakAnimation) ReadPacketData(pkt *Packet) (holder Holder, err error) {
	p.entityID, err = pkt.readVarInt()
	if err != nil {
		return
	}
	p.location, err = pkt.readBlockPos()
	if err != nil {
		return
	}
	p.destroyStage, err = pkt.readByte()
	if err != nil {
		return
	}
	return p, nil
}
