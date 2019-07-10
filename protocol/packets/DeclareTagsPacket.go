package packets

type DeclareTags struct {
}

//Name returns the name of the packet as a string
func (_ DeclareTags) Name() string { return "ServerDeclareTagsPacket" }

//ID returns the id in hex of the packet
func (_ DeclareTags) ID() int { return 0x5B }

func (p DeclareTags) ReadPacketData(pkt *Packet) (holder Holder, err error) {
	return p, nil
}
