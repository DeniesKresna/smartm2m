# How to run

- Install docker first to your pc
- Pull this repo to your pc
- Go to project root directory
- Copy env example file to .env
- Open terminal Run `docker-compose up -d`
- Open localhost:9876 in browser to open phpmyadmin (DB Client)
- import dbdata from this file to danatest DB https://drive.google.com/file/d/1e8VuQiLBXs0XfogUnuYl0DKSgK6MX4wa/view?usp=sharing
- The API ready to be executed

## Routes

[GET] http://localhost:8080/stock/{id}
[POST] http://localhost:8080/stock
[POST] http://localhost:8080/stock-bulk
postman example: https://drive.google.com/file/d/13v3EQHrf-fVjg2sR7ZtTsu4mCMK_IBf7/view?usp=sharing