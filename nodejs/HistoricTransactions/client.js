const fs = require("fs");
const grpc = require("grpc");
const protoLoader = require("@grpc/proto-loader");

const SERVER = "SERVER_PATH";
const TOKEN = 'YOUR_TOKEN';
const PROTO_FILE_PATH = "../../types.proto";
const ASSET = "BTC"
const COUNT = 10

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

const now = Date.now() / 1000 | 0; // unix timestamp
let channel = client.HistoricTransactions({
    from: {seconds: now - 7200},
    to: {seconds: now},
    type: ASSET,
    count: COUNT,
}, metadata);

channel.on("data", function (message) {
    console.log(message);
});