// Nombre: angel//
// apellido: chacon//
// C.I: 27793649//
// Seccion: 7122//
// Turno: noche//
package main

import (
	"fmt"
)

// Funcion palindrome
func palindrome(palabra string) bool {
	var palabraInversa string
	//recorre la paralabra por letra al reves
	for i := len(palabra); i > 0; i-- {
		palabraInversa = palabraInversa + palabra[i-1:i]
	}
	//valida si la palabra entrante es Palindrome
	if palabra == palabraInversa {
		return true
	} else {
		return false
	}
}

// Funcion para buscar una palabra dentro de una frace
func buscarPalabra(palabra string, frace string) bool {
	var guardarPalabra string
	cont := 0
	//recorre cada caracter para extraer cada palabra de la frace
	for i := 0; i < len(frace); i++ {
		//valida si el caracter que recorre es diferente a un espacio en blanco o si ya culmino el recorrido
		if frace[(i):(i+1)] != " " || i == (len(frace)-1) {
			//extrae una palabra de la frace
			guardarPalabra = guardarPalabra + frace[(i):(i+1)]
			//valida si ya culmino el recorrido para comparar la ultima paralabra de la prace con la palabra que se busca
			if i == (len(frace) - 1) {
				if guardarPalabra == palabra {
					cont++
				}
				guardarPalabra = ""
			}
		} else {
			//valida si la palabra estraida de la frace conincide con la palabra buscada
			if guardarPalabra == palabra {
				cont++
			}
			guardarPalabra = ""
		}
	}
	//Valida si la palabra existe dentro de la frace
	if cont > 0 {
		return true
	} else {
		return false
	}
}

// Funcion de Letras Repetidas Continuas
func parabrasContinuas(arr []string) [][]string {
	var arraytemporal []string
	arrayLetrasContinuas := [][]string{}
	//Recorre el array
	for i := 0; i < len(arr); i++ {
		var nuevoArray []string
		if len(arraytemporal) == 0 {
			arraytemporal = append(arraytemporal, arr[i])
		} else {
			if arraytemporal[(len(arraytemporal)-1)] == arr[i] {
				arraytemporal = append(arraytemporal, arr[i])
				if len(arr)-1 == i {
					for j := 0; j < len(arraytemporal); j++ {
						nuevoArray = append(nuevoArray, arraytemporal[j])
					}
					if len(nuevoArray) > 1 {
						arrayLetrasContinuas = append(arrayLetrasContinuas, nuevoArray)
					}
					arraytemporal = append(arraytemporal[:0], arraytemporal[(len(arraytemporal)):]...)
				}
			} else {
				for k := 0; k < len(arraytemporal); k++ {
					nuevoArray = append(nuevoArray, arraytemporal[k])
				}
				if len(nuevoArray) > 1 {
					arrayLetrasContinuas = append(arrayLetrasContinuas, nuevoArray)
				}
				arraytemporal = append(arraytemporal[:0], arraytemporal[(len(arraytemporal)):]...)
				arraytemporal = append(arraytemporal, arr[i])
			}

		}

	}
	return arrayLetrasContinuas
}

// funcion que genera semilla
func generarSemilla(array []string, peso []string) int {
	var suma int
	multiplicacion := 1
	for i := 0; i < len(array); i++ {
		for j := 0; j < len(peso); j++ {
			if array[i] == peso[j] {
				suma += j
				multiplicacion *= j
			}
		}
	}
	resultado := suma + multiplicacion
	return resultado

}
func main() {

	fmt.Println(palindrome("oso"))
	fmt.Println(buscarPalabra("de", "prueba de frace"))
	arr := []string{"a", "a", "c", "c", "s", "a", "a", "a", "x", "x", "x", "x"}
	fmt.Println("paralabras continuas dentro de un arry con arrays", parabrasContinuas(arr))
	array := []string{"a", "c", "x", "a"}
	peso := []string{"y", "a", "c"}
	fmt.Println("reusltado de la semilla =", generarSemilla(array, peso))

}
