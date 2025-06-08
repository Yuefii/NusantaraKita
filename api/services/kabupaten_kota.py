import aiomysql

from config import get_connection


class KabupatenKotaService:
    async def get(self, limit: int, halaman: int, pagination: bool):
        conn = await get_connection()
        async with conn.cursor(aiomysql.DictCursor) as cursor:
            try:
                if not pagination:
                    await cursor.execute("SELECT * FROM nk_kabupaten_kota")
                    data = await cursor.fetchall()
                    if not data:
                        raise Exception("tidak ditemukan data")
                    return {"data": data}

                if halaman <= 0:
                    raise Exception(
                        "nomor halaman tidak valid, halaman harus lebih besar dari 0"
                    )

                await cursor.execute("SELECT COUNT(*) AS total FROM nk_kabupaten_kota")
                total_item = (await cursor.fetchone())["total"]
                total_halaman = -(-total_item // limit)

                if halaman > total_halaman:
                    raise Exception(
                        f"nomor halaman melebihi total halaman. Halaman maksimum adalah {total_halaman}"
                    )

                offset = (halaman - 1) * limit
                await cursor.execute(
                    "SELECT * FROM nk_kabupaten_kota LIMIT %s OFFSET %s", (limit, offset)
                )
                data = await cursor.fetchall()

                if not data:
                    raise Exception("tidak ditemukan data untuk halaman yang diminta")

                result = {
                    "pagination": {
                        "total_item": total_item,
                        "total_halaman": total_halaman,
                        "halaman_saat_ini": halaman,
                        "ukuran_halaman": limit,
                    },
                    "data": data,
                }

                return result
            finally:
                conn.close()

    async def get_by_provinsi(self, kode_provinsi: str, limit: int, halaman: int, pagination: bool):
        conn = await get_connection()
        async with conn.cursor(aiomysql.DictCursor) as cursor:
            try:
                if not pagination:
                    await cursor.execute("SELECT * FROM nk_kabupaten_kota WHERE kode_provinsi = %s", 
                                         (kode_provinsi,))
                    data = await cursor.fetchall()
                    if not data:
                        raise Exception("tidak ditemukan data untuk kode provinsi tersebut")
                    return {"data": data}

                if halaman <= 0:
                    raise Exception("nomor halaman tidak valid, halaman harus lebih besar dari 0")

                await cursor.execute(
                    "SELECT COUNT(*) AS total FROM nk_kabupaten_kota WHERE kode_provinsi = %s", (kode_provinsi,)
                )
                total_item = (await cursor.fetchone())["total"]
                total_halaman = -(-total_item // limit)

                if halaman > total_halaman:
                    raise Exception(f"nomor halaman melebihi total halaman. Halaman maksimum adalah {total_halaman}")

                offset = (halaman - 1) * limit
                await cursor.execute(
                    "SELECT * FROM nk_kabupaten_kota WHERE kode_provinsi = %s LIMIT %s OFFSET %s",
                    (kode_provinsi, limit, offset)
                )
                data = await cursor.fetchall()

                if not data:
                    raise Exception("tidak ditemukan data untuk halaman yang diminta")

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
