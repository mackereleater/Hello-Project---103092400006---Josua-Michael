// KOMEN INI DIBUAT UNTUK TUGAS WAWASAN GLOBAL TIK
package main

import (
	"fmt"
)

//  Fitur wajib:
//- CRUD barang yang dijual
//- Mencatat tiap transaksi
//- Tampilan daftar transaksi dan omzet penjualan

// Kamus Global
const (
	NMAX  int = 100
	TrMAX int = 20
)

type Barang struct {
	id       string
	nama     string
	kategori string
	harga    float64
	stok     int
}
type Transaksi struct {
	idTransaksi  string
	items        [TrMAX]Barang
	jumlahBarang [TrMAX]int
	total        float64
	pembayaran   float64
	kembalian    float64
	totalItems   int
}
type (
	tabBarang    [NMAX]Barang
	tabTransaksi [NMAX]Transaksi
)

// Variabel Global
var (
	barang         tabBarang
	transaksi      tabTransaksi
	totalBarang    int
	totalTransaksi int
)

func main() {

	barang[0] = Barang{"B001", "Aqua", "Minuman", 2500.0, 50}
	barang[1] = Barang{"B002", "Indomie", "Makanan", 3000.0, 100}
	barang[2] = Barang{"B003", "Beras_5kg", "Sembako", 78500.0, 15}
	barang[3] = Barang{"B006", "KapalApi", "Minuman", 1500.0, 200}
	barang[4] = Barang{"B007", "SoKlin", "Kebersihan", 2500.0, 100}
	barang[5] = Barang{"B005", "Bimoli", "Sembako", 21500.0, 50}
	barang[6] = Barang{"B004", "Sunco2L", "Sembako", 42700.0, 40}
	totalBarang = 7
	menuUtama()
}

// MENU UTAMA KASIR serta FUNGSI & PROSEDURNYA
func menuUtama() {
	var pilih int
	for pilih != 5 {
		fmt.Println("\n=== APLIKASI KASIR MINIMART ===")
		fmt.Println("1. Kelola Data Barang")
		fmt.Println("2. Transaksi Baru")
		fmt.Println("3. Cetak Transaksi")
		fmt.Println("4. Laporan Penjualan")
		fmt.Println("5. Keluar")
		fmt.Println("6. Sort Transaksi")
		fmt.Print("Pilih menu: ")

		fmt.Scan(&pilih)

		if pilih == 1 {
			menuBarang()
		} else if pilih == 2 {
			transaksiBaru()
		} else if pilih == 3 {
			cetakTransaksi()
		} else if pilih == 4 {
			laporanPenjualan()
		} else if pilih == 6 {
			sortTransaksi()
		} else if pilih != 5 {
			fmt.Println("Pilihan tidak valid")
		}
	}
}

