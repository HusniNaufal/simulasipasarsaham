package main

import (
	"fmt"
)

func initHargaSaham(harga map[string]float64) {
	harga["BBCA"] = 8500
	harga["TLKM"] = 4200
	harga["BBRI"] = 5100
}

func tampilkanHarga(harga map[string]float64) {
	fmt.Println("Harga saham saat ini:")
	for kode, nilai := range harga {
		fmt.Printf("%s : %.0f\n", kode, nilai)
	}
}

func beliSaham(user string, saldo map[string]float64, portofolio map[string]map[string]int, harga map[string]float64, kode string, jumlah int) {
	total := harga[kode] * float64(jumlah)
	if saldo[user] >= total {
		saldo[user] = saldo[user] - total
		portofolio[user][kode] = portofolio[user][kode] + jumlah
		fmt.Println("Berhasil beli", jumlah, "lembar", kode)
	} else {
		fmt.Println("Saldo tidak cukup")
	}
}

func jualSaham(user string, saldo map[string]float64, portofolio map[string]map[string]int, harga map[string]float64, kode string, jumlah int) {
	if portofolio[user][kode] >= jumlah {
		portofolio[user][kode] = portofolio[user][kode] - jumlah
		saldo[user] = saldo[user] + harga[kode]*float64(jumlah)
		fmt.Println("Berhasil jual", jumlah, "lembar", kode)
	} else {
		fmt.Println("Saham tidak cukup")
	}
}

func tampilkanPortofolio(user string, saldo map[string]float64, portofolio map[string]map[string]int) {
	fmt.Printf("Saldo anda: %.0f\n", saldo[user])
	fmt.Println("Portofolio saham:")
	for kode, jumlah := range portofolio[user] {
		if jumlah > 0 {
			fmt.Println(kode, ":", jumlah, "lembar")
		}
	}
}

func ubahHargaSaham(harga map[string]float64, kode string, nilai float64) {
	harga[kode] = nilai
	fmt.Println("Harga saham", kode, "diubah menjadi", nilai)
}

func main() {
	var harga map[string]float64 = map[string]float64{}
	var saldo map[string]float64 = map[string]float64{
		"user":  1000000,
		"admin": 9999999,
	}
	var portofolio map[string]map[string]int = map[string]map[string]int{
		"user":  map[string]int{},
		"admin": map[string]int{},
	}

	initHargaSaham(harga)

	var login string
	fmt.Print("Login sebagai (user/admin): ")
	fmt.Scanln(&login)

	var pilihan int
	var kode string
	var jumlah int
	var nilai float64
	var selesai bool = false

	for !selesai {
		fmt.Println("\n--- Menu", login, "---")
		if login == "user" {
			fmt.Println("1. Lihat Harga Saham")
			fmt.Println("2. Beli Saham")
			fmt.Println("3. Jual Saham")
			fmt.Println("4. Lihat Portofolio")
			fmt.Println("5. Keluar")
			fmt.Print("Pilih: ")
			fmt.Scanln(&pilihan)

			if pilihan == 1 {
				tampilkanHarga(harga)
			} else if pilihan == 2 {
				fmt.Print("Kode saham: ")
				fmt.Scanln(&kode)
				fmt.Print("Jumlah: ")
				fmt.Scanln(&jumlah)
				beliSaham(login, saldo, portofolio, harga, kode, jumlah)
			} else if pilihan == 3 {
				fmt.Print("Kode saham: ")
				fmt.Scanln(&kode)
				fmt.Print("Jumlah: ")
				fmt.Scanln(&jumlah)
				jualSaham(login, saldo, portofolio, harga, kode, jumlah)
			} else if pilihan == 4 {
				tampilkanPortofolio(login, saldo, portofolio)
			} else if pilihan == 5 {
				selesai = true
			} else {
				fmt.Println("Pilihan tidak valid")
			}

		} else if login == "admin" {
			fmt.Println("1. Lihat Harga Saham")
			fmt.Println("2. Ubah Harga Saham")
			fmt.Println("3. Lihat Portofolio User")
			fmt.Println("4. Keluar")
			fmt.Print("Pilih: ")
			fmt.Scanln(&pilihan)

			if pilihan == 1 {
				tampilkanHarga(harga)
			} else if pilihan == 2 {
				fmt.Print("Kode saham: ")
				fmt.Scanln(&kode)
				fmt.Print("Harga baru: ")
				fmt.Scanln(&nilai)
				ubahHargaSaham(harga, kode, nilai)
			} else if pilihan == 3 {
				tampilkanPortofolio("user", saldo, portofolio)
			} else if pilihan == 4 {
				selesai = true
			} else {
				fmt.Println("Pilihan tidak valid")
			}
		} else {
			fmt.Println("Login tidak dikenali.")
			selesai = true
		}
	}
}
