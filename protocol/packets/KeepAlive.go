package packets

import "replay_viewer/protocol/codecs"

type KeepAlive struct {
	AliveId codecs.VarInt
}

//Name returns the name of the packet as a string
func (_ KeepAlive) Name() string { return "*KeepAlive" }

//ID returns the id in hex of the packet
func (_ KeepAlive) ID() int { return 0x20 }

func (p KeepAlive) ReadPacketData(pkt *Packet) (holder Holder, err error) {
	p.AliveId, err = pkt.readVarInt()
	if err != nil {
		return
	}
	return p, nil
}
