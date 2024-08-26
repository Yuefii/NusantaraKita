from flask import Blueprint, jsonify, request
from app.models import Districts, Provinces, Regencies, Villages
from app.utils.pagination_query import Pagination

province_bp = Blueprint("provinces", __name__)
regency_bp = Blueprint("regencies", __name__)
district_bp = Blueprint("districts", __name__)
village_bp = Blueprint("villages", __name__)


@province_bp.route("/api/provinces", methods=["GET"])
def get_provinces():
    show_all: bool = request.args.get("show_all", "false").lower() == "true"
    page: int = request.args.get("page", 1, type=int)
    per_page: int = request.args.get("per_page", 10, type=int)
    if per_page <= 0:
        return jsonify({"error": "invalid per_page value"}), 400
    if show_all:
        provinces = Provinces.query.all()
        response = {
            "data": [
                {"code": province.code, "name": province.name} for province in provinces
            ]
        }
    else:
        try:
            pagination = Pagination(Provinces.query, page, per_page)
            response = pagination.get_paginated_data()
        except ValueError as e:
            return jsonify({"error": str(e)}), 400
    return jsonify(response)


@regency_bp.route("/api/regencies", methods=["GET"])
def get_regencies():
    show_all: bool = request.args.get("show_all", "false").lower() == "true"
    page: int = request.args.get("page", 1, type=int)
    per_page: int = request.args.get("per_page", 10, type=int)
    province_code: str | None = request.args.get("province_code")

    if show_all:
        if province_code:
            regencies = Regencies.query.filter_by(province_code=province_code).all()
        else:
            regencies = Regencies.query.all()
        response = [
            {"code": regency.code, "name": regency.name} for regency in regencies
        ]
    else:
        if province_code:
            query = Regencies.query.filter_by(province_code=province_code)
        else:
            query = Regencies.query
        try:
            pagination = Pagination(query, page, per_page)
            response = pagination.get_paginated_data()
        except ValueError as e:
            return jsonify({"error": str(e)}), 400
    return jsonify(response)


@district_bp.route("/api/districts", methods=["GET"])
def get_districts():
    show_all: bool = request.args.get("show_all", "false").lower() == "true"
    page: int = request.args.get("page", 1, type=int)
    per_page: int = request.args.get("per_page", 10, type=int)
    regency_code: str | None = request.args.get("regency_code")

    if show_all:
        if regency_code:
            districts = Districts.query.filter_by(regency_code=regency_code).all()
        else:
            districts = Districts.query.all()
        response = [
            {"code": district.code, "name": district.name} for district in districts
        ]
    else:
        if regency_code:
            query = Districts.query.filter_by(regency_code=regency_code)
        else:
            query = Districts.query
        try:
            pagination = Pagination(query, page, per_page)
            response = pagination.get_paginated_data()
        except ValueError as e:
            return jsonify({"error": str(e)}), 400
    return jsonify(response)


@village_bp.route("/api/villages", methods=["GET"])
def get_villages():
    show_all: bool = request.args.get("show_all", "false").lower() == "true"
    page: int = request.args.get("page", 1, type=int)
    per_page: int = request.args.get("per_page", 10, type=int)
    district_code: str | None = request.args.get("district_code")

    if show_all:
        if district_code:
            villages = Villages.query.filter_by(district_code=district_code).all()
        else:
            villages = Villages.query.all()
        response = [
            {"code": village.code, "name": village.name} for village in villages
        ]
    else:
        if district_code:
            query = Villages.query.filter_by(district_code=district_code)
        else:
            query = Villages.query
        try:
            pagination = Pagination(query, page, per_page)
            response = pagination.get_paginated_data()
        except ValueError as e:
            return jsonify({"error": str(e)}), 400
    return jsonify(response)
