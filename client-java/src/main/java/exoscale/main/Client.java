package exoscale.main;

import exoscale.sos.GreeterGrpc;
import exoscale.sos.HelloReply;
import exoscale.sos.HelloRequest;
import io.grpc.ManagedChannel;
import io.grpc.ManagedChannelBuilder;

// https://grpc.io/docs/languages/java/basics/#client
public class Client {
    public static void main(String[] args) {
        ManagedChannel channel = ManagedChannelBuilder.forAddress("localhost", 6666)
            .usePlaintext()
            .build();

        GreeterGrpc.GreeterBlockingStub stub =
            GreeterGrpc.newBlockingStub(channel);

        HelloReply helloResponse = stub.sayHello(
            HelloRequest.newBuilder()
                .setName("Ray")
                .build());

        System.out.println(helloResponse);
        channel.shutdown();
    }
}
