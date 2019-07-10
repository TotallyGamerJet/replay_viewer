package packets

import "replay_viewer/protocol/codecs"

type UpdateLight struct {
	chunkX, chunkZ      codecs.VarInt
	skyLightMask        codecs.VarInt
	blockLightMask      codecs.VarInt
	emptySkyLightMask   codecs.VarInt
	emptyBlockLightMask codecs.VarInt
	skyLightArray       struct {
		length codecs.VarInt
		array  codecs.ByteArray
	}
	blockLightArray struct {
		length codecs.VarInt
		array  codecs.ByteArray
	}
}

//Name returns the name of the packet as a string
func (_ UpdateLight) Name() string { return "ServerUpdateLightPacket" }

//ID returns the id in hex of the packet
func (_ UpdateLight) ID() int { return 0x24 }

func (p UpdateLight) ReadPacketData(pkt *Packet) (holder Holder, err error) {
	p.chunkX, err = pkt.readVarInt()
	if err != nil {
		return
	}
	p.chunkZ, err = pkt.readVarInt()
	if err != nil {
		return
	}
	p.skyLightMask, err = pkt.readVarInt()
	if err != nil {
		return
	}
	p.blockLightMask, err = pkt.readVarInt()
	if err != nil {
		return
	}
	p.emptySkyLightMask, err = pkt.readVarInt()
	if err != nil {
		return
	}
	p.emptyBlockLightMask, err = pkt.readVarInt()
	if err != nil {
		return
	}
	p.skyLightArray.array, err = pkt.readByteArray()
	if err != nil {
		return
	}
	p.skyLightArray.length = codecs.VarInt(len(p.skyLightArray.array))
	p.blockLightArray.array, err = pkt.readByteArray()
	if err != nil {
		return
	}
	p.blockLightArray.length = codecs.VarInt(len(p.blockLightArray.array))
	return p, nil
}
