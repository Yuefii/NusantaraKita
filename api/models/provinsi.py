from typing import TypedDict, List
from .pagination import PaginationMeta


class Provinsi(TypedDict):
    """
    Struktur data untuk response daftar provinsi.

    Attributes:
        kode (str): Kode unik/primary key provinsi.
        nama (str): Nama provinsi.
        lat (float): Latitude dari lokasi provinsi.
        lng (float): Longitude dari lokasi provinsi.
        geojson_url (str): URL file GeoJSON dari provinsi.
    """

    kode: str
    nama: str
    lat: float
    lng: float
    geojson_url: str


class PaginatedProvinsiResponse(TypedDict):
    """
    Struktur data untuk response daftar provinsi dengan informasi pagination.

    Attributes:
        pagination (PaginationMeta): Informasi metadata pagination.
        data (List[Provinsi]): Daftar Provinsi.
    """

    pagination: PaginationMeta
    data: List[Provinsi]


class ProvinsiListResponse(TypedDict):
    """
    Struktur data untuk response daftar provinsi tanpa informasi pagination.

    Attributes:
        data (List[Provinsi]): Daftar Provinsi.
    """

    data: List[Provinsi]
