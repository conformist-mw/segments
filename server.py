from datetime import datetime
from flask import Flask, render_template, request, session
from flask import Response, url_for, redirect
from flask_admin import Admin
from flask_admin.contrib.sqla import ModelView
from flask_admin.base import MenuLink
from werkzeug.exceptions import HTTPException
from sqlalchemy import or_
from models import db, Segment


class AuthException(HTTPException):

    def __init__(self, message):
        super().__init__(message, Response(
            "You could not be authenticated. Please refresh the page.", 401,
            {'WWW-Authenticate': 'Basic realm="Login Required"'}
        ))


class ModelView(ModelView):

    def check_auth(self, username, password):
        return username == 'admin' and password == 'pa$$w0rd'

    def is_accessible(self):
        auth = request.authorization
        if not auth or not self.check_auth(auth.username, auth.password):
            raise AuthException('Not authenticated.')
        return True


app = Flask(__name__)
app.config.from_object('config')
db.init_app(app)
admin = Admin(app, name='Segments', url='/admin', template_mode='bootstrap3')
admin.add_link(MenuLink(name='Отрезки', url='/'))
admin.add_view(ModelView(Segment, db.session))

per_page = 10


@app.route('/')
@app.route('/<int:page>', methods=['GET'])
def segments(page=1):
    segments = Segment.query.filter(Segment.active.is_(True)).order_by(
        Segment.square).paginate(page, per_page)
    return render_template('segments.html', segments=segments)


@app.route('/search', methods=['POST'])
def search():
    session['removed'] = request.form.get('removed', False, type=bool)
    session['order_number'] = request.form.get('order_num', '%', type=str)
    session['type'] = request.form['type']
    session['color'] = request.form['color']
    session['width'] = request.form.get('width', 0, type=int)
    session['height'] = request.form.get('height', 0, type=int)
    return redirect(url_for('results'))


@app.route('/add', methods=['POST'])
def add():
    segment = Segment()
    segment.type = request.form['type']
    segment.color = request.form['color']
    segment.width = request.form.get('width', type=int)
    segment.height = request.form.get('height', type=int)
    segment.square = (segment.width * segment.height) / 10000
    db.session.add(segment)
    db.session.commit()
    return redirect('/')


@app.route('/results', methods=['GET'])
@app.route('/results/<int:page>', methods=['GET'])
def results(page=1):
    filter_conditions = [
        Segment.active.isnot(session['removed']),
        or_(Segment.width >= session['width'],
            Segment.height >= session['height']),
        or_(Segment.width >= session['height'],
            Segment.height >= session['width'])]
    if session['removed']:
        filter_conditions.append(
            Segment.order_number.like(session['order_number']))
    if session['type'] == session['color'] == 'all':
        segments = Segment.query.filter(
            *filter_conditions).order_by(
            Segment.square).paginate(page, per_page, False)
        return render_template('segments.html', segments=segments,
                               removed=session['removed'])
    elif session['type'] == 'all':
        segments = Segment.query.filter(
            Segment.color == session['color'],
            *filter_conditions).order_by(
            Segment.square).paginate(page, per_page, False)
        return render_template('segments.html', segments=segments,
                               removed=session['removed'])
    elif session['color'] == 'all':
        segments = Segment.query.filter(
            Segment.type == session['type'],
            *filter_conditions).order_by(
            Segment.square).paginate(page, per_page, False)
        return render_template('segments.html', segments=segments,
                               removed=session['removed'])
    else:
        segments = Segment.query.filter(
            Segment.type == session['type'],
            Segment.color == session['color'],
            *filter_conditions).order_by(
            Segment.square).paginate(page, per_page, False)
        return render_template('segments.html', segments=segments,
                               removed=session['removed'])


@app.route('/remove', methods=['POST'])
def remove_segment():
    segment_id = request.form['id']
    order_num = request.form['order_num']
    segment = Segment.query.filter_by(id=segment_id).first()
    segment.active = False
    segment.deleted = datetime.now()
    segment.order_number = order_num
    db.session.commit()
    return ('', 204)


if __name__ == "__main__":
    app.run(host='0.0.0.0')
