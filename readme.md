Instruções para execução:

Executar o "Docker Compose up" para subir o mysql com a estrutura da tabela de cotação

Após subir o container do MySQL, rodar o comando dentro da pasta server "go run server.go" 

Em seguida, na pasta client, rodar o comando "client.go". O client irá fazer a request no server.

Observação: Nos testes, o contexto de apenas 200 ms, como pede o desafio, ocorre, na maioria das vezes, erro por contexto de timeout, se for recorrente, sugiro aumentar o tempo para não ficar com este erro ao corrigir o desafio.
