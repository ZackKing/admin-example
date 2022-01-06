FROM admin/php

WORKDIR /work/api
VOLUME /work/api

COPY ./docker/start.sh /work/start.sh
RUN apt-get update -y && \
    apt-get install -y --no-install-recommends cron && \
    echo '* * * * * www-data /usr/local/bin/php /work/api/artisan schedule:run >> /work/api/storage/logs/crontab_`date +"\%Y\%m"`.log 2>&1' >> /etc/crontab && \
    /etc/init.d/cron restart && \
    chmod +x /work/start.sh

ENTRYPOINT [ "/bin/bash", "-c", "/work/start.sh" ]