package packets

import "replay_viewer/protocol/codecs"

type KeepAlive struct {
	AliveId codecs.VarInt
}

func (_ KeepAlive) Name() string { return "*KeepAlive" }

func (_ KeepAlive) ID() int { return 0x20 }

func (p KeepAlive) ReadPacketData(pkt *Packet) (holder Holder, err error) {
	p.AliveId, err = pkt.readVarInt()
	if err != nil {
		return
	}
	return p, nil
}
