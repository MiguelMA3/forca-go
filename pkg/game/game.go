package game

import (
	"fmt"
	//"strings"

	"github.com/MiguelMA3/forca-go/pkg/config"
	//github.com/MiguelMA3/forca-go/pkg/utils
)

var erros = 0
var chutes = []string{}

func Play() {
	secretWord, themeSelected := config.ChooseWord()
	chute, erros := Playing(secretWord)
	chutes = append(chutes, chute)

	DrawForca(chutes, themeSelected, erros)
}

func DrawForca(letrasChutadas []string, themeSelected string, erros int) {
	//chute := playing()

	fmt.Println(themeSelected)
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

func Playing(secretWord string) (string, int) {
	var chute string

	fmt.Println("Digite uma letra: ")
	fmt.Scanf(chute)

	erro := 1
	return chute, erro
}
