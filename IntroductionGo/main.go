package main

import (
	x "fmt" // -- FMT vem da Standart Lib do Go ( "O que seria o Object do Java" )

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
	x.Println(myText2) // Alien
	x.Println(STATE) // done

	

}