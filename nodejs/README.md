# Whaletrace b2b client - nodejs

# Client example

1.  ensure that you have valid api token. You won't be able to access the server without it
2.  make sure you have at least the latest stable version of Node.js
3.  type `npm install` to install dependencies
4.  Load the protobuffer definitions
    
    ```
    const proto = grpc.loadPackageDefinition(
      protoLoader.loadSync("types.proto", {
        keepCase: true,
        longs: String,
        enums: String,
        defaults: true,
        oneofs: true
      })
    );
    ``` 
    
5.  Initialize the TransactionServer service. You have to provide valid host address.
    ```
    const client = new proto.types.TransactionServer(
      SERVER,
       grpc.credentials.createInsecure())
    );
    ```
5.  Insert your token into metadata. Note, you have to provide this token in order to obtain result.
    ```
    var metadata = new grpc.Metadata();
    metadata.add('authorization', `Bearer ${YOUR_TOKEN}`);
    ```

6.  Subscribe to required stream and listen to incoming data
    ```
    let channel = client.SubscribeTransactions({
        type: ASSET,
    });
    channel.on("data", function(message) {
      console.log(message);
    });
    ```

7.  Eventually, invoke method with required parameters and obtain incoming data
    ```
    const now = Date.now() / 1000 | 0; // unix timestamp
    let channel = client.TopTransactions({
        from: {seconds: FROM_UNIX_TIMESTAMP_IN_SECONDS},
        to: {seconds: TO_UNIX_TIMESTAMP_IN_SECONDS},
        type: ASSET,
        count: COUNT,
    }, metadata);
    ```
    
For full example, see any [client.js] file in this directory.
For additional resources, see the [grpc library](https://github.com/grpc/grpc-node)
