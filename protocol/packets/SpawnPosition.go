package packets

import "replay_viewer/protocol"

type SpawnPosition struct {
	location protocol.BlockPos
}

func (_ SpawnPosition) Name() string { return "*SpawnPosition" }

func (_ SpawnPosition) ID() int { return 0x4D }

func (p SpawnPosition) ReadPacketData(pkt *Packet) (holder Holder, err error) {
	p.location, err = pkt.readBlockPos()
	if err != nil {
		return
	}
	return p, nil
}
