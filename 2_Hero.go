package main

import "fmt"

func Hero(bullets, dragons int) bool {
	// your code
	// if bullets >= dragons*2 {
	// 	return true
	// }
	// return false
	return bullets >= dragons*2
}

func main() {
	fmt.Println(Hero(10, 5))
	fmt.Println(Hero(7, 4))
	fmt.Println(Hero(4, 5))
	fmt.Println(Hero(100, 40))
	fmt.Println(Hero(1500, 751))
	fmt.Println(Hero(0, 1))
}

/*
Seorang pahlawan sedang dalam perjalanan ke kastil untuk menyelesaikan misinya.
Namun, dia diberitahu bahwa kastil tersebut dikelilingi oleh beberapa naga yang kuat!
setiap naga membutuhkan 2 peluru untuk dikalahkan, pahlawan kita tidak tahu berapa
banyak peluru yang harus dia bawa.. Dengan asumsi dia akan mengambil sejumlah peluru
tertentu dan bergerak maju untuk melawan naga lain dalam jumlah tertentu, akankah dia bertahan??
Kembalikan nilai benar jika ya, salah jika sebaliknya :)
*/
