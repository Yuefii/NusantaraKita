import pytest
from unittest.mock import AsyncMock, MagicMock, patch
from services.desa_kelurahan import DesaKelurahanService

mock_data = [
    {
        "kode": "11.01.01.2001",
        "kode_kecamatan": "11.01.01",
        "nama": "Keude Bakongan",
        "lat": 2.931094803160483,
        "lng": 97.48458404258515,
        "kode_pos": "23773",
    },
    {
        "kode": "11.01.01.2002",
        "kode_kecamatan": "11.01.01",
        "nama": "Ujong Mangki",
        "lat": 2.9527245335971086,
        "lng": 97.43761867741745,
        "kode_pos": "23773",
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

    with patch("services.desa_kelurahan.get_connection", return_value=mock_conn):
        service = DesaKelurahanService()
        result = await service.get(limit=10, halaman=1, pagination=False)

        assert "data" in result
        assert result["data"] == mock_data
        assert isinstance(result["data"][0]["kode"], str)
        assert isinstance(result["data"][0]["nama"], str)
        assert isinstance(result["data"][0]["lat"], float)
        assert isinstance(result["data"][0]["lng"], float)
        assert isinstance(result["data"][0]["kode_pos"], str)


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

    with patch("services.desa_kelurahan.get_connection", return_value=mock_conn):
        service = DesaKelurahanService()
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

    with patch("services.desa_kelurahan.get_connection", return_value=mock_conn):
        service = DesaKelurahanService()
        with pytest.raises(Exception) as exc_info:
            await service.get(limit=2, halaman=10, pagination=True)
        assert "nomor halaman melebihi total halaman" in str(exc_info.value)


@pytest.mark.asyncio
async def test_get_by_desa_kelurahan_without_pagination():
    mock_cursor = AsyncMock()
    mock_cursor.fetchall.return_value = [mock_data[0]]
    mock_cursor.__aenter__.return_value = mock_cursor
    mock_cursor.__aexit__.return_value = None

    mock_conn = AsyncMock()
    mock_conn.cursor = MagicMock(return_value=mock_cursor)

    with patch("services.desa_kelurahan.get_connection", return_value=mock_conn):
        service = DesaKelurahanService()
        result = await service.get_by_kecamatan(
            kode_kecamatan="11.01.01", limit=10, halaman=1, pagination=False
        )

        assert "data" in result
        assert all(item["kode_kecamatan"] == "11.01.01" for item in result["data"])


@pytest.mark.asyncio
async def test_get_by_desa_kelurahan_with_pagination_valid():
    paginated_data = [mock_data[0]]
    mock_cursor = AsyncMock()
    mock_cursor.fetchone.return_value = {"total": 1}
    mock_cursor.fetchall.return_value = paginated_data
    mock_cursor.__aenter__.return_value = mock_cursor
    mock_cursor.__aexit__.return_value = None

    mock_conn = AsyncMock()
    mock_conn.cursor = MagicMock(return_value=mock_cursor)

    with patch("services.desa_kelurahan.get_connection", return_value=mock_conn):
        service = DesaKelurahanService()
        result = await service.get_by_kecamatan(
            kode_kecamatan="11.01.01", limit=1, halaman=1, pagination=True
        )

        assert "pagination" in result
        assert result["pagination"]["total_item"] == 1
        assert result["pagination"]["halaman_saat_ini"] == 1
        assert all(item["kode_kecamatan"] == "11.01.01" for item in result["data"])


@pytest.mark.asyncio
async def test_get_by_kecamatan_invalid_page_number():
    mock_cursor = AsyncMock()
    mock_cursor.fetchone.return_value = {"total": 5}
    mock_cursor.__aenter__.return_value = mock_cursor
    mock_cursor.__aexit__.return_value = None

    mock_conn = AsyncMock()
    mock_conn.cursor = MagicMock(return_value=mock_cursor)

    with patch("services.desa_kelurahan.get_connection", return_value=mock_conn):
        service = DesaKelurahanService()
        with pytest.raises(Exception) as exc_info:
            await service.get_by_kecamatan(
                kode_kecamatan="11.01.01", limit=2, halaman=10, pagination=True
            )
        assert "nomor halaman melebihi total halaman" in str(exc_info.value)
