# README

## Introduce

API Server of CTEP

## Deploying

1. Install `go1.12.5` by using [gvm](https://github.com/moovweb/gvm)
2. Run `gvm use go1.12.5 --default`

## Structure

Directroy          | Description
------------------ | -----------------------------------------------------------------------------------
config/            | Configuration files
db/migrateions/    | Mysql database schema
db/dbconfig.yml    | Database connection settings for `sql-migrate`
assets/            | User uploaded files
logs/              | Log files
playground/        | Golang playground
src/assembler/     | Data transfer object, aka [DTO](https://en.wikipedia.org/wiki/Data_transfer_object)
src/errors/        | Self-defined error
src/handler/       | API handler
src/middleware/    | API middleware
src/persistence/   | Data access object, aka [DAO](https://en.wikipedia.org/wiki/Data_access_object)
src/route/         | API router
src/service/       | Bussiness logic
src/specification/ | SQL query templete
src/utils/         | Some useful functions

## Libaray

- [kataras/iris](https://github.com/kataras/iris)
- [iris-contrib/middleware](https://github.com/iris-contrib/middleware)
- [spf13/viper](https://github.com/spf13/viper)
- [jinzhu/gorm](https://github.com/jinzhu/gorm)
- [go-sql-driver/mysql](https://github.com/go-sql-driver/mysql)
- [rubenv/sql-migrate](https://github.com/rubenv/sql-migrate)
- [gomodule/redigo](https://github.com/gomodule/redigo)
- [dgrijalva/jwt-go](github.com/dgrijalva/jwt-go)
- [satori/go.uuid](github.com/satori/go.uuid)
- [golang/crypto](https://github.com/golang/crypto)
- [asaskevich/govalidator](https://github.com/asaskevich/govalidator)
