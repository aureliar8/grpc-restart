# grpc-java

## Build


Generates protobuf classes and gRPC stubs.

```shell
mvn compile
...
[INFO] Adding generated sources (java): .../grpc-restart/client-java/src/main/java
[INFO] Adding generated sources (grpc-java): .../grpc-restart/client-java/src/main/java
```

Run

```
;; separate shell on "server-go" folder
GODEBUG=http2debug=2 go run .

;; run the client
mvn compile exec:java
```
