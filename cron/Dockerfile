FROM python:3.9-alpine3.14

ENV VIRTUAL_ENV=/opt/venv
RUN python3 -m venv $VIRTUAL_ENV
ENV PATH="$VIRTUAL_ENV/bin:$PATH"

# Install dependencies:
RUN apk add build-base linux-headers
RUN pip install --upgrade pip
RUN pip install setuptools==49.6.0 
RUN pip install wheel
COPY requirements.txt .
RUN pip install -r requirements.txt

# Run the application:
COPY ./*.py /
COPY ./proto /proto

COPY crontab /etc/crontabs/root

RUN crontab -l

CMD ["crond", "-l", "2", "-f"]