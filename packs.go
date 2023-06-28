package main

func CountInstantSorceries(pack Pack) int {
	count := 0
	for _, card := range pack {
		if StringIn("instant", card.Types) || StringIn("sorcery", card.Types) {
			count++
		}
	}
	return count
}

func CountTwoTypes(pack Pack, t1, t2 string) int {
	count := 0
	for _, card := range pack {
		isT1 := false
		isT2 := false
		if StringIn(t1, card.Types) {
			isT1 = true
		} else if StringIn(t2, card.Types) {
			isT2 = true
		}
		if isT1 && isT2 {
			count++
		}
	}
	return count
}

func CountType(pack Pack, t string) int {
	count := 0
	for _, card := range pack {
		if StringIn(t, card.Types) {
			count++
		}
	}
	return count
}

func CountSuperOrSub(pack Pack, t string) int {
	count := 0
	for _, card := range pack {
		if StringIn(t, card.SuperTypes) {
			count++
		}
		if StringIn(t, card.SubTypes) {
			count++
		}
	}
	return count
}

func CountFarmAnimals(pack Pack) int {
	count := 0
	for _, card := range pack {
		if StringIn("boar", card.SubTypes) {
			count++
		}
		if StringIn("goat", card.SubTypes) {
			count++
		}
		if StringIn("horse", card.SubTypes) {
			count++
		}
		if StringIn("ox", card.SubTypes) {
			count++
		}
		if StringIn("sheep", card.SubTypes) {
			count++
		}
	}
	return count
}

func CountModified(pack Pack) int {
	count := 0
	for _, card := range pack {
		if StringIn("equipment", card.SubTypes) || StringIn("aura", card.SubTypes) ||
			StringIn("counters", card.Tags) || StringIn("+1/+1", card.Tags) {
			count++
		}
	}
	return count
}
