package exoscale.main;

import exoscale.sos.GreeterGrpc;
import exoscale.sos.HelloReply;
import exoscale.sos.HelloRequest;
import io.grpc.ManagedChannel;
import io.grpc.ManagedChannelBuilder;
import java.lang.Thread ;



// https://grpc.io/docs/languages/java/basics/#client
public class Client {
    public static void main(String[] args) {
        ManagedChannel channel = ManagedChannelBuilder.forAddress("localhost", 6666)
            .usePlaintext()
            .build();

        GreeterGrpc.GreeterBlockingStub stub =
            GreeterGrpc.newBlockingStub(channel);

	System.out.println("Starting RPC");
        HelloReply helloResponse = stub.sayHello(
            HelloRequest.newBuilder()
                .setName("Ray")
                .build());

        System.out.println("Finished RPC");
	System.out.println("Waiting 30 sec to not not close tcp connection manually too fast");
	try {
	    Thread.sleep(30*1000, 0); //30s 
	} catch (InterruptedException e) {
	}
	System.out.println("Shutting down the client channel");
        channel.shutdown();
    }
}
