package main

import "fmt"

type Address struct {
	city, province, country string
}

func main() {
	var address1 Address = Address{
		city:     "Bandung",
		province: "Jawa Barat",
		country:  "Indonesia",
	}

	// POINTER MEREFERENSI KE LOKASI TEMPAT VARIABLE ITU MENYIMPAN DATA BUKAN MENDUPLIKAT
	// JIKA VALUE TERSEBUT DIUBAH MENGGUNAKAN POINTER, MAKA VARIABLE TERSEBUT AKAN BERUBAH KESELURUHAN

	var address2 *Address = &address1
	address2.country = "Switzerland"

	fmt.Println(address1)
	fmt.Println(*address2)
}
