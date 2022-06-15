package main

import (
	"fmt"
	"math/rand"
	"time"
)

// modul random
// meminjam modul time untuk membangkitkan seed
type domino struct {
	left, right int
}

const NMAX = 28

type arrDomino [NMAX]domino
type arrInt [NMAX]int

func main() {
	// KERJAKAN MANDIRI dan BOLEH BERTANYA KE ASISTEN PRAKTIKUM
	var playerScore int = 0
	var dealerScore int = 0
	var boneyard, playerTiles, dealerTiles arrDomino
	var decision, nBoneyard, pSkor, dSkor, total int
	fmt.Println("Welcome to the Algorithm and Programming Practicum")
	rand.Seed(time.Now().UnixNano())
	// membuat 28 Tile ke dalam boneyard
	buatTile(&boneyard)
	// inisialisasi jumlah Tile di boneyard
	nBoneyard = NMAX
	// total pertandingan masih 0
	total = 0
	// menampilkan jumlah Tile di boneyard
	fmt.Println("Boneyard tiles:", nBoneyard)
	// lakukan loop selama decision adalah 0 dan di boneyard masih cukup Tile
	for decision == 0 && nBoneyard >= 4 {
		// menampilkan skor
		fmt.Printf("Your score is %d/%d \n\n", playerScore, total)
		// dealing atau melakukan pengacakan Tile
		acakTile(&boneyard, nBoneyard)
		// membagikan 2 Tile dari boneyard kepada player
		bagiTile(&boneyard, &playerTiles, &nBoneyard)
		// membagikan 2 Tile dari boneyard kepada dealer
		bagiTile(&boneyard, &dealerTiles, &nBoneyard)
		// tampilkan Tile player
		printTile(playerTiles, "player")
		// menampilkan jumlah Tile di boneyard
		fmt.Println("Boneyard tiles:", nBoneyard)
		// minta decision dengan memanggil getDecision
		getDecision(&decision)

		// lakukan loop selama pilihan adalah 1 atau 2 yaitu replace tile
		for decision == 1 || decision == 2 {
			// ganti satu kartu pemain berdasarkan nilai decision
			replaceTile(&boneyard, &playerTiles, &nBoneyard, decision)
			// tampilkan Tile player
			printTile(playerTiles, "player")
			// menampilkan jumlah Tile di boneyard
			fmt.Println("Boneyard tiles:", nBoneyard)
			// minta decision dengan memanggil getDecision
			getDecision(&decision)
		}

		// apabila pilihan adalah 9 maka keluar
		if decision == 9 {
			break
		}

		// apabila pilihan adalah 0 maka hitung skor dan pemenang
		if decision == 0 {
			// tampilan Tile dealer
			printTile(dealerTiles, "dealer")
			// cari pemenang dengan memanggil subprogram menang
			menang(playerTiles, dealerTiles, &pSkor, &dSkor)

			// tampilkan pemenang putaran ini
			if pSkor == 1 {
				fmt.Println("You won")
				playerScore++
			} else if dSkor == 1 {
				fmt.Println("You lose")
				dealerScore++
			}
			total++
		}
	}

	if total != 0 {
		fmt.Printf("Your last score is %d/%d \n", playerScore, total)
		fmt.Println("Thank you for playing with us.")
		fmt.Println("Your winning rate is", playerScore*100/total, "%")
	}
}
func getDecision(decision *int) {
	/* I.S. decision telah siap pada piranti masukan
		F.S. decision berisi nilai "0", "1", "2", atau "9", di luar itu minta masukan
	   kembali hingga valid*/
	// KERJAKAN MANDIRI
	// fmt.Println("Hallo")
	var input int
	fmt.Println("Choose your action:")
	fmt.Println("1. Replace tile")
	fmt.Println("2. Pass")
	fmt.Println("0. End turn")
	fmt.Println("9. Exit")
	fmt.Scan(&input)
	for input != 0 && input != 1 && input != 2 && input != 9 {
		fmt.Println("Invalid input, please try again")
		fmt.Scan(&input)
	}
	*decision = input

}

func buatTile(T *arrDomino) {
	/* I.S. -
	F.S. T berisi 28 Tile domino */
	var kiri, kanan, i int
	i = 0
	for kiri = 0; kiri <= 6; kiri++ {
		for kanan = 0; kanan <= kiri; kanan++ {
			T[i].left = kiri
			T[i].right = kanan
			i++
		}
	}
}

func searchInt(T arrInt, n, x int) int {
	/* mengembalikan posisi x pada array T yang berisi n bilangan bulat, -1 apabila tidak
	ditemukan */
	// KERJAKAN MANDIRI
	for i := 0; i < n; i++ {
		if T[i] == x {
			return i
		}
	}
	return -1

}

