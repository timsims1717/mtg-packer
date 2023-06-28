package main

import (
	"encoding/csv"
	"fmt"
	"github.com/pkg/errors"
	"os"
)

func createPackCSVs(packs []Pack, players int, oFilePre string) error {
	var packsA []Pack
	packNum := 1
	for _, pack := range packs {
		packsA = append(packsA, pack)
		if len(packsA) == players {
			if err := createPackCSV(packsA, oFilePre, packNum); err != nil {
				return err
			}
			packNum++
			packsA = []Pack{}
		}
	}
	return nil
}

func createPackCSV(packs []Pack, oFilePre string, packNum int) error {
	errMsg := "create pack CSV file"
	pl := len(packs[0])
	outFile, err := os.Create(fmt.Sprintf("%s-%d.csv", oFilePre, packNum))
	if err != nil {
		return errors.Wrap(err, errMsg)
	}
	csvWriter := csv.NewWriter(outFile)
	var headers []string
	for i := 0; i < len(packs); i++ {
		headers = append(headers, fmt.Sprintf("Pack %d", i+1))
	}
	csvWriter.Write(headers)
	for i := 0; i < pl; i++ {
		var row []string
		for j := 0; j < len(packs); j++ {
			row = append(row, packs[j][i].Name)
		}
		csvWriter.Write(row)
	}
	csvWriter.Flush()
	return nil
}

func createPlayersCSV(players []*Player, oFilePre string, pickNum int) error {
	errMsg := "create players CSV file"
	outFile, err := os.Create(fmt.Sprintf("%s-Draft.csv", oFilePre))
	if err != nil {
		return errors.Wrap(err, errMsg)
	}
	csvWriter := csv.NewWriter(outFile)
	headers := []string{""}
	for _, player := range players {
		headers = append(headers, fmt.Sprintf("Player %d", player.PlayerNum+1))
	}
	csvWriter.Write(headers)
	for i := 0; i < pickNum; i++ {
		row := []string{fmt.Sprintf("Pick %d", i+1)}
		for _, player := range players {
			if len(player.CurrPicks) > i {
				row = append(row, player.CurrPicks[i].Name)
			} else {
				row = append(row, "")
			}
		}
		csvWriter.Write(row)
	}
	csvWriter.Flush()
	return nil
}
