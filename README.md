## Arrancar aplicación
Levantar mongodb
```
docker compose -up
```
Asegurar que los contenedores estén RUNNING
```
docker PS
```
Iniciar Replica Set manualmente desde tu host
```
docker exec -it mongo1 mongosh
```
Dentro de mongosh
```
rs.initiate({
  _id: "rs0",
  members: [
    { _id: 0, host: "mongo1:27017" },
    { _id: 1, host: "mongo2:27018" },
    { _id: 2, host: "mongo3:27019" }
  ]
})
```
Verificar Replica Set
```
rs.status()
```
Configurar .env para conexión en la aplicación Go
```
MONGO_URI=mongodb://mongo1:27017,mongo2:27018,mongo3:27019/?replicaSet=rs0
MONGO_DB=appdb
```
Resolver esos nombre creando host aliases en /etc/hosts
```
127.0.0.1 mongo1
127.0.0.1 mongo2
127.0.0.1 mongo3
```

Conectar Go
```
export $(cat .env | xargs)
go run cmd/api/main.go
```

## Endpoints
POST /api/users
```
curl -X POST http://localhost:8080/api/users \
-H "Content-Type: application/json" \
-d '{"name":"Juan Perez","email":"juan@example.com"}'
```
GET /api/users
```
curl http://localhost:8080/api/users
```
DELETE /api/users
```
curl http://localhost:8080/api/users/698bbb8dde6aa654c804c36d
```
POST /api/products
```
curl -X POST http://localhost:8080/api/products \
-H "Content-Type: application/json" \
-d '{"name":"PC","price":1000.5,"stock":4}'
```
GET /api/products
```
curl http://localhost:8080/api/products
```
DELETE /api/products
```
curl http://localhost:8080/api/products/698bbb3ede6aa654c804c36c
```