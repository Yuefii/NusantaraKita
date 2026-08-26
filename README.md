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

- API = [https://api-nusantarakita.vercel.app](https://api-nusantarakita.vercel.app)
- DOKUMENTASI = [https://nusantarakita.vercel.app](https://nusantarakita.vercel.app)

ENDPOINT:

- `/v2/provinsi`
- `/v2/kab-kota`
- `/v2/kecamatan`
- `/v2/desa-kel`
- `/v2/{kode_provinsi}/kab-kota`
- `/v2/{kode_kabupaten_kota}/kecamatan`
- `/v2/{kode_kecamatan}/desa-kel`

API ini dihosting menggunakan `VERCEL` jadi akan terbatas untuk consume api nya jadi lebih baik kalian hosting sendiri project ini diserver kalian.

## Fitur

API ini berisi seluruh data wilayah indonesia beserta lokasinya mulai dari:

- Provinsi
- Kabupaten/Kota
- Kecamatan
- Desa/Kelurahan

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
