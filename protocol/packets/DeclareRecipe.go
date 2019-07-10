package packets

import "replay_viewer/protocol/codecs"

type DeclareRecipe struct {
	numRecipes codecs.VarInt
	Recipe     struct {
		rType codecs.Identifier
		iD    codecs.Identifier
		//data Optional
	}
}

//Name returns the name of the packet as a string
func (_ DeclareRecipe) Name() string { return "ServerDeclareRecipesPacket" }

//ID returns the id in hex of the packet
func (_ DeclareRecipe) ID() int { return 0x5A }

func (p DeclareRecipe) ReadPacketData(pkt *Packet) (holder Holder, err error) {
	p.numRecipes, err = pkt.readVarInt()
	if err != nil {
		return
	}
	p.Recipe.rType, err = pkt.readIdentifier()
	if err != nil {
		return
	}
	return p, nil
}
