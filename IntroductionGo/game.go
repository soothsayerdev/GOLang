package main

import (
	ChatGame "fmt" // -- FMT vem da Standart Lib do Go ( "O que seria o Object do Java" )
)

// Structs
type Character struct {
	Name string
	Color string
}


type GameState struct {
	Coins int
	Life int
	Blood int
	Char Character
}

func game() {   // Função main não retorna nada naturalmente
	// game

	players := map[string]GameState{
		"player1": GameState{Coins: 100, Life: 100, Blood: 100, Char: Character{Name: "Luke", Color: "Blue"}},
        "player2": GameState{Coins: 200, Life: 80, Blood: 80, Char: Character{Name: "Leia", Color: "Red"}},
	}

	if players["player1"].Char.Name == "Luke" {
		ChatGame.Println("Que a força esteja com voce")
	}else if players["player2"].Char.Name == "Leia" {
		ChatGame.Println("Me ajude Obi-Wan, você é minha unica esperança")
	}else {
		ChatGame.Print("Estou na saga errada.")
	}

	character := players["player2"].Char.Name
	switch character {
	case "Luke":
		ChatGame.Println("Que a força esteja com voce, Leia")
	case "Leia":
		ChatGame.Println("Me ajude Obi-Wan, você é minha unica esperança")
	default:
		ChatGame.Print("Estou na saga errada.")
	}
}