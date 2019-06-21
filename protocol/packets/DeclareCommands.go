package packets

import (
	"replay_viewer/protocol/codecs"
)

type DeclareCommands struct {
	count codecs.VarInt
	//nodes codecs.NodeArray
	rootIndex codecs.VarInt
}

func (_ DeclareCommands) Name() string { return "ServerDeclareCommandsPacket" }

func (_ DeclareCommands) ID() int { return 0x11 }

func (p DeclareCommands) ReadPacketData(pkt *Packet) (holder Holder, err error) {
	return p, nil
}
