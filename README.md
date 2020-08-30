# PAGES FETCHER

##Database setup:
- all connection config is hardcoded in file Kamil-Ambroziak/storage/mysql

- install mysql or use docker-compose.yml from folder Kamil-Ambroziak
    - if you are using docker-compose.yml:
    1. start container: docker-compose up -d
    2. get container id: docker ps
    3. log into container: docker exec -it <(container id)> sh
- start mysql server and log into mysql client from using terminal/cmd
    - mysql -u root -p
    - password: password
- create a scheme and tables (you can use commands included in initialization file from folder Kamil-Ambroziak/docs)  

##Application start
- go run main/main.go

##Using API:
- all requests are in requests.http file in Kamil-Ambroziak/tests folder