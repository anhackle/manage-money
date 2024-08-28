## Initialize database(Docker)
```
docker pull mysql
```

```
docker run --name some-mysql -p 3306:3306 -e MYSQL_ROOT_PASSWORD=admin -e MYSQL_USER=admin -e MYSQL_PASSWORD=admin -e MYSQL_DATABASE=moneymanage -d mysql:latest
```
## Update config filename
```
cd manage-money && mv config/prouction_example.yaml config/production.yaml
```

## Build
```
cd manage-money && go build -o moneymanage cmd/server/main.go
```

## Run
```
./moneymanage
```