## Passo a passo para rodar o projeto

1. **Iniciar o docker compose**: 
   - Execute o comando `docker-compose up` para iniciar os containers do banco de dados e do RabbitMQ.
2. **Rodar a migração**: 
   - Execute o comando `make migrateup` para criar a tabela no banco de dados.
3. **Rodar o projeto**:
   - Execute o comando `go run main.go wire_gen.go` para iniciar os servidores Web, GRPC e GraphQL.
   - O servidor web estará rodando na porta 8000.
   - Teste os endpoints do servidor Web utilizando os arquivos http da pasta `api`.
   - Acesse o GraphQL em `http://localhost:8080/`
   - Acesse o GRPC utilizando o Evans com o comando `evans -r repl`, pois a porta estará setada como 50051.
    
    