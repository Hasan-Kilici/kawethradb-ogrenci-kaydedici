package main

import (
	"fmt"
	"strconv"

	kawethradb "github.com/Hasan-Kilici/kawethradb"
)

type Ogrenci struct {
	ID    int
	Ad    string
	Soyad string
	Sinif int
}

func main() {
	ogrenciler := []Ogrenci{
		{ID: 1, Ad: "Ali", Soyad: "Veli", Sinif: 9},
		{ID: 2, Ad: "Ahmet", Soyad: "Mehmet", Sinif: 10},
		{ID: 3, Ad: "Ayşe", Soyad: "Fatma", Sinif: 11},
		{ID: 4, Ad: "Hasan", Soyad: "KILICI", Sinif: 12},
	}

	err := kawethradb.CreateDB("Ogrenciler", "./Ogrenciler.csv", ogrenciler)
	if err != nil {
		fmt.Println("err:", err)
		return
	}
	School()
}

func School() {
	for true {
		var komut string
		fmt.Println("Komut giriniz")
		fmt.Println("1: Öğrenci kayıt")
		fmt.Println("2: Öğrenci Düzenleme")
		fmt.Println("3: Öğrenci Silme")
		fmt.Println("4: Öğrenci Bulma")
		fmt.Scan(&komut)
		switch komut {
		case "1":
			var okulno int
			var ad string
			var soyad string
			var sinif int

			fmt.Println("Öğrenci numarası giriniz")
			fmt.Scan(&okulno)
			fmt.Println("Öğrenci adı giriniz")
			fmt.Scan(&ad)
			fmt.Println("Öğrenci soyadı giriniz")
			fmt.Scan(&soyad)
			fmt.Println("Öğrenci sınıfı giriniz")
			fmt.Scan(&sinif)
			yeniogrenci := Ogrenci{ID: okulno, Ad: ad, Soyad: soyad, Sinif: sinif}
			kawethradb.Insert("./Ogrenciler.csv", yeniogrenci)
			break
		case "2":
			var okulno int
			var ad string
			var soyad string
			var sinif string

			fmt.Println("Öğrenci numarası giriniz")
			fmt.Scan(&okulno)
			fmt.Println("Öğrenci adı giriniz")
			fmt.Scan(&ad)
			fmt.Println("Öğrenci soyadı giriniz")
			fmt.Scan(&soyad)
			fmt.Println("Öğrenci sınıfı giriniz")
			fmt.Scan(&sinif)
			yeniogrenci := []string{strconv.Itoa(okulno), ad, soyad, sinif}
			kawethradb.Update("./Ogrenciler.csv", "ID", okulno, yeniogrenci)
			break
		case "3":
			var okulno int

			fmt.Println("Öğrenci numarası giriniz")
			fmt.Scan(&okulno)
			kawethradb.Delete("./Ogrenciler.csv", "ID", okulno)
			break
		case "4":
			var okulno int

			fmt.Println("Öğrenci numarası giriniz")
			fmt.Scan(&okulno)
			find, _ := kawethradb.Find("./Ogrenciler.csv", "ID", okulno)
			fmt.Println(find)
			break
		}
	}
}