// Sort Berdasarkan Transaksi
func sortTransaksi() {
	var pass, j, idxMin int
	for pass = 0; pass < totalTransaksi-1; pass++ {
		idxMin = pass
		for j = pass + 1; j < totalTransaksi; j++ {

			if transaksi[j].total > transaksi[idxMin].total {
				idxMin = j
			}
			temp := transaksi[pass]
			transaksi[pass] = transaksi[idxMin]
			transaksi[idxMin] = temp
		}
	}
}
func transaksiBaru() {
	var id string
	var t Transaksi
	var jB, idx int
	var stokAwal, stokDikeranjang, stokTersediaSaatIni int
	var lanjut bool = true

	if totalTransaksi >= NMAX {
		fmt.Println("Kapasitas transaksi penuh!")
		return
	}

	t.idTransaksi = fmt.Sprintf("TRX_%d", totalTransaksi+101)
	t.totalItems = 0
	cetakBarang()

	for t.totalItems < TrMAX && lanjut {
		fmt.Print("\nMasukkan ID barang (ketik 'Y' untuk lanjut ke pembayaran, 'N' untuk batal): ")
		fmt.Scan(&id)

		if id == "Y" || id == "y" {
			lanjut = false
		} else if id == "N" || id == "n" {
			fmt.Println("Transaksi dibatalkan.")
			return
		} else {
			idx = cariBarang(id)
			if idx == -1 {
				fmt.Println("\nBarang tidak ditemukan!")
			} else {
				//Perhitungan Stok yang akan ditampilkan
				stokAwal = barang[idx].stok
				stokDikeranjang = 0

				for i := 0; i < t.totalItems; i++ {
					if t.items[i].id == id {
						stokDikeranjang += t.jumlahBarang[i]
					}
				}
				stokTersediaSaatIni = stokAwal - stokDikeranjang

				// Tampilan Stok Sementara
				fmt.Printf("Ditemukan: %s - %s (Stok Awal: %d, Tersedia untuk dibeli: %d)\n",
					barang[idx].nama,
					barang[idx].kategori,
					stokAwal,
					stokTersediaSaatIni)

				//Input JumlahBeli
				fmt.Print("Jumlah: ")
				fmt.Scan(&jB)

				//Filter Jika Stok kurang dari yang ingin dibeli
				if jB <= 0 {
					fmt.Println("\nJumlah harus lebih dari 0!")
				} else if stokTersediaSaatIni < jB {
					fmt.Printf("\nStok tidak cukup! Sisa stok yang bisa dibeli untuk item ini adalah %d.\n", stokTersediaSaatIni)
				} else {
					t.items[t.totalItems] = barang[idx]
					t.jumlahBarang[t.totalItems] = jB
					t.total += barang[idx].harga * float64(jB)
					t.totalItems++
					fmt.Println("Barang berhasil ditambahkan ke keranjang.")
				}
			}
		}
	}

	if t.totalItems >= TrMAX {
		fmt.Println("\nBatas maksimum transaksi tercapai, lanjut ke pembayaran.")
	}

	//Proses Bayar
	if t.totalItems > 0 {
		var konfirmasi string
		fmt.Printf("\nTotal Belanja Anda: Rp%.2f\n", t.total)
		fmt.Print("Lanjutkan ke pembayaran? (Y/N): ")
		fmt.Scan(&konfirmasi)

		if konfirmasi == "Y" || konfirmasi == "y" {
			fmt.Print("Jumlah Pembayaran: ")
			fmt.Scan(&t.pembayaran)

			for t.pembayaran < t.total {
				fmt.Println("Pembayaran tidak mencukupi!")
				fmt.Print("Jumlah Pembayaran: ")
				fmt.Scan(&t.pembayaran)
				if t.pembayaran < 0 {
					return
				}
			}

			// Update Stok asli setelah berhasil dibayar
			for i := 0; i < t.totalItems; i++ {
				idxBarangAsli := cariBarang(t.items[i].id)
				if idxBarangAsli != -1 {
					barang[idxBarangAsli].stok -= t.jumlahBarang[i]
				}
			}

			t.kembalian = t.pembayaran - t.total
			fmt.Printf("\nKembalian : Rp%.2f", t.kembalian)

			transaksi[totalTransaksi] = t
			totalTransaksi++
			fmt.Println("\nTransaksi berhasil dicatat!")

		} else {
			fmt.Println("Pembayaran dibatalkan. Stok tidak diubah.")
		}
	} else {
		fmt.Println("Tidak ada barang dalam transaksi.")
	}
}

