from app import db


class Provinces(db.Model):
    """
    Model untuk table "provinces" dalam database.

    Attributes:
        code (str): Code Unique untuk Provinsi (primary key).
        name (str): Nama Provinsi yang tidak boleh kosong.
    """
    code = db.Column(db.String(10), primary_key=True)
    name = db.Column(db.String(255), nullable=False)
