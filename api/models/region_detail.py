from typing import TypedDict, List, Optional


class RegionDetail(TypedDict):
    """
    Struktur data untuk detail wilayah berdasarkan kode.

    Attributes:
        kode (str): Kode wilayah.
        desa_kelurahan (Optional[str]): Nama desa/kelurahan (jika ada).
        kecamatan (Optional[str]): Nama kecamatan (jika ada).
        kabupaten_kota (Optional[str]): Nama kabupaten/kota (jika ada).
        provinsi (str): Nama provinsi.
    """

    kode: str
    desa_kelurahan: Optional[str]
    kecamatan: Optional[str]
    kabupaten_kota: Optional[str]
    provinsi: str


class RegionDetailResponse(TypedDict):
    """
    Struktur data untuk response detail wilayah.

    Attributes:
        data (List[RegionDetail]): Daftar detail wilayah.
    """

    data: List[RegionDetail]
