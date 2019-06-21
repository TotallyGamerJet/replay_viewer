package packets

import "replay_viewer/protocol/codecs"

type ServerDifficulty struct {
	difficulty codecs.UnsignedByte
	locked     codecs.Boolean
}

func (_ ServerDifficulty) Name() string { return "ServerDifficultyPacket" }

func (_ ServerDifficulty) ID() int { return 0x0D }

func (p ServerDifficulty) ReadPacketData(pkt *Packet) (holder Holder, err error) {
	p.difficulty, err = pkt.readUByte()
	if err != nil {
		return
	}
	p.locked, err = pkt.readBool()
	if err != nil {
		return
	}
	return p, nil
}
