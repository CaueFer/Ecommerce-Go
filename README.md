
# Ecommerce API With Golang

Simples REST API com Golango, um CRUD para produtos de um ecommerce.

*Simple REST API with Go, a CRUD for products of an e-commerce.*


## API Docs

#### Retorna todos os itens - Return all items

```http
  GET /api/products
```

| Parâmetro   | Tipo       | Descrição                           |
| :---------- | :--------- | :---------------------------------- |
| `api_key` | `string` | **Obrigatório**. A chave da sua API |

#### Retorna um item - Return an item

```http
  GET /api/product/${id}
```

| Parâmetro   | Tipo       | Descrição                                   |
| :---------- | :--------- | :------------------------------------------ |
| `id`      | `string` | **Obrigatório**. O ID do item que você quer |

#### Criar um item - Create an item

```http
  GET /api/createProduct
```

| Parâmetro   | Tipo       | Descrição                                   |
| :---------- | :--------- | :------------------------------------------ |
| `name`      | `string` | **Obrigatório**. Nome do item que deseja adicionar |
| `price`      | `int` | **Obrigatório**. Preço do item que deseja adicionar |

## Instalação - Instalation

Siga os passos a baixo para rodar a api.

*Follow the steps below to run the API.*

```bash
# Baixe as dependencias do Golang
go mod tidy

# Rode o main.go
go run cmd/main.go
```
## Autores - Authors

- [@cauefer](https://github.com/CaueFer)