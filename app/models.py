from app import db


class Provinces(db.Model):
    code = db.Column(db.String(10), primary_key=True)
    name = db.Column(db.String(255), nullable=False)
