package main

import "fmt"

type celcius struct {
	suhu float64
}

type reamur struct {
	suhu float64
}

type farenheit struct {
	suhu float64
}

type kelvin struct {
	suhu float64
}

// Celcius
func (c celcius) toCelcius() float64 {
	return c.suhu
}
func (c celcius) toReamur() float64 {
	return 4.0 / 5.0 * c.suhu
}
func (c celcius) toFarenheit() float64 {
	return (9.0 / 5.0 * c.suhu) + 32
}
func (c celcius) toKelvin() float64 {
	return c.suhu + 273.15
}

// Reamur
func (r reamur) toCelcius() float64 {
	return 5.0 / 4.0 * r.suhu
}
func (r reamur) toReamur() float64 {
	return r.suhu
}
func (r reamur) toFarenheit() float64 {
	return (9.0 / 4.0 * r.suhu) + 32
}
func (r reamur) toKelvin() float64 {
	return (5.0 / 4.0 * r.suhu) + 273.15
}

// Farenheit
func (f farenheit) toCelcius() float64 {
	return (f.suhu - 32) * 5.0 / 9.0
}
func (f farenheit) toReamur() float64 {
	return (f.suhu - 32) * 4.0 / 9.0
}
func (f farenheit) toFarenheit() float64 {
	return f.suhu
}
func (f farenheit) toKelvin() float64 {
	return (f.suhu + 459.67) * 5.0 / 9.0
}

// Kelvin
func (k kelvin) toCelcius() float64 {
	return k.suhu - 273.15
}
func (k kelvin) toReamur() float64 {
	return 4.0 / 5.0 * (k.suhu - 273.15)
}
func (k kelvin) toFarenheit() float64 {
	return ((9.0 / 5.0) * k.suhu) - 459.67
}
func (k kelvin) toKelvin() float64 {
	return k.suhu
}

type hitungSuhu interface {
	toCelcius() float64
	toReamur() float64
	toFarenheit() float64
	toKelvin() float64
}

func main() {
	fmt.Println("Pilihan Tipe Suhu Awal : ")
	fmt.Println("1. Celcius")
	fmt.Println("2. Reamur")
	fmt.Println("3. Farenheit")
	fmt.Println("4. Kelvin")
	fmt.Println("Masukan Tipe Suhu Awal: ")

	var tipeSuhuAwal int
	fmt.Scanln(&tipeSuhuAwal)
	for tipeSuhuAwal < 1 || tipeSuhuAwal > 4 {
		fmt.Println("Tipe suhu awal tidak valid, masukan kembali tipe yang benar: ")
		fmt.Scanln(&tipeSuhuAwal)
	}

	fmt.Println("Pilihan Tipe Suhu Akhir : ")
	fmt.Println("1. Celcius")
	fmt.Println("2. Reamur")
	fmt.Println("3. Farenheit")
	fmt.Println("4. Kelvin")
	fmt.Println("Masukan Tipe Suhu Akhir Yang Diinginkan : ")

	var tipeSuhuAkhir int
	fmt.Scanln(&tipeSuhuAkhir)
	for tipeSuhuAkhir < 1 || tipeSuhuAkhir > 4 {
		fmt.Println("Tipe suhu akhir tidak valid, masukan kembali tipe yang benar: ")
		fmt.Scanln(&tipeSuhuAkhir)
	}

	fmt.Println("Masukan nilai suhu : ")
	var suhu float64
	fmt.Scanln(&suhu)

	var interfaceSuhu hitungSuhu
	switch tipeSuhuAwal {
	case 1:
		interfaceSuhu = celcius{suhu}
	case 2:
		interfaceSuhu = reamur{suhu}
	case 3:
		interfaceSuhu = farenheit{suhu}
	case 4:
		interfaceSuhu = kelvin{suhu}
	}

	var suhuHasilKonversi float64
	switch tipeSuhuAkhir {
	case 1:
		suhuHasilKonversi = interfaceSuhu.toCelcius()
	case 2:
		suhuHasilKonversi = interfaceSuhu.toReamur()
	case 3:
		suhuHasilKonversi = interfaceSuhu.toFarenheit()
	case 4:
		suhuHasilKonversi = interfaceSuhu.toKelvin()
	}

	fmt.Printf("Suhu akhir yang didapat adalah %.2f", suhuHasilKonversi)
}
