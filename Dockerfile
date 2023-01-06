FROM python:3.10-bullseye

COPY requirements.txt /tmp/
RUN pip install --no-cache-dir -r /tmp/requirements.txt

COPY code /code/
WORKDIR /code
