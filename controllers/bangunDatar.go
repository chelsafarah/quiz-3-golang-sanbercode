package controllers

import (
	"Quiz-3/repository"
	"Quiz-3/structs"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func SegitigaSamasisi(c *gin.Context) {
	var result string

	var s structs.SegitigaSamaSisi
	hitung := c.Query("hitung")

	alas, _ := strconv.Atoi(c.Query("alas"))
	tinggi, _ := strconv.Atoi(c.Query("tinggi"))

	s.Alas = int64(alas)
	s.Tinggi = int64(tinggi)

	if hitung == "keliling" {
		keliling := repository.Keliling(s)
		result = "Keliling segitiga sama sisi dengan sisi " + strconv.Itoa(alas) + " adalah " + strconv.Itoa(int(keliling))
	} else if hitung == "luas" {
		luas := repository.Luas(s)
		result = "Keliling segitiga sama sisi dengan alas" + strconv.Itoa(alas) + " dan tinggi " + strconv.Itoa(tinggi) + "adalah" + strconv.Itoa(int(luas))
	} else {
		result = "Hanya bisa menghitung luas dan keliling bangun datar"
	}

	c.JSON(http.StatusOK, result)
}

func Persegi(c *gin.Context) {
	var result string

	var p structs.Persegi
	hitung := c.Query("hitung")

	sisi, _ := strconv.Atoi(c.Query("sisi"))

	p.Sisi = sisi

	if hitung == "keliling" {
		keliling := repository.KelilingPersegi(p)
		result = "Keliling persegi dengan sisi " + strconv.Itoa(sisi) + " adalah " + strconv.Itoa(int(keliling))
	} else if hitung == "luas" {
		luas := repository.LuasPersegi(p)
		result = "Keliling persegi dengan sisi" + strconv.Itoa(sisi) + "adalah" + strconv.Itoa(int(luas))
	} else {
		result = "Hanya bisa menghitung luas dan keliling bangun datar"
	}

	c.JSON(http.StatusOK, result)
}

func PersegiPanjang(c *gin.Context) {
	var result string

	var s structs.PersegiPanjang
	hitung := c.Query("hitung")

	panjang, _ := strconv.Atoi(c.Query("panjang"))
	lebar, _ := strconv.Atoi(c.Query("lebar"))

	s.Panjang = panjang
	s.Lebar = lebar

	if hitung == "keliling" {
		keliling := repository.KelilingPersegiPanjang(s)
		result = "Keliling persegi panjang dengan panjang " + strconv.Itoa(panjang) + " dan lebar " + strconv.Itoa(lebar) + " adalah " + strconv.Itoa(int(keliling))
	} else if hitung == "luas" {
		luas := repository.LuasPersegiPanjang(s)
		result = "Keliling persegi panjang dengan panjang " + strconv.Itoa(panjang) + " dan lebar " + strconv.Itoa(lebar) + " adalah" + strconv.Itoa(int(luas))
	} else {
		result = "Hanya bisa menghitung luas dan keliling bangun datar"
	}

	c.JSON(http.StatusOK, result)
}

func Lingkaran(c *gin.Context) {
	var result string

	var l structs.Lingkaran
	hitung := c.Query("hitung")

	jariJari, _ := strconv.Atoi(c.Query("jariJari"))

	l.JariJari = float64(jariJari)

	if hitung == "keliling" {
		keliling := repository.KelilingLingkaran(l)
		result = "Keliling lingkaran dengan jari-jari " + string(jariJari) + " adalah " + strconv.Itoa(int(keliling))
	} else if hitung == "luas" {
		luas := repository.LuasLingkaran(l)
		result = "Keliling lingkaran dengan jari-jari" + string(jariJari) + " adalah" + strconv.Itoa(int(luas))
	} else {
		result = "Hanya bisa menghitung luas dan keliling bangun datar"
	}

	c.JSON(http.StatusOK, result)
}
