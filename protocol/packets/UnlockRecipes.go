package packets

import "replay_viewer/protocol/codecs"

type UnlockRecipes struct {
	action      codecs.VarInt
	craftOpen   codecs.Boolean
	craftFilter codecs.Boolean
	smeltOpen   codecs.Boolean
	smeltFilter codecs.Boolean
	arraySize   codecs.VarInt
	recipeId    []codecs.Identifier
	arraySize2  codecs.VarInt
	recipeId2   []codecs.Identifier
}

//Name returns the name of the packet as a string
func (_ UnlockRecipes) Name() string { return "ServerUnlockRecipesPacket" }

//ID returns the id in hex of the packet
func (_ UnlockRecipes) ID() int { return 0x36 }

func (p UnlockRecipes) ReadPacketData(pkt *Packet) (holder Holder, err error) {
	p.action, err = pkt.readVarInt()
	if err != nil {
		return
	}
	p.craftOpen, err = pkt.readBool()
	if err != nil {
		return
	}
	p.craftFilter, err = pkt.readBool()
	if err != nil {
		return
	}
	p.smeltOpen, err = pkt.readBool()
	if err != nil {
		return
	}
	p.smeltFilter, err = pkt.readBool()
	if err != nil {
		return
	}
	p.arraySize, err = pkt.readVarInt()
	if err != nil {
		return
	}
	if p.arraySize > 0 {
		for i := 0; i < int(p.arraySize); i++ {
			id, err := pkt.readIdentifier()
			if err != nil {
				break
			}
			p.recipeId = append(p.recipeId, id)
		}
	}
	p.arraySize2, err = pkt.readVarInt()
	if err != nil {
		return
	}
	if p.arraySize2 > 0 {
		for i := 0; i < int(p.arraySize2); i++ {
			id, err := pkt.readIdentifier()
			if err != nil {
				break
			}
			p.recipeId = append(p.recipeId, id)
		}
	}
	return p, nil
}
