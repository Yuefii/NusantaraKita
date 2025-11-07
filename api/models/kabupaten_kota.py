from typing import TypedDict, List
from .pagination import PaginationMeta


class KabupatenKota(TypedDict):
    """
    Struktur data untuk response daftar kabupaten/kota.

    Attributes:
        kode (str): Kode unik/primary key kabupaten/kota.
        nama (str): Nama kabupaten/kota.
        lat (float): Latitude dari lokasi kabupaten/kota.
        lng (float): Longitude dari lokasi kabupaten/kota.
        kode_provinsi (str): Kode provinsi tempat kabupaten/kota berada.
        geojson_url (str): URL file GeoJSON dari kabupaten/kota.
    """

    kode: str
    nama: str
    lat: float
    lng: float
    kode_provinsi: str
    geojson_url: str


class PaginatedKabupatenKotaResponse(TypedDict):
    """
    Struktur data untuk response daftar kabupaten/kota dengan informasi pagination.

    Attributes:
        pagination (PaginationMeta): Informasi metadata pagination.
        data (List[KabupatenKota]): Daftar kabupaten/kota.
    """

    pagination: PaginationMeta
    data: List[KabupatenKota]


class KabupatenKotaListResponse(TypedDict):
    """
    Struktur data untuk response daftar kabupaten/kota tanpa informasi pagination.

    Attributes:
        data (List[KabupatenKota]): Daftar kabupaten/kota.
    """

    data: List[KabupatenKota]
