# Golang + Gin + Postgres + Docker + gRPC + NGINX

gRPC and RESTful HTTP implementation of backend server written in Go.


# Technologies used

- **`Gin`** as HTTP web framework
- **`PostgreSQL`** as database
- **`SQLC`** as code generator for SQL
- **`golang-migrate`** for database migration
- **`gRPC`** Remote procedure call framework.
- **`Docker`** for containerizing the application
- **`NGINX`** as a load balancer and reverse proxy
- **`Protoc`** as protocol buffer compiler


```
docker-compose up --build
```

*This will start the gRPC server by default on port  50051 and the HTTP Gateway server on port 8080*


**`OR`**

```
docker-compose -f docker-compose-lb.yml up --build
```

*This will run the 4 instances of our Go server in 4 different containers, each running on a different port and an NGINX load balancer to distribute the load between the 4 instances*
Each instance or container will be running two servers, one for gRPC and as an  HTTP Gateway server. The NGINX will do two things:

- Load balance the incoming HTTP requests between the 4 HTTP Gateway servers running in the 4 instances
- Load balance the incoming gRPC requests between the 4 gRPC servers running in the 4 instances


NGINX Load Balancer opens two ports:

- Port 80, which maps to port 3050 for incoming HTTP requests
- Port 9090, which maps to same port 9090  for incoming gRPC requests


To send HTTP requests to the NGINX load balancer (which will distribute the requests between the 4 HTTP Gateway servers),
just open web browser or Postman and send requests to `http://localhost:3050`

To send gRPC requests to the NGINX load balancer (which will distribute the requests between the 4 gRPC servers),
you can use the Evans CLI tool to send gRPC requests to the NGINX load balancer. Run the following command:

```
evans --port 9090 --host localhost -r repl;
```
