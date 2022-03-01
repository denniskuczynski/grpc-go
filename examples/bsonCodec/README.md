# gRPC BSON Codec

 1. Get the code:

    ```console
    $ go get google.golang.org/grpc/examples/bsonCodec/bsonCodec_client
    $ go get google.golang.org/grpc/examples/bsonCodec/bsonCodec_server
    ```

 2. Enable Logging

     ```console
     $ export GRPC_GO_LOG_VERBOSITY_LEVEL=99
     $ export GRPC_GO_LOG_SEVERITY_LEVEL=info
     ```

 3. Run the server:

    ```console
    $ $(go env GOPATH)/bin/bsonCodec_server &
    ```

    ```console
    2022/03/01 13:48:41 INFO: [core] [Server #1] Server created
    2022/03/01 13:48:41 server listening at 127.0.0.1:50051
    2022/03/01 13:48:41 INFO: [core] [Server #1 ListenSocket #2] ListenSocket created
    2022/03/01 13:48:46 Received: &[{hello world}]
    2022/03/01 13:48:46 INFO: [transport] transport: loopyWriter.run returning. connection error: desc = "transport is closing"
    ```

 4. Run the client:

    ```console
    $ $(go env GOPATH)/bin/bsonCodec_client
    Greeting: Hello world
    ```

    ```console
    2022/03/01 13:48:46 INFO: [core] [Channel #1] Channel created
    2022/03/01 13:48:46 INFO: [core] [Channel #1] original dial target is: "localhost:50051"
    2022/03/01 13:48:46 INFO: [core] [Channel #1] parsed dial target is: {Scheme:localhost Authority: Endpoint:50051 URL:{Scheme:localhost Opaque:50051 User: Host: Path: RawPath: ForceQuery:false RawQuery: Fragment: RawFragment:}}
    2022/03/01 13:48:46 INFO: [core] [Channel #1] fallback to scheme "passthrough"
    2022/03/01 13:48:46 INFO: [core] [Channel #1] parsed dial target is: {Scheme:passthrough Authority: Endpoint:localhost:50051 URL:{Scheme:passthrough Opaque: User: Host: Path:/localhost:50051 RawPath: ForceQuery:false RawQuery: Fragment: RawFragment:}}
    2022/03/01 13:48:46 INFO: [core] [Channel #1] Channel authority set to "localhost:50051"
    2022/03/01 13:48:46 INFO: [core] [Channel #1] Resolver state updated: {
    "Addresses": [
        {
        "Addr": "localhost:50051",
        "ServerName": "",
        "Attributes": null,
        "BalancerAttributes": null,
        "Type": 0,
        "Metadata": null
        }
    ],
    "ServiceConfig": null,
    "Attributes": null
    } (resolver returned new addresses)
    2022/03/01 13:48:46 INFO: [core] [Channel #1] ClientConn switching balancer to "pick_first"
    2022/03/01 13:48:46 INFO: [core] [Channel #1] Channel switches to new LB policy "pick_first"
    2022/03/01 13:48:46 INFO: [core] [Channel #1 SubChannel #2] Subchannel created
    2022/03/01 13:48:46 INFO: [core] [Channel #1 SubChannel #2] Subchannel Connectivity change to CONNECTING
    2022/03/01 13:48:46 INFO: [core] [Channel #1 SubChannel #2] Subchannel picks a new address "localhost:50051" to connect
    2022/03/01 13:48:46 INFO: [core] pickfirstBalancer: UpdateSubConnState: 0xc0001ec5d0, {CONNECTING <nil>}
    2022/03/01 13:48:46 INFO: [core] [Channel #1] Channel Connectivity change to CONNECTING
    2022/03/01 13:48:46 INFO: [core] blockingPicker: the picked transport is not ready, loop back to repick
    2022/03/01 13:48:46 INFO: [core] [Channel #1 SubChannel #2] Subchannel Connectivity change to READY
    2022/03/01 13:48:46 INFO: [core] pickfirstBalancer: UpdateSubConnState: 0xc0001ec5d0, {READY <nil>}
    2022/03/01 13:48:46 INFO: [core] [Channel #1] Channel Connectivity change to READY
    2022/03/01 13:48:46 Response: [{hello world}]
    2022/03/01 13:48:46 INFO: [core] [Channel #1] Channel Connectivity change to SHUTDOWN
    2022/03/01 13:48:46 INFO: [core] [Channel #1 SubChannel #2] Subchannel Connectivity change to SHUTDOWN
    2022/03/01 13:48:46 INFO: [core] [Channel #1 SubChannel #2] Subchannel deleted
    2022/03/01 13:48:46 INFO: [core] [Channel #1] Channel deleted
    ```
