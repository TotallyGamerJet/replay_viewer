package packets

import "replay_viewer/protocol"

type SpawnPosition struct {
	location protocol.BlockPos
}

//Name returns the name of the packet as a string
func (_ SpawnPosition) Name() string { return "*SpawnPosition" }

//ID returns the id in hex of the packet
func (_ SpawnPosition) ID() int { return 0x4D }

func (p SpawnPosition) ReadPacketData(pkt *Packet) (holder Holder, err error) {
	p.location, err = pkt.readBlockPos()
	if err != nil {
		return
	}
	return p, nil
}
