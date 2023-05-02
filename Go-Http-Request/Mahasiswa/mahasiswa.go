package Mahasiswa

type dataMahasiswa struct {
	Nama string;
	Nim int;
	Alamat string;
}

var ListMahasiswa = []dataMahasiswa{
	{Nama: "Gusti", Nim: 1234, Alamat: "Malang"},
	{Nama: "Pangestu", Nim: 321, Alamat: "Kabupaten Malang"},
}