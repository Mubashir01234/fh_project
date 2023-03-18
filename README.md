# fh_project
FH project - step by step


## env file requirement
rename env.txt to .env and add correct valuew

---

## Start a mysql project
docker run --name some-mysql -p 3306:3306 -e MYSQL_ROOT_PASSWORD=`password` -d mysql:latest

## Log into this container interactively
docker exec -it some-mysql bash

---
## Swagger
For swagger run the project and browse to http://localhost:< server-port >/swagger/index.html
