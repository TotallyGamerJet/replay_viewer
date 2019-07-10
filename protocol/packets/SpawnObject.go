package packets

import "replay_viewer/protocol/codecs"

type SpawnObject struct {
	entityID         codecs.VarInt
	objectUUID       codecs.UUID
	objectType       codecs.Byte
	x, y, z          codecs.Double
	pitch            codecs.Byte
	yaw              codecs.Byte
	data             codecs.Int
	velX, velY, velZ codecs.Short
}

//Name returns the name of the packet as a string
func (_ SpawnObject) Name() string { return "SpawnObject" }

//ID returns the id in hex of the packet
func (_ SpawnObject) ID() int { return 0x00 }

func (s SpawnObject) ReadPacketData(pkt *Packet) (holder Holder, err error) {
	s.entityID, err = pkt.readVarInt()
	if err != nil {
		return
	}
	s.objectUUID, err = pkt.readUUID()
	if err != nil {
		return
	}
	s.objectType, err = pkt.readByte()
	if err != nil {
		return
	}
	s.x, err = pkt.readDouble()
	if err != nil {
		return
	}
	s.y, err = pkt.readDouble()
	if err != nil {
		return
	}
	s.z, err = pkt.readDouble()
	if err != nil {
		return
	}
	s.pitch, err = pkt.readByte()
	if err != nil {
		return
	}
	s.yaw, err = pkt.readByte()
	if err != nil {
		return
	}
	s.data, err = pkt.readInt()
	if err != nil {
		return
	}
	if s.data > 0 {
		s.velX, err = pkt.readShort()
		if err != nil {
			return
		}
		s.velY, err = pkt.readShort()
		if err != nil {
			return
		}
		s.velZ, err = pkt.readShort()
		if err != nil {
			return
		}
	}
	return s, nil
}
