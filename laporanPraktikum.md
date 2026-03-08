# <h1 align="center">Laporan Praktikum Modul 1 - ... </h1>
<p align="center">[EDWARD ABIMAS SURYA HATTA] - [109082500171]</p>

## Unguided 

### 1. [2A]
#### 1.go

```go
package main

import "fmt"

func main() {
	var (
		satu, dua, tiga string
		temp            string
	)

	fmt.Print("Masukan input string: ")
	fmt.Scanln(&satu)
	fmt.Print("Masukan input string: ")
	fmt.Scanln(&dua)
	fmt.Print("Masukan input string: ")
	fmt.Scanln(&tiga)

	fmt.Println("Output awal = " + satu + " " + dua + " " + tiga)

	temp = satu
	satu = dua
	dua = tiga
	tiga = temp

	fmt.Println("Output akhir = " + satu + " " + dua + " " + tiga)
}
```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 1_1](https://github.com/EdwardAbimas/109082500171_EDWARD-ABIMAS-SURYA-HATTA/raw/main/PERTEMUAN1/ss1.png)
[Program pada soal Modul 2A pada dasarnya melakukan manipulasi posisi data dari tiga inputan bertipe string. Saat program dijalankan, sistem akan meminta pengguna untuk memasukkan tiga buah kata secara berurutan, yang kemudian ditampung ke dalam variabel bernama satu, dua, dan tiga. Setelah input selesai, program mencetak gabungan ketiga kata tersebut sebagai output awal untuk menunjukkan kondisi data sebelum dimanipulasi. Logika utama dari program ini terletak pada proses pertukaran nilainya, di mana digunakan sebuah variabel tambahan bernama temp sebagai tempat penyimpanan sementara. Nilai dari variabel satu dipindahkan dulu ke variabel temp agar tidak hilang saat datanya ditimpa. Setelah itu, variabel satu langsung diisi dengan nilai dari variabel dua, dan variabel dua diisi dengan nilai dari variabel tiga. Sebagai langkah terakhir, variabel tiga diisi dengan nilai awal variabel satu yang sebelumnya sudah diamankan di variabel temp. Melalui algoritma ini, program berhasil melakukan rotasi pergeseran posisi kata ke arah kiri, di mana kata yang diinputkan pertama kali akan bergeser menjadi kata terakhir ketika program menampilkan output akhirnya ke layar.]

### 2. [2B]
#### 2.go

```go
package main

import "fmt"

func main() {
	var warna1, warna2, warna3, warna4 string
	berhasil := true

	for i := 1; i <= 5; i++ {
		fmt.Printf("Percobaan %d: ", i)
		fmt.Scan(&warna1, &warna2, &warna3, &warna4)

		if warna1 != "merah" || warna2 != "kuning" || warna3 != "hijau" || warna4 != "ungu" {
			berhasil = false
		}
	}

	fmt.Println("BERHASIL:", berhasil)
}
```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 1_1](https://github.com/EdwardAbimas/109082500171_EDWARD-ABIMAS-SURYA-HATTA/raw/main/PERTEMUAN1/ss2.png)
[Untuk tugas Modul 2B, program dirancang guna menyimulasikan pencatatan hasil praktikum kimia yang menggunakan empat buah tabung reaksi. Di dalam simulasi ini, keberhasilan suatu percobaan sangat bergantung pada susunan warna cairan di setiap tabungnya. Sesuai dengan spesifikasi soal, percobaan hanya dianggap berhasil jika susunan warna pada gelas pertama hingga keempat secara berurutan adalah merah, kuning, hijau, dan ungu selama lima kali percobaan berturut-turut. Untuk merealisasikan hal ini, program diimplementasikan menggunakan struktur perulangan agar dapat meminta input urutan warna sebanyak lima kali iterasi. Pada setiap iterasinya, algoritma program akan memvalidasi urutan warna yang dimasukkan. Apabila pengguna memasukkan warna merah, kuning, hijau, dan ungu secara tepat pada kelima percobaan tersebut tanpa ada kesalahan urutan, program akan memberikan status keberhasilan berupa nilai true di akhir eksekusi. Sebaliknya, jika terdapat satu saja percobaan yang urutan warnanya salah atau sengaja ditukar posisinya, program akan langsung menandai kegagalan dan pada akhirnya menampilkan status false.]

### 3. [2C]
#### 3.go

```go
package main

import "fmt"

func main() {
	var berat int
	fmt.Print("Berat parsel (gram): ")
	fmt.Scan(&berat)

	kg := berat / 1000
	sisa := berat % 1000

	biayaKg := kg * 10000
	var biayaSisa int

	if sisa >= 500 {
		biayaSisa = sisa * 5
	} else {
		biayaSisa = sisa * 15
	}

	fmt.Printf("Detail berat: %d kg + %d gr\n", kg, sisa)
	fmt.Printf("Detail biaya: Rp. %d + Rp. %d\n", biayaKg, biayaSisa)

	totalBiaya := biayaKg

	if berat <= 10000 {
		totalBiaya += biayaSisa
	}

	fmt.Printf("Total biaya: Rp. %d\n", totalBiaya)
}
```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 1_1](https://github.com/EdwardAbimas/109082500171_EDWARD-ABIMAS-SURYA-HATTA/raw/main/PERTEMUAN1/ss3.png)
[Pada soal Modul 2C, studi kasus yang dikerjakan adalah pembuatan aplikasi kalkulasi biaya pengiriman parsel untuk PT POS yang didasarkan pada perhitungan berat barang. Alur eksekusi program dimulai dengan meminta input dari pengguna berupa total berat parsel dalam satuan gram. Input tersebut kemudian diproses secara matematis untuk memisahkan nilai berat dalam hitungan kilogram utuh dan sisa berat dalam hitungan gram. Untuk skema perhitungannya, biaya jasa pengiriman dasar dipatok sebesar sepuluh ribu rupiah per kilogramnya. Sementara itu, perhitungan biaya untuk sisa gramnya menggunakan struktur percabangan logika bersyarat. Apabila sisa berat parsel dihitung tidak kurang dari 500 gram, sistem akan membebankan tambahan biaya kirim yang relatif murah yaitu lima rupiah per gramnya. Namun, jika sisa berat tersebut ternyata di bawah 500 gram, tambahan biayanya menjadi lebih mahal yakni sebesar lima belas rupiah per gram. Aturan yang paling krusial dari penentuan tarif ini terletak pada kebijakan diskonnya, di mana biaya untuk sisa berat yang kurang dari satu kilogram tersebut akan sepenuhnya digratiskan apabila akumulasi total berat parsel melebihi sepuluh kilogram. Di bagian akhir, program akan mengkalkulasi seluruh variabel tersebut dan menampilkan rincian dari berat parsel, detail pembagian biaya, hingga total akhir yang wajib dibayarkan.]