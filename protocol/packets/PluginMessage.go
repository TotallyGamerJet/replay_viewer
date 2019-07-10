package packets

type PluginMessage struct {
}

//Name returns the name of the packet as a string
func (_ PluginMessage) Name() string { return "ServerPluginMessagePacket" }

//ID returns the id in hex of the packet
func (_ PluginMessage) ID() int { return 0x18 }

func (p PluginMessage) ReadPacketData(pkt *Packet) (holder Holder, err error) {
	return p, nil
}
