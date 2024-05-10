package module_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/qintharganteng/backn/model"
	"github.com/qintharganteng/backn/module"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// func TestInsertAnggotaPerpustakaan(t *testing.T) {
// 	jamBuka := model.JamBuka{
// 		Hari:       "Senin",
// 		JamMulai:   "08:00",
// 		JamSelesai: "16:00",
// 	}

// 	// Konversi string menjadi primitive.ObjectID
// 	anggotaID := primitive.NewObjectID()
// 	bukuID := primitive.NewObjectID()

// 	peminjaman := []model.PeminjamanBuku{
// 		model.PeminjamanBuku{
// 			AnggotaID:      anggotaID,
// 			BukuID:         bukuID,
// 			TanggalPinjam:  primitive.NewDateTimeFromTime(time.Now()),
// 			TanggalKembali: primitive.NewDateTimeFromTime(time.Now().AddDate(0, 0, 7)),
// 			Status:         "Dipinjam",
// 		},
// 	}

// 	insertedID := module.InsertAnggotaPerpustakaan("Muhammad Qinthar", "Garut, Jawa Barat", "082127854156", "member_id_1", jamBuka, peminjaman)
// 	fmt.Println("Inserted Anggota Perpustakaan ID:", insertedID)
// }

// func TestGetAllJamBuka(t *testing.T) {
// 	jamBuka := module.GetAllJamBuka()
// 	fmt.Println("All Jam Buka:", jamBuka)
// }

func TestGetAllAnggotaPerpustakaan(t *testing.T) {
	anggota := module.GetAllAnggotaPerpustakaan()
	fmt.Println("All Anggota Perpustakaan:", anggota)
}

func TestInsertAnggotaPerpustakaan(t *testing.T) {
	// Generate ObjectID secara otomatis
	objectID := primitive.NewObjectID()

	// Data untuk pengujian
	nama := "Muhammad Qinthar"
	alamat := "Garut, Jawa Barat"
	noTelp := "082127854156"
	membershipID := "member_id_1"
	status := "Dipinjam"
	jamBuka := model.JamBuka{
		Hari:       "Senin",
		JamMulai:   "08:00",
		JamSelesai: "16:00",
	}

	// Konversi string menjadi primitive.ObjectID
	anggotaID := objectID
	bukuID := primitive.NewObjectID()

	peminjaman := []model.PeminjamanBuku{
		model.PeminjamanBuku{
			AnggotaID:      anggotaID,
			BukuID:         bukuID,
			TanggalPinjam:  primitive.NewDateTimeFromTime(time.Now()),
			TanggalKembali: primitive.NewDateTimeFromTime(time.Now().AddDate(0, 0, 7)),
			Status:         status,
		},
	}

	// Lakukan insert data anggota perpustakaan
	insertedID := module.InsertAnggotaPerpustakaan(
		nama,
		alamat,
		noTelp,
		membershipID,
		jamBuka,
		peminjaman,
	)

	fmt.Println("Inserted Anggota Perpustakaan ID:", insertedID)
}

func TestGetbyid(t *testing.T) {
	id := "663da3598269dce16c485747"
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
