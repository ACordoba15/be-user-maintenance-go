# be-user-maintenance-go

- [Go Documentación](https://gobyexample.com/)

## Inicializar un proyecto Go
- `go mod init [url]/[name]`

## Ejcutar el programa
- `go build <name>.go` crea un build.
- `./<name>` ejecuta el programa.
- `go run <name>.go` Crear el build y lo ejecuta.

- `go mod tidy` Ejecuta y limpia dependencias

## Librerías a utilizar 
[Gorilla Mux](https://github.com/gorilla/mux)
- `go get -u github.com/gorilla/mux` Instalación

[Air](https://github.com/air-verse/air#installation)
- `go install github.com/air-verse/air@latest`
- `export PATH=$PATH:$(go env GOPATH)/bin`
- `air` ejecuta el proyecto y actualiza con cada cambio

[Gorm](https://gorm.io/docs/index.html)
- `go get -u gorm.io/gorm`
- `gorm.io/driver/sqlserver` Instalamos sql server

CORS
- `go get -u github.com/rs/cors`

[Swagger](https://github.com/swaggo/swag?tab=readme-ov-file#getting-started)
- `go install github.com/swaggo/swag/cmd/swag@latest`
- `go get -u github.com/swaggo/http-swagger`
- `go get -u github.com/swaggo/files`

## DB
- `docker run -e "ACCEPT_EULA=Y" -e "MSSQL_SA_PASSWORD=Prueba001." -p 1433:1433 -d mcr.microsoft.com/mssql/server:2022-latest`
- `/opt/mssql-tools/bin/sqlcmd -S localhost -U sa -P 'Prueba001.'`
## Nota
> Si una función, variable o estructura empieza con una letra mayúscula, será exportada y accesible fuera del paquete. Si comienza con minúscula, será privada del paquete.

> Los nombres de los paquetes deben ser cortos y en minúsculas, preferentemente en singular. Por ejemplo, si se tienes una función relacionada con usuarios, puedes crear un paquete user en lugar de users.