package main

import (
	"fmt"
	"strings"
)

func PrintSlots(set *Set) {
	for i, slot := range set.Slots {
		c := "any"
		r := "any"
		t := "any"
		if len(slot.Color) > 0 {
			sb := new(strings.Builder)
			for j, color := range slot.Color {
				if j != 0 {
					sb.WriteString(" or ")
				}
				sb.WriteString(color)
			}
			c = sb.String()
		}
		if len(slot.Rarity) > 0 {
			sb := new(strings.Builder)
			for j, rarity := range slot.Rarity {
				if j != 0 {
					sb.WriteString(" or ")
				}
				sb.WriteString(rarity)
			}
			r = sb.String()
		}
		if len(slot.Type) > 0 {
			sb := new(strings.Builder)
			for j, tipe := range slot.Type {
				if j != 0 {
					sb.WriteString(" or ")
				}
				sb.WriteString(tipe)
			}
			t = sb.String()
		}
		fmt.Printf("Slot %02d - color: %s, rarity: %s, type: %s\n", i+1, c, r, t)
	}
}

func PrintPacks(packs []Pack) {
	for i, pack := range packs {
		fmt.Println()
		fmt.Println("===================================")
		fmt.Printf("Pack %d\n", i+1)
		fmt.Println("===================================")
		for _, card := range pack {
			//fmt.Printf("%02d: %s\n", j+1, card.Name)
			fmt.Printf("%s\n", card.Name)
		}
	}
}
