<h3>Objetivo:</h3> Criar um sistema CLI em Go para realizar testes de carga em um serviço web. O usuário deverá fornecer a URL do serviço, o número total de requests e a quantidade de chamadas simultâneas.


O sistema deverá gerar um relatório com informações específicas após a execução dos testes.

Entrada de Parâmetros via CLI:

<b>--url:</b> URL do serviço a ser testado.

<b>--requests:</b> Número total de requests.

<b>--concurrency:</b> Número de chamadas simultâneas.


<h3>Execução do Teste:</h3>

Realizar requests HTTP para a URL especificada.
Distribuir os requests de acordo com o nível de concorrência definido.
Garantir que o número total de requests seja cumprido.
Geração de Relatório:

Apresentar um relatório ao final dos testes contendo:
Tempo total gasto na execução
Quantidade total de requests realizados.
Quantidade de requests com status HTTP 200.
Distribuição de outros códigos de status HTTP (como 404, 500, etc.).
Execução da aplicação:
Poderemos utilizar essa aplicação fazendo uma chamada via docker.

<hr>

# Instruções de uso

### Build Docker File
```bash
docker build -t stress-test .
```
```bash
Usage:
  go-stress-test stressTest [flags]

Flags:
  -c, --concurrency int   The number of requests to send concurrently (default 2)
  -h, --help              help for stressTest
  -r, --requests int      The number of requests to send (default 10)
  -u, --url string        The URL to send the requests to
```

### Run Docker Container

```bash
docker run --rm stress-test stressTest  -u http://google.com -r 11 -c 10
```

ou também sem as short flags

```bash
docker run --rm stress-test stressTest  --url http://google.com --requests 11 --concurrency 10
```


