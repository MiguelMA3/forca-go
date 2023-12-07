package config

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/BurntSushi/toml"
)

type Config struct {
	Themes map[string][]string `toml:"themes"`
}

func ChooseWord() (string, string) {
	arquivoTOML := "palavras.toml"

	var config Config

	if _, err := toml.DecodeFile(arquivoTOML, &config); err != nil {
		fmt.Println("Erro ao ler o arquivo TOML:", err)
		return "", ""
	}

	rand.Seed(time.Now().UnixNano())

	themes := make([]string, 0, len(config.Themes))
	for theme := range config.Themes {
		themes = append(themes, theme)
	}
	themeSelected := themes[rand.Intn(len(themes))]

	themeWord := config.Themes[themeSelected]
	wordSelected := themeWord[rand.Intn(len(themeWord))]

	return wordSelected, themeSelected
}
