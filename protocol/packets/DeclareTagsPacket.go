package packets

type DeclareTags struct {
}

func (_ DeclareTags) Name() string { return "ServerDeclareTagsPacket" }

func (_ DeclareTags) ID() int { return 0x5B }

func (p DeclareTags) ReadPacketData(pkt *Packet) (holder Holder, err error) {
	return p, nil
}
