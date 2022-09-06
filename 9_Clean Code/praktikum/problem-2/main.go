package main

type kendaraan struct {
	totalRoda       int
	kecepatanPerJam int
}

type mobil struct {
	kendaraan
}

func (m *mobil) tambahKecepatan(kecepatanBaru int) {
	m.kecepatanPerJam += kecepatanBaru
}

func (m *mobil) berjalan() {
	m.tambahKecepatan(10)
}

func main() {
	// define object mobilCepat dengan totalRoda mobil 4
	mobilCepat := mobil{kendaraan{totalRoda: 4}}

	for i := 0; i < 3; i++ {
		mobilCepat.berjalan()
	}

	// define object mobilLamban dengan totalRoda mobil 4
	mobilLamban := mobil{kendaraan{totalRoda: 4}}
	mobilLamban.berjalan()
}
