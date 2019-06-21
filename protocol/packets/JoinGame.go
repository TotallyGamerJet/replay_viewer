package packets

import "replay_viewer/protocol/codecs"

type JoinGame struct {
	EntityId  codecs.Int
	Gamemode  codecs.UnsignedByte
	Dimension codecs.Int
	//Difficulty codecs.UnsignedByte
	MaxPlayers   codecs.UnsignedByte
	LevelType    codecs.String
	viewDistance codecs.VarInt
	Debug        codecs.Boolean
}

func (_ JoinGame) Name() string { return "ServerJoinGamePacket" }

func (_ JoinGame) ID() int { return 0x25 }

func (p JoinGame) ReadPacketData(pkt *Packet) (holder Holder, err error) {
	p.EntityId, err = pkt.readInt()
	if err != nil {
		return
	}
	p.Gamemode, err = pkt.readUByte()
	if err != nil {
		return
	}
	p.Dimension, err = pkt.readInt()
	if err != nil {
		return
	}
	p.MaxPlayers, err = pkt.readUByte()
	if err != nil {
		return
	}
	p.LevelType, err = pkt.readString()
	if err != nil {
		return
	}
	p.viewDistance, err = pkt.readVarInt()
	if err != nil {
		return
	}
	p.Debug, err = pkt.readBool()
	if err != nil {
		return
	}
	return p, nil
}
