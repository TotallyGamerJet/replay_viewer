package main

import (
	"encoding/json"
	"replay_viewer/utils"
	"strconv"
	"time"
)

func loadMetaData(mcprFile string) MetaData {
	data, err := utils.GetZippedFile(mcprFile, "metaData.json")
	if err != nil {
		panic(err)
	}

	metaData := MetaData{}
	err = json.Unmarshal(data.Bytes(), &metaData)
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
