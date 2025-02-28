### para rodar a api com docker
```bash
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
PORT=3000 ./server
```

### swagger se encontra nesse caminho
http://127.0.0.1:9999/doc
