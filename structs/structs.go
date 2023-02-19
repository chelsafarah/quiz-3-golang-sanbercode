package structs

import "time"

type Category struct {
	ID         int       `json:"id"`
	Name       string    `json:"name"`
	Created_at time.Time `json:"created_at"`
	Updated_at time.Time `json:"updated_at"`
}

type Book struct {
	Id          int       `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	ImageUrl    string    `json:"image_url"`
	ReleaseYear int       `json:"release_year"`
	Price       string    `json:"price"`
	TotalPage   int       `json:"total_page"`
	Thickness   string    `json:"thickness"`
	Created_at  time.Time `json:"created_at"`
	Updated_at  time.Time `json:"updated_at"`
	Category_id int       `json:"category_id"`
}

type SegitigaSamaSisi struct {
	Alas   int64 `json:"alas"`
	Tinggi int64 `json:"tinggi"`
}

type Persegi struct {
	Sisi int `json:"sisi"`
}

type PersegiPanjang struct {
	Panjang int `json:"panjang"`
	Lebar   int `json:"lebar"`
}

type Lingkaran struct {
	JariJari float64 `json:"jari_jari"`
}
