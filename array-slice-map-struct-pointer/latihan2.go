package arrayslicemapstructpointer

import (
	"fmt"
)

// Barang struct untuk menyimpan data barang
type Barang struct {
	Nama  string
	Harga float64
	Stok  int
}

// Inventory struct untuk menyimpan data inventaris
type Inventory struct {
	BarangMap map[string]*Barang
}

// NewInventory membuat inventaris baru
func NewInventory() *Inventory {
	return &Inventory{BarangMap: make(map[string]*Barang)}
}

// TambahBarang menambah barang ke inventaris
func (inv *Inventory) TambahBarang(nama string, harga float64, stok int) {
	inv.BarangMap[nama] = &Barang{Nama: nama, Harga: harga, Stok: stok}
}

// TampilkanSemuaBarang menampilkan semua barang di inventaris
func (inv *Inventory) TampilkanSemuaBarang() {
	for _, barang := range inv.BarangMap {
		fmt.Printf("Nama: %s, Harga: %.2f, Stok: %d\n", barang.Nama, barang.Harga, barang.Stok)
	}
}

// UbahStokBarang mengubah jumlah stok barang
func (inv *Inventory) UbahStokBarang(nama string, stokBaru int) {
	if barang, ok := inv.BarangMap[nama]; ok {
		barang.Stok = stokBaru
	} else {
		fmt.Println("Barang tidak ditemukan")
	}
}

// CariBarang mencari barang berdasarkan nama
func (inv *Inventory) CariBarang(nama string) {
	if barang, ok := inv.BarangMap[nama]; ok {
		fmt.Printf("Nama: %s, Harga: %.2f, Stok: %d\n", barang.Nama, barang.Harga, barang.Stok)
	} else {
		fmt.Println("Barang tidak ditemukan")
	}
}

// HapusBarang menghapus barang dari inventaris
func (inv *Inventory) HapusBarang(nama string) {
	delete(inv.BarangMap, nama)
}

func Latihan2() {
	inv := NewInventory()

	// Menambah barang
	inv.TambahBarang("Laptop", 15000000, 10)
	inv.TambahBarang("Mouse", 150000, 50)

	// Menampilkan semua barang
	fmt.Println("Daftar Semua Barang:")
	inv.TampilkanSemuaBarang()

	// Mengubah stok barang
	inv.UbahStokBarang("Laptop", 5)

	// Mencari barang
	fmt.Println("\nMencari Barang 'Laptop':")
	inv.CariBarang("Laptop")

	// Menghapus barang
	inv.HapusBarang("Mouse")

	// Menampilkan semua barang setelah penghapusan
	fmt.Println("\nDaftar Semua Barang Setelah Penghapusan:")
	inv.TampilkanSemuaBarang()
}