func cetakTransaksi() {
	fmt.Println("\nDAFTAR TRANSAKSI")
	fmt.Println("===================================================================")

	for i := 0; i < totalTransaksi; i++ {
		// Header transaksi
		fmt.Printf("\n%-15s : %s\n", "ID Transaksi", transaksi[i].idTransaksi)
		fmt.Printf("%-15s : Rp%10.2f\n", "Total Belanja", transaksi[i].total)
		fmt.Printf("%-15s : Rp%10.2f\n", "Pembayaran", transaksi[i].pembayaran)
		fmt.Printf("%-15s : Rp%10.2f\n", "Kembalian", transaksi[i].kembalian)

		// Detail barang
		fmt.Println("\nDETAIL BARANG:")
		fmt.Printf("%-8s | %-20s | %-8s | %-12s\n", "ID", "NAMA BARANG", "JUMLAH", "SUB TOTAL")
		for j := 0; j < transaksi[i].totalItems; j++ {
			subTotal := transaksi[i].items[j].harga * float64(transaksi[i].jumlahBarang[j])
			fmt.Printf("%-8s | %-20s | %-8d | Rp%-12.2f\n",
				transaksi[i].items[j].id,
				transaksi[i].items[j].nama,
				transaksi[i].jumlahBarang[j],
				subTotal)
		}
		fmt.Println("-------------------------------------------------------------------")
	}
	fmt.Println("===================================================================")
}

func laporanPenjualan() {
	var total float64
	for i := 0; i < totalTransaksi; i++ {
		total += transaksi[i].total
	}
	fmt.Printf("\nTotal Omzet Penjualan: Rp%.2f\n", total)
}

// MENU CRUD BARANG serta FUNGSI & PROSEDURNYA
func menuBarang() {
	var pilih int
	for pilih != 8 {
		fmt.Println("\n=== KELOLA BARANG ===")
		fmt.Println("1. Tambah Barang")
		fmt.Println("2. Ubah Barang")
		fmt.Println("3. Hapus Barang")
		fmt.Println("4. Ubah Jumlah Stok Barang")
		fmt.Println("5. Cari Barang")
		fmt.Println("6. Sortir Barang")
		fmt.Println("7. Cetak List Barang")
		fmt.Println("8. Kembali ke Menu Utama")
		fmt.Print("Pilih menu: ")

		fmt.Scan(&pilih)

		if pilih == 1 {
			tambahBarang()
		} else if pilih == 2 {
			ubahBarang()
		} else if pilih == 3 {
			hapusBarang()
		} else if pilih == 4 {
			tambahStok()
		} else if pilih == 5 {
			menuCariBarang()
		} else if pilih == 6 {
			menuSortirBarang()
		} else if pilih == 7 {
			cetakBarang()
		} else if pilih != 8 {
			fmt.Println("Pilihan tidak valid")
		}
	}
}

// SequentialSearch
func cariBarang(id string) int {
	for i := 0; i < totalBarang; i++ {
		if barang[i].id == id {
			return i
		}
	}
	return -1
}

// BinarySearch (Panggil Fungsi Sort Sebelum Search)
func cariBarang2(id string) int {
	var kiri, kanan, tengah int
	selectionSortByID()

	kiri = 0
	kanan = totalBarang - 1

	for kiri <= kanan {
		tengah = (kiri + kanan) / 2

		if barang[tengah].id == id {
			return tengah
		} else if id < barang[tengah].id {
			kanan = tengah - 1
		} else {
			kiri = tengah + 1
		}
	}
	return -1
}

// Pakai SequentialSearch karena menampilkan list barang yang memenuhi syarat
func menuCariBarang() {
	var pilih int
	var kataKunci string
	var ditemukan bool = false
	fmt.Println("\n--- Cari Barang Berdasarkan ---")
	fmt.Println("1. ID Barang")
	fmt.Println("2. Nama Barang")
	fmt.Println("3. Kategori Barang")
	fmt.Print("Pilihan Anda: ")
	fmt.Scan(&pilih)

	fmt.Print("Masukkan kata kunci pencarian: ")
	fmt.Scan(&kataKunci)

	fmt.Println("\n--- Hasil Pencarian ---")
	fmt.Printf("%-8s | %-15s | %-15s | %-12s | %-s\n", "ID", "Nama", "Kategori", "Harga", "Stok")
	fmt.Println("----------------------------------------------------------------------")

	//Loop untuk cetak list barang yang memenuhi kata kunci
	for i := 0; i < totalBarang; i++ {
		if (pilih == 1 && barang[i].id == kataKunci) ||
			(pilih == 2 && barang[i].nama == kataKunci) ||
			(pilih == 3 && barang[i].kategori == kataKunci) {
			fmt.Printf("%-8s | %-15s | %-15s | Rp.%-10.2f | %-d\n",
				barang[i].id, barang[i].nama, barang[i].kategori, barang[i].harga, barang[i].stok)
			ditemukan = true
		}
	}
	if !ditemukan {
		fmt.Println("Barang tidak ditemukan.")
	}
	fmt.Println("----------------------------------------------------------------------")
}

