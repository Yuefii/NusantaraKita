import pytest
from unittest.mock import AsyncMock, MagicMock, patch
from services.kabupaten_kota import KabupatenKotaService

mock_data_kabupaten = [
    {
        "kode": "11.01",
        "kode_provinsi": "11",
        "nama": "Aceh Selatan",
        "lat": 3.1618538408941346,
        "lng": 97.43651771865193,
        "geojson_url": "https://cdn.jsdelivr.net/gh/yuefii/NusantaraKita@main/geojson/kabupaten_kota/11.01.geojson",
    },
    {
        "kode": "11.02",
        "kode_provinsi": "11",
        "nama": "Aceh Tenggara",
        "lat": 3.368931368634976,
        "lng": 97.69759716544476,
        "geojson_url": "https://cdn.jsdelivr.net/gh/yuefii/NusantaraKita@main/geojson/kabupaten_kota/11.02.geojson",
    },
]


@pytest.mark.asyncio
async def test_get_without_pagination():
    mock_cursor = AsyncMock()
    mock_cursor.fetchall.return_value = mock_data_kabupaten
    mock_cursor.__aenter__.return_value = mock_cursor
    mock_cursor.__aexit__.return_value = None

    mock_conn = AsyncMock()
    mock_conn.cursor = MagicMock(return_value=mock_cursor)

    with patch("services.kabupaten_kota.get_connection", return_value=mock_conn):
        service = KabupatenKotaService()
        result = await service.get(limit=10, halaman=1, pagination=False)

        assert "data" in result
        assert result["data"] == mock_data_kabupaten
        assert isinstance(result["data"][0]["kode"], str)
        assert isinstance(result["data"][0]["nama"], str)
        assert isinstance(result["data"][0]["lat"], float)
        assert isinstance(result["data"][0]["lng"], float)
        assert isinstance(result["data"][0]["geojson_url"], str)


@pytest.mark.asyncio
async def test_get_with_pagination_valid():
    paginated_data = [mock_data_kabupaten[1]]
    mock_cursor = AsyncMock()
    mock_cursor.fetchone.return_value = {"total": 1}
    mock_cursor.fetchall.return_value = paginated_data
    mock_cursor.__aenter__.return_value = mock_cursor
    mock_cursor.__aexit__.return_value = None

    mock_conn = AsyncMock()
    mock_conn.cursor = MagicMock(return_value=mock_cursor)

    with patch("services.kabupaten_kota.get_connection", return_value=mock_conn):
        service = KabupatenKotaService()
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

    with patch("services.kabupaten_kota.get_connection", return_value=mock_conn):
        service = KabupatenKotaService()
        with pytest.raises(Exception) as exc_info:
            await service.get(limit=2, halaman=10, pagination=True)
        assert "nomor halaman melebihi total halaman" in str(exc_info.value)


@pytest.mark.asyncio
async def test_get_by_provinsi_without_pagination():
    mock_cursor = AsyncMock()
    mock_cursor.fetchall.return_value = [mock_data_kabupaten[0]]
    mock_cursor.__aenter__.return_value = mock_cursor
    mock_cursor.__aexit__.return_value = None

    mock_conn = AsyncMock()
    mock_conn.cursor = MagicMock(return_value=mock_cursor)

    with patch("services.kabupaten_kota.get_connection", return_value=mock_conn):
        service = KabupatenKotaService()
        result = await service.get_by_provinsi(
            kode_provinsi="11", limit=10, halaman=1, pagination=False
        )

        assert "data" in result
        assert all(item["kode_provinsi"] == "11" for item in result["data"])


@pytest.mark.asyncio
async def test_get_by_provinsi_with_pagination_valid():
    paginated_data = [mock_data_kabupaten[0]]
    mock_cursor = AsyncMock()
    mock_cursor.fetchone.return_value = {"total": 1}
    mock_cursor.fetchall.return_value = paginated_data
    mock_cursor.__aenter__.return_value = mock_cursor
    mock_cursor.__aexit__.return_value = None

    mock_conn = AsyncMock()
    mock_conn.cursor = MagicMock(return_value=mock_cursor)

    with patch("services.kabupaten_kota.get_connection", return_value=mock_conn):
        service = KabupatenKotaService()
        result = await service.get_by_provinsi(
            kode_provinsi="11", limit=1, halaman=1, pagination=True
        )

        assert "pagination" in result
        assert result["pagination"]["total_item"] == 1
        assert result["pagination"]["halaman_saat_ini"] == 1
        assert all(item["kode_provinsi"] == "11" for item in result["data"])


@pytest.mark.asyncio
async def test_get_by_provinsi_invalid_page_number():
    mock_cursor = AsyncMock()
    mock_cursor.fetchone.return_value = {"total": 5}
    mock_cursor.__aenter__.return_value = mock_cursor
    mock_cursor.__aexit__.return_value = None

    mock_conn = AsyncMock()
    mock_conn.cursor = MagicMock(return_value=mock_cursor)

    with patch("services.kabupaten_kota.get_connection", return_value=mock_conn):
        service = KabupatenKotaService()
        with pytest.raises(Exception) as exc_info:
            await service.get_by_provinsi(
                kode_provinsi="11", limit=2, halaman=10, pagination=True
            )
        assert "nomor halaman melebihi total halaman" in str(exc_info.value)
