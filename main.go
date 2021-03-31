package main

import (
	"github.com/carlniz/golang/http"
	"github.com/carlniz/golang/model"
	uuid "github.com/satori/go.uuid"
)

//type nome3 []string  // criando um tipo novo
type Pessoa struct { // A forma de utilizar OO é usando esse struct
	Nome  string
	Idade int
}

// A a possibilidade de retornar dois parametros
// func andou(pessoa Pessoa) (string, error) { // sintaxe de uma função
// 	if pessoa.Nome != "FDP" {
// 		return "Se adianta", errors.New("Ta errado otario")
// 	}
// 	fmt.Println("Correu, otario")
// 	return "Deu certo porra", nil
// }

// func (p Pessoa) andou() (string, error) { // atribuindo a func para a struct Pessoa
// 	if p.Nome != "FDP" {
// 		return "Se adianta", errors.New("Ta errado otario")
// 	}
// 	fmt.Println("Correu, otario")
// 	return "Deu certo porra", nil
// }

// func (p *Pessoa) setNome(nome string) { // Alterar o valor de um atributo por método, não o altera fora do escopo do método
// 	p.Nome = nome // Para alterar o valor realmente, é necessario declarar o ponteiro
// 	fmt.Println(p.Nome)
// }

func main() {
	//var nome string // Declarando
	//nome = "Tenso" // atribuindo

	//nome2 := "olá" // Declarando e atribuindo

	//var nomes nome3
	//nomes = append(nomes, "oi") // append adiciona um item em um vetor
	//nomes = append(nomes, "tchau")
	// o _ é um blank identifier
	//for _, v := range nomes { // for deve ter a variavel de indice e a para cada item
	//	fmt.Println(v)
	//}

	//ser := Pessoa{  //Instanciando uma struck
	//Nome:  "FDP",
	//Idade: 23,
	//}

	//fmt.Println(ser.Idade)
	//res, err := ser.andou() //andou(ser)
	//if err != nil {         // essa é a validação de erro no go lang
	//	fmt.Println(err.Error())
	//}
	//fmt.Println(res)

	//nome := "Carlos"
	//var nome2 *string // * para instanciar ponteiros
	//nome2 = &nome
	//*nome2 = "FullCycle" // Por conta do ponteiro, ele associa o valor da variavel ponteiro para a primeira variavel

	//fmt.Println(&nome) // Acrescentar um & antes de uma variavel, faz ele printar o endereço de memoria
	//fmt.Println(nome)

	// pessoa := Pessoa{
	// 	Nome:  "Carlos",
	// 	Idade: 34,
	// }
	// pessoa.setNome("Feitosa") //Exemplos do funcionamento de ponteiro
	// fmt.Println(pessoa.Nome)

	produto1 := model.Product{
		ID:   uuid.NewV4().String(),
		Name: "Carrão",
	}

	produto2 := model.Product{
		ID:   uuid.NewV4().String(),
		Name: "Roupa",
	}

	products := model.Products{}
	products.Add(produto1)
	products.Add(produto2)

	server := http.NewWebServer()
	server.Products = &products
	server.Serve()
}
