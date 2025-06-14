package main

import (
	"fmt"
	"strings"
)

type Content struct {
	ID      int
	judul   string
	jenis   string
	media   string
	status  string
	tanggal string
}

var contents []Content
var idCounter = 1

func tambah() {
	var c Content
	var s string
	c.ID = idCounter
	idCounter++

	fmt.Print("\nJudul: ")
	fmt.Scanln(&c.judul)
	fmt.Print("Jenis Konten (Travel/Review/Tutorial/Game/Musik): ")
	fmt.Scanln(&c.jenis)
	fmt.Print("Media Konten (Blog/Video/Audio/Sosial Media): ")
	fmt.Scanln(&c.media)
	fmt.Print("Status (1.Draft/2.Scheduled/3.Published): ")
	fmt.Scanln(&s)
	switch strings.ToLower(s) {
	case "1", "draft", "1.draft":
		c.status = "Draft"
	case "2", "scheduled", "2.scheduled":
		c.status = "Scheduled"
	case "3", "published", "3.published":
		c.status = "Published"
	}
	fmt.Print("Tanggal dipublikasikan (Hari-Bulan-Tahun): ")
	fmt.Scanln(&c.tanggal)

	contents = append(contents, c)
	fmt.Println("\nKonten berhasil dibuat.")
}

func daftar() {
	fmt.Println("\nDaftar Konten:")
	for _, c := range contents {
		fmt.Printf("[%d] %s (%s)\n    Jenis: %s\n    %s: %s\n", c.ID, c.judul, c.media, c.jenis, c.status, c.tanggal)
	}
}

func hapus() {
	var id int
	fmt.Print("\nMasukkan ID konten yang akan dihapus: ")
	fmt.Scanln(&id)

	for i, c := range contents {
		if c.ID == id {
			contents = append(contents[:i], contents[i+1:]...)
			fmt.Println("\nKonten berhasil dihapus.")
			return
		}
	}
	fmt.Println("\nID tidak ditemukan.")
}

func pencarian() {
	var keyword string
	fmt.Print("\nMasukkan kata kunci pencarian: ")
	fmt.Scanln(&keyword)

	for _, c := range contents {
		if strings.Contains(strings.ToLower(c.judul), strings.ToLower(keyword)) || strings.Contains(strings.ToLower(c.jenis), strings.ToLower(keyword)) || strings.Contains(strings.ToLower(c.media), strings.ToLower(keyword)) || strings.Contains(strings.ToLower(c.status), strings.ToLower(keyword)) {
			fmt.Printf("[%d] %s (%s)\n    Jenis: %s\n    %s: %s\n", c.ID, c.judul, c.media, c.jenis, c.status, c.tanggal)
		}
	}
}

func main() {
	var pilih string

	for {
		fmt.Println("\nAplikasi Manajemen Konten")
		fmt.Println("0.Keluar\n1.Tambah-Konten\n2.Daftar-Konten\n3.Hapus-Konten\n4.Cari-Konten")
		fmt.Print("Pilihan: ")
		fmt.Scanln(&pilih)

		switch strings.ToLower(pilih) {
		case "1", "tambah-konten", "1.tambah-konten":
			tambah()
		case "2", "daftar-konten", "2.daftar-konten":
			daftar()
		case "3", "hapus-konten", "3.hapus-konten":
			hapus()
		case "4", "cari-konten", "4.cari-konten":
			pencarian()
		case "0", "keluar", "0.keluar":
			fmt.Println("\nKeluar dari aplikasi.")
			return
		default:
			fmt.Println("\nPilihan tidak valid. Ulangi!")
		}
	}
}
