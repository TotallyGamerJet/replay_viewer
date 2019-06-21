package packets

import "replay_viewer/protocol/codecs"

type PlayerChangeHeldItem struct {
	slot codecs.Byte
}

func (_ PlayerChangeHeldItem) Name() string { return "ServerPlayerChangeHeldItemPacket" }

func (_ PlayerChangeHeldItem) ID() int { return 0x3F }

func (p PlayerChangeHeldItem) ReadPacketData(pkt *Packet) (holder Holder, err error) {
	p.slot, err = pkt.readByte()
	if err != nil {
		return
	}
	return p, nil
}
