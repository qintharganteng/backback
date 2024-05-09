package planetyanglain

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type AnggotaPerpustakaan struct {
	ID           primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	Nama         string             `bson:"nama,omitempty" json:"nama,omitempty"`
	Alamat       string             `bson:"alamat,omitempty" json:"alamat,omitempty"`
	NoTelp       string             `bson:"no_telp,omitempty" json:"no_telp,omitempty"`
	MembershipID string             `bson:"membership_id,omitempty" json:"membership_id,omitempty"`
	JamBuka      JamBuka            `bson:"jam_buka,omitempty" json:"jam_buka,omitempty"`
	Peminjaman   []PeminjamanBuku   `bson:"peminjaman,omitempty" json:"peminjaman,omitempty"`
}

type JamBuka struct {
	Hari       string `bson:"hari,omitempty" json:"hari,omitempty"`
	JamMulai   string `bson:"jam_mulai,omitempty" json:"jam_mulai,omitempty"`
	JamSelesai string `bson:"jam_selesai,omitempty" json:"jam_selesai,omitempty"`
}

type PeminjamanBuku struct {
	ID             primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	AnggotaID      primitive.ObjectID `bson:"anggota_id,omitempty" json:"anggota_id,omitempty"`
	BukuID         primitive.ObjectID `bson:"buku_id,omitempty" json:"buku_id,omitempty"`
	TanggalPinjam  primitive.DateTime `bson:"tanggal_pinjam,omitempty" json:"tanggal_pinjam,omitempty"`
	TanggalKembali primitive.DateTime `bson:"tanggal_kembali,omitempty" json:"tanggal_kembali,omitempty"`
	Status         string             `bson:"status,omitempty" json:"status,omitempty"`
}
