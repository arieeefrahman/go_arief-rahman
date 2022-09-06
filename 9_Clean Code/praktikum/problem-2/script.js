class Kendaraan {
    constructor(totalRoda, kecepetanPerJam) {
        this.totalRoda = 0;
        this.kecepatanPerJam = 0;        
    }
}

class Mobil extends Kendaraan {
    tambahKecepatan(kecepatanBaru) {
        this.kecepatanPerJam += kecepatanBaru;
    }

    berjalan() {
        this.tambahKecepatan(10);
    }
}

let mobilCepat = new Mobil();

for (let i = 0; i < 3; i++) {
    mobilCepat.berjalan();
}

let mobilLamban = new Mobil();
mobilLamban.berjalan();