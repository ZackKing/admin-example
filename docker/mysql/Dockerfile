FROM mysql:8.0

RUN mv /etc/localtime /etc/localtime.bak && \
    ln -sf /usr/share/zoneinfo/Asia/Hong_Kong  /etc/localtime

ENTRYPOINT [ "mysqld", "--default-authentication-plugin=mysql_native_password" ]