# Admin example api

## About

    api for admin-example with laravel and frankenphp

## Dir

- app/Http/Controllers - control flow, session http response
- app/Http/Logic - logic of service
- app/Model - data model and simple handle
- app/Components/HttpError.php - all http logic error code
- config/ - config dir, load in bootstrap/app.php
- routes/web.php - define http router

## apidoc

    chmod +x ./apidoc.sh
    ./apidoc.sh

## dev

    # should has php and nodejs for local

    chmod +x dev.sh
    ./dev.sh
