package main

import (
	"fmt"
	"mclib-go/protocol"
	"os"
	"replay_viewer/protocol/packets"
	"replay_viewer/utils"
)

func main() {
	utils.Unzip("test_recording.mcpr", "./output")
	metaData := loadMetaData("output")
	fmt.Println(metaData)
	readRecording("./output/recording.tmcpr")
	for i, p := range packets.PacketList {
		fmt.Println(i, p.ID())
	}
}

func readRecording(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	for i := 0; i < 5; i++ {
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
		fmt.Printf("#%d t:%d l:%d id:%d\n", i, t, l, pkt.ID)

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
	//fmt.Println(holder)
}
