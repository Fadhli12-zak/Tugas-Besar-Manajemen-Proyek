package main

import (
	"fmt"
)

const NMAXp int = 100
const NMAXt int = 100

type pegawai struct {
	nama, jobdesk, domisili string
	umur                    int
}
type proyek struct {
	namaP, tugas, deadline, prioritas, status, waktu             string
	anggaran, tanggal, bulan, tahun, tanggalDl, bulanDl, tahunDl int
}
type tabInt [NMAXp]pegawai
type tgsInt [NMAXt]proyek

func main() {
	var B tgsInt
	var A tabInt
	var x, y, w, v int
	var z, menu, tampilan string

	tampilan = "=================================\n \n   Selamat Datang  \n  Aplikasi Manajemen Proyek \n \n================================= "
	menu = "Menu:\n1. Isi Data Pegawai\n2. Isi Data Tugas\n3. Cari Data\n4. Edit Data\n5. Tampilkan Data\n6. Hapus data pegawai\n7. Tampilkan total Anggaran\n0. Exit\nMasukan Nomor Pilihan Anda: "
	fmt.Println(tampilan)
	fmt.Print(menu)
	fmt.Scan(&y)
	for y != 0 {
		if y == 1 {
			isiNama(&A, &x)
		}
		if y == 2 {
			isiTugas(&B, &A, &w, &x, z)
		}
		if y == 3 {
			fmt.Println("****************\n\nCari Data\n\n****************")
			fmt.Println("\n1.Nama\n2.Proyek\n3.jobdesk\ncari berdasarkan:")
			fmt.Scan(&v)

			if v == 1 {
				fmt.Print("Masukan nama yang ingin dicari: ")
				fmt.Scan(&z)
				cariNama(A, x, z)
			}
			if v == 2 {
				fmt.Print("Masukan tugas yang ingin dicari: ")
				fmt.Scan(&z)
				cariTugas(B, w, z)
			}
			if v == 3 {
				fmt.Print("Masukan jobdesk yang anda cari: ")
				fmt.Scan(&z)
				cariJobdesk(A, x, z)
			}
		}
		if y == 4 {
			fmt.Println("\n1. nama\n2. tugas\npilih data yang ingin anda edit: ")
			var pilih int
			fmt.Scan(&pilih)
			if pilih == 1 {
				cetakNama(A, x)
				fmt.Println("masukan nama yang ingin di rubah:")
				fmt.Scan(&z)
				editNama(&A, x, z)
			}
			if pilih == 2 {

				fmt.Println("Masukan tugas yang ingin di edit:")
				srtStatus(&B, w)
				cetakTugas(B, A, w)
				fmt.Scan(&z)
				editTugas(&B, w, z)
			}
		}
		if y == 5 {
			fmt.Println("1. Data pegawai\n2. Data tugas\ntampilkan data:")
			var pilih int
			fmt.Scan(&pilih)
			if pilih == 1 {
				cetakNama(A, x)
			}
			if pilih == 2 {
				fmt.Println("Tampilkan terurut berdasarkan\n1. Prioritas\n2. Tanggal\n3. Status\n4.Deadline\nPilih:")
				fmt.Scan(&pilih)
				if pilih == 1 {
					srtPrioritas(&B, w)
					cetakTugas(B, A, w)
				}
				if pilih == 2 {
					SortTgl(&B, w)
					cetakTugas(B, A, w)
				}
				if pilih == 3 {
					srtStatus(&B, w)
					cetakTugas(B, A, w)
				}
				if pilih == 4 {
					SortDl(&B, w)
					cetakTugas(B, A, w)
				}
			}
		}
		if y == 6 {
			fmt.Println("pilih data yang ingin anda hapus: \n1. nama\n2. tugas")
			var pilih int
			fmt.Scan(&pilih)
			if pilih == 1 {
				cetakNama(A, x)
				fmt.Println("hapus data:")
				fmt.Scan(&z)
				deleteNama(&A, &x, z)
			}
			if pilih == 2 {
				srtStatus(&B, w)
				cetakTugas(B, A, w)
				fmt.Println("hapus data:")
				fmt.Scan(&z)
				deleteTugas(&B, &w, z)
			}
		}
		if y == 7 {
			fmt.Println("Total anggaran dari seluruh tugas adalah: ")
			hitungTotal(&B, w)
		}
		fmt.Print(menu)
		fmt.Scan(&y)
	}
}

