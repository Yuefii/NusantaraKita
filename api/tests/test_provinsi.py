import pytest
from unittest.mock import AsyncMock, MagicMock, patch
from services.provinsi import ProvinsiService

mock_data = [
    {
        "kode": "11",
        "nama": "Aceh",
        "lat": 4.225728583038235,
        "lng": 96.9063499723947,
        "geojson_url": "https://cdn.jsdelivr.net/gh/yuefii/NusantaraKita@main/geojson/provinsi/11.geojson",
    },
    {
        "kode": "12",
        "nama": "Sumatera Utara",
        "lat": 2.1884379790819697,
        "lng": 99.63567390209315,
        "geojson_url": "https://cdn.jsdelivr.net/gh/yuefii/NusantaraKita@main/geojson/provinsi/12.geojson",
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

    with patch("services.provinsi.get_connection", return_value=mock_conn):
        service = ProvinsiService()
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

    with patch("services.provinsi.get_connection", return_value=mock_conn):
        service = ProvinsiService()
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

    with patch("services.provinsi.get_connection", return_value=mock_conn):
        service = ProvinsiService()
        with pytest.raises(Exception) as exc_info:
            await service.get(limit=2, halaman=10, pagination=True)
        assert "nomor halaman melebihi total halaman" in str(exc_info.value)
