package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
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

func PrintDraft(players []*Player) {
	for _, player := range players {
		PrintPackFull(player.CurrPicks, fmt.Sprintf("Player %d\n", player.PlayerNum+1))
	}
}

func PrintPackFull(cards []Card, head string) {
	fmt.Println("===================================")
	fmt.Println(head)
	fmt.Println("-----------------------------------")
	for _, card := range cards {
		fmt.Println(FullCardString(card, -1))
	}
	fmt.Println("===================================")
}

func PrintCardList(cards []Card, head string) {
	if head != "" {
		fmt.Println("===================================")
		fmt.Println(head)
		fmt.Println("-----------------------------------")
	}
	if len(cards) < 9 {
		for i, card := range cards {
			s := FullCardString(card, i+1)
			fmt.Println(s)
		}
	} else if len(cards) < 17 {
		half := len(cards)/2 + 1
		for i := 0; i < half; i++ {
			var s2 string
			c1 := cards[i]
			s1 := FullCardString(c1, i+1)
			if i+half < len(cards) {
				c2 := cards[i+half]
				s2 = FullCardString(c2, i+half+1)
			}
			fmt.Printf("%54s | %54s\n", s1, s2)
		}
	} else {
		third := len(cards)/3 + 1
		for i := 0; i < third; i++ {
			var s2, s3 string
			c1 := cards[i]
			s1 := FullCardString(c1, i+1)
			if i+third < len(cards) {
				c2 := cards[i+third]
				s2 = FullCardString(c2, i+third+1)
			}
			if i+third+third < len(cards) {
				c3 := cards[i+third+third]
				s3 = FullCardString(c3, i+third+1)
			}
			fmt.Printf("%54s | %54s | %54s\n", s1, s2, s3)
		}
	}
	if head != "" {
		fmt.Println("===================================")
	}
}

func FullCardString(card Card, i int) string {
	var sb strings.Builder
	sb.WriteString(card.Name)
	sb.WriteString(" ~")
	for _, st := range card.SuperTypes {
		sb.WriteString(" ")
		sb.WriteString(strings.ToTitle(st))
	}
	for _, st := range card.Types {
		sb.WriteString(" ")
		sb.WriteString(strings.ToTitle(st))
	}
	start := true
	for _, st := range card.SubTypes {
		if start {
			sb.WriteString(" - ")
		} else {
			sb.WriteString(" ")
		}
		sb.WriteString(strings.ToTitle(st))
		start = false
	}
	title := Truncate(sb.String(), 40)
	mv := card.Pips
	if mv == "" {
		if !StringIn("land", card.Types) {
			mv = strconv.Itoa(card.ManaValue)
		} else {
			mv = "LAND"
		}
	}
	if i > 0 {
		return fmt.Sprintf("%03d: %6s - %-40s", i, mv, title)
	} else {
		return fmt.Sprintf("%6s - %-40s", mv, title)
	}
}

func ConsoleInput(msg string) string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print(msg)
	input, _ := reader.ReadString('\n')
	return strings.Replace(input, "\n", "", -1)
}

func Truncate(str string, length int) string {
	if length <= 0 {
		return ""
	}

	orgLen := len(str)
	if orgLen <= length {
		return str
	}
	return str[:length]
}
