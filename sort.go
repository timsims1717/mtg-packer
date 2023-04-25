package main

import "strings"

func Sort(set *Set) {
	for _, card := range set.Cards {
		// color
		if strings.ToLower(card.Color) == "white" {
			set.White = append(set.White, card)
		} else if strings.ToLower(card.Color) == "blue" {
			set.Blue = append(set.Blue, card)
		} else if strings.ToLower(card.Color) == "black" {
			set.Black = append(set.Black, card)
		} else if strings.ToLower(card.Color) == "red" {
			set.Red = append(set.Red, card)
		} else if strings.ToLower(card.Color) == "green" {
			set.Green = append(set.Green, card)
		} else if strings.ToLower(card.Color) == "multicolor" {
			set.Multicolor = append(set.Multicolor, card)
		} else if strings.ToLower(card.Color) == "colorless" {
			set.Colorless = append(set.Colorless, card)
		}
		// rarity
		if strings.ToLower(card.Rarity) == "mythic" {
			set.Mythic = append(set.Mythic, card)
		} else if strings.ToLower(card.Rarity) == "rare" {
			set.Rare = append(set.Rare, card)
		} else if strings.ToLower(card.Rarity) == "uncommon" {
			set.Uncommon = append(set.Uncommon, card)
		} else if strings.ToLower(card.Rarity) == "common" {
			set.Common = append(set.Common, card)
		}
		// types
		for _, tipe := range card.Types {
			if strings.ToLower(tipe) == "artifact" {
				set.Artifact = append(set.Artifact, card)
			} else if strings.ToLower(tipe) == "battle" {
				set.Battle = append(set.Battle, card)
			} else if strings.ToLower(tipe) == "creature" {
				set.Creature = append(set.Creature, card)
			} else if strings.ToLower(tipe) == "enchantment" {
				set.Enchantment = append(set.Enchantment, card)
			} else if strings.ToLower(tipe) == "instant" {
				set.Instant = append(set.Instant, card)
			} else if strings.ToLower(tipe) == "land" {
				set.Land = append(set.Land, card)
			} else if strings.ToLower(tipe) == "planeswalker" {
				set.Planeswalker = append(set.Planeswalker, card)
			} else if strings.ToLower(tipe) == "sorcery" {
				set.Sorcery = append(set.Sorcery, card)
			} else if strings.ToLower(tipe) == "tribal" {
				set.Tribal = append(set.Tribal, card)
			}
		}
	}
}

func AllowedInSlots(set *Set) {
	for i, slot := range set.Slots {
		for _, card := range set.Cards {
			isLegal := false
			if legal(card.Color, slot.Color) && legal(card.Rarity, slot.Rarity) {
				for _, tipe := range card.Types {
					if legal(tipe, slot.Type) {
						isLegal = true
					}
				}
			}
			if isLegal {
				set.Slots[i].Cards = append(set.Slots[i].Cards, card)
			}
		}
	}
}

func legal(s string, a []string) bool {
	if len(a) == 0 {
		return true
	}
	for _, as := range a {
		if as == s {
			return true
		}
	}
	return false
}

func StringIn(s string, a []string) bool {
	for _, as := range a {
		if as == s {
			return true
		}
	}
	return false
}
