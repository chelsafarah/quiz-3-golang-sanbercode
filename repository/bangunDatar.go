package repository

import (
	"Quiz-3/structs"
	"math"
)

func Luas(s structs.SegitigaSamaSisi) int64 {
	return (s.Alas * s.Tinggi) / 2
}

func Keliling(s structs.SegitigaSamaSisi) int64 {
	return 3 * s.Alas
}

func LuasPersegi(s structs.Persegi) int {
	return (s.Sisi * 4)
}

func KelilingPersegi(s structs.Persegi) int {
	return s.Sisi * s.Sisi
}

func LuasPersegiPanjang(p structs.PersegiPanjang) int {
	return p.Panjang * p.Lebar
}

func KelilingPersegiPanjang(s structs.PersegiPanjang) int {
	return (s.Panjang + s.Lebar) * 2
}

func LuasLingkaran(l structs.Lingkaran) float64 {
	return math.Pi * l.JariJari * l.JariJari
}

func KelilingLingkaran(l structs.Lingkaran) float64 {
	return 2 * math.Pi * l.JariJari
}