func isiTugas(P *tgsInt, Q *tabInt, q, n *int, x string) {
	var A tgsInt
	var find int
	var berhenti bool = true
	if *q >= NMAXt || *n >= NMAXp {
		berhenti = false
	}

	for *q < NMAXt && berhenti && *n < NMAXp {
		fmt.Print("Masukan nama anda (masukan x jika selesai): ")
		fmt.Scan(&x)
		if x == "x" {
			berhenti = false
		} else {
			find = cariNama(*Q, *n, x)
			if find != -1 {
				fmt.Print("Masukan ulang nama anda: ")
				fmt.Scan(&A[*q].namaP)
				fmt.Print("Masukan tugas anda: ")
				fmt.Scan(&A[*q].tugas)
				fmt.Print("Masukan anggaran tugas anda: ")
				fmt.Scan(&A[*q].anggaran)
				fmt.Print("Tipe prioritas tugas:")
				fmt.Scan(&A[*q].prioritas)
				fmt.Print("Masukan tanggal tugas:")
				fmt.Scan(&A[*q].tanggal)
				fmt.Print("Masukan bulan tugas:")
				fmt.Scan(&A[*q].bulan)
				fmt.Print("Masukan tahun tugas:")
				fmt.Scan(&A[*q].tahun)
				fmt.Print("Masukan Tanggal Deadline tugas: ")
				fmt.Scan(&A[*q].tanggalDl)
				fmt.Print("Masukan bulan Deadline tugas: ")
				fmt.Scan(&A[*q].bulanDl)
				fmt.Print("Masukan Tahun Deadline tugas: ")
				fmt.Scan(&A[*q].tahunDl)
				fmt.Print("Masukan status tugas anda(selesai/belum): ")
				fmt.Scan(&A[*q].status)
				P[*q] = A[*q]
				*q++
			}
		}
	}
}

func isiNama(P *tabInt, n *int) {
	var A tabInt
	var berhenti bool = true
	if *n >= NMAXp && berhenti {
		berhenti = false
	}
	for *n < NMAXp && berhenti {
		fmt.Print("Masukan nama anda (masukan x jika selesai): ")
		fmt.Scan(&A[*n].nama)
		if A[*n].nama == "x" {
			fmt.Println("Data selesai ditambahkan")
			berhenti = false
		} else {
			fmt.Print("Masukan jobdesk anda: ")
			fmt.Scan(&A[*n].jobdesk)
			fmt.Print("Masukan domisili anda: ")
			fmt.Scan(&A[*n].domisili)
			fmt.Print("Masukan umur anda: ")
			fmt.Scan(&A[*n].umur)
			P[*n] = A[*n]
			*n++
		}
	}
}

func cetakNama(P tabInt, n int) {
	var i int
	fmt.Println("--------------------------------------------------------------------------")
	fmt.Printf("%-5s %-15s %-15s %-15s %15s\n", "No", "Nama", "Jobdesk", "Domisili", "Umur")
	fmt.Println("--------------------------------------------------------------------------")
	for i = 0; i < n; i++ {
		fmt.Printf("%-5d %-15s %-15s %-15s %15d\n", i+1, P[i].nama, P[i].jobdesk, P[i].domisili, P[i].umur)
	}
}
func cetakTugas(P tgsInt, Q tabInt, q int) {
	var i int
	fmt.Println("--------------------------------------------------------------------------------------------------------")
	fmt.Printf("%-5s %-13s %-20s %-13s %-13s %-13s %-13s %-13s\n", "no", "nama", "tugas", "anggaran", "Prioritas", "waktu", "deadline", "status")
	fmt.Println("--------------------------------------------------------------------------------------------------------")
	for i = 0; i < q; i++ {
		P[i].waktu = fmt.Sprintf("%d-%d-%d", P[i].tanggal, P[i].bulan, P[i].tahun)
		P[i].deadline = fmt.Sprintf("%d-%d-%d", P[i].tanggalDl, P[i].bulanDl, P[i].tahunDl)
		fmt.Printf("%-5d %-13s %-20s %-13d %-13s %-13s %-13s %-13s\n ", i+1, P[i].namaP, P[i].tugas, P[i].anggaran, P[i].prioritas, P[i].waktu, P[i].deadline, P[i].status)
	}
}

