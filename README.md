# Nusantara Kita

![GitHub stars](https://img.shields.io/github/stars/Yuefii/NusantaraKita.svg?style=social)
![GitHub forks](https://img.shields.io/github/forks/Yuefii/NusantaraKita.svg?style=social)

![Vercel](https://img.shields.io/badge/Vercel-000000?style=flat-square&logo=vercel&logoColor=white)
![PostgreSQL](https://img.shields.io/badge/PostgreSQL-336791?style=flat-square&logo=postgresql&logoColor=white)
![React](https://img.shields.io/badge/React-61DAFB?style=flat-square&logo=react&logoColor=white)
![TypeScript](https://img.shields.io/badge/TypeScript-3178C6?style=flat-square&logo=typescript&logoColor=white)
![Vite](https://img.shields.io/badge/Vite-646CFF?style=flat-square&logo=vite&logoColor=white)
![Go](https://img.shields.io/badge/Go-00ADD8?style=flat-square&logo=go&logoColor=white)

<img src=".github/assets/logo.png" width="200" alt="logo">

Nusantara Kita adalah sebuah API yang menyediakan data wilayah Indonesia. Project ini dirancang untuk memudahkan akses dan penggunaan data geospasial terkait wilayah-wilayah di Indonesia. API ini dapat digunakan untuk berbagai aplikasi yang memerlukan informasi seperti batas wilayah, data administratif, dan lain-lain. Semoga API ini bisa bermanfaat untuk kalian, dan selamat mencoba.

URL:

- API = [https://api.nusantarakita.web.id](https://api.nusantarakita.web.id)
- DOKUMENTASI = [https://nusantarakita.web.id](https://nusantarakita.web.id)

ENDPOINT:

- **List Wilayah:**
  - `/v2/provinsi`
  - `/v2/kab-kota`
  - `/v2/kecamatan`
  - `/v2/desa-kel`

- **Filter Berdasarkan Wilayah Induk:**
  - `/v2/provinsi/{kode}/kab-kota` (atau `/v2/{kode_provinsi}/kab-kota`)
  - `/v2/kab-kota/{kode}/kecamatan` (atau `/v2/{kode_kabupaten_kota}/kecamatan`)
  - `/v2/kecamatan/{kode}/desa-kel` (atau `/v2/{kode_kecamatan}/desa-kel`)

- **Detail Wilayah (Get by Kode):**
  - `/v2/provinsi/{kode}`
  - `/v2/kab-kota/{kode}`
  - `/v2/kecamatan/{kode}`
  - `/v2/desa-kel/{kode}`

- **Pencarian / Autocomplete:**
  - `/v2/search?q={keyword}&tipe={provinsi|kab-kota|kecamatan|desa-kel}&limit=10&halaman=1`

- **Reverse Geocoding:**
  - `/v2/reverse?lat={latitude}&lng={longitude}`

- **System:**
  - `/health`

API ini sudah dihosting di domain sendiri. Namun jika kamu ingin menjalankannya di server sendiri, kamu tetap bisa melakukan self-hosting sesuai panduan instalasi di bawah.

## Fitur

API ini berisi seluruh data wilayah indonesia beserta lokasinya mulai dari:

- Provinsi (38 Provinsi)
- Kabupaten/Kota (514 Kab/Kota)
- Kecamatan (7.277 Kecamatan)
- Desa/Kelurahan (82.978 Desa/Kelurahan)
- **Search & Autocomplete:** Pencarian pintar nama wilayah dan kode pos dengan skor relevansi
- **Detail by Kode:** Pengambilan 1 wilayah langsung berdasarkan kode unik
- **Reverse Geocoding:** Konversi koordinat GPS (`lat`, `lng`) menjadi hierarki wilayah Indonesia terdekat
- **GeoJSON Boundaries:** Batas wilayah geospasial untuk visualisasi peta

## Contoh Penggunaan API

### 1. Pencarian & Autocomplete
```bash
# Cari wilayah umum (Provinsi, Kab/Kota, Kecamatan, Desa)
curl "https://api.nusantarakita.web.id/v2/search?q=bandung"

# Filter berdasarkan jenis wilayah (tipe: provinsi | kab-kota | kecamatan | desa-kel)
curl "https://api.nusantarakita.web.id/v2/search?q=bandung&tipe=kab-kota"

# Cari berdasarkan 5 digit Kode Pos
curl "https://api.nusantarakita.web.id/v2/search?q=40151"
```

### 2. Detail Wilayah Berdasarkan Kode
```bash
# Detail Provinsi
curl "https://api.nusantarakita.web.id/v2/provinsi/32"

# Detail Kabupaten/Kota
curl "https://api.nusantarakita.web.id/v2/kab-kota/32.73"

# Detail Kecamatan
curl "https://api.nusantarakita.web.id/v2/kecamatan/32.73.01"

# Detail Desa/Kelurahan
curl "https://api.nusantarakita.web.id/v2/desa-kel/32.73.01.1001"
```

### 3. Reverse Geocoding (Koordinat GPS ke Wilayah)
```bash
# Mengembalikan Provinsi, Kab/Kota, Kecamatan, Desa/Kelurahan, Kode Pos, & Jarak dalam meter
curl "https://api.nusantarakita.web.id/v2/reverse?lat=-6.874112&lng=107.589123"
```

### 4. Health Check
```bash
curl "https://api.nusantarakita.web.id/health"
```

## Instalasi

Langkah-langkah cara install project ini dilokal :

**Prasyarat:** Pastikan kamu sudah menginstall [Go](https://go.dev/doc/install) (minimal versi 1.22) di komputermu.

Git clone API nya:

```bash
git clone --filter=blob:none --no-checkout https://github.com/Yuefii/NusantaraKita.git && cd NusantaraKita && git sparse-checkout set api data && git checkout
```

Setup Environmentnya:

```bash
cd api
```

```bash
cp .env.example .env
```
*(Jangan lupa untuk mengisi `DATABASE_URL` dengan koneksi PostgreSQL kamu di file `.env`)*

Menjalankan Projectnya menggunakan `Go`

- Unduh dependencies:
  ```bash
  go mod download
  ```
- Menjalankan server:
  ```bash
  go run cmd/server/main.go
  ```
- Build project (opsional):
  ```bash
  go build -o server ./cmd/server
  ```

Menjalankan Projectnya menggunakan `docker`

- setelah clone project kamu bisa langsung menjalankan projectnya dengan mudah menggunakan perintah:
  ```bash
  docker-compose up -d
  ```
  atau gunakan perintah ini jika ada perubahan kode dan kamu ingin melakukan *build* ulang image API-nya:
  ```bash
  docker-compose up -d --build
  ```


## References

Dataset yang digunakan dalam project ini berasal dari:

- [https://github.com/cahyadsn/wilayah_boundaries](https://github.com/cahyadsn/wilayah_boundaries)
- [https://github.com/cahyadsn/wilayah_kodepos](https://github.com/cahyadsn/wilayah_kodepos)

Kami sangat menghargai pembuat dataset ini. Tanpa adanya data tersebut, project ini tidak akan bisa dikembangkan.

## Berkontribusi & Kode Etik

Baca [panduan kontribusi](./CONTRIBUTING.md) kami untuk mempelajari cara berkontribusi pada proyek kami.

Pastikan untuk mematuhi [kode etik](./CODE_OF_CONDUCT.md) kami.

## License

Project ini dilisensikan di bawah Lisensi MIT - lihat [LICENSE](/LICENSE) untuk detailnya.
