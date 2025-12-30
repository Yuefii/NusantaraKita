import aiomysql
from typing import List
from config import get_connection
from models.region_detail import RegionDetail, RegionDetailResponse


class RegionDetailService:
    def _determine_region_level(self, code: str) -> str:
        """
        Determine the region level based on code format.

        Format:
        - Provinsi: XX (e.g., "11", "32")
        - Kabupaten/Kota: XX.XX (e.g., "11.01", "32.11")
        - Kecamatan: XX.XX.XX (e.g., "11.01.01")
        - Desa/Kelurahan: XX.XX.XX.XXXX (e.g., "11.01.01.2001")
        """
        parts = code.split(".")

        if len(parts) == 1:
            return "provinsi"
        elif len(parts) == 2:
            return "kabupaten_kota"
        elif len(parts) == 3:
            return "kecamatan"
        elif len(parts) == 4:
            return "desa_kelurahan"
        else:
            raise ValueError(f"Invalid code format: {code}")

    async def _get_provinsi_name(self, cursor, kode_provinsi: str) -> str:
        """Get provinsi name by code."""
        query = "SELECT nama FROM nk_provinsi WHERE kode = %s"
        await cursor.execute(query, (kode_provinsi,))
        result = await cursor.fetchone()
        return result["nama"] if result else None

    async def _get_kabupaten_kota_name(self, cursor, kode_kabupaten_kota: str) -> str:
        """Get kabupaten/kota name by code."""
        query = "SELECT nama FROM nk_kabupaten_kota WHERE kode = %s"
        await cursor.execute(query, (kode_kabupaten_kota,))
        result = await cursor.fetchone()
        return result["nama"] if result else None

    async def _get_kecamatan_name(self, cursor, kode_kecamatan: str) -> str:
        """Get kecamatan name by code."""
        query = "SELECT nama FROM nk_kecamatan WHERE kode = %s"
        await cursor.execute(query, (kode_kecamatan,))
        result = await cursor.fetchone()
        return result["nama"] if result else None

    async def _get_desa_kelurahan_name(self, cursor, kode_desa_kelurahan: str) -> str:
        """Get desa/kelurahan name by code."""
        query = "SELECT nama FROM nk_desa_kelurahan WHERE kode = %s"
        await cursor.execute(query, (kode_desa_kelurahan,))
        result = await cursor.fetchone()
        return result["nama"] if result else None

    async def get_region_details(self, codes: List[str]) -> RegionDetailResponse:
        """
        Get region details for multiple codes.

        Args:
            codes: List of region codes

        Returns:
            RegionDetailResponse containing list of region details
        """
        conn = await get_connection()
        async with conn.cursor(aiomysql.DictCursor) as cursor:
            try:
                results: List[RegionDetail] = []

                for code in codes:
                    code = code.strip()
                    region_level = self._determine_region_level(code)

                    detail: RegionDetail = {
                        "kode": code,
                        "desa_kelurahan": None,
                        "kecamatan": None,
                        "kabupaten_kota": None,
                        "provinsi": None,
                    }

                    if region_level == "provinsi":
                        # Get provinsi name
                        detail["provinsi"] = await self._get_provinsi_name(cursor, code)
                        if not detail["provinsi"]:
                            raise Exception(
                                f"Provinsi dengan kode {code} tidak ditemukan"
                            )

                    elif region_level == "kabupaten_kota":
                        # Get kabupaten/kota name
                        detail["kabupaten_kota"] = await self._get_kabupaten_kota_name(
                            cursor, code
                        )
                        if not detail["kabupaten_kota"]:
                            raise Exception(
                                f"Kabupaten/Kota dengan kode {code} tidak ditemukan"
                            )

                        # Get provinsi name (first part of code)
                        kode_provinsi = code.split(".")[0]
                        detail["provinsi"] = await self._get_provinsi_name(
                            cursor, kode_provinsi
                        )

                    elif region_level == "kecamatan":
                        # Get kecamatan name
                        detail["kecamatan"] = await self._get_kecamatan_name(
                            cursor, code
                        )
                        if not detail["kecamatan"]:
                            raise Exception(
                                f"Kecamatan dengan kode {code} tidak ditemukan"
                            )

                        # Get kabupaten/kota name (first two parts)
                        parts = code.split(".")
                        kode_kabupaten_kota = f"{parts[0]}.{parts[1]}"
                        detail["kabupaten_kota"] = await self._get_kabupaten_kota_name(
                            cursor, kode_kabupaten_kota
                        )

                        # Get provinsi name
                        detail["provinsi"] = await self._get_provinsi_name(
                            cursor, parts[0]
                        )

                    elif region_level == "desa_kelurahan":
                        # Get desa/kelurahan name
                        detail["desa_kelurahan"] = await self._get_desa_kelurahan_name(
                            cursor, code
                        )
                        if not detail["desa_kelurahan"]:
                            raise Exception(
                                f"Desa/Kelurahan dengan kode {code} tidak ditemukan"
                            )

                        # Get kecamatan name (first three parts)
                        parts = code.split(".")
                        kode_kecamatan = f"{parts[0]}.{parts[1]}.{parts[2]}"
                        detail["kecamatan"] = await self._get_kecamatan_name(
                            cursor, kode_kecamatan
                        )

                        # Get kabupaten/kota name
                        kode_kabupaten_kota = f"{parts[0]}.{parts[1]}"
                        detail["kabupaten_kota"] = await self._get_kabupaten_kota_name(
                            cursor, kode_kabupaten_kota
                        )

                        # Get provinsi name
                        detail["provinsi"] = await self._get_provinsi_name(
                            cursor, parts[0]
                        )

                    results.append(detail)

                if not results:
                    raise Exception("Tidak ditemukan data")

                return {"data": results}

            except Exception as e:
                raise e
            finally:
                conn.close()
