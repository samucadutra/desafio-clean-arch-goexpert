## Passo a passo para rodar o projeto

1. **Iniciar o docker compose**: 
   - Execute o comando `docker-compose up` para iniciar os containers do banco de dados, do RabbitMQ, do app G e a migração. Tudo será feito com apenas esse comando.
2. **Informações do projeto**:
   - O servidor web estará rodando na porta 8000.
   - Teste os endpoints do servidor Web utilizando os arquivos http da pasta `api`.
   - Acesse o GraphQL em `http://localhost:8080/`
   - Acesse o GRPC utilizando o Evans com o comando `evans -r repl`, pois a porta estará setada como 50051.
    
    