func cariNama(P tabInt, n int, x string) int {
	var index int
	index = -1
	var status bool = true
	var i int = 0
	for i < n && status {
		if P[i].nama == x {
			status = false
			index = i
			fmt.Println(P[i].nama, P[i].jobdesk, P[i].domisili, P[i].umur)
		}
		i++
	}
	if index == -1 {
		fmt.Println("Nama tidak ditemukan")
	}
	return index
}
func cariTugas(P tgsInt, q int, m string) int {
	var index int
	index = -1
	var status bool = true
	var i int = 0
	for i < q && status {
		if P[i].tugas == m {
			status = false
			index = i
			fmt.Println(P[i].namaP, P[i].tugas, P[i].anggaran, P[i].prioritas, P[i].waktu, P[i].deadline, P[i].status)
		}
		i++
	}
	if index == -1 {
		fmt.Println("tugas tidak ditemukan")
	}
	return index
}

func cariJobdesk(P tabInt, n int, x string) int {
	var index int
	index = -1
	var status bool = true
	var i int = 0
	for i < n && status {
		if P[i].jobdesk == x {
			status = false
			index = i
			fmt.Println(P[i].nama, P[i].jobdesk, P[i].domisili, P[i].umur)
		}
		i++
	}
	if index == -1 {
		fmt.Println("Jobdes tidak ditemukan")
	}
	return index
}
func editNama(P *tabInt, n int, x string) {
	var finds int
	var datInt int
	var datStr string
	finds = cariNama(*P, n, x)
	if finds != -1 {
		fmt.Print("Msukan nama / (-) jika data tidak ingin diubah: ")
		fmt.Scan(&datStr)
		if datStr != "-" {
			P[finds].nama = datStr
		}
		fmt.Print("Masukan jobdesk /(-) jika data tidak ingin diubah: ")
		fmt.Scan(&datStr)
		if datStr != "-" {
			P[finds].jobdesk = datStr
		}

		fmt.Print("Masukan Domisili /(-) jika data tidak ingin diubah: ")
		fmt.Scan(&datStr)
		if datStr != "-" {
			P[finds].domisili = datStr
		}

		fmt.Print("Masukan umur /(0) jika data tidak ingin diubah: ")
		fmt.Scan(&datInt)
		if datInt != 0 {
			P[finds].umur = datInt
		}
	}
}
func editTugas(P *tgsInt, q int, m string) {
	var ketemu int
	var integer int
	var stringdata string
	ketemu = cariTugas(*P, q, m)
	if ketemu != -1 {
		fmt.Print("Msukan Nama /(-) jika data tidak ingin diubah: ")
		fmt.Scan(&stringdata)
		if stringdata != "-" {
			P[ketemu].namaP = stringdata
		}
		fmt.Print("Msukan tugas /(-) jika data tidak ingin diubah: ")
		fmt.Scan(&stringdata)
		if stringdata != "-" {
			P[ketemu].tugas = stringdata
		}
		fmt.Print("Masukan Prioritas /(-) jika data tidak ingin diubah: ")
		fmt.Scan(&stringdata)
		if stringdata != "-" {
			P[ketemu].prioritas = stringdata
		}

		fmt.Print("Masukan anggaran /(0) jika data tidak ingin diubah: ")
		fmt.Scan(&integer)
		if integer != 0 {
			P[ketemu].anggaran = integer
		}
		fmt.Print("Masukan tanggal  /(0) jika data tidak ingin diubah: ")
		fmt.Scan(&integer)
		if integer != 0 {
			P[ketemu].tanggal = integer
		}
		fmt.Print("Masukan bulan  /(0) jika data tidak ingin diubah: ")
		fmt.Scan(&integer)
		if integer != 0 {
			P[ketemu].bulan = integer
		}
		fmt.Print("Masukan tahun  /(0) jika data tidak ingin diubah: ")
		fmt.Scan(&integer)
		if integer != 0 {
			P[ketemu].tahun = integer
		}
		fmt.Print("Masukan tanggal deadline  /(0) jika data tidak ingin diubah: ")
		fmt.Scan(&integer)
		if integer != 0 {
			P[ketemu].tanggalDl = integer
		}
		fmt.Print("Masukan bulan deadline  /(0) jika data tidak ingin diubah: ")
		fmt.Scan(&integer)
		if integer != 0 {
			P[ketemu].bulanDl = integer
		}
		fmt.Print("Masukan tahun deadline  /(0) jika data tidak ingin diubah: ")
		fmt.Scan(&integer)
		if integer != 0 {
			P[ketemu].tahunDl = integer
		}
		fmt.Print("data telah diubah")
	} else {
		fmt.Println("Data Tidak ditemukan")
	}
}

