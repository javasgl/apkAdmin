# apkAdmin [![Build Status](https://travis-ci.org/javasgl/apkAdmin.svg?branch=master)](https://travis-ci.org/javasgl/apkAdmin)

## how to fly

use docker & docker-compose, your just need to `docker-compose up -d`, then open your browser access `http://localhost:8888`

## development

rebuild mysql:
- docker-compose stop mysql	
- docker-compose rm mysql	
- docker-compose build mysql
- docker-compose up -d mysql
