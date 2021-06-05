/* Nama : Afrizal Syahruluddin Yusuf
   NIM	: 1301194288
   Nama	: Habib Alfarabi
   NIM	: 1301194170
*/
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"time"
)
type pegawai struct {
	kode_pegawai string
	nama         string
	golongan     int
	umur         int
	alamat       string
	reward       string
}
type absensi struct {
	kode_pegawai string
	jam_masuk    string
	jam_keluar   string
	jam_lembur   int
	pulang_cepat bool
}

var datapegawai []pegawai
var histori_absensi []absensi

func tambahpegawai() {
	var kode_pegawai, nama, alamat, input string
	var golongan, umur int
	var err error
	for true {
		scanner := bufio.NewScanner(os.Stdin)
		fmt.Println("Masukkan kode pegawai")
		fmt.Scanln(&kode_pegawai)
		fmt.Println("Masukkan nama pegawai")
		scanner.Scan()
		nama = scanner.Text()
		fmt.Println("Masukkan golongan")
		fmt.Scanln(&input)
		golongan, err = strconv.Atoi(input)
		if err != nil {
			fmt.Println("Masukan angka dengan benar")
			continue
		}
		fmt.Println("Masukkan umur")
		fmt.Scanln(&input)
		umur, err = strconv.Atoi(input)
		if err != nil {
			fmt.Println("Masukan angka dengan benar")
			continue
		}
		fmt.Println("Masukkan alamat")
		scanner.Scan()
		alamat = scanner.Text()
		break
	}
	data := pegawai{kode_pegawai: kode_pegawai, nama: nama, golongan: golongan, umur: umur, alamat: alamat, reward: ""}
	datapegawai = append(datapegawai, data)
}

func carikode(kode string) int {
	for i := 0; i < len(datapegawai); i++ {
		if kode == datapegawai[i].kode_pegawai {
			return i
		}
	}
	return -1
}
func caripegawai() {
	var kode string
	var index int
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Masukkan kode pegawai")
	scanner.Scan()
	kode = scanner.Text()
	index = carikode(kode)
	if index >= 0 {
		fmt.Println("Kode pegawai:", datapegawai[index].kode_pegawai)
		fmt.Println("Nama pegawai:", datapegawai[index].nama)
		fmt.Println("Golongan:", datapegawai[index].golongan)
		fmt.Println("Umur:", datapegawai[index].umur)
		fmt.Println("Alamat:", datapegawai[index].alamat)
		fmt.Println("Reward:", datapegawai[index].reward)
	} else {
		fmt.Println("Pegawai tidak temukan")
	}
}
func hitungjamkerja(masuk, keluar string) int {

	var datetime1, datetime2 time.Time
	datetime1, _ = time.Parse("15.04", masuk)
	datetime2, _ = time.Parse("15.04", keluar)
	return int(datetime2.Sub(datetime1).Hours())
}
func berireward(kode_pegawai, masuk, keluar string) {
	var total_waktu int
	for i := 0; i < len(histori_absensi); i++ {
		if kode_pegawai == histori_absensi[i].kode_pegawai {
			total_waktu += histori_absensi[i].jam_lembur
		}
	}
	if total_waktu >= 10 {
		datapegawai[carikode(kode_pegawai)].reward = "Pekerja keras"
		fmt.Println("Selamat, ", kode_pegawai, " reward menjadi pekerja keras")
	} else if masuk == "08.00" && keluar == "16.00" {
		datapegawai[carikode(kode_pegawai)].reward = "Menghargai waktu"
		fmt.Println("Selamat, ", kode_pegawai, " reward menjadi menghargai waktu")
	}
}
func inputAbsensi() {
	var kode_pegawai, jam_masuk, jam_keluar string
	var jam_lembur, jam_kerja int
	var pulang_cepat bool
	fmt.Println("Masukkan kode pegawai")
	fmt.Scanln(&kode_pegawai)
	fmt.Println("Jam masuk")
	fmt.Scanln(&jam_masuk)
	fmt.Println("Jam keluar")
	fmt.Scanln(&jam_keluar)
	jam_kerja = hitungjamkerja(jam_masuk, jam_keluar)
	if jam_kerja >= 9 {
		jam_lembur = jam_kerja - 9
		pulang_cepat = false
	} else {
		jam_lembur = 0
		pulang_cepat = true
	}
	data := absensi{kode_pegawai: kode_pegawai, jam_masuk: jam_masuk, jam_keluar: jam_keluar, jam_lembur: jam_lembur, pulang_cepat: pulang_cepat}
	histori_absensi = append(histori_absensi, data)
	berireward(kode_pegawai, jam_masuk, jam_keluar)
}
func pekerjakeras() {
	for i := 0; i < len(datapegawai); i++ {
		if "Pekerja keras" == datapegawai[i].reward {
			fmt.Println(datapegawai[i].nama)
		}
	}
}
func tampilhistori() {
	for i := 0; i < len(histori_absensi); i++ {
		fmt.Println(histori_absensi[i].kode_pegawai)
		fmt.Println(histori_absensi[i].jam_masuk)
		fmt.Println(histori_absensi[i].jam_keluar)
		fmt.Println(histori_absensi[i].jam_lembur)
		fmt.Println(histori_absensi[i].pulang_cepat)
	}
}
func sorting(databaru []pegawai) []pegawai {
	for i := len(databaru); i > 0; i-- {
		for j := 1; j < i; j++ {
			if databaru[j-1].nama > databaru[j].nama {
				c := databaru[j]
				databaru[j] = databaru[j-1]
				databaru[j-1] = c
			}
		}
	}
	return databaru
}
func terurut() {
	dataterurut := sorting(datapegawai)
	for i := 0; i < len(dataterurut); i++ {
		fmt.Println("==============================")
		fmt.Println(dataterurut[i].kode_pegawai)
		fmt.Println(dataterurut[i].nama)
		fmt.Println(dataterurut[i].golongan)
		fmt.Println(dataterurut[i].umur)
		fmt.Println(dataterurut[i].alamat)
		fmt.Println(dataterurut[i].reward)
	}
}
func exit() {
	fmt.Println("===========================================================")
	fmt.Println("			TERIMA KASIH	")
	fmt.Println("===========================================================")
}
func mainmenu() {
	var pilih int
	for true {
		fmt.Println("===========================================================")
		fmt.Println("	SELAMAT DATANG DI APLIKASI ABSENSI PEGAWAI	")
		fmt.Println("===========================================================")
		fmt.Println("1. Tambah Pegawai")
		fmt.Println("2. Cari Pegawai")
		fmt.Println("3. Absen Pegawai")
		fmt.Println("4. Histori Pegawai")
		fmt.Println("5. Tampil Pekerja Keras")
		fmt.Println("6. Tampil Pegawai")
		fmt.Println("7. Exit")
		fmt.Println("")
		fmt.Println("===========================================================")
		fmt.Println("Pilihan Anda")
		fmt.Scanln(&pilih)
		switch pilih {
		case 1:
			tambahpegawai()
		case 2:
			caripegawai()
		case 3:
			inputAbsensi()
		case 4:
			tampilhistori()
		case 5:
			pekerjakeras()
		case 6:
			terurut()
		case 7:
			exit()
			break
		default:
			fmt.Println("Silahkan pilih 1 sampai 7")
		}
		break
	}
}
func main() {
	mainmenu()
}