package materi3

import (
	"fmt"
)

// Latihan:
// 1. Buat program yang mencetak angka dari 1 hingga 10, tetapi lewati angka 5 menggunakan `continue`.
// 2. Buat program yang keluar dari loop setelah mencapai angka 7 menggunakan `break`.
// Tugas Mandiri:
// 1. Buat program yang mencetak tabel perkalian 5.
// 2. Buat program yang mencetak bilangan prima dari 1 hingga 50.

func No1() {
	for i := 1; i <= 10; i++ {
		if i == 5 {
			continue
		}
		fmt.Println(i)
	}
}

func No2() {
	for i := 1; i <= 10; i++ {
		if i == 7 {
			break
		}
		fmt.Println(i)
	}
}

func TugasMandiri1() {
	for i := 1; i <= 10; i++ {
		fmt.Println(i * 5)
	}
}

func TugasMandiri2() {
	for i := 1; i <= 50; i++ {
		if i == 1 {
			continue
		}
		if i == 2 {
			fmt.Println(i)
			continue
		}
		isPrime := true
		for j := 2; j < i; j++ {
			if i%j == 0 {
				isPrime = false
				break
			}
		}
		if isPrime {
			fmt.Println(i)
		}
	}
}