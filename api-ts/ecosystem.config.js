module.exports = {
  apps : [{
    name: 'admin',
    script: 'dist/main.js',
    watch: false,
    autorestart: true,
    env: {
      NODE_ENV: 'prod',
    },
    out_file: 'storages/logs/out.log'
  }]
}
