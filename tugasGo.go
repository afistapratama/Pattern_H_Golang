package main

import (
	"fmt"
)


func cetak_gambar(number int) {
	if number%2 != 1 {
		fmt.Println("angka yang dimasukkan harus bilangan ganjil")
	} else {
		for i := 0; i < number; i++ {
			for j := 0; j < numb; j++ {
				if j == 0 || i == number/2 || j == number-1 {
					fmt.Print("* ")
				} else {
					fmt.Print("= ")
				}
			}
			fmt.Println("\n")
		}
	}

}

func main() {
	cetak_gambar(5)
}
