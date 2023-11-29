package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/BurntSushi/toml"
)

var letrasChutadas []string
var erros = 0
var chutes []string

type Config struct {
	Themes map[string][]string `toml:"themes"`
}

func chooseWord() string {
	arquivoTOML := "palavras.toml"

	var config Config

	if _, err := toml.DecodeFile(arquivoTOML, &config); err != nil {
		fmt.Println("Erro ao ler o arquivo TOML:", err)
		return ""
	}

	rand.Seed(time.Now().UnixNano())

	themes := make([]string, 0, len(config.Themes))
	for theme := range config.Themes {
		themes = append(themes, theme)
	}
	temaEscolhido := themes[rand.Intn(len(themes))]

	themeWord := config.Themes[temaEscolhido]
	wordSelected := themeWord[rand.Intn(len(themeWord))]

	return wordSelected
}

func drawForca() {
	chute := playing()

	fmt.Println(letrasChutadas)

	fmt.Printf("  _______       \n")
	fmt.Printf(" |/      |      \n")
	fmt.Printf(" |       %s     \n", func() string {
		if erros > 0 {
			return "(_)"
		} else {
			return "   "
		}
	}())
	fmt.Printf(" |      %s      \n", func() string {
		if erros == 2 {
			return " |  "
		} else if erros == 3 {
			return "/| "
		} else if erros > 3 {
			return "/|\\"
		} else {
			return "   "
		}
	}())
	fmt.Printf(" |       %s     \n", func() string {
		if erros > 1 {
			return "|"
		} else {
			return " "
		}
	}())
	fmt.Printf(" |      %s   \n", func() string {
		if erros == 5 {
			return "/ "
		} else if erros > 5 {
			return "/  \\"
		} else {
			return "   "
		}
	}())
	fmt.Printf(" |              \n")
	fmt.Printf("_|___           \n")
	fmt.Printf("\n\n")

}

func playing() {
	var chute string

	fmt.Println("Digite uma letra: ")
	fmt.Scanf(chute)

	for _, char := range secretWord {
		if char == chute {
			return 
		}
	}
	if chute

	erros++
	chutes.append(chute)
}

func menu(status int) int {
	option := status

	fmt.Println("|====================|")
	switch status {
	default:
		fmt.Println("|-------Forca--------|")
	case 1:
		fmt.Println("|-------Winner-------|")
	case 2:
		fmt.Println("|------GameOver------|")
	}
	fmt.Println("|====================|")

	return option
}

func winner() {
	fmt.Println("\nParabéns, você ganhou!")
	fmt.Println()

	fmt.Println("       ___________      ")
	fmt.Println("      '._==_==_=_.'     ")
	fmt.Println("     .-||:      |-.    ")
	fmt.Println("     | (|:.     |) |    ")
	fmt.Println("      '-|:.     |-'     ")
	fmt.Println("        ||::.   |      ")
	fmt.Println("         '::. .'        ")
	fmt.Println("           ) (          ")
	fmt.Println("         _.' '._        ")
	fmt.Println("        '-------'       ")

	erros = 0
}

func loser() {
	fmt.Println("Lamento, vocÊ foi enforcado!")
	fmt.Println("A palavra era")
	fmt.Println()

	fmt.Println("     _______________       ")
	fmt.Println("    |               |      ")
	fmt.Println("   |                 |     ")
	fmt.Println("  |                   |    ")
	fmt.Println("  |   XXXX     XXXX   |    ")
	fmt.Println("  |   XXXX     XXXX   |    ")
	fmt.Println("  |   XXX       XXX   |    ")
	fmt.Println("  |                   |    ")
	fmt.Println("  |__      XXX      __|    ")
	fmt.Println("    ||     XXX     ||      ")
	fmt.Println("    | |           | |      ")
	fmt.Println("    | I I I I I I I |      ")
	fmt.Println("    |  I I I I I I  |      ")
	fmt.Println("    ||_            _|      ")
	fmt.Println("     ||_         _|        ")
	fmt.Println("       ||_______|          ")

	erros = 0
}
