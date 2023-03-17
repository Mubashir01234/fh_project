# fh_project
FH project - step by step


---

## Start a mysql project
docker run --name some-mysql -p 3306:3306 -e MYSQL_ROOT_PASSWORD=`password` -d mysql:latest

## Log into this container interactively
docker exec -it some-mysql bash
