FROM python:3.9-slim-buster

ENV VIRTUAL_ENV=/opt/venv
RUN python3 -m venv $VIRTUAL_ENV
ENV PATH="$VIRTUAL_ENV/bin:$PATH"

RUN apt-get update && apt-get install -y cron --no-install-recommends

RUN /opt/venv/bin/python3 -m pip install --upgrade pip
COPY requirements.txt .
RUN pip install -r requirements.txt

# Run the application:
COPY ./*.py /
COPY ./proto /proto

COPY crontab /etc/cron.d/default-cron

RUN crontab /etc/cron.d/default-cron

CMD ["cron", "-f"]