package materi3

import (
	"fmt"
)

// Deskripsi Kasus: Buat program untuk menghitung nilai rata-rata siswa berdasarkan nilai ujian dan
// menentukan lulus atau tidak. Program harus menggunakan loop untuk menerima nilai dari banyak
// siswa dan memutuskan kelulusan berdasarkan kondisi if-elseif.
// Spesifikasi Penilaian:
// • Jika rata-rata nilai >= 80: Lulus dengan predikat "Sangat Memuaskan".
// • Jika rata-rata nilai >= 60 dan < 80: Lulus dengan predikat "Memuaskan".
// • Jika rata-rata nilai >= 50 dan < 60: Lulus dengan predikat "Cukup".
// • Jika rata-rata nilai < 50: Tidak Lulus.
// Langkah-langkah Implementasi:
// • Gunakan looping untuk menerima input nilai dari beberapa siswa.
// • Hitung rata-rata nilai.
// • Gunakan if-elseif untuk menentukan predikat kelulusan siswa.

func StudiKasus1() {
	var banyakSiswa int
	var nilai, totalNilai float64
	var rataRata float64

	fmt.Print("Masukkan banyak siswa: ")
	fmt.Scanln(&banyakSiswa)

	for i := 1; i <= banyakSiswa; i++ {
		fmt.Printf("Masukkan nilai siswa %d: ", i)
		fmt.Scanln(&nilai)
		totalNilai += nilai
	}

	if banyakSiswa > 0 {
		rataRata = totalNilai / float64(banyakSiswa)

		if rataRata >= 80 {
			fmt.Println("Sangat Memuaskan")
		} else if rataRata >= 60 {
			fmt.Println("Memuaskan")
		} else if rataRata >= 50 {
			fmt.Println("Cukup")
		} else {
			fmt.Println("Tidak Lulus")
		}
	} else {
		fmt.Println("Tidak ada siswa")
	}
}
