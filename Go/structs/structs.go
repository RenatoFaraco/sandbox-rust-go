package main

import (
	"fmt"
)

// Produto representa um item do mercado
type Produto struct {
	Nome       string
	Preco      float64
	Quantidade int
}

// Método para calcular o valor total em estoque
func (p *Produto) ValorTotal() float64 {
	return p.Preco * float64(p.Quantidade)
}

// Método para atualizar o preço
func (p *Produto) AtualizarPreco(novoPreco float64) {
	if novoPreco > 0 {
		p.Preco = novoPreco
	}
}

// Estoque representa o conjunto de produtos
type Estoque struct {
	Produtos []*Produto
}

// Método para adicionar um produto ao estoque
func (e *Estoque) AdicionarProduto(p *Produto) {
	e.Produtos = append(e.Produtos, p)
}

// Método para calcular o valor total do estoque
func (e *Estoque) ValorTotalEstoque() float64 {
	total := 0.0
	for _, produto := range e.Produtos {
		total += produto.ValorTotal()
	}
	return total
}

func main() {
	// Criando alguns produtos
	arroz := &Produto{"s", 25.90, 50}
	feijao := &Produto{"a", 8.50, 30}
	oleo := &Produto{"a", 7.99, 20}

	// Atualizando preço de um produto
	oleo.AtualizarPreco(8.49)

	// Criando e populando o estoque
	estoque := Estoque{}
	estoque.AdicionarProduto(arroz)
	estoque.AdicionarProduto(feijao)
	estoque.AdicionarProduto(oleo)

	// Exibindo informações
	fmt.Println("=== Produtos em Estoque ===")
	for _, p := range estoque.Produtos {
		fmt.Printf("%s - Preço: R$%.2f - Quantidade: %d - Total: R$%.2f\n",
			p.Nome, p.Preco, p.Quantidade, p.ValorTotal())
	}

	fmt.Printf("\n Total do estoque: R$%.2f\n", estoque.ValorTotalEstoque())
}
