package planetyanglain

import (
	"context"
	"errors"
	"fmt"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var MongoString string = os.Getenv("MONGOSTRING")

func MongoConnect(dbname string) (db *mongo.Database) {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(MongoString))
	if err != nil {
		fmt.Printf("MongoConnect: %v\n", err)
	}
	return client.Database(dbname)
}

func InsertOneDoc(db string, collection string, doc interface{}) (insertedID interface{}) {
	insertResult, err := MongoConnect(db).Collection(collection).InsertOne(context.TODO(), doc)
	if err != nil {
		fmt.Printf("InsertOneDoc: %v\n", err)
	}
	return insertResult.InsertedID
}

func InsertPeminjamanBuku(anggotaID primitive.ObjectID, bukuID primitive.ObjectID, tanggalPinjam time.Time, tanggalKembali time.Time, status string) (insertedID interface{}) {
	var peminjaman PeminjamanBuku
	peminjaman.AnggotaID = anggotaID
	peminjaman.BukuID = bukuID
	peminjaman.TanggalPinjam = primitive.NewDateTimeFromTime(tanggalPinjam.UTC())
	peminjaman.TanggalKembali = primitive.NewDateTimeFromTime(tanggalKembali.UTC())
	peminjaman.Status = status
	return InsertOneDoc("tesdb2024", "peminjaman_buku", peminjaman)
}

func InsertJamBuka(hari string, jamMulai string, jamSelesai string) (insertedID interface{}) {
	var jamBuka JamBuka
	jamBuka.Hari = hari
	jamBuka.JamMulai = jamMulai
	jamBuka.JamSelesai = jamSelesai
	return InsertOneDoc("tesdb2024", "jam_buka", jamBuka)
}

func InsertAnggotaPerpustakaan(nama string, alamat string, noTelp string, membershipID string, jamBuka JamBuka, peminjaman []PeminjamanBuku) (insertedID interface{}) {
	var anggota AnggotaPerpustakaan
	anggota.ID = primitive.NewObjectID()
	anggota.Nama = nama
	anggota.Alamat = alamat
	anggota.NoTelp = noTelp
	anggota.MembershipID = membershipID
	anggota.JamBuka = jamBuka
	anggota.Peminjaman = peminjaman

	return InsertOneDoc("tesdb2024", "anggota_perpustakaan", anggota)
}

func GetAllPeminjamanBuku() (data []PeminjamanBuku) {
	karyawan := MongoConnect("tesdb2024").Collection("peminjaman_buku")
	filter := bson.M{}
	cursor, err := karyawan.Find(context.TODO(), filter)
	if err != nil {
		fmt.Println("GetAllPeminjamanBuku:", err)
	}
	err = cursor.All(context.TODO(), &data)
	if err != nil {
		fmt.Println(err)
	}
	return
}

func GetAllJamBuka() (data []JamBuka) {
	koleksi := MongoConnect("tesdb2024").Collection("jam_buka")
	filter := bson.M{}
	cursor, err := koleksi.Find(context.TODO(), filter)
	if err != nil {
		fmt.Println("GetAllJamBuka:", err)
	}
	err = cursor.All(context.TODO(), &data)
	if err != nil {
		fmt.Println(err)
	}
	return
}

func GetAllAnggotaPerpustakaan() (data []AnggotaPerpustakaan) {
	koleksi := MongoConnect("tesdb2024").Collection("anggota_perpustakaan")
	filter := bson.M{}
	cursor, err := koleksi.Find(context.TODO(), filter)
	if err != nil {
		fmt.Println("GetAllAnggotaPerpustakaan:", err)
	}
	err = cursor.All(context.TODO(), &data)
	if err != nil {
		fmt.Println(err)
	}
	return
}

func GetAllPeminjaman(db *mongo.Database, col string) (data []AnggotaPerpustakaan, jam []JamBuka, peminjaman []PeminjamanBuku) {
	koleksi := db.Collection(col)
	filter := bson.M{}
	cursor, err := koleksi.Find(context.TODO(), filter)
	if err != nil {
		fmt.Println("GetAllPeminjamanBuku:", err)
		return nil, nil, nil
	}
	err = cursor.All(context.TODO(), &data)
	if err != nil {
		fmt.Println(err)
	}
	return data, nil, nil
}


func GetAnggotaPerpustakaanByID(_id primitive.ObjectID, db *mongo.Database, col string) (anggota AnggotaPerpustakaan, errs error) {
	perpustakaan := db.Collection(col)
	filter := bson.M{"_id": _id}
	err := perpustakaan.FindOne(context.TODO(), filter).Decode(&anggota)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return anggota, fmt.Errorf("no data found for ID %s", _id)
		}
		return anggota, fmt.Errorf("error retrieving data for ID %s: %s", _id, err.Error())
	}
	return anggota, nil
}

//ghhgh
