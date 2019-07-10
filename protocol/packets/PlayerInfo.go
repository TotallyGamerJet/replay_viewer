package packets

import (
	"replay_viewer/protocol/codecs"
)

type PlayerInfo struct {
	action       codecs.VarInt
	numOfPlayers codecs.VarInt
	player       struct {
		uuid codecs.UUID
		//Action == 0 (Add Player)
		action0 struct {
			name       codecs.String
			numOfProps codecs.VarInt
			property   struct {
				name      codecs.String
				value     codecs.String
				isSigned  codecs.Boolean
				signature codecs.String
			}
			gamemode       codecs.VarInt
			ping           codecs.VarInt
			hasDisplayName codecs.Boolean
			displayName    codecs.String //Optional Chat
		}
		//Action == 1 (Update Gamemode)
		action1 struct {
			gameMode codecs.VarInt
		}
		//Action == 2 (Update Latency)
		action2 struct {
			ping codecs.VarInt
		}
		//Action == 3 (Update Display Name)
		action3 struct {
			hasDisplayName codecs.Boolean
			displayName    codecs.String //Optional Chat
		}
		//Action == 4 (Remove Player)
		action4 struct {
			//no fields
		}
	}
}

//Name returns the name of the packet as a string
func (_ PlayerInfo) Name() string { return "ServerPlayerListEntryPacket" }

//ID returns the id in hex of the packet
func (_ PlayerInfo) ID() int { return 0x33 }

func (p PlayerInfo) ReadPacketData(pkt *Packet) (holder Holder, err error) {
	p.action, err = pkt.readVarInt()
	if err != nil {
		return
	}
	p.numOfPlayers, err = pkt.readVarInt()
	if err != nil {
		return
	}
	switch p.action {
	case 0:
		action0 := p.player.action0
		action0.name, err = pkt.readString()
		action0.numOfProps, err = pkt.readVarInt()
		property := action0.property
		property.name, err = pkt.readString()
		property.value, err = pkt.readString()
		property.isSigned, err = pkt.readBool()
		if property.isSigned {
			property.signature, err = pkt.readString()
		}
		action0.property = property
		action0.gamemode, err = pkt.readVarInt()
		action0.ping, err = pkt.readVarInt()
		action0.hasDisplayName, err = pkt.readBool()
		if action0.hasDisplayName {
			action0.displayName, err = pkt.readString()
		}
		p.player.action0 = action0
	case 1:
		p.player.action1.gameMode, err = pkt.readVarInt()
	case 2:
		p.player.action2.ping, err = pkt.readVarInt()
	case 3:
		p.player.action3.hasDisplayName, err = pkt.readBool()
		if p.player.action3.hasDisplayName {
			p.player.action3.displayName, err = pkt.readString()
		}
	case 4:
		//Nothing happens
	}
	return p, nil
}
