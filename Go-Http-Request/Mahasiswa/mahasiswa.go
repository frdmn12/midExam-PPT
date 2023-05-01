package Mahasiswa

type dataMahasiswa struct {
	Nama string;
	Nim int;
	Alamat string;
}

var ListMahasiswa = []dataMahasiswa{
	{Nama: "Gusti", Nim: 111, Alamat: "Jl. Sudirman No.1"},
	{Nama: "Pangestu", Nim: 222, Alamat: "Jl. Gatot Subroto No.10"},
}