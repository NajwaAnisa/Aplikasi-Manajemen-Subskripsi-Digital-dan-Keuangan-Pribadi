package main

import "fmt"

// Deklarasi konstanta ukuran maksimum array pelanggan
const NMAX int = 1000

// Tipe data array untuk menyimpan data langganan pelanggan
type tabPelanggan [NMAX]langganan

// Tipe data bentukan untuk menyimpan informasi satu langganan pelanggan
type langganan struct {
	nama            string  // Nama pelanggan atau langganan
	kategori        string  // Kategori layanan langganan
	biaya_langganan float64 // Biaya yang harus dibayar untuk langganan
	tanggal_bayar   waktu   // Tanggal saat pembayaran dilakukan
	tenggat         waktu   // Batas waktu pembayaran berikutnya (jatuh tempo)
	metode          string  // Metode pembayaran (misal: transfer, kartu, dll)
	status          string  // Status langganan (aktif, nonaktif, dll)
}

// Tipe data bentukan untuk menyimpan informasi tenggat_bayar dan tenggat
type waktu struct {
	tanggal, bulan, tahun int
}

// Fungsi untuk menambahkan data langganan baru ke dalam array
func tambahLangganan(A *tabPelanggan, n *int) {
	// Meminta input data langganan secara berurutan dari pengguna
	fmt.Print("Nama Langganan: ")
	fmt.Scan(&A[*n].nama)
	fmt.Print("Kategori: ")
	fmt.Scan(&A[*n].kategori)
	fmt.Print("Biaya langganan: ")
	fmt.Scan(&A[*n].biaya_langganan)
	fmt.Print("Tanggal bayar: ")
	fmt.Scan(&A[*n].tanggal_bayar.tanggal, &A[*n].tanggal_bayar.bulan, &A[*n].tanggal_bayar.tahun)
	
	// Inisialisasi tanggal tenggat sama dengan tanggal bayar (jatuh tempo awal)
	A[*n].tenggat.tanggal = A[*n].tanggal_bayar.tanggal
	A[*n].tenggat.bulan = A[*n].tanggal_bayar.bulan
	A[*n].tenggat.tahun = A[*n].tanggal_bayar.tahun
	
	// Membaca metode pembayaran, dibuat demikian agar Metode dan Status terdapat pada line berbeda
	fmt.Print("Metode: ")
	fmt.Scanln()                          
	fmt.Scanln(&A[*n].metode)            
	
	// Membaca status langganan
	fmt.Print("Status: ")
	fmt.Scanln(&A[*n].status)
	
	// Menambah jumlah langganan yang tersimpan
	*n = *n + 1
	fmt.Println(" - - - - - - - - - - - - ")
}

// Fungsi untuk menampilkan seluruh data langganan yang tersimpan
func tampilkanLangganan(A *tabPelanggan, n *int) {
	// Mulai dari indeks 1 sampai jumlah langganan (karena indeks 0 belum dipakai)
	for i := 1; i < *n; i++ {
		fmt.Println(" - - - - - - - - - - - - ")
		fmt.Println("Langganan", i)
		fmt.Println("1. Nama:", A[i].nama)
		fmt.Println("2. Kategori:", A[i].kategori)
		fmt.Println("3. Biaya Langganan:", A[i].biaya_langganan)
		fmt.Println("3. Tanggal Bayar:", A[i].tanggal_bayar.tanggal, A[i].tanggal_bayar.bulan, A[i].tanggal_bayar.tahun)
		fmt.Print("4. Jatuh tempo: ")
		// Tampilkan tanggal jatuh tempo jika bulan jatuh tempo berbeda dengan bulan tanggal bayar
		if A[i].tanggal_bayar.bulan != A[i].tenggat.bulan {
			fmt.Print(A[i].tenggat.tanggal, A[i].tenggat.bulan, A[i].tenggat.tahun)
		}
		fmt.Println()
		fmt.Println("5. Metode: ", A[i].metode)
		fmt.Println("6. Status: ", A[i].status)
		fmt.Println(" - - - - - - - - - - - - ")
	}
}

