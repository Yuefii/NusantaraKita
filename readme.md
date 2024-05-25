# API NUSANTARA-KITA

API ini menyediakan data lengkap mengenai wilayah-wilayah di Indonesia yang telah dideploy menggunakan `VERCEL`, semoga API ini bisa bermanfaat untuk kalian, dan selamat mencoba.

Demo: [https://nusantara-kita.yuefii.my.id](https://nusantara-kita.yuefii.my.id)

- `/api/nusantara/provinces`
- `/api/nusantara/{provinces_code}/regencies`
- `/api/nusantara/{regency_code}/districts`
- `/api/nusantara/{district_code}/villages`

`noted` Karna Api ini dihosting menggunakan hosting yang free jadi akan terbatas untuk consume api nya jadi lebih baik kalian hosting sendiri project ini.

`Fork` project ini ke repostory kalian lalu kemudian deploy API nya.

## Fitur

API ini berisi seluruh data wilayah indonesia mulai dari:

- Provinsi
- Kabupaten/Kota
- Kecamatan
- Desa

## Dokumentasi Project

Untuk dokumentasi lengkap dan ingin mencoba apinya kalian bisa cek [yuefii.my.id/nusantara-kita](https://yuefii.my.id/nusantara-kita)

## Database API

Data yang kami gunakan didasarkan pada data resmi dari pemerintah dan disimpan dalam file CSV yang dipisahkan berdasarkan tingkatan dalam direktori [data](./data/).

## Instalasi

Langkah-langkah cara install project ini dilokal :

Git clone

```
git clone https://github.com/Yuefii/api-nusantara-kita.git
```

Masuk ke directori project

```
cd api-nusantara-kita
```

Menjalankan Projectnya menggunakan `YARN`

- install dependencies
  ```
  yarn install
  ```
- jalankan projectnya
  ```
  yarn dev
  ```

#### Jika kalian ingin menggunakan `Docker`

Untuk Menjalankannya di `Development`

```
docker-compose -f docker-compose.yaml up --build
```

`noted`
Untuk menggunakan `Docker` pada development kalian harus `yarn install` atau `npm install` terlebih dahulu, untuk menginstall `node_modules` tidak seperti `Production` yang bisa langsung berjalan.

Untuk Menjalankannya di `Production`

```
docker-compose -f docker-compose.prod.yaml up --build -d
```

Jika kalian ingin lebih mudah lagi bisa menggunakan `Make`

Caranya kalian ketika gunakan command ini :

- `make dev`
- `make dev-down`
- `make prod`
- `make prod-down`

Untuk selengkapnya kalian bisa cek di pada [Makefile](./Makefile) disana akan ada komentar yang menjelaskan kegunaannya.

`noted`
Untuk menggunakan `Make` kalian harus menginstall nya terlebih dahulu.

## Endpoint

#### 1. Mengambil Data Provinsi

```
GET http://localhost:8080/api/nusantara/provinces
```

#### 2. Mengambil Data Kabupaten/Kota berdasarkan kode Provinsi

```
GET http://localhost:8080/api/nusantara/{provinces_code}/regencies
```

#### 3. Mengambil Data Kecamatan berdasarkan kode Kabupaten/kota

```
GET http://localhost:8080/api/nusantara/{regency_code}/districts
```

#### 4. Mengambil Data Desa berdasarkan kode Kecamatan

```
GET http://localhost:8080/api/nusantara/{district_code}/villages
```
