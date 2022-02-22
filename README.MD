# mutant-go
go project

## 1. Configuración del entorno.

El proyecto usa las siguientes variables de entorno:

`Solicitar archivo .env`

Ejemplo del archivo .env
```sh
PORT=":4000"
ALLOW_URL="http://localhost:4000;http://localhost:4001"
DB_URL= // url de la DB con secretos
```

### 2. Ejecución de tests.

Para ejecutar los tests correr

```sh
$ go test ./... -cover
```

### 3. End-points.
GET http://{URL}/stats


POST http://{URL}/mutant
```sh
payload
{
    "dna": [
        "ATGCGA",
        "CAGTGC",
        "TTATGT",
        "AGAAGG",
        "CCCCTA",
        "TCACTC"
    ]
}
```