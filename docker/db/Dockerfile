FROM mysql:5.7

# documentation: https://hub.docker.com/_/mysql

# root password
ENV MYSQL_ROOT_PASSWORD h65FZU5gtW7EdPTHtCejf49VmyKAPRwByRR4n6hnQCqY6dxUTv2L7v3j7LLPugYW

# add a database
ENV MYSQL_DATABASE teacher_education

# set timezone
ENV TZ=Asia/Taipei
RUN ln -snf /usr/share/zoneinfo/$TZ /etc/localtime && echo $TZ > /etc/timezone
