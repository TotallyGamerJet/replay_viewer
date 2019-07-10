package packets

import (
	"replay_viewer/protocol/codecs"
)

type ChunkData struct {
	chunkX, chunkZ codecs.Int
	fullChunk      codecs.Boolean
	primaryBitMask codecs.VarInt
	//heightMaps         NBT
	size               codecs.VarInt
	data               codecs.ByteArray
	numOfBlockEntities codecs.VarInt
	//blockEntities      []NBT
}

//Name returns the name of the packet as a string
func (_ ChunkData) Name() string { return "ServerChunkDataPacket" }

//ID returns the id in hex of the packet
func (_ ChunkData) ID() int { return 0x21 }

func (p ChunkData) ReadPacketData(pkt *Packet) (holder Holder, err error) {
	panic("not implemented")
	if p.chunkX, err = pkt.readInt(); err != nil {
		return
	}
	if p.chunkZ, err = pkt.readInt(); err != nil {
		return
	}
	if p.fullChunk, err = pkt.readBool(); err != nil {
		return
	}
	if p.primaryBitMask, err = pkt.readVarInt(); err != nil {
		return
	}
	//TODO: NBT
	if p.size, err = pkt.readVarInt(); err != nil {
		return
	}
	if p.data, err = pkt.readByteArray(); err != nil {
		return
	}
	if p.numOfBlockEntities, err = pkt.readVarInt(); err != nil {
		return
	}
	//TODO: NBT Array
	return p, nil
}