// Fungsi untuk mengubah data langganan tertentu berdasarkan indeks yang diminta
func ubahLangganan(A *tabPelanggan, n *int) {
	var i int
	fmt.Print("Langganan ke:")
	fmt.Scan(&i) // Input nomor langganan yang ingin diubah
	if i < *n {
		// Meminta input data baru untuk langganan tersebut
		fmt.Println(" - - - - - - - - - - - - ")
		fmt.Print("Nama Langganan baru: ")
		fmt.Scan(&A[i].nama)
		fmt.Print("Kategori baru: ")
		fmt.Scan(&A[i].kategori)
		fmt.Print("Biaya Langganan baru: ")
		fmt.Scan(&A[i].biaya_langganan)
		fmt.Print("Tanggal bayar baru: ")
		fmt.Scan(&A[i].tenggat.tanggal, &A[i].tenggat.bulan, &A[i].tenggat.tahun)
		fmt.Print("Metode baru: ")
		fmt.Scan(&A[i].metode)
		fmt.Print("Status baru: ")
		fmt.Scan(&A[i].status)
		fmt.Println(" - - - - - - - - - - - - ")
	}
}

// Fungsi untuk menghapus data langganan berdasarkan indeks yang diminta
func hapusLangganan(A *tabPelanggan, n *int) {
	var i, hapus int
	fmt.Println(" - - - - - - - - - - - - ")
	fmt.Print("Langganan yang akan dihapus: ")
	fmt.Scan(&hapus) // Input indeks langganan yang akan dihapus
	
	// Kurangi jumlah langganan karena ada yang dihapus
	*n = *n - 1
	
	// Geser semua data setelah indeks hapus ke depan, menimpa data yang dihapus
	for i = hapus; i < *n; i++ {
		A[i] = A[i+1]
	}
	fmt.Println(" - - - - - - - - - - - - ")
}

// Fungsi untuk mencari langganan berdasarkan kategori secara sequential (berurutan)
func carikategori(A *tabPelanggan, n *int, x string) {
	var i int
	var ditemukan bool = false
	
	fmt.Println(" - - - - - - - - - - - - ")
	fmt.Print("kategori: ")
	fmt.Scan(&x) // Input kategori yang dicari
	
	// Telusuri satu per satu data langganan
	for i = 1; i < *n; i++ {
		if A[i].kategori == x {
			// Jika ketemu, tampilkan data langganan tersebut
			fmt.Println("Langganan", i)
			fmt.Println("1. Nama: ", A[i].nama)
			fmt.Println("2. Kategori: ", A[i].kategori)
			fmt.Println("3. Biaya Langganan: ", A[i].biaya_langganan)
			fmt.Println("4. Tanggal Bayar: ", A[i].tenggat.tanggal, A[i].tenggat.bulan, A[i].tenggat.tahun)
			fmt.Println("5. Metode: ", A[i].metode)
			fmt.Println("6. Status: ", A[i].status)
			fmt.Println(" - - - - - - - - - - - - ")
			ditemukan = true
		}
	}
	if !ditemukan {
		// Jika tidak ketemu sama sekali, beri tahu pengguna
		fmt.Println("Data tidak ditemukan")
	}
		fmt.Println(" - - - - - - - - - - - - ")
}

// Fungsi pengurutan data langganan berdasarkan nama secara ascending menggunakan metode insertion sort
func sortnama(A *tabPelanggan, n int) {
	var pass, i int
	var temp langganan
	pass = 1
	for pass < n {
		temp = A[pass]
		i = pass - 1
		// Geser data yang lebih besar ke kanan agar tempat untuk data yang sedang dibandingkan terbuka
		for i >= 0 && temp.nama < A[i].nama {
			A[i+1] = A[i]
			i = i - 1
		}
		A[i+1] = temp
		pass = pass + 1
	}
}

