package main

import "github.com/aleluchesi/cursodego/mapas/model"
import "fmt"

var cache map[string]model.Imovel

func main() {

	cache = make(map[string]model.Imovel, 0)

	casa := model.Imovel{}
	casa.Nome = "Casa Amarela"
	casa.X = 10
	casa.Y = 25
	casa.SetValor(60000)

	cache["Casa Amarela"] = casa

	fmt.Println("Há uma casa amarela no cache?")
	imovel, achou := cache["Casa Amarela"]
	if achou {
		fmt.Printf("Sim, e o que achei foi %+v\r\n", imovel)
	}

	apto := model.Imovel{}
	apto.Nome = "Apto Azul"
	apto.X = 19
	apto.Y = 06
	apto.SetValor(70000)

	cache[apto.Nome] = apto

	fmt.Println("Quantos itens há no cache? ", len(cache))

	for chave, imovel := range cache {
		fmt.Printf("Chaves[%s] = %+v\r\n", chave, imovel)
	}

	imovel, achou = cache["Casa Amarela"]
	if achou {
		delete(cache, imovel.Nome)
	}

	fmt.Println("Quantos itens há no cache? ", len(cache))


}