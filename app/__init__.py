from flask import Flask, render_template
from app.config import Config
from flask_sqlalchemy import SQLAlchemy

db = SQLAlchemy()


def create_app():
    app = Flask(__name__)
    app.config.from_object(Config)
    db.init_app(app)

    @app.route("/")
    def home():
        return render_template("index.html")

    with app.app_context():
        from app.routes import province_bp

        app.register_blueprint(province_bp)
        db.create_all()
    return app
