# CEP Weather API

## Descrição Geral

Este repositório contém dois microserviços desenvolvidos em Go, estruturados conforme os princípios da Clean Architecture. Ambos estão integrados com OpenTelemetry e Zipkin para rastreamento distribuído:

- **Serviço A**: Recebe um CEP via requisição `POST`, realiza a validação do formato e encaminha para o Serviço B.
- **Serviço B**: Ao receber o CEP, consulta o ViaCEP para descobrir a cidade correspondente e, em seguida, utiliza a WeatherAPI para buscar a temperatura atual. A resposta inclui as temperaturas em Celsius, Fahrenheit e Kelvin.

---

## Estrutura da Arquitetura

Cada microserviço segue a seguinte organização:

- O arquivo `main.go` realiza a configuração inicial, como setup do tracer OTEL, definição de middlewares do Gin e inicialização do servidor.
- A separação de responsabilidades é feita através dos diretórios `cmd`, `internal` e `cfg`, que lidam respectivamente com rotas, lógica de negócio, comunicação HTTP, estrutura de dados e carregamento de configurações por ambiente.
- A instrumentação de observabilidade é feita com OTEL, incluindo middleware no Gin para capturar spans de entrada e saída de requisições.

---

## Configuração

Todo o ecossistema (os dois serviços e os componentes de observabilidade como OTEL Collector e Zipkin) é gerenciado por Docker Compose. Algumas variáveis importantes:

- `SERVICE_A`, `SERVICE_B`: definem as portas e os endpoints internos de cada microserviço.
- `FREEWEATHER_API_KEY`: chave de acesso à WeatherAPI.
- `OTEL_ENDPOINT`: URL de destino para o coletor OTEL (normalmente `otel-collector:4317`).
- `OTEL_SERVICE_NAME`: identifica cada serviço dentro do Zipkin.

---

## Como Executar

1. Verifique se você tem Docker e Docker Compose instalados em sua máquina.
2. Edite o valor de `WEATHER_API_KEY` no `docker-compose.yml` (dentro do serviço `api2`).
3. No terminal, dentro da raiz do projeto, execute:

   ```bash
   docker-compose up --build
   ```

4. Aguarde até que todos os containers estejam ativos, saudáveis (`healthy`) e sem apresentar erros.

---

## Endpoints Disponíveis

- **Serviço A**
  - `POST /cep`
    - Envie um JSON no formato `{ "cep": "XXXXXXXX" }`.
    - Retorna erro 422 se o CEP for inválido, ou os dados de cidade e temperatura via resposta do Serviço B.

- **Serviço B**
  - `GET /weather/:cep`
    - Recebe o CEP diretamente como parâmetro na URL.
    - Retorna:
      - 422 se o formato do CEP for inválido,
      - 404 se o CEP não existir no ViaCEP,
      - ou um JSON com nome da cidade e temperaturas nos três formatos (C, F e K).

---

## Observabilidade e Tracing

- A interface do Zipkin pode ser acessada em: [http://localhost:9411/zipkin/](http://localhost:9411/zipkin/)
- Dentro do Zipkin:
  1. Escolha o serviço (`api1` ou `api2`) no campo **Service Name**.
  2. Ajuste o intervalo de tempo para incluir as requisições feitas.
  3. Clique em **Find Traces** para visualizar as execuções recentes.

Cada etapa relevante — como validação de CEP, consulta ao ViaCEP e obtenção da temperatura — é representada por um span, facilitando a identificação de gargalos e falhas.

---

## Documentação via Swagger

- A documentação interativa do Serviço A está disponível em: [http://localhost:8080/swagger/index.html](http://localhost:8080/swagger/index.html)
- A documentação do Serviço B pode ser acessada em: [http://localhost:8090/swagger/index.html](http://localhost:8090/swagger/index.html)
