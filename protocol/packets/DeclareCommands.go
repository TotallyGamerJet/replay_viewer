package packets

import (
	"replay_viewer/protocol/codecs"
)

type DeclareCommands struct {
	count codecs.VarInt
	//nodes codecs.NodeArray
	rootIndex codecs.VarInt
}

//Name returns the name of the packet as a string
func (_ DeclareCommands) Name() string { return "ServerDeclareCommandsPacket" }

//ID returns the id in hex of the packet
func (_ DeclareCommands) ID() int { return 0x11 }

func (p DeclareCommands) ReadPacketData(pkt *Packet) (holder Holder, err error) {
	return p, nil
}
