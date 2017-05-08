FROM mysql:5.7

ARG MYSQL_ROOT_PASSWORD
ENV MYSQL_ROOT_PASSWORD=${MYSQL_ROOT_PASSWORD}

CMD ["mysqld", "--init-file=/etc/mysql/conf.d/init.sql"]