// Fungsi pencarian nama langganan menggunakan binary search setelah data diurutkan
func binaryUrutkannama(A *tabPelanggan, n int) {
	var kiri, kanan, tengah, idx int
	var x string
	kiri = 0
	kanan = n - 1
	idx = -1
	
	// Input nama yang dicari
	fmt.Println(" - - - - - - - - - - - - ")
	fmt.Print("Nama yang dicari (setelah diurutkan): ")
	fmt.Scan(&x)
	
	// Proses binary search: cari di tengah-tengah, perkecil pencarian berdasarkan perbandingan
	for (kiri <= kanan) && (idx == -1) {
		tengah = (kiri + kanan) / 2
		if x < A[tengah].nama {
			kanan = tengah - 1
		} else if x > A[tengah].nama {
			kiri = tengah + 1
		} else {
			idx = tengah
		}
	}
		fmt.Println(" - - - - - - - - - - - - ")
	
	// Jika ditemukan, tampilkan data langganan
	if idx != -1 {
		fmt.Println(" - - - - - - - - - - - - ")
		fmt.Println("Nama Langganan", idx)
		fmt.Println("1. Nama: ", A[idx].nama)
		fmt.Println("2. Kategori:", A[idx].kategori)
		fmt.Println("3. Biaya Langganan: ", A[idx].biaya_langganan)
		fmt.Println("4. Tanggal Bayar: ", A[idx].tenggat)
		fmt.Println("5. Metode: ", A[idx].metode)
		fmt.Println("6. Status: ", A[idx].status)
		fmt.Println("---")
	} else {
		// Jika tidak ditemukan, beri tahu pengguna
		fmt.Println("Data tidak ditemukan")
	}
		fmt.Println(" - - - - - - - - - - - - ")
}

// Fungsi pengurutan data langganan berdasarkan biaya langganan secara ascending menggunakan selection sort
func urutkanBiaya(A *tabPelanggan, n int) {
	var pass, i, idx int
	var temp langganan
	
	for pass = 0; pass < n-1; pass++ {
		idx = pass
		for i = pass + 1; i < n; i++ {
			if A[i].biaya_langganan < A[idx].biaya_langganan {
				idx = i
			}
		}
		// Tukar data agar biaya yang terkecil ada di posisi pass
		temp = A[pass]
		A[pass] = A[idx]
		A[idx] = temp
	}
}

// Fungsi untuk menghitung total biaya bulanan dari semua langganan yang statusnya aktif
func biayaBulanan(A *tabPelanggan, n *int) {
	var i int
	var jumlah float64 = 0
	
	// Jumlahkan biaya langganan yang statusnya aktif
	for i = 0; i < *n; i++ {
		if A[i].status == "Aktif" {
			jumlah = jumlah + A[i].biaya_langganan
		}
	}
	fmt.Println(" - - - - - - - - - - - - ")
	fmt.Println("Total biaya bulanan dari langganan aktif:", jumlah)
	fmt.Println(" - - - - - - - - - - - - ")
}

// Fungsi untuk mencari indeks langganan dengan biaya termahal
func biayaTermahal(A tabPelanggan, n int) int {
	var i, max int
	if n <= 1 {
		return 1
	}
	max = 1
	for i = 1; i < n; i++ {
		if A[i].biaya_langganan > A[max].biaya_langganan && A[i].status == "Aktif" {
			max = i
		}
	}
	return max
}

// Fungsi untuk memperbarui tanggal jatuh tempo dengan menambah satu bulan ke tanggal tenggat
func jatuhTempo(A *tabPelanggan, n int) {
	var i int
	for i = 0; i < n; i++ {
		// Cek apakah tanggal tenggat valid (tanggal <= 30, bulan <= 12, tahun positif)
		if A[i].tenggat.tanggal <= 30 && A[i].tenggat.bulan <= 12 && A[i].tenggat.tahun > 0 {
			// Tambahkan satu bulan ke tanggal tenggat
			A[i].tenggat.bulan = A[i].tenggat.bulan + 1
		}
	}
}

// Fungsi memberikan rekomendasi hemat dengan menyarankan langganan termahal untuk dihentikan
func rekomendasiHemat(A *tabPelanggan, n int) {
	var total float64 = 0
	var i, max int
	
	// Hitung total pengeluaran semua langganan
	for i = 1; i < n; i++ {
		if A[i].status == "Aktif" {
			total = total + A[i].biaya_langganan
		}
	}
	
	// Cari langganan dengan biaya termahal
	max = biayaTermahal(*A, n)
	
	// Tampilkan rekomendasi kepada pengguna
	fmt.Println(" - - - - - - - - - - - - ")
	fmt.Println("Total pengeluaran bulanan: ", total)
	fmt.Println("Langganan ke", max)
	fmt.Println("Untuk menghemat, pertimbangkan hentikan langganan: ")
	fmt.Println("Nama: ", A[max].nama)
	fmt.Println("Biaya: ", A[max].biaya_langganan)
	fmt.Println(" - - - - - - - - - - - - ")
}