func deleteNama(P *tabInt, n *int, x string) {
	var ketemu int

	ketemu = cariNama(*P, *n, x)
	if ketemu != -1 {
		for i := ketemu; i < *n-1; i++ {
			P[i] = P[i+1]
		}
		*n = *n - 1
	}
}

func deleteTugas(P *tgsInt, q *int, m string) {
	var found int
	found = cariTugas(*P, *q, m)
	if found == 0 {
		for i := found; i < *q-1; i++ {
			P[i] = P[i+1]
		}
		*q--
	}
}
func srtPrioritas(P *tgsInt, q int) {
	var i, idx int
	var temp proyek
	if q > NMAXt {
		q = NMAXt
	}
	i = 1
	for i < q {
		var j int
		idx = i - 1
		j = i
		for j < q {
			if P[idx].prioritas < P[j].prioritas {
				idx = j
			}
			j++
		}
		temp = P[idx]
		P[idx] = P[i-1]
		P[i-1] = temp
		i++
	}
}
func srtStatus(P *tgsInt, q int) {
	var i, idx int
	var temp proyek
	if q > NMAXt {
		q = NMAXt
	}
	i = 1
	for i < q {
		var j int
		idx = i - 1
		j = i
		for j < q {
			if P[idx].status < P[j].status {
				idx = j
			}
			j++
		}
		temp = P[idx]
		P[idx] = P[i-1]
		P[i-1] = temp
		i++
	}
}

func SortTgl(Q *tgsInt, n int) {
	var i, IDX, j int
	var temp proyek
	for i = 1; i < n; i++ {
		IDX = i - 1
		for j = i; j < n; j++ {
			if Q[IDX].tahun > Q[j].tahun {
				IDX = j
			} else if Q[IDX].tahun == Q[j].tahun {
				if Q[IDX].bulan > Q[j].bulan {
					IDX = j
				} else if Q[IDX].bulan == Q[j].bulan {
					if Q[IDX].tanggal > Q[j].tanggal {
						IDX = j
					}
				}
			}
		}
		if IDX != i-1 {
			temp = Q[i-1]
			Q[i-1] = Q[IDX]
			Q[IDX] = temp
		}
	}
}

func SortDl(Q *tgsInt, n int) {
	var i, IDX, j int
	var temp proyek
	for i = 1; i < n; i++ {
		IDX = i - 1
		for j = i; j < n; j++ {
			if Q[IDX].tahunDl > Q[j].tahunDl {
				IDX = j
			} else if Q[IDX].tahunDl == Q[j].tahunDl {
				if Q[IDX].bulanDl > Q[j].bulanDl {
					IDX = j
				} else if Q[IDX].bulanDl == Q[j].bulanDl {
					if Q[IDX].tanggalDl > Q[j].tanggalDl {
						IDX = j
					}
				}
			}
		}
		if IDX != i-1 {
			temp = Q[i-1]
			Q[i-1] = Q[IDX]
			Q[IDX] = temp
		}
	}
}

func hitungTotal(Q *tgsInt, n int) {
	var i, total int
	for i = 0; i < n; i++ {
		total += Q[i].anggaran
	}
	fmt.Println("Rp.", total)
}
