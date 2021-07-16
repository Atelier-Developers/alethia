FROM alpine:3.14

COPY script.sh /usr/local/bin/my_script

RUN chmod +x /usr/local/bin/my_script

COPY crontab /etc/crontabs/root

RUN crontab -l

CMD ["crond", "-l", "2", "-f"]