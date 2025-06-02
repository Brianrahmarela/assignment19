package main

import "fmt"

func reverseString(s string) string {
	//di go, string adalah slice of byte, maka hasil Halo => angka o=111, l=108, h=104, a=97
	fmt.Print("Input string terbaca int di go: ", s, "\n")
	// Ubah string menjadi slice of rune(int32) agar jika ada karakter non-ASCII (Unicode) tdk error
	// characters := []rune(s)
	// fmt.Print("Karakter: ", characters, "\n")
	// fmt.Print("len: ", len(characters)-1)
	var reversed string

	// atau lgsg Ambil karakter dari belakang ke depan tanpa conversi ke rune
	for i := len(s) - 1; i >= 0; i-- {
		fmt.Print("i: ", i, " karakter: ", s[i], "\n")
		fmt.Print("i: ", i, " string: ", string(s[i]), "\n")
		// ini convert int ke string
		reversed += string(s[i])
	}

	return reversed
}

func main() {
	input := "Halo"
	fmt.Println("Hasil:", reverseString(input))
}
