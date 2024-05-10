package _test

import (
	"fmt"
	"testing"
	"time"
	"github.com/qintharganteng/backn/model"
	"github.com/qintharganteng/backn/module"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestGetAllAnggotaPerpustakaan(t *testing.T) {
	data := module.GetAllAnggotaPerpustakaan() // Panggil fungsi menggunakan nama paket
	fmt.Println(data)
}

func TestInsertSemua(t *testing.T) {
	nama := "Muhammad Qinthar"
	alamat := "Garut, Jawa Barat"
	noTelp := "082127854156"
	membershipID := "714220058"
	status := "Dipinjam"
	jamBuka := model.JamBuka{
		Hari:       "Senin",
		JamMulai:   "08:00",
		JamSelesai: "16:00",
	}

	// Tanggal pinjam dan tanggal kembali
	tanggalPinjam := time.Now()
	tanggalKembali := tanggalPinjam.Add(24 * time.Hour)

	// Lakukan insert data anggota perpustakaan
	insertedID := module.InsertAnggotaPerpustakaan(
		nama,
		alamat,
		noTelp,
		membershipID,
		jamBuka,
		status,
		tanggalPinjam,
		tanggalKembali,
	)

	fmt.Println("Inserted ID:", insertedID)

	// Lakukan pengujian terhadap fungsi InsertAnggotaPerpustakaan di sini
}

func TestGetbyid(t *testing.T) {
	id := "663c7b36df80b142c3ba2fc5"
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		t.Fatalf("error converting id to ObjectID: %v", err)
	}

	biodata, err := module.GetAnggotaPerpustakaanByID(objectID, module.MongoConn, "anggota_perpustakaan")
	if err != nil {
		t.Fatalf("error calling GetAnggotaPerpustakaanByID: %v", err)
	}

	fmt.Println(biodata)
}


   