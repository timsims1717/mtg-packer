package main

import (
	"fmt"
	"math/rand"
	"strconv"
)

func simulateDraft(packs []Pack, playerCount, packCount, packSize int, set *Set, alg bool) []*Player {
	var players []*Player
	for i := 0; i < playerCount; i++ {
		players = append(players, &Player{
			PlayerNum:  i,
			Strats:     make(map[string]int),
			ManaValues: make(map[int]int),
		})
	}
	offset := 0
	mod := playerCount
	left := true
	n := 0
	pickNum := 0
	for n < packCount {
		m := 0
		if !left {
			m = packSize - 1
		}
		for (left && m < packSize) || (!left && m >= 0) {
			for i, player := range players {
				packI := (i+m)%mod + offset
				var pickI int
				if alg {
					pickI = pickCardAlg(player, packs[packI], packCount*packSize, set)
				} else {
					pickI = pickCardPerson(player, packs[packI], packI+1, pickNum+1)
				}
				if len(packs[packI]) > 1 {
					packs[packI] = append(packs[packI][:pickI], packs[packI][pickI+1:]...)
				} else {
					packs[packI] = []Card{}
				}
			}
			pickNum++
			if left {
				m++
			} else {
				m--
			}
		}
		n++
		left = !left
		offset += playerCount
	}
	return players
}

func pickCardPerson(player *Player, currPack Pack, packNum, pickNum int) int {
	fmt.Println("===================================")
	fmt.Printf("== Player %d - Pack %d - Pick %d ==\n", player.PlayerNum+1, packNum, pickNum)
	if len(player.CurrPicks) > 0 {
		PrintCardList(player.CurrPicks, "Current Picks")
	} else {
		fmt.Println("===================================")
	}
	for {
		PrintCardList(currPack, "")
		fmt.Println()
		var pickMsg string
		if len(currPack) > 1 {
			pickMsg = fmt.Sprintf("Pick (%d - %d): ", 1, len(currPack))
		} else {
			pickMsg = fmt.Sprint("Pick (1): ")
		}
		in := ConsoleInput(pickMsg)
		inI, err := strconv.Atoi(in)
		if err != nil || inI < 1 || inI > len(currPack) {
			PrintPackFull(player.CurrPicks, "Current Picks")
			ConsoleInput("Press any key ...")
		} else {
			fmt.Printf("Player %d picked %s\n\n", player.PlayerNum+1, currPack[inI-1].Name)
			player.CurrPicks = append(player.CurrPicks, currPack[inI-1])
			return inI - 1
		}
	}
}

