package main

import "fmt"

func MakeNegative(x int) int {
	if x >= 0 {
		return -x
	}
	return x
}

func main() {
	fmt.Println(MakeNegative(-1))
}

/*
Dalam tugas sederhana ini Anda diberi nomor dan harus menjadikannya negatif. Tapi mungkin angkanya sudah negatif?

Contoh

MakeNegative(1)    // return -1
MakeNegative(-5)   // return -5
MakeNegative(0)    // return 0
Catatan
Angkanya mungkin sudah negatif, sehingga tidak diperlukan perubahan.
Nol (0) tidak dicentang untuk tanda tertentu. Angka nol negatif tidak masuk akal secara matematis.
*/
