## Desafio #05 - Open Telemetry e Zipkin - Golang

O sistema deve receber um CEP, identificar a cidade e retornar o clima atual (temperatura em graus celsius, fahrenheit e kelvin).

---
#### 🖥️ Detalhes Gerais:

- O sistema deve receber um input de 8 dígitos via POST, através do schema: `{ "cep": "29902555" }`
- O sistema deve validar se o input é valido (contem 8 dígitos) e é uma `STRING`
- Caso seja válido, será encaminhado para o Serviço B via HTTP
- Caso seja inválido, deve retornar:
  - Código HTTP: 422
  - Mensagem: invalid zipcode

> 💡 Dica:<br/>
> - A conversão de Celsius para Fahrenheit é: **F = C * 9/5 + 32**
> - A conversão de Celsius para Kelvin é: **K = C + 273.15**

#### 🗂️ Estrutura do Projeto
    .
    ├── cmd                  # Entrypoints da aplicação
    │    └── app_a
    │           └── main.go  ### Entrypoint principal
    ├── config               # helpers para configuração da aplicação (viper)
    ├── internal
    │    ├── application     # Implementações de casos de uso e utilitários
    │    │      ├── helper        ### Funções utilitárias
    │    │      └── usecase       ### Casos de uso da aplicação
    │    └── infra           # Implementações de repositórios e conexões com serviços externos
    │           └── web           ### Implementações e códigos gerados para a API Rest
    ├── pkg                  # Pacotes reutilizáveis utilizados na aplicação
    ├── test                 # Testes automatizados
    ├── Dockerfile           # Arquivo de configuração do Docker
    ├── .env                 # Arquivo de parametrizações globais
    └── README.md

#### 🧭 Parametrização
A aplicação servidor possui um arquivo de configuração `.env` onde é possível definir as URL's das API's para busca de cep e informações sobre temperatura, além da porta padrão da aplicação.

```
API_SERVICE=http://service-b:8081/{ZIP}
WEB_SERVER_PORT=8080
SERVICE_NAME=service-a
SERVICE_NAME_REQUEST=service-a-request
COLLECTOR_URL=otel-collector:4317
```

> 💡 **Importante:**<br/>
> Para executar a aplicação localmente, é necessário criar um arquivo `.env` (baseado no `.env.example`) na raiz do projeto com as informações acima.

#### 🚀 Execução:
> Idealmente, o serviço deverá ser executado em conjunto com o serviço B e outros serviços que compõem a aplicação. Para isso, é possível utilizar o Docker Compose que está na raíz do projeto para subir todos os serviços de forma orquestrada.

### 📝 Usando a API:

- **Buscar temperatura baseada no CEP informado:**

```bash
$ # POST
$ curl --location 'http://localhost:8080' \
--data '{
    "cep": "96215300"
}'
```

---
#### Exemplo de resposta de sucesso (status code 200):
```json
{
  "city": "Rio Grande",
  "temp_C": 19.3,
  "temp_F": 66.74000000000001,
  "temp_K": 292.45
}
```

#### Exemplo de resposta de falha - CEP inválido (status code 422):
```
invalid zipcode
```

#### Exemplo de resposta de falha - CEP não encontrado (status code 404):
```
can not find zipcode
```