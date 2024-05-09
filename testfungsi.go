package planetyanglain

import (
	"fmt"
	"testing"
)


func TestInsertAnggotaPerpustakaan(t *testing.T) {
	nama := "Muhammad Qinthar"
	alamat := "Garut"
	noTelp := "081234567890"
	membershipID := "714220058"
	hasil := InsertAnggotaPerpustakaan(nama, alamat, noTelp, membershipID)
	fmt.Println(hasil)
}

func TestGetAnggotaPerpustakaanFromPhoneNumber(t *testing.T) {
	noTelp := "081234567890"
	anggota := GetAnggotaPerpustakaanFromPhoneNumber(noTelp)
	fmt.Println(anggota)
}

func TestGetAllAnggotaPerpustakaan(t *testing.T) {
	data := GetAllAnggotaPerpustakaan()
	fmt.Println(data)
}
