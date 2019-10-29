FROM nginx:1.17.3

ENV TZ=Asia/Taipei

ENV DEBIAN_FRONTEND noninteractive

WORKDIR /usr/share/nginx/html/front/

RUN ln -snf /usr/share/zoneinfo/$TZ /etc/localtime && echo $TZ > /etc/timezone

RUN apt-get update \
    && apt-get -y upgrade

RUN apt-get install -y apt-utils \
    software-properties-common \
    net-tools \
    gnupg2 \
    curl \
    vim \
    gcc \
    g++ \
    make

# install nodejs and npm
RUN add-apt-repository -y -r ppa:chris-lea/node.js \
    && curl -sL https://deb.nodesource.com/setup_10.x | bash - \
    && apt-get install -y nodejs

# install yarn
RUN curl -sL https://dl.yarnpkg.com/debian/pubkey.gpg | apt-key add - \
    && echo "deb https://dl.yarnpkg.com/debian/ stable main" | tee /etc/apt/sources.list.d/yarn.list \
    && apt-get update \
    && apt-get install -y yarn

COPY config/nginx.conf /etc/nginx/nginx.conf

RUN mkdir -p /etc/nginx/logs
