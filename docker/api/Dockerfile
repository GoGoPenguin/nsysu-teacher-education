FROM golang:1.14.4

ENV TZ=Asia/Taipei

RUN ln -snf /usr/share/zoneinfo/$TZ /etc/localtime && echo $TZ > /etc/timezone

RUN apt-get update \
    && apt-get -y upgrade

RUN apt-get install -y apt-utils \
    supervisor \
    net-tools \
    vim

RUN mkdir -p /var/www/html/nsysu-teacher-education/

COPY ./config/supervisord.conf /etc/supervisor/conf.d/supervisord.conf

WORKDIR /var/www/html/api

EXPOSE 80 443

CMD ["/usr/bin/supervisord"]