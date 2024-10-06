package materi1

import (
	"fmt"
)

func HelloWorld() {
	fmt.Println("Hello, World!")
}

func DataType() {
	fmt.Println("Satu = ", 1)
	fmt.Println("Dua = ", 2)
	fmt.Println("Tiga Koma Lima = ", 3.5)

	// Bilangan Integer
	var positiveNumber uint8 = 89
	var negativeNumber = -1243423644

	fmt.Printf("Bilangan positif = %d\n ", positiveNumber)
	fmt.Printf("Bilangan negatif = %d\n ", negativeNumber)

	// Bilangan Floating Point
	var decimalNumber = 2.62

	fmt.Printf("Bilangan desimal: %f\n", decimalNumber)
	fmt.Printf("Bilangan desimal: %.3f\n", decimalNumber)

	// Boolean
	fmt.Println("Benar = ", true)
	fmt.Println("Salah = ", false)

	// String
	fmt.Println("Pemerograman")
	fmt.Println("Pemerograman Dasar 2")

	// Fungsi String
	fmt.Println(len("Pemerograman"))
	fmt.Println("Pemerograman Dasar 2"[0])
	fmt.Println("Pemerograman Dasar 2"[4])
}

func Variable() {
	var nama string

	nama = "Pemerograman Dasar 2"
	fmt.Println(nama)

	nama = "Bukan Pemrograman Dasar 2"
	fmt.Println(nama)

	// Tanpa deklarasi tipe data
	var teks = "Ini contoh string"
	fmt.Println(teks)

	// Tanpa kata kunci var, untuk deklarasi variabel
	teks1 := "Tanpa kata kunci var"
	fmt.Println(teks1)

	// Deklarasi variabel lebih dari satu
	var namaDepan, namaBelakang string
	namaDepan = "Muhammad"
	namaBelakang = "Fauzan"
	fmt.Println(namaDepan, namaBelakang)

	var (
		namaDepan1    = "Muhammad"
		namaBelakang1 = "Fauzan"
	)
	fmt.Println(namaDepan1, namaBelakang1)

	// Konstanta
	const (
		phi = 3.14
	)
	fmt.Println(phi)

}

func InputTerminal() {
	fmt.Print("Masukkan Angka: ")
	var input float64
	fmt.Scanf("%f", &input)

	output := input * 2

	fmt.Println(output)
}

func ArithmeticOperator() {
	var a = 10
	var b = 5

	var hasil = a + b
	fmt.Println("Hasil dari a + b = ", hasil)

	hasil = a - b
	fmt.Println("Hasil dari a - b = ", hasil)

	hasil = a * b
	fmt.Println("Hasil dari a * b = ", hasil)

	hasil = a / b
	fmt.Println("Hasil dari a / b = ", hasil)

	hasil = a % b
	fmt.Println("Hasil dari a % b = ", hasil)

	a++
	fmt.Println("Nilai dari a++ = ", a)

	b--
	fmt.Println("Nilai dari b-- = ", b)
}

func ComparisonOperator() {
	var a = 10
	var b = 5

	var hasil = a == b
	fmt.Println("Apakah a == b: ", hasil)

	hasil = a != b
	fmt.Println("Apakah a != b: ", hasil)

	hasil = a > b
	fmt.Println("Apakah a > b: ", hasil)

	hasil = a < b
	fmt.Println("Apakah a < b: ", hasil)

	hasil = a >= b
	fmt.Println("Apakah a >= b: ", hasil)

	hasil = a <= b
	fmt.Println("Apakah a <= b: ", hasil)
}

func LogicalOperator() {
	var a = true
	var b = false

	var hasil = a && b
	fmt.Println("Apakah a && b: ", hasil)

	hasil = a || b
	fmt.Println("Apakah a || b: ", hasil)

	hasil = !a
	fmt.Println("Apakah !a: ", hasil)
}

func DataTypeArray() {
	var names [4]string

	names[0] = "Muhammad"
	names[1] = "Fauzan"
	names[2] = "Rizki"
	names[3] = "Nur"

	fmt.Println(names[0], names[1], names[2], names[3])

	var numbers = [4]int{
		90,
		95,
		80,
		85,
	}

	fmt.Println(numbers[0], numbers[1], numbers[2], numbers[3])
}
