import os
from typing import Optional
from dotenv import load_dotenv

load_dotenv()


class Config:
    """
    Konfigurasi untuk koneksi ke database.

    Attributes:
        SQLALCHEMY_DATABASE_URI (str): URL untuk koneksi pada database yang diambil dari variabel "DATABASE_URL" yang ada di .env.
        SQLALCHEMY_TRACK_MODIFICATIONS (bool): Menonaktifkan pelacakan pada perubahan object untuk mencegah terjadinya overhead.
    """

    SQLALCHEMY_DATABASE_URI: str | None = os.getenv("DATABASE_URL")
    SQLALCHEMY_TRACK_MODIFICATIONS: bool = False
