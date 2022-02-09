FROM php:8.1-fpm

ADD https://raw.githubusercontent.com/mlocati/docker-php-extension-installer/master/install-php-extensions /usr/local/bin/
RUN mv /etc/localtime /etc/localtime.bak && \
    ln -sf /usr/share/zoneinfo/Asia/Hong_Kong  /etc/localtime && \
    apt-get update -y && \
    apt-get install -y --no-install-recommends zip unzip && \
    # php extensions
    chmod uga+x /usr/local/bin/install-php-extensions && sync && \
    install-php-extensions pdo_mysql redis opcache apcu && \
    # composer
    mkdir /work && \
    cd /work && \
    php -r "copy('https://install.phpcomposer.com/installer', 'composer-setup.php');" && \
    php composer-setup.php && \
    php -r "unlink('composer-setup.php');" && \
    mv composer.phar /usr/local/bin/composer && \
    ln -s /usr/local/sbin/php-fpm /etc/init.d/php-fpm && \
    chmod +x /etc/init.d/php-fpm && \
    # clean
    apt-get clean && apt-get autoclean && apt-get autoremove && \
    rm -rf /tmp/* && \
    rm -rf /var/tmp/*

COPY ./etc/php/php.ini /usr/local/etc/php/php.ini
COPY ./etc/php-fpm.d/www.conf /usr/local/etc/php-fpm.d/www.conf