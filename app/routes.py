from flask import Blueprint, jsonify
from app.models import Provinces

province_bp = Blueprint("provinces", __name__)


@province_bp.route("/api/provinces", methods=["GET"])
def get_provinces():
    provinces = Provinces.query.all()
    return jsonify(
        [{"code": province.code, "name": province.name} for province in provinces]
    )
