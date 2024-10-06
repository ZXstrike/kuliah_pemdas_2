package materi1

import (
	"fmt"
)

// Latihan:
// 1. Buat program yang mencetak angka dari 1 hingga 10, tetapi lewati angka 5 menggunakan `continue`.
// 2. Buat program yang keluar dari loop setelah mencapai angka 7 menggunakan `break`.
// Tugas Mandiri:
// 1. Buat program yang mencetak tabel perkalian 5.
// 2. Buat program yang mencetak bilangan prima dari 1 hingga 50.

// Dua cara untuk membuat variabel baru
// 1. Menggunakan deklarasi var
//var a int = 10  note: dikomen karenan pesan error nya ganggu

// 2. Menggunakan deklarasi singkat
// b := 20 note: hanya bisa digunakan di dalam fungsi

func Main() {
	// Nilai x setelah menjalankan x := 5; x += 1 adalah 6
	x := 5
	x += 1
	fmt.Println("Nilai x:", x)

	// Program konversi Fahrenheit ke Celsius
	var fahrenheit float64
	fmt.Print("Masukkan suhu dalam Fahrenheit: ")
	fmt.Scanf("%f", &fahrenheit)
	celsius := (fahrenheit - 32) * 5 / 9
	fmt.Printf("Suhu dalam Celsius: %.2f\n", celsius)

	// Program konversi kaki ke meter
	var feet float64
	fmt.Print("Masukkan panjang dalam kaki: ")
	fmt.Scanf("%f", &feet)
	meters := feet * 0.3048
	fmt.Printf("Panjang dalam meter: %.2f\n", meters)
}

// Scope adalah cakupan di mana variabel dapat diakses. Scope variabel di Go ditentukan oleh tempat deklarasinya. Variabel yang dideklarasikan di dalam fungsi hanya dapat diakses di dalam fungsi tersebut (local scope), sedangkan variabel yang dideklarasikan di luar fungsi dapat diakses oleh semua fungsi dalam paket yang sama (package scope).

// Perbedaan antara var dan const adalah:
// - var digunakan untuk mendeklarasikan variabel yang nilainya dapat berubah.
// - const digunakan untuk mendeklarasikan konstanta yang nilainya tidak dapat berubah setelah didefinisikan.
