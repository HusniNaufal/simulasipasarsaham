package main

import (
	"fmt"
	"os"
	"os/exec"
)

var jumlahUser int = 0
type user struct {
	Username     string
	Password     string
	Saldo        float64
	Portofolio   [100]string
	Jumlah       [100]int
	JumData      int
}

type arrUser [100]user

type saham struct {
	Kode  string
	Harga float64
}

type arrSaham [100]saham


//Data  statis untuk awal, namun bsa ditambahkan
func initSaham(T *arrSaham, n *int) {
	T[0] = saham{"BBCA", 8500}
	T[1] = saham{"TLKM", 4200}
	T[2] = saham{"BBRI", 5100}
	*n = 3
}

//Berfungsi untuk mengClear CMD supaya lebih rapi
func bersihLayar() {
	cmd := exec.Command("cmd", "/c", "cls")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

//Menampilkan harga saham yg ada pada data
func tampilkanHarga(T arrSaham, n int) {
	bersihLayar()
	fmt.Println("Harga saham saat ini:")
	for i := 0; i < n; i++ {
		fmt.Printf("%s : %.0f\n", T[i].Kode, T[i].Harga)
	}
}

//Untuk mengubah harga dari saham
func ubahHargaSaham(T *arrSaham, n int, kode string, hargaBaru float64) {
	for i := 0; i < n; i++ {
		if T[i].Kode == kode {
			T[i].Harga = hargaBaru
			fmt.Println("Harga saham", kode, "diubah menjadi", hargaBaru)
			return
		}
	}
	fmt.Println("Kode saham tidak ditemukan.")
}

//Menambahkan saham baru
func tambahSahamBaru(T *arrSaham, n *int, kode string, harga float64) {

	for i:=0; i<*n; i++ {
		if T[i].Kode == kode {
			fmt.Println("Saham", kode, "sudah ada di daftar.")
			return
		}
	}
	if *n >= 100 {
		fmt.Println("Daftar saham sudah penuh.")
		return
	}
	T[*n] = saham{Kode: kode, Harga: harga}
	*n++
	fmt.Println("Saham", kode, "telah berhasil ditambahkan ke daftar.")
}

func cariUser(T arrUser, n int, username string) int {
	for i := 0; i < n; i++ {
		if T[i].Username == username {
			return i
		}
	}
	return -1
}


//Berfungsi untuk membeli saham
func beliSaham(u *user, daftarSaham arrSaham, nSaham int, kode string, jumlah int) {
	var ketemu bool = false
	bersihLayar()
	for i := 0; i < nSaham; i++ {
		if daftarSaham[i].Kode == kode {
			total := daftarSaham[i].Harga * float64(jumlah)
			if u.Saldo >= total {
				u.Saldo -= total
				index := -1
				for j := 0; j < u.JumData && !ketemu; j++ {
					if u.Portofolio[j] == kode {
						index = j
						ketemu = true
					}
				}
				if index == -1 {
					u.Portofolio[u.JumData] = kode
					u.Jumlah[u.JumData] = jumlah
					u.JumData++
				} else {
					u.Jumlah[index] += jumlah
				}
				fmt.Println("Berhasil beli", jumlah, "lembar", kode)
			} else {
				fmt.Println("Saldo tidak cukup")
			}
			return
		}
	}
	fmt.Println("Kode saham tidak ditemukan")
}


//Berfungsi untuk menjual saham
func jualSaham(u *user, daftarSaham arrSaham, nSaham int, kode string, jumlah int) {
	bersihLayar()
	for i := 0; i < nSaham; i++ {
		if daftarSaham[i].Kode == kode {
			for j := 0; j < u.JumData; j++ {
				if u.Portofolio[j] == kode {
					if u.Jumlah[j] >= jumlah {
						u.Jumlah[j] -= jumlah
						u.Saldo += daftarSaham[i].Harga * float64(jumlah)
						fmt.Println("Berhasil jual", jumlah, "lembar", kode)
					} else {
						fmt.Println("Saham tidak cukup")
					}
					return
				}
			}
			fmt.Println("Kamu tidak memiliki saham ini")
			return
		}
	}
	fmt.Println("Kode saham tidak ditemukan")
}


//Berfungsi untuk menampilkan portofolio masing-masing user
func tampilkanPortofolio(u user) {
	bersihLayar()
	fmt.Printf("Saldo: %.0f\n", u.Saldo)
	fmt.Println("Portofolio saham:")
	for i := 0; i < u.JumData; i++ {
		if u.Jumlah[i] > 0 {
			fmt.Println("-", u.Portofolio[i], ":", u.Jumlah[i], "lembar")
		}
	}
}


//Befungsi untuk menampilkan daftar user dan saldonya
func tampilkanDaftarUser(T arrUser, n int) {
	bersihLayar()
	if jumlahUser == 0{
		fmt.Println("Tidak ada user")
	}else {
		fmt.Println("Daftar User Terdaftar:")
		for i := 0; i < n; i++ {
			fmt.Printf("- %s (Saldo: %.0f)\n", T[i].Username, T[i].Saldo)
		}
	}
}


//Berfungsi untuk menampilkan portofolio user pada laman admin
func tampilkanPortofolioSemuaUser(T arrUser, n int) {
	bersihLayar()
	fmt.Println("Portofolio Semua User:")
	if jumlahUser == 0 {
		fmt.Println("Tidak ada user")
	}else {

		for i := 0; i < n; i++ {
			fmt.Println("User:", T[i].Username)
			fmt.Printf("  Saldo: %.0f\n", T[i].Saldo)
			for j := 0; j < T[i].JumData; j++ {
				if T[i].Jumlah[j] > 0 {
					fmt.Printf("  - %s: %d lembar\n", T[i].Portofolio[j], T[i].Jumlah[j])
				}
			}
		}
	}
}

func main() {
	var pengguna arrUser
	var sahamList arrSaham
	var jumlahSaham, jumlah, pilihan int
	var selesai bool = false
	var username, password, kode string
	var nilai float64

	initSaham(&sahamList, &jumlahSaham)

	for !selesai {
		fmt.Println("\n=== Menu Awal ===")
		fmt.Println("1. Login")
		fmt.Println("2. Register")
		fmt.Println("3. Lihat harga saham")
		fmt.Println("4. Keluar")
		fmt.Print("Pilih (1/2/3/4): ")
		fmt.Scan(&pilihan)
		fmt.Println("=================")

		if pilihan == 1 {
			fmt.Print("Username: ")
			fmt.Scan(&username)
			fmt.Print("Password: ")
			fmt.Scan(&password)

		if username == "admin" && password == "123456" {
			var lanjutAdmin bool = true
			for lanjutAdmin {
				fmt.Println("\n========= Menu Admin =========")
				fmt.Println("1. Lihat Harga Saham")
				fmt.Println("2. Ubah Harga Saham")
				fmt.Println("3. Lihat Semua User")
				fmt.Println("4. Lihat Portofolio Semua User")
				fmt.Println("5. Logout")
				fmt.Print("Pilih (1/2/3/4/5): ")
				fmt.Scan(&pilihan)
				fmt.Println("=======================")

				if pilihan == 1 {
					tampilkanHarga(sahamList, jumlahSaham)
				} else if pilihan == 2 {
					fmt.Print("Kode saham: ")
					fmt.Scan(&kode)
					fmt.Print("Harga baru: ")
					fmt.Scan(&nilai)
					ubahHargaSaham(&sahamList, jumlahSaham, kode, nilai)
				} else if pilihan == 3 {
					tampilkanDaftarUser(pengguna, jumlahUser)
				} else if pilihan == 4 {
					tampilkanPortofolioSemuaUser(pengguna, jumlahUser)
				} else if pilihan == 5 {
					lanjutAdmin = false  // Logout admin
				} else {
					fmt.Println("Pilihan tidak valid")
				}
			}
		} else {
			index := cariUser(pengguna, jumlahUser, username)
			if index != -1 && pengguna[index].Password == password {
				var lanjutUser bool = true
				for lanjutUser {
					fmt.Println("\n=== Menu User ===")
					fmt.Println("1. Lihat Harga Saham")
					fmt.Println("2. Beli Saham")
					fmt.Println("3. Jual Saham")
					fmt.Println("4. Lihat Portofolio")
					fmt.Println("5. Logout")
					fmt.Print("Pilih (1/2/3/4/5): ")
					fmt.Scan(&pilihan)
					fmt.Println("=================")

					if pilihan == 1 {
						tampilkanHarga(sahamList, jumlahSaham)
					} else if pilihan == 2 {
						bersihLayar()
						tampilkanHarga(sahamList, jumlahSaham)
						fmt.Println("=== Pembelian saham ===")
						fmt.Print("Kode saham: ")
						fmt.Scan(&kode)
						fmt.Print("Jumlah: ")
						fmt.Scan(&jumlah)
						beliSaham(&pengguna[index], sahamList, jumlahSaham, kode, jumlah)
					} else if pilihan == 3 {
						fmt.Print("Kode saham: ")
						fmt.Scan(&kode)
						fmt.Print("Jumlah: ")
						fmt.Scan(&jumlah)
						jualSaham(&pengguna[index], sahamList, jumlahSaham, kode, jumlah)
					} else if pilihan == 4 {
						tampilkanPortofolio(pengguna[index])
					} else if pilihan == 5 {
						lanjutUser = false  // Logout user
					} else {
						fmt.Println("Pilihan tidak valid")
					}
				}
			}
		}

		} else if pilihan == 2 {
			if jumlahUser >= 100 {
				fmt.Println("User maksimal tercapai.")
			} else {
				fmt.Print("Masukkan username baru: ")
				fmt.Scan(&username)
				if cariUser(pengguna, jumlahUser, username) == -1 {
					fmt.Print("Masukkan password: ")
					fmt.Scan(&password)
					pengguna[jumlahUser] = user{Username: username, Password: password, Saldo: 1000000}
					jumlahUser++
					bersihLayar()
					fmt.Println("Berhasil registrasi!")
				} else {
					fmt.Println("Username sudah digunakan")
				}
			}
		} else if pilihan == 3 {
			tampilkanHarga(sahamList, jumlahSaham)
		} else if pilihan == 4{
			selesai = true
		}else {
			fmt.Println("Pilihan tidak valid")
		}
	}
}
