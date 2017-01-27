from flask_sqlalchemy import SQLAlchemy
db = SQLAlchemy()


class Segment(db.Model):
    id = db.Column(db.Integer, primary_key=True)
    type = db.Column(db.String(15))
    color = db.Column(db.String(15))
    width = db.Column(db.Integer)
    height = db.Column(db.Integer)
    square = db.Column(db.Float)
    created = db.Column(db.DateTime)
    deleted = db.Column(db.DateTime)
    active = db.Column(db.Boolean, default=True)
    order_number = db.Column(db.String(15))
    rack = db.Column(db.String(20))
