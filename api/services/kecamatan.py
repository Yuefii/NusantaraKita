import aiomysql
from typing import List, Union
from config import get_connection
from helpers.cdn import CDN_PATHS
from models.kecamatan import (
    Kecamatan,
    KecamatanListResponse,
    PaginatedKecamatanResponse,
)


class KecamatanService:
    async def get(
        self, limit: int, halaman: int, pagination: bool, search: str = None
    ) -> Union[KecamatanListResponse, PaginatedKecamatanResponse]:
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
                    query = f"SELECT kode, nama, lat, lng, kode_kabupaten_kota FROM nk_kecamatan{where_clause}"
                    await cursor.execute(query, params)
                    data: List[Kecamatan] = await cursor.fetchall()
                    if not data:
                        raise Exception("tidak ditemukan data")

                    for kec in data:
                        kec["geojson_url"] = (
                            f"{CDN_PATHS['kabupaten_kota']}/{kec['kode']}.geojson"
                        )

                    return {"data": data}

                if halaman <= 0:
                    raise Exception(
                        "nomor halaman tidak valid, halaman harus lebih besar dari 0"
                    )

                count_query = (
                    f"SELECT COUNT(*) AS total FROM nk_kecamatan{where_clause}"
                )
                await cursor.execute(count_query, params)
                total_item: int = (await cursor.fetchone())["total"]
                total_halaman: int = -(-total_item // limit) if total_item > 0 else 1

                if halaman > total_halaman:
                    raise Exception(
                        f"nomor halaman melebihi total halaman. Halaman maksimum adalah {total_halaman}"
                    )

                offset: int = (halaman - 1) * limit
                query = f"SELECT kode, nama, lat, lng, kode_kabupaten_kota FROM nk_kecamatan{where_clause} LIMIT %s OFFSET %s"
                params.extend([limit, offset])
                await cursor.execute(query, params)
                data: List[Kecamatan] = await cursor.fetchall()

                if not data:
                    raise Exception("tidak ditemukan data untuk halaman yang diminta")

                for kec in data:
                    kec["geojson_url"] = (
                        f"{CDN_PATHS['kabupaten_kota']}/{kec['kode']}.geojson"
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

    async def get_by_kabupaten_kota(
        self,
        kode_kabupaten_kota: str,
        limit: int,
        halaman: int,
        pagination: bool,
        search: str = None,
    ) -> Union[KecamatanListResponse, PaginatedKecamatanResponse]:
        conn = await get_connection()
        async with conn.cursor(aiomysql.DictCursor) as cursor:
            try:
                # Build WHERE clause
                where_clause = " WHERE kode_kabupaten_kota = %s"
                params = [kode_kabupaten_kota]
                if search:
                    where_clause += " AND nama LIKE %s"
                    params.append(f"%{search}%")

                if not pagination:
                    query = f"SELECT kode, nama, lat, lng, kode_kabupaten_kota FROM nk_kecamatan{where_clause}"
                    await cursor.execute(query, params)
                    data: List[Kecamatan] = await cursor.fetchall()
                    if not data:
                        raise Exception(
                            "tidak ditemukan data untuk kode provinsi tersebut"
                        )

                    for kec in data:
                        kec["geojson_url"] = (
                            f"{CDN_PATHS['kecamatan']}/{kec['kode']}.geojson"
                        )

                    return {"data": data}

                if halaman <= 0:
                    raise Exception(
                        "nomor halaman tidak valid, halaman harus lebih besar dari 0"
                    )

                count_query = (
                    f"SELECT COUNT(*) AS total FROM nk_kecamatan{where_clause}"
                )
                await cursor.execute(count_query, params)
                total_item: int = (await cursor.fetchone())["total"]
                total_halaman: int = -(-total_item // limit) if total_item > 0 else 1

                if halaman > total_halaman:
                    raise Exception(
                        f"nomor halaman melebihi total halaman. Halaman maksimum adalah {total_halaman}"
                    )

                offset: int = (halaman - 1) * limit
                query = f"SELECT kode, nama, lat, lng, kode_kabupaten_kota FROM nk_kecamatan{where_clause} LIMIT %s OFFSET %s"
                params.extend([limit, offset])
                await cursor.execute(query, params)
                data: List[Kecamatan] = await cursor.fetchall()

                if not data:
                    raise Exception("tidak ditemukan data untuk halaman yang diminta")

                for kec in data:
                    kec["geojson_url"] = (
                        f"{CDN_PATHS['kecamatan']}/{kec['kode']}.geojson"
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
