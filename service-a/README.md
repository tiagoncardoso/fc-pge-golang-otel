## Go Template

O sistema deve receber um CEP, identificar a cidade e retornar o clima atual (temperatura em graus celsius, fahrenheit e kelvin).

---
#### 🖥️ Detalhes Gerais:

Especificações e detalhes gerais do projeto. 
- O sistema deve receber um CEP válido de 8 digitos
- Será utilizada a API viaCEP para encontrar a localização que deseja consultar a temperatura: https://viacep.com.br/
- Será utilizada a API WeatherAPI para consultar a temperatura da cidade: https://www.weatherapi.com/
- O sistema deve responder adequadamente nos seguintes cenários:
  - Em caso de sucesso:
    - Código HTTP: **200**
    - Response Body: **{ "temp_C": 28.5, "temp_F": 28.5, "temp_K": 28.5 }**
  - Em caso de falha, caso o CEP seja inválido (com formato correto):
    - Código HTTP: **422**
    - Mensagem: **invalid zipcode**
  - Em caso de falha, caso o CEP não seja encontrado:
    -  Código HTTP: **404**
    - Mensagem: **can not find zipcode**
- Desenvolver testes automatizados para garantir a qualidade do código
- Utilizar docker-compose para subir a aplicação
- O sistema deverá ser publicado no Google Cloud Run

> 💡 Dica:<br/>
> - A conversão de Celsius para Fahrenheit é: **F = C * 9/5 + 32**
> - A conversão de Celsius para Kelvin é: **K = C + 273.15**

#### 🗂️ Estrutura do Projeto
    .
    ├── cmd                  # Entrypoints da aplicação
    │    └── weather_zip_app
    │           └── main.go       ### Entrypoint principal
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
    ├── docker-compose.yaml  # Arquivo de configuração do Docker Compose
    ├── .env                 # Arquivo de parametrizações globais
    └── README.md

#### 🧭 Parametrização
A aplicação servidor possui um arquivo de configuração `.env` onde é possível definir as URL's das API's para busca de cep e informações sobre temperatura, além da porta padrão da aplicação.

```
API_URL_ZIP = http://viacep.com.br/ws/{ZIP}/json/
API_URL_WEATHER = https://api.weatherapi.com/v1/current.json?q={CITY}&key=
API_KEY_WEATHER = b*********************1
WEB_SERVER_PORT = 8080
```

> 💡 **Importante:**<br/>
> Para executar a aplicação localmente, é necessário criar um arquivo `.env` na raiz do projeto com as informações acima. E adicionar a chave da API WeatherAPI no campo `API_KEY_WEATHER`.

#### 🚀 Execução:
Para executar a aplicação em ambiente local, basta utilizar o docker-compose disponível na raiz do projeto. Para isso, execute o comando abaixo:
```bash
$ docker-compose up
```

> 💡 O comando acima poderá falhar caso a porta da aplicação esteja em uso. Caso isso ocorra, será necessário alterar o valor da variável **WEB_SERVER_PORT** no arquivo `.env` ou encerrar o processo que utiliza a porta (por padrão) 8080.

### 📝 Usando a API:

- **Buscar temperatura baseada no CEP informado:**

#### 🖥️ Em ambiente local (utilizando o docker compose):
```bash
$ curl --location 'http://localhost:8000/temperature/{zipCode}' \
```

#### 🌐 Em ambiente remoto (Google Cloud Run):
```bash
$ curl --location 'https://temperatura-cep-mcaf4qqlxq-uc.a.run.app/temperature/{zipCode}' \
```
---
#### Exemplo de resposta de sucesso (status code 200):
```json
{
  "temp_C": 28.5,
  "temp_F": 83.3,
  "temp_K": 301.6
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