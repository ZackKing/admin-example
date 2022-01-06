# Admin example

## Start

```bash
    export ADMIN_DIR=/path/to/here
    cd ${ADMIN_DIR}/docker/php & chmod +x ./build.sh & ./build.sh

    # for dev, has mysql, redis container, if remove it, edit dev.yml
    cd ${ADMIN_DIR} & docker-compose -f dev.yml up -d --build
    # view 127.0.0.1:8081

    # for prod
    cd ${ADMIN_DIR}/web & chmod +x ./docker-build.sh & ./docker-build.sh
    cd ${ADMIN_DIR} & docker-compose -f docker-compose.yml up -d --build
    # view 127.0.0.1
```

## database

+ init sql

```bash
    cd api/sql/0.0.1
```
