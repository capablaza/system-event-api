package main

import "strings"

func translateDate(dateInput string) string {
	r := strings.NewReplacer(
		"Jan", "Ene",
		"Feb", "Feb",
		"Mar", "Mar",
		"Apr", "Abr",
		"May", "May",
		"Jun", "Jun",
		"Jul", "Jul",
		"Aug", "Ago",
		"Sep", "Sep",
		"Oct", "Oct",
		"Nov", "Nov",
		"Dec", "Dic")

	return r.Replace(dateInput)
}
