Exercise description:

1. Create a slice using `make` which can contains all states of Brasil
1. Show the `len` and `cap` of the slice
1. Show all the slice values without use `range`

> Code:
```go
package main

import "fmt"

func main() {
	states := make([]string, 26, 26)
	states = []string{"Acre", "Alagoas", "Amapá", "Amazonas", "Bahia", "Ceará", "Espírito Santo", "Goiás", "Maranhão", "Mato Grosso", "Mato Grosso do Sul", "Minas Gerais", "Pará", "Paraíba", "Paraná", "Pernambuco", "Piauí", "Rio de Janeiro", "Rio Grande do Norte", "Rio Grande do Sul", "Rondônia", "Roraima", "Santa Catarina", "São Paulo", "Sergipe", "Tocantins"}

	fmt.Println(len(states), cap(states))

	for i := 0; i < len(states); i++ {
		fmt.Println(states[i])
	}

}

```

> Output:
```console
26 26
Acre
Alagoas
Amapá
Amazonas
Bahia
Ceará
Espírito Santo
Goiás
Maranhão
Mato Grosso
Mato Grosso do Sul
Minas Gerais
Pará
Paraíba
Paraná
Pernambuco
Piauí
Rio de Janeiro
Rio Grande do Norte
Rio Grande do Sul
Rondônia
Roraima
Santa Catarina
São Paulo
Sergipe
Tocantins
```