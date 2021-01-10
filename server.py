import locale

from flask import (
    Flask,
    Response,
    jsonify,
    redirect,
    render_template,
    request,
    session,
    url_for,
)
from flask_admin import Admin
from flask_admin.base import MenuLink
from flask_admin.contrib.sqla import ModelView
from sqlalchemy import and_, or_
from werkzeug.exceptions import HTTPException

from models import Segment, db

locale.setlocale(locale.LC_TIME, 'ru_UA.UTF-8')


class SegmentModelView(ModelView):
    def __init__(self, model, session, *args, **kwargs):
        super(SegmentModelView, self).__init__(model, session, *args, **kwargs)
        self.static_folder = 'static'
        self.endpoint = 'admin'
        self.name = 'Segments admin'


class AuthException(HTTPException):

    def __init__(self, message):
        super().__init__(message, Response(
            'You could not be authenticated. Please refresh the page.', 401,
            {'WWW-Authenticate': 'Basic realm="Login Required"'},
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
admin = Admin(
    app,
    name='Segments',
    template_mode='bootstrap3',
    index_view=SegmentModelView(Segment, db.session, url='/admin'))
admin.add_link(MenuLink(name='Отрезки', url='/'))

per_page = 10


def bad_request(message, input_field):
    response = jsonify({'message': message, 'field': input_field})
    response.status_code = 400
    return response


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
    segment.rack = request.form['rack']
    segment.square = (segment.width * segment.height) / 10000
    db.session.add(segment)
    db.session.commit()
    return redirect('/')


@app.route('/results', methods=['GET'])
@app.route('/results/<int:page>', methods=['GET'])
def results(page=1):
    filter_conditions = [
        Segment.active.isnot(session['removed']),
        or_(and_(Segment.width >= session['width'],
                 Segment.height >= session['height']),
            and_(Segment.height >= session['width'],
                 Segment.width >= session['height']))]
    if session['removed']:
        filter_conditions.append(
            Segment.order_number.like(session['order_number']))
    if session['type'] == session['color'] == 'все':
        segments = Segment.query.filter(
            *filter_conditions).order_by(
            Segment.square).paginate(page, per_page, False)
        return render_template('segments.html', segments=segments,
                               removed=session['removed'])
    elif session['type'] == 'все':
        segments = Segment.query.filter(
            Segment.color == session['color'],
            *filter_conditions).order_by(
            Segment.square).paginate(page, per_page, False)
        return render_template('segments.html', segments=segments,
                               removed=session['removed'])
    elif session['color'] == 'все':
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
    defect = request.form.get('defect', False, type=bool)
    description = request.form.get('description', '')
    if not order_num:
        if not (defect and description):
            return bad_request(
                'Это обязательное поле для дефекта', 'input.description',
            )
    if order_num:
        db_order = db.session.query(Segment).filter(
            Segment.order_number == order_num).first()
        if db_order:
            return bad_request(
                'Такой номер заказа уже есть в базе', 'input.order')
    segment = Segment.query.filter_by(id=segment_id).first()
    segment.active = False
    segment.order_number = order_num
    segment.defect = defect
    segment.description = description
    db.session.add(segment)
    db.session.commit()
    return ('', 204)


@app.route('/activate', methods=['POST'])
def activate_segment():
    segment_id = request.form['id']
    segment = db.session.query(Segment).get(segment_id)
    segment.active = True
    segment.order_number = None
    db.session.add(segment)
    db.session.commit()
    return ('', 204)


@app.route('/replace', methods=['POST'])
def replace():
    segment_id = request.form['id']
    segment = db.session.query(Segment).get(segment_id)
    segment.rack = request.form['rack']
    db.session.add(segment)
    db.session.commit()
    return ('', 204)


@app.route('/print_segments', methods=['POST'])
def print_segments():
    rack = request.form['rack']
    if rack == 'Все':
        segments = db.session.query(Segment).filter(
            Segment.active.is_(True)).order_by(
                Segment.type, Segment.color, Segment.width).all()
    else:
        segments = db.session.query(Segment).filter(
            and_(Segment.active.is_(True), Segment.rack.is_(rack))).order_by(
                Segment.type, Segment.color, Segment.width).all()
    return render_template('table.html', segments=segments)


if __name__ == '__main__':
    app.run(host='0.0.0.0', debug=True)
