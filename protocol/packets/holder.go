package packets

type Holder interface {
	ID() int
	ReadPacketData(pkt *Packet) (Holder, error)
	Name() string
}