// Fungsi memberikan rekomendasi hemat dengan menyarankan langganan termahal untuk dihentikan
func main() {
	var n int = 1             // Inisialisasi jumlah langganan, mulai dari 1 karena indeks awal pakai 1
	var p tabPelanggan        // Deklarasi variabel array untuk menampung data langganan
	var x string              // Variabel sementara untuk pencarian kategori
	var pilihan int = -1      // Variabel untuk menyimpan pilihan menu user, -1 agar masuk ke loop pertama

	for pilihan != 0 {       // Loop terus berjalan selama pilihan bukan 0 (keluar)
		// Menampilkan menu utama program
		fmt.Println("\n=== Manajemen Langganan Digital ===")
		fmt.Println(" - - - - - -  M E N U - - - - - - - ")
		fmt.Println("1. Tambah Langganan")
		fmt.Println("2. Tampilkan Langganan")
		fmt.Println("3. Ubah Langganan")
		fmt.Println("4. Hapus Langganan")
		fmt.Println("5. Cari layanan berdasarkan kategori") //sequential
		fmt.Println("6. Urutkan Nama Layanan") //insertion sort
		fmt.Println("7. Cari layanan setelah diurutkan") // binary search
		fmt.Println("8. Urutkan Biaya") // selection sort
		fmt.Println("9. Cek Jatuh Tempo")
		fmt.Println("10. Total Biaya Bulanan")
		fmt.Println("11. Biaya Termahal")
		fmt.Println("12. Rekomendasi langganan hemat")
		fmt.Println("0. Keluar")
		fmt.Println(" - - - - - -  M E N U - - - - - - - ")
		fmt.Print("Pilih menu: ")
		fmt.Scan(&pilihan)   // Membaca input pilihan dari user
		// Mengecek pilihan user dan memanggil fungsi sesuai menu yang dipilih
		if pilihan == 1 {
			tambahLangganan(&p, &n)          // Menambah data langganan baru
		} else if pilihan == 2 {
			tampilkanLangganan(&p, &n)       // Menampilkan semua data langganan yang sudah ada
		} else if pilihan == 3 {
			ubahLangganan(&p, &n)            // Mengubah data langganan tertentu
		} else if pilihan == 4 {
			hapusLangganan(&p, &n)           // Menghapus data langganan tertentu
		} else if pilihan == 5 {
			carikategori(&p, &n, x)          // Mencari langganan berdasarkan kategori secara sequential
		} else if pilihan == 6 {
			sortnama(&p, n)                  // Mengurutkan data langganan berdasarkan nama secara ascending
			tampilkanLangganan(&p, &n)       // Menampilkan data yang sudah diurutkan
		} else if pilihan == 7 {
			sortnama(&p,n)                   // Mengurutkan data langganan berdasarkan nama secara ascending
			binaryUrutkannama(&p,n)          // Mencari data langganan yang sudah terurut berdasarkan nama
		} else if pilihan ==  8{
			urutkanBiaya(&p, n)             // Mengurutkan data langganan berdasarkan biaya langganan secara ascending
			tampilkanLangganan(&p, &n)       // Menampilkan data yang sudah diurutkan
		} else if pilihan == 9 {
			jatuhTempo(&p, n)               // Mengecek dan memperbarui tanggal jatuh tempo langganan
			tampilkanLangganan(&p, &n)       // Menampilkan data setelah update jatuh tempo
		} else if pilihan == 10 {
			biayaBulanan(&p, &n)            // Menghitung total biaya langganan aktif per bulan
		} else if pilihan == 11 {
			max := biayaTermahal(p, n)       // Mencari langganan dengan biaya termahal
			fmt.Println(" - - - - - - - - - - - - ")
			fmt.Println("Langganan ke-", max)
			fmt.Println("Nama: ", p[max].nama)
			fmt.Println("Biaya: ", p[max].biaya_langganan)
			fmt.Println(" - - - - - - - - - - - - ")
		} else if pilihan == 12 {
			rekomendasiHemat(&p, n)         // Memberikan rekomendasi langganan mana yang bisa dihentikan untuk menghemat biaya
		} else if pilihan == 0 {
			fmt.Println("Keluar dari program") // Pesan keluar program
		} else {
			fmt.Println("Pilihan tidak valid") // Pesan jika input pilihan tidak dikenali
		}
	}
}

