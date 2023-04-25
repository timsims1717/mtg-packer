package main

type Card struct {
	Color  string   `json:"color"`
	Name   string   `json:"name"`
	Rarity string   `json:"rarity"`
	Types  []string `json:"types"`
}

type Slot struct {
	Color  []string `json:"color"`
	Rarity []string `json:"rarity"`
	Type   []string `json:"type"`

	Cards []Card `json:"-"`
}

type Pack []Card

type Set struct {
	Cards []Card `json:"cards"`
	Slots []Slot `json:"slots"`

	Mythic   []Card `json:"-"`
	Rare     []Card `json:"-"`
	Uncommon []Card `json:"-"`
	Common   []Card `json:"-"`

	White      []Card `json:"-"`
	Blue       []Card `json:"-"`
	Black      []Card `json:"-"`
	Red        []Card `json:"-"`
	Green      []Card `json:"-"`
	Multicolor []Card `json:"-"`
	Colorless  []Card `json:"-"`

	Artifact     []Card `json:"-"`
	Battle       []Card `json:"-"`
	Creature     []Card `json:"-"`
	Enchantment  []Card `json:"-"`
	Instant      []Card `json:"-"`
	Land         []Card `json:"-"`
	Planeswalker []Card `json:"-"`
	Sorcery      []Card `json:"-"`
	Tribal       []Card `json:"-"`
}
