from fastapi import APIRouter, HTTPException, Query

from services.provinsi import ProvinsiService
from services.kabupaten_kota import KabupatenKotaService
from services.kecamatan import KecamatanService
from services.desa_kelurahan import DesaKelurahanService

router = APIRouter()
provinsi = ProvinsiService()
kabupaten_kota = KabupatenKotaService()
kecamatan = KecamatanService()
desa_kelurahan = DesaKelurahanService()

@router.get("/v2/provinsi")
async def get_provinsi(
    limit: int = Query(10, ge=1),
    halaman: int = Query(1, ge=1),
    pagination: bool = Query(True),
):
    try:
        return await provinsi.get(limit, halaman, pagination)
    except Exception as e:
        raise HTTPException(status_code=500, detail=str(e))


@router.get("/v2/kab-kota")
async def get_kabupatan_kota(
    limit: int = Query(10, ge=1),
    halaman: int = Query(1, ge=1),
    pagination: bool = Query(True),
):
    try:
        return await kabupaten_kota.get(limit, halaman, pagination)
    except Exception as e:
        raise HTTPException(status_code=500, detail=str(e))

@router.get("/v2/{kode_provinsi}/kab-kota")
async def get_by_provinsi(
    kode_provinsi: str,
    limit: int = Query(10, gt=0),
    halaman: int = Query(1, gt=0),
    pagination: bool = Query(True)
):
    try:
        result = await kabupaten_kota.get_by_provinsi(
            kode_provinsi, limit, halaman, pagination
        )
        return result
    except Exception as e:
        raise HTTPException(status_code=500, detail=str(e))

@router.get("/v2/kecamatan")
async def get_kecamatan(
    limit: int = Query(10, ge=1),
    halaman: int = Query(1, ge=1),
    pagination: bool = Query(True),
):
    try:
        return await kecamatan.get(limit, halaman, pagination)
    except Exception as e:
        raise HTTPException(status_code=500, detail=str(e))

@router.get("/v2/{kode_kabupaten_kota}/kecamatan")
async def get_by_kabupaten_kota(
    kode_kabupaten_kota: str,
    limit: int = Query(10, gt=0),
    halaman: int = Query(1, gt=0),
    pagination: bool = Query(True)
):
    try:
        result = await kecamatan.get_by_kabupaten_kota(
            kode_kabupaten_kota, limit, halaman, pagination
        )
        return result
    except Exception as e:
        raise HTTPException(status_code=500, detail=str(e))


@router.get("/v2/desa-kel")
async def get_desa_kelurahan(
    limit: int = Query(10, ge=1),
    halaman: int = Query(1, ge=1),
    pagination: bool = Query(True),
):
    try:
        return await desa_kelurahan.get(limit, halaman, pagination)
    except Exception as e:
        raise HTTPException(status_code=500, detail=str(e))

@router.get("/v2/{kode_kecamatan}/desa-kel")
async def get_by_desa_kelurahan(
    kode_kecamatan: str,
    limit: int = Query(10, gt=0),
    halaman: int = Query(1, gt=0),
    pagination: bool = Query(True)
):
    try:
        result = await desa_kelurahan.get_by_kecamatan(
            kode_kecamatan, limit, halaman, pagination
        )
        return result
    except Exception as e:
        raise HTTPException(status_code=500, detail=str(e))
