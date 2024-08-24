from app import db


class Provinces(db.Model):
    """
    Model untuk table "provinces" dalam database.

    Attributes:
        code (str): Code Unique untuk Provinsi (primary key).
        name (str): Nama Provinsi yang tidak boleh kosong.
    """

    code = db.Column(db.String(2), primary_key=True)
    name = db.Column(db.String(255), nullable=False)


class Regencies(db.Model):
    """
    Model untuk table "regencies" dalam database.

    Attributes:
        code (str): Code Unique untuk Kabupaten/Kota (primary key).
        province_code (str): Kode Provinsi yang merupakan foreign key dari tabel "provinces".
        name (str): Nama Kabupaten/Kota yang tidak boleh kosong.
    """

    code = db.Column(db.String(5), primary_key=True)
    province_code = db.Column(
        db.String(2), db.ForeignKey("provinces.code"), nullable=False
    )
    name = db.Column(db.String(255), nullable=False)
    province = db.relationship("Provinces", backref=db.backref("regencies", lazy=True))


class Districts(db.Model):
    code = db.Column(db.String(8), primary_key=True)
    regency_code = db.Column(
        db.String(5), db.ForeignKey("regencies.code"), nullable=False
    )
    name = db.Column(db.String(255), nullable=False)
    regency = db.relationship("Regencies", backref=db.backref("districts", lazy=True))
