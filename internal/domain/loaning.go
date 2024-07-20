package domain

import "time"

type Loaning struct {
	Id        string
	Book      Book
	Peminjam  User
	Status    string //DIPINJAM atau DIKEMBALIKAN
	CreatedAt time.Time
	UpdatedAt time.Time
}
