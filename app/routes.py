from flask import Blueprint, jsonify, request
from app.models import Districts, Provinces, Regencies, Villages
from app.utils.pagination_query import Pagination

province_bp = Blueprint("provinces", __name__)
regency_bp = Blueprint("regencies", __name__)
district_bp = Blueprint("districts", __name__)
village_bp = Blueprint("villages", __name__)


@province_bp.route("/api/provinces", methods=["GET"])
def get_provinces():
    show_all = request.args.get("show_all", "false").lower() == "true"
    page = request.args.get("page", 1, type=int)
    per_page = request.args.get("per_page", 10, type=int)
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
    show_all = request.args.get("show_all", "false").lower() == "true"
    page = request.args.get("page", 1, type=int)
    per_page = request.args.get("per_page", 10, type=int)
    province_code = request.args.get("province_code")

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
            total_regencies = Regencies.query.filter_by(
                province_code=province_code
            ).count()
            regencies_query = Regencies.query.filter_by(
                province_code=province_code
            ).paginate(page=page, per_page=per_page, error_out=False)
        else:
            total_regencies = Regencies.query.count()
            regencies_query = Regencies.query.paginate(
                page=page, per_page=per_page, error_out=False
            )
        total_pages = (total_regencies + per_page - 1) // per_page
        if page > total_pages:
            return jsonify({"error": "page number exceeds total pages"}), 400
        response = {
            "pagination": {
                "total_items": regencies_query.total,
                "total_pages": regencies_query.pages,
                "current_page": regencies_query.page,
                "per_page": regencies_query.per_page,
            },
            "data": [
                {"code": regency.code, "name": regency.name}
                for regency in regencies_query.items
            ],
        }
    return jsonify(response)


@district_bp.route("/api/districts", methods=["GET"])
def get_districts():
    show_all = request.args.get("show_all", "false").lower() == "true"
    page = request.args.get("page", 1, type=int)
    per_page = request.args.get("per_page", 10, type=int)
    regency_code = request.args.get("regency_code")

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
            total_districts = Districts.query.filter_by(
                regency_code=regency_code
            ).count()
            districts_query = Districts.query.filter_by(
                regency_code=regency_code
            ).paginate(page=page, per_page=per_page, error_out=False)
        else:
            total_districts = Districts.query.count()
            districts_query = Districts.query.paginate(
                page=page, per_page=per_page, error_out=False
            )
        total_pages = (total_districts + per_page - 1) // per_page
        if page > total_pages:
            return jsonify({"error": "page number exceeds total pages"}), 400
        response = {
            "pagination": {
                "total_items": districts_query.total,
                "total_pages": districts_query.pages,
                "current_page": districts_query.page,
                "per_page": districts_query.per_page,
            },
            "data": [
                {"code": district.code, "name": district.name}
                for district in districts_query.items
            ],
        }

    return jsonify(response)


@village_bp.route("/api/villages", methods=["GET"])
def get_villages():
    show_all = request.args.get("show_all", "false").lower() == "true"
    page = request.args.get("page", 1, type=int)
    per_page = request.args.get("per_page", 10, type=int)
    district_code = request.args.get("district_code")

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
            total_villages = Villages.query.filter_by(
                district_code=district_code
            ).count()
            villages_query = Villages.query.filter_by(
                district_code=district_code
            ).paginate(page=page, per_page=per_page, error_out=False)
        else:
            total_villages = Villages.query.count()
            villages_query = Villages.query.paginate(
                page=page, per_page=per_page, error_out=False
            )
        total_pages = (total_villages + per_page - 1) // per_page
        if page > total_pages:
            return jsonify({"error": "page number exceeds total pages"}), 400
        response = {
            "pagination": {
                "total_items": villages_query.total,
                "total_pages": villages_query.pages,
                "current_page": villages_query.page,
                "per_page": villages_query.per_page,
            },
            "data": [
                {"code": village.code, "name": village.name}
                for village in villages_query.items
            ],
        }

    return jsonify(response)