func menuSortirBarang() {
	var pilihAlgo int

	if totalBarang < 2 {
		fmt.Println("\nTidak cukup barang untuk disortir (minimal 2).")
		return
	}

	fmt.Println("\n=== Menu Sortir Berdasarkan ID & Nama ===")
	fmt.Println("Pilih Algoritma Sortir:")
	fmt.Println("1. Selection Sort (Ascending)")
	fmt.Println("2. Insertion Sort (Descending)")
	fmt.Print("Pilih: ")
	fmt.Scan(&pilihAlgo)

	if pilihAlgo == 1 {
		selectionSortByID()
		fmt.Println("\nBarang berhasil disortir dengan Selection Sort.")
	} else if pilihAlgo == 2 {
		insertionSortByID()
		fmt.Println("\nBarang berhasil disortir dengan Insertion Sort.")
	} else {
		fmt.Println("Pilihan tidak valid.")
		return
	}
	cetakBarang()
}

func selectionSortByID() {
	var pass, j, idxMin int
	for pass = 0; pass < totalBarang-1; pass++ {
		idxMin = pass
		for j = pass + 1; j < totalBarang; j++ {

			if barang[j].id < barang[idxMin].id {
				idxMin = j
			} else if barang[j].id == barang[idxMin].id && barang[j].nama < barang[idxMin].nama {
				idxMin = j
			}

			temp := barang[pass]
			barang[pass] = barang[idxMin]
			barang[idxMin] = temp
		}
	}
}
func insertionSortByID() {
	var pass, i int
	for pass = 1; pass < totalBarang; pass++ {
		temp := barang[pass]
		i = pass - 1
		for i >= 0 && (barang[i].id < temp.id || (barang[i].id == temp.id && barang[i].nama < temp.nama)) {
			barang[i+1] = barang[i]
			i = i - 1
		}
		barang[i+1] = temp
	}
}

func tambahBarang() {
	var A Barang
	var n int
	var idValid bool

	fmt.Print("Berapa jenis barang yang ingin ditambahkan: ")
	fmt.Scan(&n)

	for i := 0; i < n; i++ {
		if totalBarang >= NMAX {
			fmt.Println("Kapasitas barang penuh!")
			return
		}

		fmt.Printf("\n--- Menambahkan Barang ke-%d dari %d ---\n", i+1, n)

		// Cek Kesamaan ID
		idValid = false
		for !idValid {
			fmt.Print("ID Barang (ketik 'N' atau 'n' untuk keluar): ")
			fmt.Scan(&A.id)

			if A.id == "N" || A.id == "n" {
				fmt.Println("Keluar dari proses input..")
				return
			}
			if cariBarang(A.id) != -1 {
				fmt.Println("ID barang sudah ada, gunakan ID yang lain.")
			} else {
				idValid = true
			}
		}

		// Lanjutkan input
		fmt.Print("Nama Barang: ")
		fmt.Scan(&A.nama)
		fmt.Print("Kategori: ")
		fmt.Scan(&A.kategori)
		fmt.Print("Harga: ")
		fmt.Scan(&A.harga)
		fmt.Print("Stok: ")
		fmt.Scan(&A.stok)

		// Simpan data ke array
		barang[totalBarang] = A
		totalBarang++
		fmt.Printf("Barang '%s' (%s) berhasil ditambahkan.\n", A.nama, A.id)
	}
}

