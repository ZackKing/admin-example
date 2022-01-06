#!/bin/bash

cd /work/api
composer install
chmod -R 777 ./storage

if [ ! -d /home/www-data  ];then
  mkdir /home/www-data # create user dir
  chown -R www-data:www-data /home/www-data
fi

cron -l
php-fpm
