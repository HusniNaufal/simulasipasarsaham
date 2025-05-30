/*
Anggota :
1. Muhammad Kresna
2. M Husni Naufal K
*/
package main

import (
	"fmt"
	"os"
	"os/exec"
)

var jumlahUser int = 0

type user struct {
	Username   string
	Password   string
	Saldo      float64
	Portofolio [100]string
	Jumlah     [100]int
	JumData    int
}

type arrUser [100]user

type saham struct {
	Kode             string
	NamaPerusahaan   string
	Harga            float64
	TotalKepemilikan string
}

type arrSaham [100]saham

func initSaham(T *arrSaham, n *int) {
	T[0] = saham{"BBCA", "Bank Central Asia", 8500, "75%"}
	T[1] = saham{"TLKM", "Telkom Indonesia", 4200, "65%"}
	T[2] = saham{"BBRI", "Bank Rakyat Indonesia", 5100, "74%"}
	*n = 3
}

func bersihLayar() {
	cmd := exec.Command("cmd", "/c", "cls")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func tampilkanHarga(T arrSaham, n int) {
	bersihLayar()
	fmt.Println("Harga saham saat ini:")
	for i := 0; i < n; i++ {
		fmt.Printf("%s : %.0f\n", T[i].Kode, T[i].Harga)
	}
}


//melihat profil perusahan (menggunakan linear search)
func lihatProfilPerusahaan(T arrSaham, n int, kode string) {
	var i int
	for i = 0; i < n; i++ {
		if T[i].Kode == kode {
			bersihLayar()
			fmt.Println("===== Profil Perusahaan =====")
			fmt.Println("Kode Saham       :", T[i].Kode)
			fmt.Println("Harga Saham 	 :", T[i].Harga)
			fmt.Println("Nama Perusahaan  :", T[i].NamaPerusahaan)
			fmt.Println("Total Kepemilikan:", T[i].TotalKepemilikan)
			return
		}
	}
	fmt.Println("Kode saham tidak ditemukan.")
}

func selectionSortTermurah(T *arrSaham, n int) {
	var i, j, min int
	var temp saham
	for i = 0; i < n-1; i++ {
		min = i
		for j = i + 1; j < n; j++ {
			if T[j].Harga < T[min].Harga {
				min = j
			}
		}
		temp = T[i]
		T[i] = T[min]
		T[min] = temp
	}
}

func insertionSortTermahal(T *arrSaham, n int) {
	var i, j int
	var key saham
	for i = 1; i < n; i++ {
		key = T[i]
		j = i - 1
		for j >= 0 && T[j].Harga < key.Harga {
			T[j+1] = T[j]
			j--
		}
		T[j+1] = key
	}
}

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

func tambahSahamBaru(T *arrSaham, n *int, kode string, nama string, kepemilikan string, nilai float64) {
	for i := 0; i < *n; i++ {
		if T[i].Kode == kode {
			fmt.Println("Saham", kode, "sudah ada di daftar.")
			return
		}
	}
	if *n >= 100 {
		fmt.Println("Daftar saham sudah penuh.")
		return
	}
	T[*n] = saham{
		Kode:             kode,
		NamaPerusahaan:   nama,
		TotalKepemilikan: kepemilikan,
		Harga:            nilai,
	
	}
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

func beliSaham(u *user, daftarSaham arrSaham, nSaham int, kode string, jumlah int) {
	var i, j, index int
	var ketemu bool = false
	bersihLayar()
	for i = 0; i < nSaham; i++ {
		if daftarSaham[i].Kode == kode {
			total := daftarSaham[i].Harga * float64(jumlah)
			if u.Saldo >= total {
				u.Saldo -= total
				index = -1
				for j = 0; j < u.JumData && !ketemu; j++ {
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

func tampilkanDaftarUser(T arrUser, n int) {
	bersihLayar()
	if jumlahUser == 0 {
		fmt.Println("Tidak ada user")
	} else {
		fmt.Println("Daftar User Terdaftar:")
		for i := 0; i < n; i++ {
			fmt.Printf("- %s (Saldo: %.0f)\n", T[i].Username, T[i].Saldo)
		}
	}
}

func tampilkanPortofolioSemuaUser(T arrUser, n int) {
	bersihLayar()
	fmt.Println("Portofolio Semua User:")
	if jumlahUser == 0 {
		fmt.Println("Tidak ada user")
	} else {
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
	var jumlahSaham, jumlah, pilihan, index int
	var selesai bool = false
	var username, password, kode, pilih string
	var nama, kepemilikan string
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
				bersihLayar()
				fmt.Println("Login sebagai ADMIN berhasil!")
				for lanjutAdmin {
					fmt.Println("\n========= Menu Admin =========")
					fmt.Println("1. Lihat Harga Saham")
					fmt.Println("2. Ubah Harga Saham")
					fmt.Println("3. Tambah Saham Baru")
					fmt.Println("4. Lihat Semua User")
					fmt.Println("5. Lihat Portofolio Semua User")
					fmt.Println("6. Logout")
					fmt.Print("Pilih (1/2/3/4/5/6): ")
					fmt.Scan(&pilihan)

					switch pilihan {
					case 1:
						bersihLayar()
						fmt.Println("1. Termahal ke Termurah (Descending)")
						fmt.Println("2. Termurah ke Termahal (Ascending)")
						fmt.Println("3. Liat profil perusahaan")
								fmt.Print("Pilihan: ")
								fmt.Scan(&pilihan)
							if pilihan == 1 {
								insertionSortTermahal(&sahamList, jumlahSaham)
								tampilkanHarga(sahamList, jumlahSaham)
							} else if pilihan == 2 {
								selectionSortTermurah(&sahamList, jumlahSaham)
								tampilkanHarga(sahamList, jumlahSaham)
							} else if pilihan == 3 {
								fmt.Print("Masukkan kode saham: ")
								fmt.Scan(&kode)
								lihatProfilPerusahaan(sahamList, jumlahSaham, kode)
							} 
					case 2:
						bersihLayar()
						tampilkanHarga(sahamList, jumlahSaham)
						fmt.Print("Kode saham: ")
						fmt.Scan(&kode)
						fmt.Print("Harga baru: ")
						fmt.Scan(&nilai)
						ubahHargaSaham(&sahamList, jumlahSaham, kode, nilai)
					case 3:
						fmt.Print("Kode saham baru: ")
						fmt.Scan(&kode)
						fmt.Print("Nama perusahaan: ")
						fmt.Scan(&nama)
						fmt.Print("Total kepemilikan (misal 75%): ")
						fmt.Scan(&kepemilikan)
						fmt.Print("Harga saham: ")
						fmt.Scan(&nilai)

						tambahSahamBaru(&sahamList, &jumlahSaham, kode, nama, kepemilikan, nilai)
					case 4:
						tampilkanDaftarUser(pengguna, jumlahUser)
					case 5:
						tampilkanPortofolioSemuaUser(pengguna, jumlahUser)
					case 6:
						lanjutAdmin = false
						bersihLayar()
					default:
						fmt.Println("Pilihan tidak valid")
					}
				}
			} else {
				index = cariUser(pengguna, jumlahUser, username)
				if index != -1 && pengguna[index].Password == password {
					var lanjutUser bool = true
					bersihLayar()
					fmt.Println("Hai ", username)
					for lanjutUser {
						fmt.Println("\n=== Menu User ===")
						fmt.Println("1. Lihat Harga Saham")
						fmt.Println("2. Beli Saham")
						fmt.Println("3. Jual Saham")
						fmt.Println("4. Lihat Portofolio")
						fmt.Println("5. Logout")
						fmt.Print("Pilih (1/2/3/4/5): ")
						fmt.Scan(&pilihan)
						switch pilihan {
							case 1:
								tampilkanHarga(sahamList, jumlahSaham)
								fmt.Println("1. Termahal ke Termurah (Descending)")
								fmt.Println("2. Termurah ke Termahal (Ascending)")
								fmt.Println("3. Liat profil perusahaan")
								fmt.Print("Pilihan: ")
								fmt.Scan(&pilihan)
							if pilihan == 1 {
								insertionSortTermahal(&sahamList, jumlahSaham)
								tampilkanHarga(sahamList, jumlahSaham)
							} else if pilihan == 2 {
								selectionSortTermurah(&sahamList, jumlahSaham)
								tampilkanHarga(sahamList, jumlahSaham)
							} else if pilihan == 3 {
								fmt.Print("Masukkan kode saham: ")
								fmt.Scan(&kode)
								lihatProfilPerusahaan(sahamList, jumlahSaham, kode)
							} 
							
						case 2:
							tampilkanHarga(sahamList, jumlahSaham)
							fmt.Print("Kode saham: ")
							fmt.Scan(&kode)
							fmt.Print("Jumlah: ")
							fmt.Scan(&jumlah)
							beliSaham(&pengguna[index], sahamList, jumlahSaham, kode, jumlah)
						case 3:
							fmt.Print("Kode saham: ")
							fmt.Scan(&kode)
							fmt.Print("Jumlah: ")
							fmt.Scan(&jumlah)
							jualSaham(&pengguna[index], sahamList, jumlahSaham, kode, jumlah)
						case 4:
							tampilkanPortofolio(pengguna[index])
						case 5:
							lanjutUser = false
						default:
							fmt.Println("Pilihan tidak valid")
						}
					}
				}
			}
		} else {
			switch pilihan {
			case 2:
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
			case 3:
				fmt.Println("1. Termahal ke Termurah (Descending)")
				fmt.Println("2. Termurah ke Termahal (Ascending)")
				fmt.Print("Pilihan: ")
				fmt.Scan(&pilihan)
				if pilihan == 1 {
					insertionSortTermahal(&sahamList, jumlahSaham)
				} else if pilihan == 2 {
					selectionSortTermurah(&sahamList, jumlahSaham)
				}
				tampilkanHarga(sahamList, jumlahSaham)
				fmt.Println("\nIngin lihat profil perusahaan? (y/n): ")
				fmt.Scan(&pilih)
				if pilih == "y" || pilih == "Y" {
					fmt.Print("Masukkan kode saham: ")
					fmt.Scan(&kode)
					lihatProfilPerusahaan(sahamList, jumlahSaham, kode)
				}
			case 4:
				selesai = true
				bersihLayar()
			default:
				fmt.Println("Pilihan tidak valid")
			}
		}
	}
}
