package packets

import (
	"replay_viewer/protocol/codecs"
)

type PlayerAbilities struct {
	flags    codecs.Byte
	flySpeed codecs.Float
	fov      codecs.Float
}

//Name returns the name of the packet as a string
func (_ PlayerAbilities) Name() string { return "ServerPlayerAbilitiesPacket" }

//ID returns the id in hex of the packet
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
