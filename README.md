<div align="center">
<img src=".github/assets/logo.png" width="150" alt="logo">
</div>

<div align="center">

![Vercel](https://img.shields.io/badge/Vercel-000000?style=flat-square&logo=vercel&logoColor=white)
![MySQL](https://img.shields.io/badge/MySQL-4479A1?style=flat-square&logo=mysql&logoColor=white)
![Vue.js](https://img.shields.io/badge/Vue.js-4FC08D?style=flat-square&logo=vuedotjs&logoColor=white)
![Vite](https://img.shields.io/badge/Vite-646CFF?style=flat-square&logo=vite&logoColor=white)
![Python](https://img.shields.io/badge/Python-3776AB?style=flat-square&logo=python&logoColor=white)
![FastAPI](https://img.shields.io/badge/FastAPI-009688?style=flat-square&logo=fastapi&logoColor=white)

</div>

<p align="center">
Nusantara Kita adalah sebuah API yang menyediakan data wilayah Indonesia. Project ini dirancang untuk memudahkan akses dan penggunaan data geospasial terkait wilayah-wilayah di Indonesia. API ini dapat digunakan untuk berbagai aplikasi yang memerlukan informasi seperti batas wilayah, data administratif, dan lain-lain. Semoga API ini bisa bermanfaat untuk kalian, dan selamat mencoba.
</p>

URL:

- API = [https://api.nusakita.yuefii.site](https://api.nusakita.yuefii.site)
- DOKUMENTASI = [https://nusakita.yuefii.site](https://nusakita.yuefii.site)

ENDPOINT:

- `/v2/provinsi`
- `/v2/kab-kota`
- `/v2/kecamatan`
- `/v2/desa-kel`
- `/v2/{kodeProvinsi}/kab-kota`
- `/v2/{kodeKabKota}/kecamatan`
- `/v2/{kodeKecamatan}/desa-kel`

API ini dihosting menggunakan `VERCEL` jadi akan terbatas untuk consume api nya jadi lebih baik kalian hosting sendiri project ini diserver kalian.

## Fitur

API ini berisi seluruh data wilayah indonesia beserta lokasinya mulai dari:

- Provinsi
- Kabupaten/Kota
- Kecamatan
- Desa/Kelurahan

## Instalasi

Langkah-langkah cara install project ini dilokal :

Git clone API nya:

```bash
git clone --filter=blob:none --no-checkout https://github.com/Yuefii/NusantaraKita.git
cd NusantaraKita
git sparse-checkout set api
git checkout
```

Masuk ke directori API:

```bash
cd api
```

Setup Environmentnya:

```bash
cp env.example .env
```

Menjalankan Projectnya menggunakan `pip`

- create virtual environment:
  ```bash
  python3 -m venv venv
  ```
- activate virtual environment:
  ```bash
  source venv/bin/activate
  ```
- install dependencies:
  ```bash
  pip install -r requirements.txt
  ```
- Running project:
  ```
  uvicorn generate:app --reload
  ```

## ⚙️ License

Project ini dilisensikan di bawah Lisensi MIT - lihat [LICENSE](/LICENSE) untuk detailnya.
