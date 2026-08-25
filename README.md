# reniec-sunat-client-go

Cliente Go para consumir el microservicio `reniec-sunat`.

## Instalacion

```bash
go get github.com/DLabsPeru/reniec-sunat-client-go
```

## Uso rapido

```go
package main

import (
	"context"
	"log"
	"time"

	reniecsunatclient "github.com/DLabsPeru/reniec-sunat-client-go"
)

func main() {
	client := reniecsunatclient.New(
		reniecsunatclient.WithTimeout(15*time.Second),
	)

	health, err := client.Health(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	log.Println("health:", health.Status)

	person, err := client.GetPersonByDNI(context.Background(), "71101328")
	if err != nil {
		log.Fatal(err)
	}

	company, err := client.GetCompanyByRUC(context.Background(), "20604633070")
	if err != nil {
		log.Fatal(err)
	}

	log.Println(person.FirstNames, person.LastNames)
	log.Println(company.BusinessName)
}
```

Si necesitas otro ambiente, puedes sobreescribir la URL:

```go
client := reniecsunatclient.NewWithBaseURL("http://localhost:8080")
```

## Opciones utiles

```go
client := reniecsunatclient.New(
    reniecsunatclient.WithTimeout(10*time.Second),
    reniecsunatclient.WithHeader("X-Request-ID", "demo-123"),
)
```

Si luego el microservicio requiere autenticacion:

```go
client := reniecsunatclient.New(
    reniecsunatclient.WithBearerToken("tu-token"),
)
```

## Manejo de errores

```go
company, err := client.GetCompanyByRUC(ctx, "00000000000")
if err != nil {
    if apiErr, ok := err.(*reniecsunatclient.Error); ok {
        if apiErr.IsNotFound() {
            log.Println("empresa no encontrada")
        }
        log.Println(apiErr.StatusCode, apiErr.Code, apiErr.Message)
    }
}
```

## Metodos disponibles

- `Health(ctx)`
- `GetPersonByDNI(ctx, dni)`
- `GetCompanyByRUC(ctx, ruc)`

## Versionado y releases

El proyecto utiliza Release Please y Conventional Commits para generar las
versiones automaticamente:

- `fix:` incrementa la version patch.
- `feat:` incrementa la version minor.
- `feat!:` o `BREAKING CHANGE:` incrementa la version major.

Cada push a `main` actualiza el PR de release. Al fusionar ese PR se genera el
tag `vX.Y.Z`, el `CHANGELOG.md` y el GitHub Release. Como esta es una libreria
Go, los consumidores obtienen las versiones directamente desde los tags del
repositorio; no se publica una imagen Docker.

El workflow usa `RELEASE_PLEASE_TOKEN` cuando el repositorio lo tiene
configurado y recurre a `GITHUB_TOKEN` en caso contrario.
