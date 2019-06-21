package packets

import (
	"replay_viewer/protocol/codecs"
)

type PlayerAbilities struct {
	flags    codecs.Byte
	flySpeed codecs.Float
	fov      codecs.Float
}

func (_ PlayerAbilities) Name() string { return "ServerPlayerAbilitiesPacket" }

func (_ PlayerAbilities) ID() int { return 0x18 }

func (p PlayerAbilities) ReadPacketData(pkt *Packet) (holder Holder, err error) {
	p.flags, err = pkt.readByte()
	if err != nil {
		return
	}
	p.flySpeed, err = pkt.readFloat()
	if err != nil {
		return
	}
	p.fov, err = pkt.readFloat()
	if err != nil {
		return
	}
	return p, nil
}
