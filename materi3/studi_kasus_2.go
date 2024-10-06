package materi3

import (
	"fmt"
)

// Kita akan membuat simulasi sistem manajemen pemesanan hotel yang mencakup beberapa faktor
// seperti jenis kamar, durasi menginap, jumlah tamu, dan status keanggotaan pelanggan. Sistem ini akan
// menentukan harga akhir berdasarkan kombinasi kondisi tersebut.

// Tujuan:
// • Menentukan harga kamar berdasarkan jenisnya (Standard, Deluxe, Suite).
// • Menerapkan diskon khusus jika pelanggan adalah anggota.
// • Menerapkan biaya tambahan jika tamu melebihi kapasitas standar kamar.
// • Menghitung total harga berdasarkan durasi menginap.
// • Menampilkan ringkasan total biaya yang harus dibayar oleh pelanggan.
// Spesifikasi Kamar dan Harga:
// • Standard Room: Rp 500.000 per malam, kapasitas 2 orang.
// • Deluxe Room: Rp 750.000 per malam, kapasitas 3 orang.
// • Suite Room: Rp 1.200.000 per malam, kapasitas 4 orang.
// Biaya tambahan: Jika jumlah tamu melebihi kapasitas kamar, biaya tambahan sebesar Rp 200.000 per
// orang per malam dikenakan.
// Diskon keanggotaan: 10% diskon untuk pelanggan yang terdaftar sebagai anggota.

func StudiKasus2() {
	for {
		// Input data dari pengguna
		var roomType string
		var nights, guests int
		var isMember string
		var isMemberBool bool

		// Mengambil input jenis kamar
		fmt.Print("Pilih jenis kamar (Standard, Deluxe, Suite): ")
		fmt.Scan(&roomType)

		// Mengambil input jumlah malam menginap
		fmt.Print("Masukkan jumlah malam menginap: ")
		fmt.Scan(&nights)

		// Mengambil input jumlah tamu
		fmt.Print("Masukkan jumlah tamu: ")
		fmt.Scan(&guests)

		// Mengambil status keanggotaan
		fmt.Print("Apakah Anda anggota? (yes/no): ")
		fmt.Scan(&isMember)

		// Konversi status keanggotaan ke boolean
		if isMember == "yes" {
			isMemberBool = true
		} else {
			isMemberBool = false
		}

		// Inisialisasi variabel harga dan kapasitas
		var roomPrice, extraCharge, totalPrice int
		var capacity int

		// Struktur kontrol untuk menentukan harga berdasarkan jenis kamar
		if roomType == "Standard" {
			roomPrice = 500000
			capacity = 2
		} else if roomType == "Deluxe" {
			roomPrice = 750000
			capacity = 3
		} else if roomType == "Suite" {
			roomPrice = 1200000
			capacity = 4
		} else {
			fmt.Println("Jenis kamar tidak valid.")
			continue
		}

		// Hitung total harga kamar
		totalPrice = roomPrice * nights

		// Cek apakah ada kelebihan tamu
		if guests > capacity {
			extraGuests := guests - capacity
			extraCharge = extraGuests * 200000 * nights
			totalPrice += extraCharge
		}

		// Jika anggota, berikan diskon 10%
		if isMemberBool {
			discount := totalPrice * 10 / 100
			totalPrice -= discount
		}

		// Cetak ringkasan biaya
		fmt.Println("\n=== Ringkasan Pemesanan ===")
		fmt.Printf("Jenis Kamar: %s\n", roomType)
		fmt.Printf("Jumlah Malam: %d\n", nights)
		fmt.Printf("Jumlah Tamu: %d\n", guests)
		fmt.Printf("Total Harga Kamar: Rp %d\n", roomPrice*nights)
		if extraCharge > 0 {
			fmt.Printf("Biaya Tambahan Tamu: Rp %d\n", extraCharge)
		}
		if isMemberBool {
			fmt.Println("Diskon Anggota: 10%")
		}
		fmt.Printf("Total Harga Akhir: Rp %d\n", totalPrice)

		// Tanya apakah ingin melanjutkan atau keluar
		var continueInput string
		fmt.Print("Apakah Anda ingin memasukkan data lagi? (yes/no): ")
		fmt.Scan(&continueInput)
		if continueInput != "yes" {
			break
		}
	}
}
