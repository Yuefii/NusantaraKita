# API NUSANTARA-KITA

API ini menyediakan data lengkap mengenai wilayah-wilayah di Indonesia yang telah dideploy menggunakan `VERCEL`, semoga API ini bisa bermanfaat untuk kalian, dan selamat mencoba.

Demo: [https://nusantara-kita.yuefii.my.id](https://nusantara-kita.yuefii.my.id)

- `/api/nusantara/provinces`
- `/api/nusantara/{provinces_code}/regencies`
- `/api/nusantara/{regency_code}/districts`
- `/api/nusantara/{district_code}/villages`

## Fitur

Api ini berisi seluruh data wilayah indonesia mulai dari:

- Provinsi
- Kabupaten/Kota
- Kecamatan
- Desa

## Dokumentasi Project

Untuk dokumentasi lengkap dan ingin mencoba apinya kalian bisa cek [yuefii.my.id/nusantara-kita](https://yuefii.my.id/nusantara-kita)

## Database API

Data yang kami gunakan didasarkan pada data resmi dari pemerintah dan disimpan dalam file CSV yang dipisahkan berdasarkan tingkatan dalam direktori [data](./data/).

## Instalasi

Langkah-langkah Cara install project ini dilokal

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
