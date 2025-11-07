from typing import TypedDict, List
from .pagination import PaginationMeta


class Kecamatan(TypedDict):
    """
    Struktur data untuk response daftar kecamatan.

    Attributes:
        kode (str): Kode unik/primary key kecamatan.
        nama (str): Nama kecamatan.
        lat (float): Latitude dari lokasi kecamatan.
        lng (float): Longitude dari lokasi kecamatan.
        kode_kabupaten_kota (str): Kode kabupaten/kota tempat kecamatan berada.
        geojson_url (str): URL file GeoJSON dari kecamatan.
    """

    kode: str
    nama: str
    lat: float
    lng: float
    kode_kabupaten_kota: str
    geojson_url: str


class PaginatedKecamatanResponse(TypedDict):
    """
    Struktur data untuk response daftar kecamatan dengan informasi pagination.

    Attributes:
        pagination (PaginationMeta): Informasi metadata pagination.
        data (List[Kecamatan]): Daftar kecamatan.
    """

    pagination: PaginationMeta
    data: List[Kecamatan]


class KecamatanListResponse(TypedDict):
    """
    Struktur data untuk response daftar kecamatan tanpa informasi pagination.

    Attributes:
        data (List[Kecamatan]): Daftar kecamatan.
    """

    data: List[Kecamatan]
