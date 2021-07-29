package piscine

func UltimateDivMod(a *int, b *int) {
	tmpa := *a
	tmpb := *b

	*a = tmpa / tmpb
	*b = tmpa % tmpb
}
