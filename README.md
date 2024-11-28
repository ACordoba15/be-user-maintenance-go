# be-user-maintenance-go

- [Go Documentación](https://gobyexample.com/)

## Inicializar un proyecto Go
- `go mod init [url]/[name]`

## Librerías a utilizar 
[Gorilla Mux](https://github.com/gorilla/mux)
- `go get -u github.com/gorilla/mux` Instalación

[Air](https://github.com/air-verse/air#installation)
- `go install github.com/air-verse/air@latest`
- `export PATH=$PATH:$(go env GOPATH)/bin`
- `air` Ejecuta el proyecto y actualiza con cada cambio

[Gorm](https://gorm.io/docs/index.html)
- `go get -u gorm.io/gorm`
- `gorm.io/driver/sqlserver` Instalamos sql server

[CORS](https://github.com/rs/cors)
- `go get -u github.com/rs/cors`

[Swagger](https://github.com/swaggo/swag?tab=readme-ov-file#getting-started) Documentación del API
- `go install github.com/swaggo/swag/cmd/swag@latest`
- `go get -u github.com/swaggo/http-swagger`
- `go get -u github.com/swaggo/files`
- `swag init` Actualiza el swagger

[GoDotEnv](https://github.com/joho/godotenv) Variables de ambiente
- `go get github.com/joho/godotenv`

[Testify](https://github.com/stretchr/testify) Pruebas unitarias
- `go get github.com/stretchr/testify`

## Ejcutar el programa
- `go mod tidy` Ejecuta y limpia dependencias

- `go build <name>.go` Crea un build.
- `./<name>` Ejecuta el programa.

- `go run <name>.go` Crear el build y lo ejecuta.

- `air` Ejecuta el proyecto y lee los cambios


## ENV
DB_SERVER=localhost
DB_PORT=1433
DB_USER=sa
DB_PASSWORD=Prueba001.
DB_NAME=tdusers-go

## DB
- `docker run -e "ACCEPT_EULA=Y" -e "MSSQL_SA_PASSWORD=password" -p 1433:1433 -d mcr.microsoft.com/mssql/server:2022-latest`
- `/opt/mssql-tools/bin/sqlcmd -S localhost -U sa -P 'password'`

## Nota
> Si una función, variable o estructura empieza con una letra mayúscula, será exportada y accesible fuera del paquete. Si comienza con minúscula, será privada del paquete.

> Los nombres de los paquetes deben ser cortos y en minúsculas, preferentemente en singular. Por ejemplo, si se tienes una función relacionada con usuarios, puedes crear un paquete user en lugar de users.