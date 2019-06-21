package packets

import (
	"replay_viewer/protocol/codecs"
)

type PlayKeepAlive struct {
	AliveId codecs.VarInt
}

func (_ PlayKeepAlive) ID() int { return 0x1F }

type PlayJoinGame struct {
	EntityId   codecs.Int
	Gamemode   codecs.UnsignedByte
	Dimension  codecs.Int
	Difficulty codecs.UnsignedByte
	MaxPlayers codecs.UnsignedByte
	LevelType  codecs.String
	Debug      codecs.Boolean
}

func (_ PlayJoinGame) ID() int { return 0x23 }

type PlaySpawnPosition struct {
	Location codecs.Long
}

func (_ PlaySpawnPosition) ID() int { return 0x43 }

type PlayPositionAndLook struct {
	X     codecs.Double
	Y     codecs.Double
	Z     codecs.Double
	Yaw   codecs.Float
	Pitch codecs.Float
	Flags codecs.Byte
	Data  codecs.VarInt
}

func (_ PlayPositionAndLook) ID() int { return 0x2E }
