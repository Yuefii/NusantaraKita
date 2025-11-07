import aiomysql
from typing import List, Union
from config import get_connection
from helpers.cdn import CDN_PATHS
from models.desa_kelurahan import (
    DesaKelurahan,
    DesaKelurahanListResponse,
    PaginatedDesaKelurahanResponse,
)


class DesaKelurahanService:
    async def get(
        self, limit: int, halaman: int, pagination: bool
    ) -> Union[DesaKelurahanListResponse, PaginatedDesaKelurahanResponse]:
        conn = await get_connection()
        async with conn.cursor(aiomysql.DictCursor) as cursor:
            try:
                if not pagination:
                    await cursor.execute(
                        "SELECT kode, nama, lat, lng, kode_kecamatan, kode_pos FROM nk_desa_kelurahan"
                    )
                    data: List[DesaKelurahan] = await cursor.fetchall()
                    if not data:
                        raise Exception("tidak ditemukan data")
                    for desa_kel in data:
                        desa_kel["geojson_url"] = (
                            f"{CDN_PATHS['desa_kelurahan']}/{desa_kel['kode']}.geojson"
                        )
                    return {"data": data}

                if halaman <= 0:
                    raise Exception(
                        "nomor halaman tidak valid, halaman harus lebih besar dari 0"
                    )

                await cursor.execute("SELECT COUNT(*) AS total FROM nk_desa_kelurahan")
                total_item: int = (await cursor.fetchone())["total"]
                total_halaman: int = -(-total_item // limit)

                if halaman > total_halaman:
                    raise Exception(
                        f"nomor halaman melebihi total halaman. Halaman maksimum adalah {total_halaman}"
                    )

                offset: int = (halaman - 1) * limit
                await cursor.execute(
                    "SELECT kode, nama, lat, lng, kode_kecamatan, kode_pos FROM nk_desa_kelurahan LIMIT %s OFFSET %s",
                    (limit, offset),
                )
                data: List[DesaKelurahan] = await cursor.fetchall()

                if not data:
                    raise Exception("tidak ditemukan data untuk halaman yang diminta")

                for desa_kel in data:
                    desa_kel["geojson_url"] = (
                        f"{CDN_PATHS['desa_kelurahan']}/{desa_kel['kode']}.geojson"
                    )

                return {
                    "pagination": {
                        "total_item": total_item,
                        "total_halaman": total_halaman,
                        "halaman_saat_ini": halaman,
                        "ukuran_halaman": limit,
                    },
                    "data": data,
                }
            finally:
                conn.close()

    async def get_by_kecamatan(
        self, kode_kecamatan: str, limit: int, halaman: int, pagination: bool
    ) -> Union[DesaKelurahanListResponse, PaginatedDesaKelurahanResponse]:
        conn = await get_connection()
        async with conn.cursor(aiomysql.DictCursor) as cursor:
            try:
                if not pagination:
                    await cursor.execute(
                        "SELECT kode, nama, lat, lng, kode_kecamatan, kode_pos FROM nk_desa_kelurahan WHERE kode_kecamatan = %s",
                        (kode_kecamatan,),
                    )
                    data: List[DesaKelurahan] = await cursor.fetchall()
                    if not data:
                        raise Exception(
                            "tidak ditemukan data untuk kode provinsi tersebut"
                        )
                    for desa_kel in data:
                        desa_kel["geojson_url"] = (
                            f"{CDN_PATHS['desa_kelurahan']}/{desa_kel['kode']}.geojson"
                        )
                    return {"data": data}

                if halaman <= 0:
                    raise Exception(
                        "nomor halaman tidak valid, halaman harus lebih besar dari 0"
                    )

                await cursor.execute(
                    "SELECT COUNT(*) AS total FROM nk_desa_kelurahan WHERE kode_kecamatan = %s",
                    (kode_kecamatan,),
                )
                total_item: int = (await cursor.fetchone())["total"]
                total_halaman: int = -(-total_item // limit)

                if halaman > total_halaman:
                    raise Exception(
                        f"nomor halaman melebihi total halaman. Halaman maksimum adalah {total_halaman}"
                    )

                offset: int = (halaman - 1) * limit
                await cursor.execute(
                    "SELECT kode, nama, lat, lng, kode_kecamatan, kode_pos FROM nk_desa_kelurahan WHERE kode_kecamatan = %s LIMIT %s OFFSET %s",
                    (kode_kecamatan, limit, offset),
                )
                data: List[DesaKelurahan] = await cursor.fetchall()

                if not data:
                    raise Exception("tidak ditemukan data untuk halaman yang diminta")

                for desa_kel in data:
                    desa_kel["geojson_url"] = (
                        f"{CDN_PATHS['desa_kelurahan']}/{desa_kel['kode']}.geojson"
                    )

                return {
                    "pagination": {
                        "total_item": total_item,
                        "total_halaman": total_halaman,
                        "halaman_saat_ini": halaman,
                        "ukuran_halaman": limit,
                    },
                    "data": data,
                }
            finally:
                conn.close()
