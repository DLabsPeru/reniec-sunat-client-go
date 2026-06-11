# reniec-sunat-client

Cliente Go para consumir el microservicio `reniec-sunat`.

## Instalacion

```bash
go get github.com/Destiny-Peru/reniec-sunat/reniec-sunat-client
```

## Uso

```go
package main

import (
	"context"
	"log"

	reniecsunatclient "github.com/Destiny-Peru/reniec-sunat/reniec-sunat-client"
)

func main() {
	client := reniecsunatclient.New("https://api-reniec-sunat.destiny-peru.com")

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
