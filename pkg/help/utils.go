package utils

import (
	"fmt"
)

func Menu(status int) uint {
	var option uint

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

	fmt.Println("Selecione uma Opção:")
	fmt.Println("1-Jogar       2-Sair")
	fmt.Scan(option)

	return option
}

func Winner(erros int) int {
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
	return erros
}

func Loser(secretWord string) {
	fmt.Println("Lamento, você foi enforcado!")
	fmt.Println("A palavra era", secretWord)
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
}
