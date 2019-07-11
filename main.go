package main

import (
	"fmt"
	"mclib-go/protocol"
	"replay_viewer/protocol/packets"
	"replay_viewer/utils"
)

func main() {
	metaData := loadMetaData("test_recording.mcpr")
	fmt.Println(metaData)
	readRecording("test_recording.mcpr")
}

func readRecording(mcprFile string) {
	file, err := utils.GetZippedFile(mcprFile, "recording.tmcpr")
	if err != nil {
		panic(err)
	}
	//file := bytes.NewBuffer(data)
	for i := 0; i < 18; i++ {
		t, err := packets.ReadInt(file)
		if err != nil {
			panic(err)
		}
		l, err := packets.ReadInt(file)
		if err != nil {
			panic(err)
		}
		//file.Seek(int64(l), 1)
		pkt, err := packets.ReadPacket(file, l)
		if err != nil {
			panic(err)
		}
		fmt.Printf("#%d t:%d l:%d id:%d ", i, t, l, pkt.ID)

		readPkt(pkt)
	}
}

func readPkt(pkt *packets.Packet) {
	holder, ok := packets.PacketList[pkt.ID]
	if !ok {
		panic(protocol.UnknownPacketType)
	}
	holder, err := holder.ReadPacketData(pkt)
	if err != nil {
		panic(err)
	}
	fmt.Printf("n:%s\n", holder.Name())
	//fmt.Println(holder)
}
