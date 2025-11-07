from typing import TypedDict, List
from .pagination import PaginationMeta


class DesaKelurahan(TypedDict):
    """
    Struktur data untuk response daftar desa/kelurahan.

    Attributes:
        kode (str): Kode unik/primary key desa/kelurahan.
        nama (str): Nama desa/kelurahan.
        lat (float): Latitude dari lokasi desa/kelurahan.
        lng (float): Longitude dari lokasi desa/kelurahan.
        kode_kecamatan (str): Kode kecamatan tempat desa/kelurahan berada.
    """

    kode: str
    nama: str
    lat: float
    lng: float
    kode_kecamatan: str
    kode_pos: str


class PaginatedDesaKelurahanResponse(TypedDict):
    """
    Struktur data untuk response daftar desa/kelurahan dengan informasi pagination.

    Attributes:
        pagination (PaginationMeta): Informasi metadata pagination.
        data (List[DesaKelurahan]): Daftar desa/kelurahan.
    """

    pagination: PaginationMeta
    data: List[DesaKelurahan]


class DesaKelurahanListResponse(TypedDict):
    """
    Struktur data untuk response daftar desa/kelurahan tanpa informasi pagination.

    Attributes:
        data (List[DesaKelurahan]): Daftar desa/kelurahan.
    """

    data: List[DesaKelurahan]
