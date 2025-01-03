package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Band struct {
	Name    string `json:"name"`
	Live    bool   `json:"live"`
	Members int8   `json:"members"`
}

func NewBand(name string, live bool, members int8) Band {
	result := Band{
		name,
		live,
		members,
	}
	return result
}

func (b *Band) ToString() string {
	playing := ""
	if !b.Live {
		playing = "not "
	}
	return fmt.Sprintf("the band name is %s and is %splaying with his %d members", b.Name, playing, b.Members)
}

func main() {
	fmt.Println("hello")
	acdc := NewBand("AC/DC", true, 5)
	fmt.Println(acdc.ToString())
	dataRow, err := os.ReadFile("band.json")
	if err != nil {
		os.Exit(1)
	}

	aero := Band{}
	err = json.Unmarshal(dataRow, &aero)
	if err != nil {
		os.Exit(1)
	}

	fmt.Println(aero.ToString())

	aero.Members = 44

	result, err := json.Marshal(&aero)
	if err != nil {
		os.Exit(1)
	}

	os.WriteFile("aero.json", result, 0777)

	var bands []Band
	bandsRow, err := os.ReadFile("bands.json")

	if err != nil {
		os.Exit(1)
	}

	json.Unmarshal(bandsRow, &bands)
	fmt.Println(bands)
}
