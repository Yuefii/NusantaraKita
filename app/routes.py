from flask import Blueprint, jsonify, request
from app.models import Districts, Provinces, Regencies, Villages

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
        response = [
            {"code": province.code, "name": province.name} for province in provinces
        ]
    else:
        total_provinces = Provinces.query.count()
        total_pages = (total_provinces + per_page - 1) // per_page

        if page > total_pages:
            return jsonify({"error": "page number exceeds total pages"}), 400

        provinces_query = Provinces.query.paginate(
            page=page, per_page=per_page, error_out=False
        )

        response = {
            "pagination": {
                "total_items": provinces_query.total,
                "total_pages": provinces_query.pages,
                "current_page": provinces_query.page,
                "per_page": provinces_query.per_page,
            },
            "data": [
                {"code": province.code, "name": province.name}
                for province in provinces_query.items
            ],
        }
    return jsonify(response)


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


@district_bp.route("/api/districts", methods=["GET"])
def get_districts():
    regency_code = request.args.get("regency_code")
    if not regency_code:
        return jsonify({"error": "regency code is required"}), 400
    districts = Districts.query.filter_by(regency_code=regency_code).all()
    if not districts:
        return jsonify({"message": "no districts found for given regency code"}), 404
    response = [
        {"code": district.code, "name": district.name} for district in districts
    ]
    return jsonify({"data": response})


@village_bp.route("/api/villages", methods=["GET"])
def get_villages():
    district_code = request.args.get("district_code")
    if not district_code:
        return jsonify({"error": "district code is required"}), 400
    villages = Villages.query.filter_by(district_code=district_code).all()
    if not villages:
        return jsonify({"message": "no villages found for given district code"}), 404
    response = [{"code": village.code, "name": village.name} for village in villages]
    return jsonify({"data": response})
