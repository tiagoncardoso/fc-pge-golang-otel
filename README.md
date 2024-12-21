## Desafio #05 - Open Telemetry e Zipkin - Golang

**Objetivo:** Desenvolver um sistema em Go que receba um CEP, identifica a cidade e retorna o clima atual (temperatura em graus celsius, fahrenheit e kelvin) juntamente com a cidade. Esse sistema deverá implementar OTEL(Open Telemetry) e Zipkin.

Baseado no cenário conhecido "Sistema de temperatura por CEP" denominado Serviço B, será incluso um novo projeto, denominado Serviço A.

### Requisitos - Serviço A (responsável pelo input):

- O sistema deve receber um input de 8 dígitos via POST, através do schema: `{ "cep": "29902555" }`
- O sistema deve validar se o input é valido (contem 8 dígitos) e é uma STRING
- Caso seja válido, será encaminhado para o Serviço B via HTTP
- Caso não seja válido, deve retornar:
  - Código HTTP: 422
  - Mensagem: invalid zipcode

### Requisitos - Serviço B (responsável pela orquestração):

- O sistema deve receber um CEP válido de 8 digitos
- O sistema deve realizar a pesquisa do CEP e encontrar o nome da localização, a partir disso, deverá retornar as temperaturas e formata-lás em: Celsius, Fahrenheit, Kelvin juntamente com o nome da localização.
- O sistema deve responder adequadamente nos seguintes cenários:
  - Em caso de sucesso:
    - Código HTTP: 200
    - Response Body: `{ "city: "São Paulo", "temp_C": 28.5, "temp_F": 28.5, "temp_K": 28.5 }`
  - Em caso de falha, caso o CEP não seja válido (com formato correto):
    - Código HTTP: 422
    - Mensagem: invalid zipcode
  - Em caso de falha, caso o CEP não seja encontrado:
    - Código HTTP: 404
    - Mensagem: can not find zipcode

> #### Após a implementação dos serviços, adicione a implementação do OTEL + Zipkin:
>> - Implementar tracing distribuído entre Serviço A - Serviço B
>> - Utilizar span para medir o tempo de resposta do serviço de busca de CEP e busca de temperatura

### 🗂️ Estrutura do Projeto
    .
    ├── .docker              # Arquivos de configuração utilizados pelo Docker
    ├── service-a            # Serviço A (responsável pelo input)    
    ├── service-b            # Serviço B (responsável pela orquestração)    
    └── docker-compose.yaml  # Arquivo de configuração do Docker Compose    
    
**Service A**

Estrutura e detalhes do projeto aqui: [service-a/README.md](service-a/README.md)

**Service B**

Estrutura e detalhes do projeto aqui: [service-b/README.md](service-b/README.md)

#### 🧭 Parametrização
As aplicações possuem arquivo de configuração independentes (usar o `.env.example` por referência).

```dotenv
##> Service-a [service-a/.env]

API_SERVICE=http://service-b:8081/temperature/{ZIP}
WEB_SERVER_PORT=8080
SERVICE_NAME=service-a
SERVICE_NAME_REQUEST=service-a-request
COLLECTOR_URL=otel-collector:4317
```
```dotenv
##> Service-b [service-b/.env]

API_URL_ZIP=https://viacep.com.br/ws/{ZIP}/json/
API_URL_WEATHER=https://api.weatherapi.com/v1/current.json?q={CITY}&key=
API_KEY_WEATHER=b*********************1
WEB_SERVER_PORT=8081
SERVICE_NAME=service-b
SERVICE_NAME_REQUEST=service-b-request
COLLECTOR_URL=otel-collector:4317
```

> 💡 **Importante:**<br/>
> Para executar a aplicação localmente, é necessário criar um arquivo `.env` na raiz do projeto com as informações acima. E adicionar a chave da API WeatherAPI no campo `API_KEY_WEATHER`.

#### 🚀 Execução:
Para executar a aplicação em ambiente local, basta utilizar o docker-compose disponível na raiz do projeto. Para isso, execute o comando abaixo:
```bash
$ docker-compose up
```

> 💡 **Portas necessárias:**
> - Service A: 8080
> - Service B: 8081
> - Zipkin: 9411
> - Jaeger: 16686
> - Health check: 13133
> - OpenTelemetry gRPC receiver: 4317
> - zPages extension: 55679
