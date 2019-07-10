package packets

import "replay_viewer/protocol/codecs"

type UpdateViewPosition struct {
	chunkX, chunkZ codecs.VarInt
}

//Name returns the name of the packet as a string
func (_ UpdateViewPosition) Name() string { return "ServerUpdateViewPositionPacket" }

//ID returns the id in hex of the packet
func (_ UpdateViewPosition) ID() int { return 0x40 }

func (p UpdateViewPosition) ReadPacketData(pkt *Packet) (holder Holder, err error) {
	p.chunkX, err = pkt.readVarInt()
	if err != nil {
		return
	}
	p.chunkZ, err = pkt.readVarInt()
	if err != nil {
		return
	}
	return p, nil
}
