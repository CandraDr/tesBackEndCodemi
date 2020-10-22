package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type identitas struct {
	tipe  string
	nomor int
}

func main() {
	var n int

	//Menentukan jumlah loker
	fmt.Print("init: ")

	//Jika inputan user bukan integer maka akan muncul error dan program akan berhenti
	if _, err := fmt.Scanln(&n); err != nil {
		fmt.Println("Input jumlah loker gagal, karena", err)
		return
	}

	//Memanggil fungsi inputCommand dengan passing parameter n, yaitu jumlah loker
	inputCommand(n)
}

func inputCommand(n int) {
	scanner := bufio.NewScanner(os.Stdin)

	//Membuat slice struct dengan panjang slice sesuai dengan jumlah loker
	loker := make([]identitas, n)
	//Membuat struct identitas kosong yang akan digunakan untuk perbandingan dan mengkosongkan slice struct
	var emptyIdentitas identitas

	fmt.Println("Selamat datang. Ketik perintah help untuk melihat daftar perintah.")

	//looping terus menerus sampai perintah exit dipilih
	for scanner.Scan() {
		//Menerima inputan user sebagai string
		command := scanner.Text()
		//Pisah string berdasarkan whitespace
		commandSplit := strings.Split(command, " ")
		//Panjang command yang sudah di split
		commandLength := (len(commandSplit))

		switch commandSplit[0] {
		case "status", "STATUS":
			fmt.Print("No loker \t Tipe identitas \t No identitas\n")
			for i := 0; i < n; i++ {
				fmt.Printf("%d \t\t %s \t\t\t %d\n", i+1, loker[i].tipe, loker[i].nomor)
			}
		case "input", "INPUT":
			//Jika format input user ada 3 buah seperti input sim 12345
			if commandLength == 3 {
				//Inisialisasi lokerIndex dengan panjang loker
				var lokerIndex int = n

				//Looping untuk mencari index loker yang datanya masih kosong
			searchEmptyIndex:
				for i := 0; i < n; i++ {
					if loker[i] == emptyIdentitas {
						//Jika ketemu assign lokerIndex dengan index sekarang dan keluar dari loop searchEmptyIndex
						lokerIndex = i
						break searchEmptyIndex
					}
				}

				//Jika lokerIndex sama dengan panjang loker artinya nilai belum berubah dari loop searchEmptyIndex berarti loker sudah penuh
				if lokerIndex == n {
					fmt.Println("Maaf loker sudah penuh")
				} else {
					//Input data identitas kedalam loker
					convStrToInt, err := strconv.Atoi(commandSplit[2])
					if err != nil {
						fmt.Println("Input nomor identitas gagal, karena", err)
						continue
					}

					loker[lokerIndex].tipe = commandSplit[1]
					loker[lokerIndex].nomor = convStrToInt
					fmt.Println("Kartu identitas tersimpan di loker no", lokerIndex+1)
				}
			} else {
				fmt.Println("perintah input tidak valid.")
				fmt.Println("format: input [tipe identitas] [nomor identitas].")
			}
		case "leave", "LEAVE":
			if commandLength == 2 {

				//Jika inputan bukan integer akan muncul peringatan
				convStrToInt, err := strconv.Atoi(commandSplit[1])
				if err != nil {
					fmt.Println("Input nomor identitas gagal, karena", err)
					continue
				} else if convStrToInt > n {
					//Jika inputan melebihi jumlah loker akan muncul peringatan
					fmt.Println("Nomor loker tidak ditemukan")
					continue
				}

				//Mengkosongkan data loker sesuai index
				loker[convStrToInt-1] = emptyIdentitas
				fmt.Printf("Loker nomer %d berhasil dikosongkan \n", convStrToInt)
			} else {
				fmt.Println("perintah leave tidak valid.")
				fmt.Println("format: leave [nomor loker].")
			}
		case "find", "FIND":
			if commandLength == 2 {
				convStrToInt, err := strconv.Atoi(commandSplit[1])
				if err != nil {
					fmt.Println("Input nomor identitas gagal, karena", err)
					continue
				}

				var lokerIndex int = n

				//Loop untuk mencari index yang sama dengan nomor identitas yang diinput
			findByNomor:
				for i := 0; i < n; i++ {
					if loker[i].nomor == convStrToInt {
						lokerIndex = i
						break findByNomor
					}
				}

				//Jika lokerIndex belum berubah nilainya maka nomor identitas tidak ditemukan
				if lokerIndex == n {
					fmt.Print("Nomor identitas tidak ditemukan\n")
				} else {
					fmt.Printf("Kartu identitas tersebut berada di loker no %d\n", lokerIndex+1)
				}
			} else {
				fmt.Println("perintah find tidak valid.")
				fmt.Println("format: find [nomor identitas].")
			}
		case "search", "SEARCH":
			if commandLength == 2 {
				//Membuat slice baru untuk menampung hasil pencarian
				searchByType := []int{}

				//Loop untuk mencari identitas yang tipenya sama dengan inputan
				for i := 0; i < n; i++ {
					if loker[i].tipe == commandSplit[1] {
						searchByType = append(searchByType, loker[i].nomor)
					}
				}

				//Jika panjang slice searchByType masih 0 maka pencarian tidak ditemukan
				if len(searchByType) == 0 {
					fmt.Printf("loker dengan tipe identitas %s tidak ditemukan\n", commandSplit[1])
				} else {
					fmt.Println(searchByType)
				}
			} else {
				fmt.Println("perintah search tidak valid.")
				fmt.Println("format: leave [tipe identitas].")
			}
		case "help", "HELP":
			help()
		case "exit", "EXIT":
			//Keluar dari program
			os.Exit(0)
		default:
			//Jika inputan user tidak sama dengan daftar perintah
			fmt.Println("perintah tidak ditemukan")
		}
	}
}

//Menampilkan daftar perintah yang ada
func help() {
	fmt.Println("\n--------------Command--------------")
	fmt.Println("1. status")
	fmt.Println("2. input [tipe identitas: string] [nomor identitas: integer]")
	fmt.Println("3. leave [nomor loker: integer]")
	fmt.Println("4. find [nomor identitas: integer]")
	fmt.Println("5. search [tipe identitas: string]")
	fmt.Println("6. exit")
	fmt.Println("--------------End Command--------------\n")
}
