package main 

import (
	x "fmt" // -- FMT vem da Standart Lib do Go ( "O que seria o Object do Java" )
	"strconv"
)

var boolean bool // naturalmente false
var mtString string = "Gandalf"

var a, b int // nautralmente starta em 0

// limitar o tamanho de uso caso necessario
var myint8 int8 // armazena 8bits por numero
var myint16 int16 // armazena 16bits por numero
var myint64 int64 // armazena 64bits por numero

var myfloat32 float32
var myfloat64 float64

var myText = "Luke Skywalker"

// Constantes
const STATE = "done"

func main() {   // Função main não retorna nada naturalmente

	myText2 := "Alien" // inicialização rapida ( somente dentro do escopo de uma função )
	
	x.Println("Hello world")
	x.Println(boolean) // output - false
	x.Println(mtString) // output - Gandalf
	x.Println(a, b) // output - 0 0
	x.Println(myText) // output - Luke Skywalker
	x.Println(myText2) // output - Alien
	x.Println(STATE) // output -done

	a = 10 
	a = a + 10
	a += 1

	boolean = a < b
	x.Println(a) // output - 21
	x.Println(boolean) // output - false


	film := "Laranja"
	sum_string := film + " " + "Mecanica"
	x.Println(sum_string) // output - LaranjaMecanica

	x.Printf("%s is a good film", sum_string) // output - Laranja Mecanica is a good film
	x.Println("")

	value := "123"
	s, err := strconv.Atoi(value)
	if err != nil {
		panic("oh no!")
	}

	s += 1
	x.Println(s) // output - 124
	// x.Println(strconv.Atoi(value)) retorna o valor convertido e <nil> ( nao desejado que apareça, tratamento disso acima )


	invertValue := 123
	s2 := strconv.Itoa(invertValue) // de string para inteiro não é necessario tratamento
	s2 += "456"
	x.Println(s2) // output - 123456

	boolValue := "true"
	s3, err := strconv.ParseBool(boolValue)
	if err != nil { // nil -> valor nulo
		panic("oh no!")
	}
	x.Println(s3) // output - true

	// array
	var array [2]string
	array[0] = "Boba Fett"
	array[1] = "Dr Strange"
	array2 := [3]string{"Sparrow", "Darth", "Frodo"}
	x.Println(array)
	x.Println(array2)

	// slices
	slice := []string{"Leia", "R2-D2", "C-3PO"}
	x.Println(slice) // output - [Leia R2-D2 C-3PO]

	slice2 := slice[:2]
	x.Println(slice2) // output - [Leia R2-D2]

	slice = append(slice, "Gollum")
	x.Println(slice) // output - [Leia R2-D2 C-3PO Gollum]


	// Map
	myMap := map[string]int{"Luke": 1, "Leia": 2, "Han": 3}
    x.Println(myMap["Luke"]) // output - 1

    myMap["Leia"] = 4
    x.Println(myMap["Leia"]) // output - 4

    delete(myMap, "Leia")
    x.Println(myMap["Leia"]) // output - 0 ( pois foi deletado)

    // structs
    type Person struct {
        Name string
        Age  int
    }

    var p Person
    p.Name = "Luke"
    p.Age = 22
    x.Println(p) // output - {Luke 22}

    p2 := Person{Name: "Leia", Age: 21}
    x.Println(p2)

	gameState := map[string]int{
		"coins": 22,
		"life": 2,
		"blood": 50,
	}
	x.Println(gameState) // output - map[blood:50 coins:22 life:2]
	delete(gameState, "coins")
	x.Println(gameState) // output - map[blood:50  life:2]
}