# Concurrent CEP Fetcher

Este projeto em Go é uma aplicação que busca informações de endereço através de duas APIs distintas de forma simultânea. O objetivo é determinar a API que entrega a resposta mais rápida e descartar a resposta mais lenta.

## Requisitos

- Go 1.21
- Acesso à Internet para fazer chamadas às APIs

## Funcionalidades

- **Busca de CEP:**
    - Duas chamadas simultâneas às seguintes APIs:
        1. https://brasilapi.com.br/api/cep/v1/{cep}
        2. http://viacep.com.br/ws/{cep}/json/
    - A aplicação acata a resposta mais rápida e descarta a resposta mais lenta.

- **Limite de Tempo:**
    - O tempo de resposta é limitado a 1 segundo. Se uma API não responder dentro desse prazo, um erro de timeout será exibido.

- **Saída no Command Line:**
    - O resultado da busca, incluindo os dados do endereço e a indicação da API que forneceu a resposta mais rápida, é exibido no command line.

## Execução do Aplicativo

**Exemplo de uso no terminal:**
```bash
go run main.go 01153000
```

**Saída Esperada:**
```
CEP obtido da https://viacep.com.br: 
{
  "cep": "01153-000",
  "logradouro": "Rua Vitorino Carmilo",
  "complemento": "",
  "bairro": "Barra Funda",
  "localidade": "São Paulo",
  "uf": "SP",
  "ibge": "3550308",
  "gia": "1004",
  "ddd": "11",
  "siafi": "7107"
}
```
Ou
```
CEP obtido da https://brasilapi.com.br: 
{"cep":"01153000","state":"SP","city":"São Paulo","neighborhood":"Barra Funda","street":"Rua Vitorino Carmilo","service":"correios"}
```

**Nota:** Certifique-se de substituir "01153000" pelo CEP desejado ao executar o comando.