func acakTile(T *arrDomino, n int) {
	/* I.S. terdefinisi array T berisi sejumlah n Tile domino
	F.S. array T tersusun acak */
	// BOLEH DIBANTU ASISTEN PRAKTIKUM
	fmt.Println("Dealing ...")
	var temp arrInt
	var temp2 arrDomino = *T
	// 1. buat array of integer (temp) yang berisi bilangan acak dari 1 hingga n
	// 2. bilangan acak yang dibuat hanya akan ditambahkan ke array (temp) apabila belum ada di array
	// 3. temp akan berisi bilangan 1 hingga n dengan susunan acak dan tidak ada duplikat
	var x, i int

	for i = 0; i < n; i++ {
		x = rand.Intn(n)
		for searchInt(temp, i, x) != -1 {
			x = rand.Intn(n)
		}
		temp[i] = x
	}
	// 4. gunakan temp sebagai indeks dari temp2 sehingga tile pada T akan tersusun acak
	for i = 0; i < n; i++ {
		T[i] = temp2[temp[i]]
	}

}

func bagiTile(T, p *arrDomino, n *int) {
	/* I.S. terdefinisi array T yang berisi sejumlah n Tile domino
	F.S. pemain p memperoleh 2 Tile terakhir dari T, nilai n berkurang
	catatan: ambil Tile sejumlah m yang terakhir dari T, nilai n berkurang*/
	// KERJAKAN MANDIRI
	p[0] = T[*n-1]
	p[1] = T[*n-2]
	*n -= 2
}

func replaceTile(T, p *arrDomino, n *int, posisi int) {
	/* I.S. terdefinisi array T yang berisi sejumlah n Tile domino, dan array p yang berisi
	sejumlah Tile pemain
	 F.S. mengambil satu Tile dari T dan mengganti sesuai nomor posisi Tile pada array
	p*/
	// KERJAKAN MANDIRI
	if posisi == 1 {
		p[0] = T[*n-1]
	} else {
		p[1] = T[*n-1]
	}
	*n -= 1

}

func haveBalak(p domino) bool {
	/* Mengembalikan true apabila p adalah balak, false untuk kondisi sebaliknya */
	// KERJAKAN MANDIRI
	if p.left == p.right {
		return true
	} else {
		return false
	}
}

func tilePoin(p domino) int {
	/* Mengembalikan total poin sisi left dan right dari p*/
	// KERJAKAN MANDIRI
	return p.left + p.right

}

func haveTwoBalak(p arrDomino) bool {
	/* Mengembalikan true apabila pemain p memiliki 2 Tile balak. */
	// KERJAKAN MANDIRI
	return haveBalak(p[0]) && haveBalak(p[1])

}

func menang(p1, p2 arrDomino, p1Skor, p2Skor *int) {
	/* I.S. terdefinisi 2 Tile yang masing-masing dimiliki oleh pemain p1 dan p2
	 F.S. p1Skor bernilai 1 dan p2Skor bernilai 0 apabila p1 menang, atau p1Skor bernilai
	0 dan p2Skor bernilai 1 apabila p2 menang */
	// KERJAKAN MANDIRI dan BOLEH BERTANYA KE ASISTEN PRAKTIKUM
	var poin1, poin2 int
	*p1Skor = 0
	*p2Skor = 0
	// hitung total poin setiap pemain
	poin1 = tilePoin(p1[0]) + tilePoin(p1[1])
	poin2 = tilePoin(p2[0]) + tilePoin(p2[1])
	// hanya p1 yang memiliki 2 balak
	if haveTwoBalak(p1) && !haveTwoBalak(p2) {
		*p1Skor = 1
		// hanya p2 yang memilikiki 2 balak
	} else if !haveTwoBalak(p1) && haveTwoBalak(p2) {
		*p2Skor = 1
		// untuk kondisi yang lain
	} else {
		if poin1 > poin2 {
			*p1Skor = 1
		} else if poin2 > poin1 {
			*p2Skor = 1
		} else {
			fmt.Println("Draw!!")
		}
	}
}

func printTile(p arrDomino, s string) {
	/* I.S. terdefinisi 2 Tile pada array p, dan sting s yang berisi nama pemain
	F.S. menampilkan isi dari p dan s sesuai contoh masukan dan keluaran */
	var teks string
	if s == "player" {
		teks = "Your tiles"
	} else { // s == "dealer"
		teks = "Dealer tiles"
	}
	fmt.Printf("%s: (%d,%d) (%d,%d)\n", teks, p[0].left, p[0].right, p[1].left,
		p[1].right)
}
