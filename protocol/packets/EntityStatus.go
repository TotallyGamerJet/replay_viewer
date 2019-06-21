package packets

import "replay_viewer/protocol/codecs"

type EntityStatus struct {
	entityID codecs.Int
	status   codecs.Byte
}

func (_ EntityStatus) Name() string { return "ServerEntityStatusPacket" }

func (_ EntityStatus) ID() int { return 0x1B }

func (p EntityStatus) ReadPacketData(pkt *Packet) (holder Holder, err error) {
	p.entityID, err = pkt.readInt()
	if err != nil {
		return
	}
	p.status, err = pkt.readByte()
	if err != nil {
		return
	}
	return p, nil
}
