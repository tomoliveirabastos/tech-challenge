### para rodar a api com docker
```bash
docker-compose build
docker-compose up
```

### para rodar a api sem docker
```bash
go run .
ou
go build . -o server && ./server 
```

### você também pode ajustar a variavel PORT para definir em qual porta a api vai rodar
```bash
PORT=9999 ./server
```

### para executar os testes unitários
```bash
go test *.go
```

### swagger se encontra nesse caminho
http://127.0.0.1:9999/doc/index.html


### chamada na API
http://127.0.0.1:9999/?escolha=pedra