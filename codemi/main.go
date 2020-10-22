package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type indentitas struct {
	tipe  string
	nomor int
}

func main() {
	var n int

	fmt.Print("init: ")
	if _, err := fmt.Scanln(&n); err != nil {
		fmt.Println("Input jumlah loker gagal, karena", err)
		return
	}

	inputCommand(n)
}

func inputCommand(n int) {
	scanner := bufio.NewScanner(os.Stdin)
	loker := make([]indentitas, n)
	var emptyIdentitas indentitas

	for scanner.Scan() {
		command := scanner.Text()
		commandSplit := strings.Split(command, " ")
		commandLength := (len(commandSplit))

		switch commandSplit[0] {
		case "status", "STATUS":
			fmt.Print("No loker \t Tipe identitas \t No identitas\n")
			for i := 0; i < n; i++ {
				fmt.Printf("%d \t\t %s \t\t\t %d\n", i+1, loker[i].tipe, loker[i].nomor)
			}
		case "input", "INPUT":
			if commandLength == 3 {
				var lokerIndex int = n

			searchEmptyIndex:
				for i := 0; i < n; i++ {
					if loker[i] == emptyIdentitas {
						lokerIndex = i
						break searchEmptyIndex
					}
				}

				if lokerIndex == n {
					fmt.Println("Maaf loker sudah penuh")
				} else {
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

				convStrToInt, err := strconv.Atoi(commandSplit[1])
				if err != nil {
					fmt.Println("Input nomor identitas gagal, karena", err)
					continue
				}

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

			findByNomor:
				for i := 0; i < n; i++ {
					if loker[i].nomor == convStrToInt {
						lokerIndex = i
						break findByNomor
					}
				}

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
				searchByTipe := []int{}
				for i := 0; i < n; i++ {
					if loker[i].tipe == commandSplit[1] {
						searchByTipe = append(searchByTipe, loker[i].nomor)
					}
				}

				if len(searchByTipe) == 0 {
					fmt.Printf("loker dengan tipe identitas %s tidak ditemukan\n", commandSplit[1])
				} else {
					fmt.Println(searchByTipe)
				}
			} else {
				fmt.Println("perintah search tidak valid.")
				fmt.Println("format: leave [tipe identitas].")
			}
		case "exit", "EXIT":
			os.Exit(0)
		default:
			fmt.Println("perintah tidak ditemukan")
		}
	}
}
