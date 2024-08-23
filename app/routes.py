from flask import Blueprint, jsonify, request
from app.models import Provinces, Regencies

province_bp = Blueprint("provinces", __name__)
regency_bp = Blueprint("regencies", __name__)


@province_bp.route("/api/provinces", methods=["GET"])
def get_provinces():
    provinces = Provinces.query.all()
    response = [
        {"code": province.code, "name": province.name} for province in provinces
    ]
    return jsonify({"data": response})


@regency_bp.route("/api/regencies", methods=["GET"])
def get_regencies():
    province_code = request.args.get("province_code")
    if not province_code:
        return jsonify({"error": "province code is required"}), 400
    regencies = Regencies.query.filter_by(province_code=province_code).all()
    if not regencies:
        return jsonify({"message": "no regencies found for given province code"}), 404
    response = [{"code": regency.code, "name": regency.name} for regency in regencies]
    return jsonify({"data": response})
