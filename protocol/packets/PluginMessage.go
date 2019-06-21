package packets

type PluginMessage struct {
}

func (_ PluginMessage) Name() string { return "ServerPluginMessagePacket" }

func (_ PluginMessage) ID() int { return 0x18 }

func (p PluginMessage) ReadPacketData(pkt *Packet) (holder Holder, err error) {
	return p, nil
}