func pickCardAlg(player *Player, currPack Pack, totalSize int, set *Set) int {
	var picks []int
	for _, card := range currPack {
		cardName := card.Name
		pickAdj := Max(len(player.CurrPicks), 10)

		score := 0
		rarityScore := 0
		colorScore := 0
		tagScore := 0
		typeScore := 0
		mvScore := 0

		// rarity
		switch card.Rarity {
		case "mythic":
			rarityScore += 4
		case "rare":
			rarityScore += 3
		case "uncommon":
			rarityScore += 1
		}
		if len(player.CurrPicks) == 0 {
			rarityScore *= 2
		}

		// card types
		isCreature := StringIn("isCreature", card.Tags)
		isLand := false
		isArtifact := false
		isEnchantment := false
		for _, t := range card.Types {
			if t == "creature" {
				isCreature = true
			} else if t == "land" {
				isLand = true
			} else if t == "artifact" {
				isArtifact = true
			} else if t == "enchantment" {
				isEnchantment = true
			}
			tagScore += Min(player.Strats[t], 4)
		}
		if isCreature {
			typeScore += Min(Max(0, pickAdj/(totalSize/15)-player.CreatCnt), 4)
			if isArtifact {
				tagScore += Min(player.Strats["artifactCreature"], 4)
			}
			if isEnchantment {
				tagScore += Min(player.Strats["enchantmentCreature"], 4)
			}
		}
		if isLand {
			typeScore += 2
		}

		// card subtypes
		for _, t := range card.SubTypes {
			tagScore += Min(player.Strats[t], 4)
		}

		// card supertypes
		for _, t := range card.SuperTypes {
			tagScore += Min(player.Strats[t], 4)
		}

		// current colors
		if card.Color != "multicolor" && (card.Color != "colorless" || !(isLand || isArtifact)) {
			colorScore += player.Strats[card.Color]
		}

		// mana curve
		tn := 0
		if !isLand {
			switch card.ManaValue {
			case 0, 1:
				tn = 4
			case 2:
				tn = 15
			case 3:
				tn = 12
			case 4:
				tn = 5
			case 5:
				tn = 3 + player.Strats["ramp"]/3
			default:
				tn = 2 + player.Strats["ramp"]/2
			}
			ctn := totalSize / tn
			mvScore += Max(0, pickAdj/ctn-player.ManaValues[card.ManaValue])
		}

		// current tags
		for _, tag := range card.Tags {
			switch tag {
			case "azorius":
				colorScore += Min(player.Strats["white"], player.Strats["blue"])
				tagScore += player.Strats[tag]
			case "dimir":
				colorScore += Min(player.Strats["blue"], player.Strats["black"])
				tagScore += player.Strats[tag]
			case "rakdos":
				colorScore += Min(player.Strats["black"], player.Strats["red"])
				tagScore += player.Strats[tag]
			case "gruul":
				colorScore += Min(player.Strats["red"], player.Strats["green"])
				tagScore += player.Strats[tag]
			case "selesnya":
				colorScore += Min(player.Strats["green"], player.Strats["white"])
				tagScore += player.Strats[tag]
			case "orzhov":
				colorScore += Min(player.Strats["white"], player.Strats["black"])
				tagScore += player.Strats[tag]
			case "izzet":
				colorScore += Min(player.Strats["blue"], player.Strats["red"])
				tagScore += player.Strats[tag]
			case "golgari":
				colorScore += Min(player.Strats["black"], player.Strats["green"])
				tagScore += player.Strats[tag]
			case "boros":
				colorScore += Min(player.Strats["red"], player.Strats["white"])
				tagScore += player.Strats[tag]
			case "simic":
				colorScore += Min(player.Strats["green"], player.Strats["blue"])
				tagScore += player.Strats[tag]
			case "bomb":
				tagScore += 4
			case "advantage":
				tagScore += 3
			case "removal":
				tagScore += 2
			case "protection":
				tagScore += 2
			case "aggro":
				tagScore += 2
			case "evasion":
				tagScore += 2
			case "cantrip":
				tagScore += 1
			case "fixing":
				colors := -1
				if player.Strats["white"] > 3 {
					colors++
				}
				if player.Strats["blue"] > 3 {
					colors++
				}
				if player.Strats["black"] > 3 {
					colors++
				}
				if player.Strats["red"] > 3 {
					colors++
				}
				if player.Strats["green"] > 3 {
					colors++
				}
				tagScore += Max(0, colors*2)
			case "modified":
				tagScore += CountModified(player.CurrPicks)
			case "spellslinger":
				tagScore += CountInstantSorceries(player.CurrPicks)
			case "artifact", "creature", "enchantment", "instant", "land", "planeswalker", "sorcery":
				tagScore += CountType(player.CurrPicks, tag)
			case "artifactCreature":
				tagScore += CountTwoTypes(player.CurrPicks, "artifact", "creature")
			case "enchantmentCreature":
				tagScore += CountTwoTypes(player.CurrPicks, "enchantment", "creature")
			case "farmAnimals":
				tagScore += CountFarmAnimals(player.CurrPicks)
			default:
				if StringIn(tag, set.SuperOrSub) {
					tagScore += CountSuperOrSub(player.CurrPicks, tag)
				} else {
					tagScore += player.Strats[tag]
				}
			}
		}

		// adjust score
		score += rarityScore
		score += mvScore
		score += typeScore
		score += tagScore
		score += colorScore

		// record
		picks = append(picks, score)
		fmt.Sprintln(cardName)
	}
	var bestI []int
	best := -1
	for i, score := range picks {
		if score > best {
			bestI = []int{i}
			best = score
		} else if score == best {
			bestI = append(bestI, i)
		}
	}
	pickI := bestI[rand.Intn(len(bestI))]
	pick := currPack[pickI]
	if pick.Color != "multicolor" {
		if player.Strats[pick.Color] == 0 {
			player.Strats[pick.Color]++
		}
		player.Strats[pick.Color]++
	}
	for _, tag := range pick.Tags {
		switch tag {
		case "advantage", "removal", "bomb", "aggro", "evasion", "isCreature", "protection", "cantrip", "fixing":
		default:
			player.Strats[tag]++
		}
		switch tag {
		case "azorius":
			player.Strats["white"]++
			player.Strats["blue"]++
		case "dimir":
			player.Strats["blue"]++
			player.Strats["black"]++
		case "rakdos":
			player.Strats["black"]++
			player.Strats["red"]++
		case "gruul":
			player.Strats["red"]++
			player.Strats["green"]++
		case "selesnya":
			player.Strats["green"]++
			player.Strats["white"]++
		case "orzhov":
			player.Strats["white"]++
			player.Strats["black"]++
		case "izzet":
			player.Strats["blue"]++
			player.Strats["red"]++
		case "golgari":
			player.Strats["black"]++
			player.Strats["green"]++
		case "boros":
			player.Strats["red"]++
			player.Strats["white"]++
		case "simic":
			player.Strats["green"]++
			player.Strats["blue"]++
		case "modified":
			player.Strats["equipment"]++
			player.Strats["aura"]++
			player.Strats["counters"]++
			player.Strats["+1/+1"]++
		case "spellslinger":
			player.Strats["instant"]++
			player.Strats["sorcery"]++
		case "farmAnimals":
			player.Strats["boar"]++
			player.Strats["goat"]++
			player.Strats["horse"]++
			player.Strats["ox"]++
			player.Strats["sheep"]++
		}
	}
	isCreature := false
	for _, t := range pick.Types {
		if t == "creature" {
			isCreature = true
		}
	}
	if isCreature {
		player.CreatCnt++
	}
	player.ManaValues[pick.ManaValue]++
	player.CurrPicks = append(player.CurrPicks, pick)
	return pickI
}
