# 🚀 Plano de Estudos em Go (Golang) - Foco Backend

**Meta Final**: Construir uma API escalável com Go, PostgreSQL, autenticação JWT e deploy em cloud.

---

## 📅 Cronograma (12 Semanas)

### 🟢 Fase 1: Fundamentos (Semanas 1-3)
| Semana | Tópicos               | Projeto Prático               | Recursos |
|--------|-----------------------|-------------------------------|----------|
| 1      | Sintaxe básica        | CLI de cálculos matemáticos   | [Tour of Go](https://go.dev/tour/) |
| 2      | Structs e JSON        | Conversor CSV para JSON       | [Go by Example](https://gobyexample.com/) |
| 3      | Testes e concorrência | Web crawler simples           | [Learn Go with Tests](https://quii.gitbook.io/learn-go-with-tests/) |

**Exercício Chave**:
```go
package main

import (
    "encoding/csv"
    "encoding/json"
    "os"
)

func main() {
    // Implemente um conversor CSV→JSON
}
```

### 🟢 Fase 2: Backend Básico (Semanas 4-6)
| Semana | Tópicos               | Projeto Prático               | Recursos |
|--------|-----------------------|-------------------------------|----------|
| 4      | HTTP Server Nativo	 | API de clima (OpenWeather)    | [`net/http`](https://pkg.go.dev/net/http) |
| 5      | Gin Framework         | To-Do List API                | [Gin guide](https://gin-gonic.com/en/) |
| 6      | SQLite/PostgreSQL     | Sistema de usuários           | [GROM](https://gorm.io) |


**Exemplo de endpoint Gin**:

```go
router := gin.Default()
router.GET("/users", func(c *gin.Context) {
    c.JSON(200, gin.H{"data": "lista de usuários"})
})
```
### 🟢 Fase 3: Sistemas profissionais (Semanas 7-12)
| Semana | Tópicos               | Projeto Prático               
|--------|-----------------------|-------------------------------
| 7-8    | Autenticação JWT	     | API com login    
| 9-10   | GRPC + Protobuf       | Microserviço de chat               
| 11-12  | Deploy em Cloud       | API no AWS/GCP           


