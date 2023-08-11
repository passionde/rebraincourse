package random

import (
	"github.com/passionde/rebraincourse/GO-03/wordz"
	"strings"
)

var citiesMap = map[string]string{}

func init() {
	cityNames := []string{
		"Лондон",
		"Москва",
		"Париж",
		"Сургут",
		"Великий Новгород",
	}

	for i, v := range wordz.Words {
		citiesMap[v] = cityNames[i%len(cityNames)]
	}
}

func City() string {
	return citiesMap[strings.ReplaceAll(wordz.Random(), wordz.Prefix, "")]
}

func Digit() string {
	return strings.ReplaceAll(wordz.Random(), wordz.Prefix, "")
}
