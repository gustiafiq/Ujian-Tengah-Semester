package main

import (
	"fmt"
)

type namapengguna struct {
	Username string
	Password string
}

type Buku struct {
	judul  string
	jumlah int
}

var nama = []nama{
	{"Afiq", "2406499065"},
}

var buku = []Buku{
	{"Pemrograman", 10},
	{"Film", 5},
	{"Printing", 20},
}

var historiPeminjaman []string

func main() {
	fmt.Println("
	===========================================
	=Selamat datang di Peminjaman Buku Vokasi!=
	===========================================
	")
	if !login() {
		fmt.Println("
		=================================================
		Username atau Password salah. Program dihentikan.
		=================================================
		")
		return
	}

	for {
		fmt.Println("\nMenu:")
		fmt.Println("=1. Lihat Informasi Pengguna=")
		fmt.Println("==2. Lihat Daftar Buku==")
		fmt.Println("===3. Tambah Daftar Buku===")
		fmt.Println("====4. Tambah Peminjaman Buku====")
		fmt.Println("=====5. Histori Peminjaman Buku=====")
		fmt.Println("======6. Keluar dari Program======")
		fmt.Print("Pilih menu (1-6): ")

		var choice int
		fmt.Scan(&choice)

		switch choice {
		case 1:
			lihatinfopengguna()
		case 2:
			lihatBuku()
		case 3:
			tambahBuku()
		case 4:
			pinjamBuku()
		case 5:
			lihatHistoriPeminjaman()
		case 6:
			fmt.Println("
			==============================
			Terima kasih! Silahkan Keluar.
			==============================
			")
			return
		default:
			fmt.Println("
			=======================================
			Pilihan tidak valid. Silakan coba lagi.
			=======================================
			")
		}
	}
}

func Masuk() bool {
	var nama, password string
	fmt.Print("Masukkan Username: ")
	fmt.Scan(&nama)
	fmt.Print("Masukkan Password: ")
	fmt.Scan(&password)

	for _, user := range users {
		if user.Username == username && user.Password == password {
			return true
		}
	}
	return false
}

func lihatInfoPengguna() {
	fmt.Println("Informasi Pengguna:")
	for _, user := range users {
		fmt.Printf("Username: %s, Password: %s\n", user.Username, user.Password)
	}
}

func lihatBuku() {
	fmt.Println("Daftar Buku:")
	for _, book := range books {
		fmt.Printf("Nama Buku: %s, Jumlah: %d\n", book.Title, book.Amount)
	}
}

func TambahBuku() {
	var title string
	var amount int
	fmt.Print("Masukkan judul buku yang akan ditambahkan: ")
	fmt.Scan(" %[^\n]", &title)
	fmt.Print("Masukkan jumlah buku: ")
	fmt.Scan(&amount)

	if amount <= 0 {
		fmt.Println("Jumlah buku harus lebih dari 0.")
		return
	}

	books = append(books, Book{Title: title, Amount: amount})
	fmt.Println("Buku berhasil ditambahkan.")
}

func pinjamBuku() {
	var title string
	var quantity int
	fmt.Print("Masukkan judul buku yang ingin dipinjam: ")
	fmt.Scan(" %[^\n]", &title)
	fmt.Print("Masukkan jumlah buku yang ingin dipinjam: ")
	fmt.Scan(&quantity)

	if quantity <= 0 {
		fmt.Println("Jumlah peminjaman tidak valid (harus lebih dari 0).")
		return
	}

	for i, book := range books {
		if book.Title == title {
			if book.Amount >= quantity {
				books[i].Amount -= quantity
				borrowingHistory = append(borrowingHistory, fmt.Sprintf("Dipinjam: %s, Jumlah: %d", title, quantity))
				fmt.Println("Buku berhasil dipinjam.")
				return
			}
			fmt.Println("Jumlah buku yang dipinjam melebihi jumlah yang tersedia.")
			return
		}
	}
	fmt.Println("Buku tidak ditemukan.")
}

func lihatHistoriPeminjaman() {
	fmt.Println("Histori Peminjaman Buku:")
	if len(borrowingHistory) == 0 {
		fmt.Println("Tidak ada histori peminjaman.")
	} else {
		for _, record := range borrowingHistory {
			fmt.Println(record)
		}
	}
}
