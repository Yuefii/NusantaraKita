from fastapi import APIRouter, HTTPException, Query
from typing import List

from services.provinsi import ProvinsiService
from services.kabupaten_kota import KabupatenKotaService
from services.kecamatan import KecamatanService
from services.desa_kelurahan import DesaKelurahanService
from services.region_detail import RegionDetailService

router = APIRouter()
provinsi = ProvinsiService()
kabupaten_kota = KabupatenKotaService()
kecamatan = KecamatanService()
desa_kelurahan = DesaKelurahanService()
region_detail = RegionDetailService()


@router.get("/v2/provinsi")
async def get_provinsi(
    limit: int = Query(10, ge=1),
    halaman: int = Query(1, ge=1),
    pagination: bool = Query(True),
    search: str = Query(None),
):
    try:
        return await provinsi.get(limit, halaman, pagination, search)
    except Exception as e:
        raise HTTPException(status_code=500, detail=str(e))


@router.get("/v2/kab-kota")
async def get_kabupatan_kota(
    limit: int = Query(10, ge=1),
    halaman: int = Query(1, ge=1),
    pagination: bool = Query(True),
    search: str = Query(None),
):
    try:
        return await kabupaten_kota.get(limit, halaman, pagination, search)
    except Exception as e:
        raise HTTPException(status_code=500, detail=str(e))


@router.get("/v2/{kode_provinsi}/kab-kota")
async def get_by_provinsi(
    kode_provinsi: str,
    limit: int = Query(10, gt=0),
    halaman: int = Query(1, gt=0),
    pagination: bool = Query(True),
    search: str = Query(None),
):
    try:
        result = await kabupaten_kota.get_by_provinsi(
            kode_provinsi, limit, halaman, pagination, search
        )
        return result
    except Exception as e:
        raise HTTPException(status_code=500, detail=str(e))


@router.get("/v2/kecamatan")
async def get_kecamatan(
    limit: int = Query(10, ge=1),
    halaman: int = Query(1, ge=1),
    pagination: bool = Query(True),
    search: str = Query(None),
):
    try:
        return await kecamatan.get(limit, halaman, pagination, search)
    except Exception as e:
        raise HTTPException(status_code=500, detail=str(e))


@router.get("/v2/{kode_kabupaten_kota}/kecamatan")
async def get_by_kabupaten_kota(
    kode_kabupaten_kota: str,
    limit: int = Query(10, gt=0),
    halaman: int = Query(1, gt=0),
    pagination: bool = Query(True),
    search: str = Query(None),
):
    try:
        result = await kecamatan.get_by_kabupaten_kota(
            kode_kabupaten_kota, limit, halaman, pagination, search
        )
        return result
    except Exception as e:
        raise HTTPException(status_code=500, detail=str(e))


@router.get("/v2/desa-kel")
async def get_desa_kelurahan(
    limit: int = Query(10, ge=1),
    halaman: int = Query(1, ge=1),
    pagination: bool = Query(True),
    search: str = Query(None),
):
    try:
        return await desa_kelurahan.get(limit, halaman, pagination, search)
    except Exception as e:
        raise HTTPException(status_code=500, detail=str(e))


@router.get("/v2/{kode_kecamatan}/desa-kel")
async def get_by_desa_kelurahan(
    kode_kecamatan: str,
    limit: int = Query(10, gt=0),
    halaman: int = Query(1, gt=0),
    pagination: bool = Query(True),
    search: str = Query(None),
):
    try:
        result = await desa_kelurahan.get_by_kecamatan(
            kode_kecamatan, limit, halaman, pagination, search
        )
        return result
    except Exception as e:
        raise HTTPException(status_code=500, detail=str(e))


@router.get("/v2/region/details")
async def get_region_details(
    ids: List[str] = Query(..., description="List of region codes (e.g., ['11', '32.11', '11.01.01.2004'])"),
):
    """
    Get region details by codes.
    
    Returns hierarchical region names based on the code format:
    - Provinsi code (XX): Returns provinsi name
    - Kabupaten/Kota code (XX.XX): Returns kabupaten/kota and provinsi names
    - Kecamatan code (XX.XX.XX): Returns kecamatan, kabupaten/kota, and provinsi names
    - Desa/Kelurahan code (XX.XX.XX.XXXX): Returns desa/kelurahan, kecamatan, kabupaten/kota, and provinsi names
    
    Example:
    - /v2/region/details?ids=32.11 -> Returns Kabupaten: Indramayu, Provinsi: Jawa Barat
    - /v2/region/details?ids=11.01.01.2004 -> Returns Desa/Kelurahan: Gampong Drien, Kecamatan: Bakongan, Kabupaten: Aceh Selatan, Provinsi: Aceh
    """
    try:
        return await region_detail.get_region_details(ids)
    except Exception as e:
        raise HTTPException(status_code=500, detail=str(e))
