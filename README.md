# Nusantara Kita

<p align="center"><img src=".github/assets/logo.png" width="150" alt="logo"></p>

Nusantara Kita adalah sebuah API yang menyediakan data wilayah Indonesia. Proyek ini dirancang untuk memudahkan akses dan penggunaan data geospasial terkait wilayah-wilayah di Indonesia. API ini dapat digunakan untuk berbagai aplikasi yang memerlukan informasi seperti batas wilayah, data administratif, dan lain-lain.

### 🧞 Running Project

#### bagaimana menjalankan projectnya dilokal?

- Setup Enviromentnya:
  ```bash
  mv env.example .env
  ```
  lalu kemudian kamu isi sesuai dengan database kalian.

- Install Requirements:
  ```bash
  pip install -r requirements.txt
  ```
- Menjalankan Jupyter Notebook nya:
  ```bash
   jupyter notebook
  ```
  jalankan semua code yang ada di notebook untuk merubah dataset csv menjadi sql yang akan langsung diinsert ke database kalian.
- Menjalankan API nya:
  ```bash
   python main.py
  ```

#### Ingin lebih mudah menjalankan databasenya, kamu bisa gunakan docker compose yang kami sediakan?
   - Menjalankan Docker
     ```bash
      docker compose up -d
      ```
   - Menghentikan Docker
     ```bash
      docker compose down
      ```
## 📖 Documentation
Untuk Dokumentasi cara menggunakan API nya, sudah kami sediakan ketika kamu sudah menjalankan project ini, langsung ajh buka endpoint url nya di browser kalian.

## 🚀 Contribute

Untuk berkontribusi di project ini, kamu bisa cek [disini](/CONTRIBUTE.md). Ayo berkontribusi bersama kami untuk membuat project ini menjadi lebih baik lagi.
  
## ⚙️ License

Project ini dilisensikan di bawah Lisensi MIT - lihat [LICENSE](/LICENSE) untuk detailnya.