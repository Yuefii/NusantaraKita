import aiomysql
from typing import List, Union
from config import get_connection
from helpers.cdn import CDN_PATHS
from models.kabupaten_kota import (
    KabupatenKota,
    PaginatedKabupatenKotaResponse,
    KabupatenKotaListResponse,
)


class KabupatenKotaService:
    async def get(
        self, limit: int, halaman: int, pagination: bool, search: str = None
    ) -> Union[KabupatenKotaListResponse, PaginatedKabupatenKotaResponse]:
        conn = await get_connection()
        async with conn.cursor(aiomysql.DictCursor) as cursor:
            try:
                # Build WHERE clause for search
                where_clause = ""
                params = []
                if search:
                    where_clause = " WHERE nama LIKE %s"
                    params.append(f"%{search}%")

                if not pagination:
                    query = f"SELECT kode, nama, lat, lng, kode_provinsi FROM nk_kabupaten_kota{where_clause}"
                    await cursor.execute(query, params)
                    data: List[KabupatenKota] = await cursor.fetchall()
                    if not data:
                        raise Exception("tidak ditemukan data")

                    for kab_kota in data:
                        kab_kota["geojson_url"] = (
                            f"{CDN_PATHS['kabupaten_kota']}/{kab_kota['kode']}.geojson"
                        )

                    return {"data": data}

                if halaman <= 0:
                    raise Exception(
                        "nomor halaman tidak valid, halaman harus lebih besar dari 0"
                    )

                count_query = (
                    f"SELECT COUNT(*) AS total FROM nk_kabupaten_kota{where_clause}"
                )
                await cursor.execute(count_query, params)
                total_item: int = (await cursor.fetchone())["total"]
                total_halaman: int = -(-total_item // limit) if total_item > 0 else 1

                if halaman > total_halaman:
                    raise Exception(
                        f"nomor halaman melebihi total halaman. Halaman maksimum adalah {total_halaman}"
                    )

                offset: int = (halaman - 1) * limit
                query = f"SELECT kode, nama, lat, lng, kode_provinsi FROM nk_kabupaten_kota{where_clause} LIMIT %s OFFSET %s"
                params.extend([limit, offset])
                await cursor.execute(query, params)
                data: List[KabupatenKota] = await cursor.fetchall()

                if not data:
                    raise Exception("tidak ditemukan data untuk halaman yang diminta")

                for kab_kota in data:
                    kab_kota["geojson_url"] = (
                        f"{CDN_PATHS['kabupaten_kota']}/{kab_kota['kode']}.geojson"
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

    async def get_by_provinsi(
        self,
        kode_provinsi: str,
        limit: int,
        halaman: int,
        pagination: bool,
        search: str = None,
    ) -> Union[KabupatenKotaListResponse, PaginatedKabupatenKotaResponse]:
        conn = await get_connection()
        async with conn.cursor(aiomysql.DictCursor) as cursor:
            try:
                # Build WHERE clause
                where_clause = " WHERE kode_provinsi = %s"
                params = [kode_provinsi]
                if search:
                    where_clause += " AND nama LIKE %s"
                    params.append(f"%{search}%")

                if not pagination:
                    query = f"SELECT kode, nama, lat, lng, kode_provinsi FROM nk_kabupaten_kota{where_clause}"
                    await cursor.execute(query, params)
                    data: List[KabupatenKota] = await cursor.fetchall()
                    if not data:
                        raise Exception(
                            "tidak ditemukan data untuk kode provinsi tersebut"
                        )

                    for kab_kota in data:
                        kab_kota["geojson_url"] = (
                            f"{CDN_PATHS['kabupaten_kota']}/{kab_kota['kode']}.geojson"
                        )

                    return {"data": data}

                if halaman <= 0:
                    raise Exception(
                        "nomor halaman tidak valid, halaman harus lebih besar dari 0"
                    )

                count_query = (
                    f"SELECT COUNT(*) AS total FROM nk_kabupaten_kota{where_clause}"
                )
                await cursor.execute(count_query, params)
                total_item: int = (await cursor.fetchone())["total"]
                total_halaman: int = -(-total_item // limit) if total_item > 0 else 1

                if halaman > total_halaman:
                    raise Exception(
                        f"nomor halaman melebihi total halaman. Halaman maksimum adalah {total_halaman}"
                    )

                offset: int = (halaman - 1) * limit
                query = f"SELECT kode, nama, lat, lng, kode_provinsi FROM nk_kabupaten_kota{where_clause} LIMIT %s OFFSET %s"
                params.extend([limit, offset])
                await cursor.execute(query, params)
                data: List[KabupatenKota] = await cursor.fetchall()

                if not data:
                    raise Exception("tidak ditemukan data untuk halaman yang diminta")

                for kab_kota in data:
                    kab_kota["geojson_url"] = (
                        f"{CDN_PATHS['kabupaten_kota']}/{kab_kota['kode']}.geojson"
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
