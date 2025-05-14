package main

import "fmt"

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

func initSaham(T *arrSaham, n *int) {
	T[0] = saham{"BBCA", 8500}
	T[1] = saham{"TLKM", 4200}
	T[2] = saham{"BBRI", 5100}
	*n = 3
}

func tampilkanHarga(T arrSaham, n int) {
	fmt.Println("Harga saham saat ini:")
	for i := 0; i < n; i++ {
		fmt.Printf("%s : %.0f\n", T[i].Kode, T[i].Harga)
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

func cariUser(T arrUser, n int, username string) int {
	for i := 0; i < n; i++ {
		if T[i].Username == username {
			return i
		}
	}
	return -1
}

func beliSaham(u *user, daftarSaham arrSaham, nSaham int, kode string, jumlah int) {
	for i := 0; i < nSaham; i++ {
		if daftarSaham[i].Kode == kode {
			total := daftarSaham[i].Harga * float64(jumlah)
			if u.Saldo >= total {
				u.Saldo -= total
				index := -1
				for j := 0; j < u.JumData; j++ {
					if u.Portofolio[j] == kode {
						index = j
						break
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
	fmt.Printf("Saldo: %.0f\n", u.Saldo)
	fmt.Println("Portofolio saham:")
	for i := 0; i < u.JumData; i++ {
		if u.Jumlah[i] > 0 {
			fmt.Println(u.Portofolio[i], ":", u.Jumlah[i], "lembar")
		}
	}
}

func tampilkanDaftarUser(T arrUser, n int) {
	var jumlahUser int
	if jumlahUser == 0{
		fmt.Println("Tidak ada user")
	}else {
		fmt.Println("Daftar User Terdaftar:")
		for i := 0; i < n; i++ {
			fmt.Printf("- %s (Saldo: %.0f)\n", T[i].Username, T[i].Saldo)

	}
	}
}

func tampilkanPortofolioSemuaUser(T arrUser, n int) {
	fmt.Println("Portofolio Semua User:")
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

func main() {
	var pengguna arrUser
	var jumlahUser int = 0
	var sahamList arrSaham
	var jumlahSaham int
	var selesai bool = false
	var username, password, kode string
	var jumlah int
	var nilai float64
	var pilihan int

	initSaham(&sahamList, &jumlahSaham)

	for !selesai {
		fmt.Println("\n=== Menu Awal ===")
		fmt.Println("1. Login")
		fmt.Println("2. Register")
		fmt.Println("3. Keluar")
		fmt.Print("Pilih: ")
		fmt.Scan(&pilihan)

		if pilihan == 1 {
			fmt.Print("Username: ")
			fmt.Scan(&username)
			fmt.Print("Password: ")
			fmt.Scan(&password)

			if username == "admin" && password == "123456" {
				// Menu Admin
				for {
					fmt.Println("\n=== Menu Admin ===")
					fmt.Println("1. Lihat Harga Saham")
					fmt.Println("2. Ubah Harga Saham")
					fmt.Println("3. Lihat Semua User")
					fmt.Println("4. Lihat Portofolio Semua User")
					fmt.Println("5. Logout")
					fmt.Print("Pilih: ")
					fmt.Scan(&pilihan)

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
						break
					} else {
						fmt.Println("Pilihan tidak valid")
					}
				}
			} else {
				index := cariUser(pengguna, jumlahUser, username)
				if index != -1 && pengguna[index].Password == password {
					// Menu User
					for {
						fmt.Println("\n=== Menu User ===")
						fmt.Println("1. Lihat Harga Saham")
						fmt.Println("2. Beli Saham")
						fmt.Println("3. Jual Saham")
						fmt.Println("4. Lihat Portofolio")
						fmt.Println("5. Logout")
						fmt.Print("Pilih: ")
						fmt.Scan(&pilihan)

						if pilihan == 1 {
							tampilkanHarga(sahamList, jumlahSaham)
						} else if pilihan == 2 {
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
							break
						} else {
							fmt.Println("Pilihan tidak valid")
						}
					}
				} else {
					fmt.Println("Username atau password salah")
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
					fmt.Println("Berhasil registrasi!")
				} else {
					fmt.Println("Username sudah digunakan")
				}
			}
		} else if pilihan == 3 {
			selesai = true
		} else {
			fmt.Println("Pilihan tidak valid")
		}
	}
}
