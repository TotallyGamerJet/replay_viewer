package packets

import "replay_viewer/protocol/codecs"

type PlayerChangeHeldItem struct {
	slot codecs.Byte
}

//Name returns the name of the packet as a string
func (_ PlayerChangeHeldItem) Name() string { return "ServerPlayerChangeHeldItemPacket" }

//ID returns the id in hex of the packet
func (_ PlayerChangeHeldItem) ID() int { return 0x3F }

func (p PlayerChangeHeldItem) ReadPacketData(pkt *Packet) (holder Holder, err error) {
	p.slot, err = pkt.readByte()
	if err != nil {
		return
	}
	return p, nil
}
