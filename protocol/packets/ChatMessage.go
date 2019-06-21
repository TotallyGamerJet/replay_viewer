package packets

import (
	"replay_viewer/protocol/codecs"
)

type ChatMessage struct {
	Chat     codecs.String //chat.TextComponent
	Position codecs.Byte
}

func (_ ChatMessage) Name() string { return "ChatMessage" }

func (_ ChatMessage) ID() int { return 0x0F }

func (p ChatMessage) ReadPacketData(pkt *Packet) (holder Holder, err error) {
	p.Chat, err = pkt.readString()
	if err != nil {
		return
	}
	p.Position, err = pkt.readByte()
	if err != nil {
		return
	}
	return p, nil
}