func ubahBarang() {
	var A Barang
	var id string
	var idValid bool
	var yakin string

	fmt.Print("Input id yang mau diubah: ")
	fmt.Scan(&id)

	idx := cariBarang(id)
	if idx == -1 {
		fmt.Println("Barang tidak ditemukan!")
		return
	}

	fmt.Print("Data saat ini: \n")
	fmt.Printf("%-8s | %-15s | %-18s | %-12s   | %-8s\n",
		"ID", "Nama", "Kategori", "Harga", "Stok")
	fmt.Printf("%-8s | %-15s | %-18s | Rp%-12.2f | %-8d\n",
		barang[idx].id, barang[idx].nama, barang[idx].kategori, barang[idx].harga, barang[idx].stok)

	A = barang[idx]

	//Filter jika ID sama dengan ID lain
	idValid = false
	for !idValid {
		fmt.Print("ID barang: ")
		fmt.Scan(&id)

		if cariBarang(id) != -1 && A.id != id {
			fmt.Println("ID barang sudah ada, gunakan ID yang lain.")
		} else {
			idValid = true
		}
	}

	A.id = id
	fmt.Print("Nama Barang: ")
	fmt.Scan(&A.nama)
	fmt.Print("Kategori: ")
	fmt.Scan(&A.kategori)
	fmt.Print("Harga: ")
	fmt.Scan(&A.harga)
	fmt.Print("Stok: ")
	fmt.Scan(&A.stok)

	fmt.Printf("Apakah anda ingin mengubah barang %s %s\n", barang[idx].id, barang[idx].nama)
	fmt.Print("Input 'Y' atau 'y' untuk konfirmasi: ")
	fmt.Scan(&yakin)

	if yakin == "Y" || yakin == "y" {
		barang[idx] = A
		fmt.Println("Data barang berhasil diubah!")
	}
}

func hapusBarang() {
	var id string
	var yakin string

	fmt.Print("Input id yang mau dihapus: ")
	fmt.Scan(&id)

	idx := cariBarang(id)
	if idx == -1 {
		fmt.Println("Barang tidak ditemukan!")
		return
	}
	fmt.Printf("Apakah anda ingin menghapus barang %s %s\n", barang[idx].id, barang[idx].nama)
	fmt.Print("Input 'Y' atau 'y' untuk konfirmasi: ")
	fmt.Scan(&yakin)
	if yakin == "Y" || yakin == "y" {
		items := barang[idx].nama
		for i := idx; i < totalBarang-1; i++ {
			barang[i] = barang[i+1]
		}
		totalBarang--
		fmt.Printf("Barang %s %s berhasil dihapus!\n", id, items)
	}
}

func tambahStok() {
	var stokTambah int
	var id string
	fmt.Print("Input id barang yang mau ditambah atau dikurangi stok: ")
	fmt.Scan(&id)

	idx := cariBarang(id)
	if idx == -1 {
		fmt.Println("Barang tidak ditemukan!")
		return
	}

	fmt.Print("Data saat ini: \n")
	fmt.Printf("%-8s | %-15s | %-18s | %-12s   | %-8s\n",
		"ID", "Nama", "Kategori", "Harga", "Stok")
	fmt.Printf("%-8s | %-15s | %-18s | Rp%-12.2f | %-8d\n",
		barang[idx].id, barang[idx].nama, barang[idx].kategori, barang[idx].harga, barang[idx].stok)

	fmt.Print("Input jumlah barang: ")
	fmt.Scan(&stokTambah)
	barang[idx].stok += stokTambah

	fmt.Printf("Barang %s %s berhasil ditambahkan\n", barang[idx].id, barang[idx].nama)

}

func cetakBarang() {
	fmt.Println("\nDaftar Barang:")
	fmt.Printf("%-8s | %-15s | %-10s | %-10s   | %-s\n", "ID", "Nama", "Kategori", "Harga", "Stok")
	for i := 0; i < totalBarang; i++ {
		fmt.Printf("%-8s | %-15s | %-10s | Rp%-10.2f | %-d\n",
			barang[i].id,
			barang[i].nama,
			barang[i].kategori,
			barang[i].harga,
			barang[i].stok)
	}
}

