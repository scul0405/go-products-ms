# go-products-ms
Go products microservice

### Jaeger UI:

http://localhost:16686

### Prometheus UI:

http://localhost:9090

### Grafana UI:

http://localhost:3000

For local development: \
Note: rename config.yaml.example to config.yaml and change mongodb uri before start server
```
make tidy // add libraries
make local // runs docker-compose.local.yml 
make run // run grpc server at localhost:8080
```
