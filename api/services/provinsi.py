import aiomysql
from typing import List, Union
from config import get_connection
from helpers.cdn import CDN_PATHS
from models.provinsi import Provinsi, ProvinsiListResponse, PaginatedProvinsiResponse


class ProvinsiService:
    async def get(
        self, limit: int, halaman: int, pagination: bool
    ) -> Union[ProvinsiListResponse, PaginatedProvinsiResponse]:
        conn = await get_connection()
        async with conn.cursor(aiomysql.DictCursor) as cursor:
            try:
                if not pagination:
                    await cursor.execute(
                        "SELECT kode, nama, lat, lng  FROM nk_provinsi"
                    )
                    data: List[Provinsi] = await cursor.fetchall()
                    if not data:
                        raise Exception("tidak ditemukan data")

                    for prov in data:
                        prov["geojson_url"] = (
                            f"{CDN_PATHS['provinsi']}/{prov['kode']}.geojson"
                        )

                    return {"data": data}

                if halaman <= 0:
                    raise Exception(
                        "nomor halaman tidak valid, halaman harus lebih besar dari 0"
                    )

                await cursor.execute("SELECT COUNT(*) AS total FROM nk_provinsi")
                total_item: int = (await cursor.fetchone())["total"]
                total_halaman: int = -(-total_item // limit)

                if halaman > total_halaman:
                    raise Exception(
                        f"nomor halaman melebihi total halaman. Halaman maksimum adalah {total_halaman}"
                    )

                offset: int = (halaman - 1) * limit
                await cursor.execute(
                    "SELECT kode, nama, lat, lng FROM nk_provinsi LIMIT %s OFFSET %s",
                    (limit, offset),
                )
                data: List[Provinsi] = await cursor.fetchall()

                if not data:
                    raise Exception("tidak ditemukan data untuk halaman yang diminta")

                for prov in data:
                    prov["geojson_url"] = (
                        f"{CDN_PATHS['provinsi']}/{prov['kode']}.geojson"
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
