package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"os"
	"strconv"

	"github.com/ahl5esoft/golang-underscore"
)

// RentalInfo is struct
type RentalInfo struct {
	BikeID           string `json:"bikeid"`
	StartTime        string `json:"starttime"`
	StartStationID   string `json:"startstationid"`
	StartStationName string `json:"startstationname"`
	EndTime          string `json:"endtime"`
	EndStationID     string `json:"endstationid"`
	Duration         string `json:"duration"`
	Distance         string `json:"distance"`
}

func main() {
	filename := "sample"
	csvFile, _ := os.Open(filename + ".csv")
	defer csvFile.Close()

	// Read File into a Variable
	lines, err := csv.NewReader(bufio.NewReader(csvFile)).ReadAll()
	if err != nil {
		panic(err)
	}

	fmt.Println("total lines:", len(lines))

	// Loop through lines & turn into object
	var data []RentalInfo
	for _, line := range lines[1:] {
		data = append(data, RentalInfo{
			BikeID:           line[0],
			StartTime:        line[1],
			StartStationID:   line[2],
			StartStationName: line[3],
			EndTime:          line[5],
			EndStationID:     line[6],
			Duration:         line[9],
			Distance:         line[10],
		})
	}
	// dataJSON, _ := json.Marshal(data)
	// fmt.Println(string(dataJSON))

	v := underscore.GroupBy(data, "StartStationID")
	dict, ok := v.(map[string][]RentalInfo)
	if !ok {
		panic("error...")
	}

	fmt.Println("total station:", len(dict))

	file, err := os.Create(filename + "_output.csv")
	if err != nil {
		panic(err)
	}

	wr := csv.NewWriter(bufio.NewWriter(file))
	for _, g := range dict {
		wr.Write([]string{g[0].StartStationID, g[0].StartStationName, strconv.Itoa(len(g))})
		// fmt.Println(g[0].StartStationID, g[0].StartStationName, len(g))
	}
	wr.Flush()
}
