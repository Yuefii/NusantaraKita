const MainFeatures = () => {
  const fiturUtamaData = [
    'Data batas wilayah administratif Indonesia, seperti provinsi, kabupaten/kota, dan kecamatan.',
    'Informasi geospasial terkait koordinat wilayah yang dapat diintegrasikan dengan aplikasi peta.',
    'Menyediakan endpoint untuk mendapatkan data berdasarkan jenis wilayah atau nama wilayah tertentu.',
    'Dukungan untuk format data JSON yang mudah digunakan di berbagai platform dan bahasa pemrograman.',
    'Akses ke data yang selalu diperbarui untuk memastikan akurasi informasi yang diberikan.',
    'Mendukung fitur GeoJSON wilayah dari tingkat provinsi hingga desa/kelurahan untuk visualisasi peta interaktif.',
  ];

  const kelebihanApiData = [
    'API yang ringan dan cepat, ideal untuk aplikasi yang membutuhkan akses cepat ke data wilayah.',
    'Dokumentasi API yang jelas dan mudah diikuti untuk pengembang.',
    'Didukung oleh sistem yang scalable, cocok untuk aplikasi besar dengan kebutuhan data tinggi.',
    'Penggunaan yang mudah di berbagai platform, termasuk mobile dan web applications.',
  ];

  const informasiText =
    'API Nusantara Kita dapat digunakan untuk berbagai kebutuhan aplikasi yang memerlukan data terkait wilayah Indonesia. Pengguna dapat mengakses data melalui berbagai endpoint yang telah disediakan, baik untuk mendapatkan informasi batas wilayah atau data administratif lainnya.';

  return (
    <section className="flex max-w-5xl flex-col gap-6 p-6">
      <div className="rounded-lg bg-white p-6 shadow-md">
        <h3 className="mb-4 text-2xl font-semibold text-gray-800">
          Fitur Utama
        </h3>
        <ul className="list-outside list-disc space-y-2 pl-5 text-gray-700">
          {fiturUtamaData.map((fitur, index) => (
            <li key={`fitur-${index}`}>{fitur}</li>
          ))}
        </ul>
      </div>

      <div className="rounded-lg bg-white p-6 shadow-md">
        <h3 className="mb-4 text-2xl font-semibold text-gray-800">
          Kelebihan API
        </h3>
        <ul className="list-outside list-disc space-y-2 pl-5 text-gray-700">
          {kelebihanApiData.map((fitur, index) => (
            <li key={`fitur-${index}`}>{fitur}</li>
          ))}
        </ul>
      </div>

      <div className="rounded-lg bg-white p-6 shadow-md">
        <h3 className="mb-4 text-2xl font-semibold text-gray-800">Informasi</h3>
        <p className="leading-relaxed text-gray-700">{informasiText}</p>
      </div>
    </section>
  );
};

export default MainFeatures;
