package planetyanglain

import (
	"fmt"
	"testing"
	"time"

	"github.com/qintharganteng/backn/module"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestGetAllAnggotaPerpustakaan(t *testing.T) {
	data := GetAllAnggotaPerpustakaan()
	fmt.Println(data)
}

func TestInsertSemua(t *testing.T) {
	// Data untuk pengujian
	nama := "Muhammad Qinthar"
	alamat := "Garut,Jawa Barat"
	noTelp := "082127854156"
	membershipID := "714220058"
	hari := "Senin"
	jamMulai := "08:00"
	jamSelesai := "16:00"
	tanggalPinjam := time.Now()
	tanggalKembali := time.Now().Add(24 * time.Hour)
	status := "Dipinjam"

	// Membuat data jam buka
	jamBuka := JamBuka{
		Hari:       hari,
		JamMulai:   jamMulai,
		JamSelesai: jamSelesai,
	}

	// Membuat data peminjaman buku
	peminjaman := []PeminjamanBuku{
		PeminjamanBuku{
			BukuID:         primitive.NewObjectID(),
			TanggalPinjam:  primitive.NewDateTimeFromTime(tanggalPinjam.UTC()),
			TanggalKembali: primitive.NewDateTimeFromTime(tanggalKembali.UTC()),
			Status:         status,
		},
	}

	// Memanggil fungsi untuk memasukkan data
	insertedID := InsertAnggotaPerpustakaan(nama, alamat, noTelp, membershipID, jamBuka, peminjaman)
	fmt.Println(insertedID)
}

func TestGetbyid(t *testing.T) {
	id := "663c7b36df80b142c3ba2fc5"
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		t.Fatalf("error converting id to ObjectID: %v", err)
	}
	biodata, err := GetAnggotaPerpustakaanByID(objectID, module.MongoConn, "anggota_perpustakaan")
	if err != nil {
		t.Fatalf("error calling GetAnggotaPerpustakaanByID: %v", err)
	}
	fmt.Println(biodata)
}
