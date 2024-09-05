# admin server

## run

### dev

```shell
    ./docker-dev.sh
```

### prod

```shell
    npm install -g yarn pm2
    yarn
    yarn build
    
    # run and log with pm2
    pm2 install pm2-logrotate
    pm2 set pm2-logrotate:dateFormat 'YYYYMMDD'
    pm2 set pm2-logrotate:rotateInterval '0 0 * * *'
    pm2 set pm2-logrotate:max_size '1G'
    pm2 set pm2-logrotate:rotateModule true
    pm2 start --name admin
    
    # for pm2 controller 
    pm2 restart admin
    pm2 stop admin
```
