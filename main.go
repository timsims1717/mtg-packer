package main

import (
	"encoding/json"
	"io/ioutil"
	"math/rand"
	"os"
	"strconv"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	argArray := os.Args
	if len(argArray) < 2 {
		panic("no filename included")
	}
	setFName := os.Args[1]
	packCount := 8
	if len(argArray) > 2 {
		i, err := strconv.Atoi(argArray[2])
		if err != nil {
			panic(err)
		}
		packCount = i
	}
	contents, err := ioutil.ReadFile(setFName)
	if err != nil {
		panic(err)
	}
	var set Set
	err = json.Unmarshal(contents, &set)
	if err != nil {
		panic(err)
	}
	PrintSlots(&set)
	Sort(&set)
	AllowedInSlots(&set)
	var packs []Pack
	for i := 0; i < packCount; i++ {
		var usedNames []string
		var pack Pack
		for _, slot := range set.Slots {
			for {
				card := slot.Cards[rand.Intn(len(slot.Cards))]
				if !StringIn(card.Name, usedNames) {
					usedNames = append(usedNames, card.Name)
					pack = append(pack, card)
					break
				}
			}
		}
		packs = append(packs, pack)
	}
	PrintPacks(packs)
}
