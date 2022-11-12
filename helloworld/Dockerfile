# syntax=docker/dockerfile:1

FROM python:latest

WORKDIR /python-docker

COPY flask-listener.py flask-listener.py
COPY function.py function.py
COPY requirements.txt requirements.txt

RUN pip3 install -r requirements.txt
RUN pip3 install requests
RUN pip3 install Flask

CMD [ "python3", "flask-listener.py"]