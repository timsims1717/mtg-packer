package main

import (
	"encoding/json"
	"flag"
	"math/rand"
	"os"
	"time"
)

func main() {
	flag.Parse()
	rand.New(rand.NewSource(time.Now().UnixNano()))
	if *args.filename == "" {
		flag.PrintDefaults()
		panic("no set input file specified")
	}
	contents, err := os.ReadFile(*args.filename)
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
	for j := 0; j < *args.playerCount; j++ {
		for i := 0; i < *args.packCount; i++ {
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
	}
	PrintPacks(packs)
	if err != nil {
		panic(err)
	}
	if !*args.noDraft {
		players := simulateDraft(packs, *args.playerCount, *args.packCount, len(set.Slots), &set, *args.simulate)
		PrintDraft(players)
		err = createPlayersCSV(players, *args.outFile, *args.packCount*len(set.Slots))
		if err != nil {
			panic(err)
		}
	}
	err = createPackCSVs(packs, *args.playerCount, *args.outFile)
}

var args struct {
	packCount   *int
	playerCount *int
	noDraft     *bool
	simulate    *bool
	filename    *string
	outFile     *string
}

func init() {
	args.packCount = flag.Int("packs", 3, "number of packs for each player")
	args.playerCount = flag.Int("players", 8, "number of players")
	args.noDraft = flag.Bool("nodraft", false, "turn off the draft")
	args.simulate = flag.Bool("simulate", false, "turn on the draft algorithm")
	args.filename = flag.String("setfile", "", "name of the mtg set input file")
	args.outFile = flag.String("outfile", "pack", "prefix of the output files")
}
