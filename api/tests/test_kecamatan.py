import pytest
from unittest.mock import AsyncMock, MagicMock, patch
from services.kecamatan import KecamatanService

mock_data = [
    {
        "kode": "11.01.01",
        "kode_kabupaten_kota": "11.01",
        "nama": "Bakongan",
        "lat": 2.960325743420683,
        "lng": 97.46087307098534,
        "geojson_url": "https://cdn.jsdelivr.net/gh/yuefii/NusantaraKita@main/geojson/kecamatan/11.01.01.geojson",
    },
    {
        "kode": "11.02",
        "kode_kabupaten_kota": "11.01",
        "nama": "Kluet Utara",
        "lat": 3.123248075034301,
        "lng": 97.34619985236354,
        "geojson_url": "https://cdn.jsdelivr.net/gh/yuefii/NusantaraKita@main/geojson/kecamatan/11.01.02.geojson",
    },
]


@pytest.mark.asyncio
async def test_get_without_pagination():
    mock_cursor = AsyncMock()
    mock_cursor.fetchall.return_value = mock_data
    mock_cursor.__aenter__.return_value = mock_cursor
    mock_cursor.__aexit__.return_value = None

    mock_conn = AsyncMock()
    mock_conn.cursor = MagicMock(return_value=mock_cursor)

    with patch("services.kecamatan.get_connection", return_value=mock_conn):
        service = KecamatanService()
        result = await service.get(limit=10, halaman=1, pagination=False)

        assert "data" in result
        assert result["data"] == mock_data
        assert isinstance(result["data"][0]["kode"], str)
        assert isinstance(result["data"][0]["nama"], str)
        assert isinstance(result["data"][0]["lat"], float)
        assert isinstance(result["data"][0]["lng"], float)
        assert isinstance(result["data"][0]["geojson_url"], str)


@pytest.mark.asyncio
async def test_get_with_pagination_valid():
    paginated_data = [mock_data[1]]
    mock_cursor = AsyncMock()
    mock_cursor.fetchone.return_value = {"total": 1}
    mock_cursor.fetchall.return_value = paginated_data
    mock_cursor.__aenter__.return_value = mock_cursor
    mock_cursor.__aexit__.return_value = None

    mock_conn = AsyncMock()
    mock_conn.cursor = MagicMock(return_value=mock_cursor)

    with patch("services.kecamatan.get_connection", return_value=mock_conn):
        service = KecamatanService()
        result = await service.get(limit=1, halaman=1, pagination=True)

        assert "pagination" in result
        assert result["pagination"]["total_item"] == 1
        assert result["pagination"]["halaman_saat_ini"] == 1
        assert result["data"] == paginated_data


@pytest.mark.asyncio
async def test_get_invalid_page_number():
    mock_cursor = AsyncMock()
    mock_cursor.fetchone.return_value = {"total": 5}
    mock_cursor.__aenter__.return_value = mock_cursor
    mock_cursor.__aexit__.return_value = None

    mock_conn = AsyncMock()
    mock_conn.cursor = MagicMock(return_value=mock_cursor)

    with patch("services.kecamatan.get_connection", return_value=mock_conn):
        service = KecamatanService()
        with pytest.raises(Exception) as exc_info:
            await service.get(limit=2, halaman=10, pagination=True)
        assert "nomor halaman melebihi total halaman" in str(exc_info.value)


@pytest.mark.asyncio
async def test_get_by_kabupaten_kota_without_pagination():
    mock_cursor = AsyncMock()
    mock_cursor.fetchall.return_value = [mock_data[0]]
    mock_cursor.__aenter__.return_value = mock_cursor
    mock_cursor.__aexit__.return_value = None

    mock_conn = AsyncMock()
    mock_conn.cursor = MagicMock(return_value=mock_cursor)

    with patch("services.kecamatan.get_connection", return_value=mock_conn):
        service = KecamatanService()
        result = await service.get_by_kabupaten_kota(
            kode_kabupaten_kota="11.01", limit=10, halaman=1, pagination=False
        )

        assert "data" in result
        assert all(item["kode_kabupaten_kota"] == "11.01" for item in result["data"])


@pytest.mark.asyncio
async def test_get_by_kabupaten_kota_with_pagination_valid():
    paginated_data = [mock_data[0]]
    mock_cursor = AsyncMock()
    mock_cursor.fetchone.return_value = {"total": 1}
    mock_cursor.fetchall.return_value = paginated_data
    mock_cursor.__aenter__.return_value = mock_cursor
    mock_cursor.__aexit__.return_value = None

    mock_conn = AsyncMock()
    mock_conn.cursor = MagicMock(return_value=mock_cursor)

    with patch("services.kecamatan.get_connection", return_value=mock_conn):
        service = KecamatanService()
        result = await service.get_by_kabupaten_kota(
            kode_kabupaten_kota="11.01", limit=1, halaman=1, pagination=True
        )

        assert "pagination" in result
        assert result["pagination"]["total_item"] == 1
        assert result["pagination"]["halaman_saat_ini"] == 1
        assert all(item["kode_kabupaten_kota"] == "11.01" for item in result["data"])


@pytest.mark.asyncio
async def test_get_by_kabupaten_kota_invalid_page_number():
    mock_cursor = AsyncMock()
    mock_cursor.fetchone.return_value = {"total": 5}
    mock_cursor.__aenter__.return_value = mock_cursor
    mock_cursor.__aexit__.return_value = None

    mock_conn = AsyncMock()
    mock_conn.cursor = MagicMock(return_value=mock_cursor)

    with patch("services.kecamatan.get_connection", return_value=mock_conn):
        service = KecamatanService()
        with pytest.raises(Exception) as exc_info:
            await service.get_by_kabupaten_kota(
                kode_kabupaten_kota="11.01", limit=2, halaman=10, pagination=True
            )
        assert "nomor halaman melebihi total halaman" in str(exc_info.value)
