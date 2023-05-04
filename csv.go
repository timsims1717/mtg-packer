package main

import (
	"encoding/csv"
	"fmt"
	"github.com/pkg/errors"
	"os"
)

func createCSVs(packs []Pack, players int, oFilePre string) error {
	var packsA []Pack
	packNum := 1
	for _, pack := range packs {
		packsA = append(packsA, pack)
		if len(packsA) == players {
			if err := createCSV(packsA, oFilePre, packNum); err != nil {
				return err
			}
			packNum++
			packsA = []Pack{}
		}
	}
	return nil
}

func createCSV(packs []Pack, oFilePre string, packNum int) error {
	errMsg := "create CSV file"
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
