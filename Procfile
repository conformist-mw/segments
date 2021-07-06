release: python manage.py collectstatic --no-post-process --noinput && python manage.py migrate
web: gunicorn config.wsgi
