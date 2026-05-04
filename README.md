# Tucil 3 Strategi Algoritma

Program ini adalah penyelesai (solver) untuk permainan Ice Sliding Puzzle dengan antarmuka grafis (GUI) berbasis Wails.
Program membaca peta dari file teks, kemudian mencari jalur penyelesaian menggunakan salah satu dari lima algoritma pencarian (UCS, BFS, GBFS, A*, IDA*) dengan tiga pilihan heuristik (Manhattan, Euclidean, Chebyshev), lalu menampilkan visualisasi langkah demi langkah dari solusi yang ditemukan.

## Requirement Program
- Go (versi 1.25.5 dari `go.mod`)
- Wails CLI v2 (`go install github.com/wailsapp/wails/v2/cmd/wails@latest`)
- Xcode Command Line Tools (macOS) atau WebView2 Runtime (Windows)

Verifikasi instalasi dengan `wails doctor`.

## Compile and Run
Dari root project, jalankan:

- Linux/macOS:

	```bash
	wails build
	```

- Windows:

	```powershell
	wails build
	```

Output binary akan tersimpan di `build/bin/Solver.app` (macOS) atau `build/bin/Solver.exe` (Windows).

## Cara Menjalankan dan Menggunakan Program
Buka binary hasil kompilasi:

- macOS:

	```bash
	open build/bin/Solver.app
	```

- Linux:

	```bash
	./build/bin/Solver
	```

- Windows:

	```powershell
	./build/bin/Solver.exe
	```

Atau jalankan dalam mode pengembangan (dengan hot-reload):

```bash
wails dev
```

Penggunaan:
- Klik **Browse...** untuk memilih file peta `.txt` (contoh: `test/input.txt`), lalu klik **Load**.
- Pilih algoritma (`UCS` / `BFS` / `GBFS` / `A*` / `IDA*`) dan heuristik (`H1` Manhattan / `H2` Euclidean / `H3` Chebyshev).
- Klik **Solve** untuk menjalankan pencarian; status menampilkan path, cost, jumlah iterasi, dan waktu eksekusi.
- Gunakan tombol **< Prev**, **Next >**, atau **Jump** untuk memutar ulang langkah solusi pada papan.
- Klik **Save** untuk menyimpan solusi ke file `test/output/<nama>.txt`.

Format file peta:
- Baris pertama: dimensi `<baris> <kolom>`
- Baris berikutnya: grid peta (`X` tembok, `*` es, `Z` posisi awal, `O` tujuan, `L` lava, `0`-`9` angka berurutan)
- Lalu matriks cost per cell, dengan dimensi yang sama dengan grid

## Author
|        Nama         |    NIM   |
|---------------------|----------|
| Nathaniel Christian | 13524122 |
| Ahmad Fauzan Putra  | 13524141 |