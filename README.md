This aims at reproducing a weird shutdown behaviour when using a java grpc client and a go grpc server. 

To reproduce, shut down the server with Ctrl-C while it is processing the client RPC. 
- With the go client: the RPC finishes correctly, and the the server stops. 
- With the java client: the RPC finishes correctly, adn the sever hangs until the client closes the connection manually 

Run the go server with 
```sh
cd ./server-go
go run . 
```

Run the java client with: (Requires maven & jdk 17+)
```sh
cd client-java
mvn compile exec:java
```

Run the go client:
```sh
cd ./client-go
go run . 
```
