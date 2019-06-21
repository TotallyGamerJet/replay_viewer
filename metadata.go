package main

import (
	"encoding/json"
	"io/ioutil"
	"strconv"
	"time"
)

func loadMetaData(foldername string) MetaData {
	file, err := ioutil.ReadFile(foldername + "/metaData.json")
	if err != nil {
		panic("file doesn't exist")
	}
	/* FAILED ATTEMPT to read files inside the .mcpr
	file, err := utils.GetFile("test_recording.mcpr", "metaData.json")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	var readData []byte
	n, err := file.Read(readData)
	if err != nil || n == 0 {
		panic(err)
	}
	fmt.Println(readData)*/

	metaData := MetaData{}
	err = json.Unmarshal(file, &metaData)
	if err != nil {
		panic("metadata format incorrect?")
	}
	return metaData
}

func (data MetaData) getDate() time.Time {
	return time.Unix(0, int64(data.Date)*int64(time.Millisecond))
}

func (data MetaData) getDuration() time.Duration {
	d, err := time.ParseDuration(strconv.Itoa(data.Duration) + "ms")
	if err != nil {
		panic(err)
	}
	return d
}

//MetaData the metadata of the replay
type MetaData struct {
	SinglePlayer      bool     `json:"singleplayer"`
	ServerName        string   `json:"serverName"`
	Duration          int      `json:"duration"`
	Date              int      `json:"date"`
	McVersion         string   `json:"mcversion"`
	FileFormat        string   `json:"fileFormat"`
	FileFormatVersion int      `json:"fileFormatVersion"`
	Protocol          int      `json:"protocol"`
	Generator         string   `json:"generator"`
	SelfID            int      `json:"selfId"`
	Players           []string `json:"players"`
}
