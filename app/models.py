from app import db


class Provinces(db.Model):
    """
    Model untuk table "provinces" dalam database.

    Attributes:
        code (str): Code Unique untuk Provinsi (primary key).
        name (str): Nama Provinsi yang tidak boleh kosong.
    """

    code: str = db.Column(db.String(2), primary_key=True)
    name: str = db.Column(db.String(255), nullable=False)


class Regencies(db.Model):
    """
    Model untuk table "regencies" dalam database.

    Attributes:
        code (str): Code Unique untuk Kabupaten/Kota (primary key).
        province_code (str): Kode Provinsi yang merupakan foreign key dari tabel "provinces".
        name (str): Nama Kabupaten/Kota yang tidak boleh kosong.
    """

    code: str = db.Column(db.String(5), primary_key=True)
    province_code: str = db.Column(
        db.String(2), db.ForeignKey("provinces.code"), nullable=False
    )
    name: str = db.Column(db.String(255), nullable=False)
    province = db.relationship("Provinces", backref=db.backref("regencies", lazy=True))


class Districts(db.Model):
    """
    Model untuk table "districts" dalam database.

    Attributes:
        code (str): Code Unique untuk Kecamatan (primary key).
        regency_code (str): Kode Kabupaten/Kota yang merupakan foreign key dari tabel "regencies".
        name (str): Nama Kecamatan yang tidak boleh kosong.
    """

    code: str = db.Column(db.String(8), primary_key=True)
    regency_code: str = db.Column(
        db.String(5), db.ForeignKey("regencies.code"), nullable=False
    )
    name: str = db.Column(db.String(255), nullable=False)
    regency = db.relationship("Regencies", backref=db.backref("districts", lazy=True))


class Villages(db.Model):
    """
    Model untuk table "villages" dalam database.

    Attributes:
        code (str): Code Unique untuk Desa (primary key).
        district_code (str): Kode Kecamatan yang merupakan foreign key dari tabel "districts".
        name (str): Nama Desa yang tidak boleh kosong.
    """

    code: str = db.Column(db.String(13), primary_key=True)
    district_code: str = db.Column(
        db.String(8), db.ForeignKey("districts.code"), nullable=False
    )
    name: str = db.Column(db.String(255), nullable=False)
    district = db.relationship("Districts", backref=db.backref("villages", lazy=True))
