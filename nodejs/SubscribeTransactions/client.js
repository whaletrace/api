const fs = require("fs");
const grpc = require("grpc");
const protoLoader = require("@grpc/proto-loader");

const SERVER = "grpc.whaletrace.com:30000";
const TOKEN = 'YOUR_TOKEN';
const PROTO_FILE_PATH = "../../types.proto";
const ASSET = "BTC"

const proto = grpc.loadPackageDefinition(
    protoLoader.loadSync(PROTO_FILE_PATH, {
        keepCase: true,
        longs: String,
        enums: String,
        defaults: true,
        oneofs: true
    })
);

var types = proto.types;
const client =  new types.TransactionServer(SERVER, grpc.credentials.createInsecure());

var metadata = new grpc.Metadata();

metadata.add('authorization', `Bearer ${TOKEN}`);

let channel = client.SubscribeTransactions({
    type: ASSET,
}, metadata);

channel.on("data", function (message) {
    console.log(message);
});