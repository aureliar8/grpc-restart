package exoscale.main;

import exoscale.sos.GreeterGrpc;
import exoscale.sos.HelloReply;
import exoscale.sos.HelloRequest;
import io.grpc.ManagedChannel;
import io.grpc.netty.NettyChannelBuilder;
import io.grpc.okhttp.OkHttpChannelBuilder;
import io.grpc.stub.StreamObserver;

import java.util.concurrent.TimeUnit;
public class Client {

    public static void main(String[] args) throws InterruptedException {

        // final ManagedChannel channel = ManagedChannelBuilder
        //final ManagedChannel channel = OkHttpChannelBuilder
        final ManagedChannel channel = NettyChannelBuilder
            .forAddress("localhost", 6666)
            .usePlaintext()
            .build();

        GreeterGrpc.GreeterStub astub =
            GreeterGrpc.newStub(channel);

        HelloRequest req = HelloRequest.newBuilder()
            .setName("Ray")
            .build();
        astub.sayHello(req, new StreamObserver<>() {

            @Override
            public void onNext(HelloReply helloReply) {
                System.out.println(Thread.currentThread().getName());
                System.out.println("GOT:" + helloReply);
            }

            @Override
            public void onError(Throwable throwable) {
                System.out.println("Error");
            }

            @Override
            public void onCompleted() {
                System.out.println(Thread.currentThread().getName());
                System.out.println("Completed");
                try {
                    Thread.sleep(1000);
                } catch (InterruptedException e) {
                    throw new RuntimeException(e);
                }
            }
        });

        for (int i = 0; i < 20; i++) {
            System.out.println("[Main]Waiting " + i + " ..." + channel.getState(false));
            Thread.sleep(1000);
        }
        System.out.println("Exit");
        channel.shutdownNow();
    }

}
