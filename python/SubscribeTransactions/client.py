import sys; sys.path.append('../') # for correct types inclusion,

import grpc

import types_pb2
import types_pb2_grpc
from google.protobuf import timestamp_pb2

SERVER = "grpc.whaletrace.com:30000"
TOKEN = 'YOUR_TOKEN'
ASSET = "BTC"

def main():

    # connect the the server
    channel = grpc.insecure_channel(SERVER)
    stub = types_pb2_grpc.TransactionServerStub(channel)

    #add your token into the header
    metadata = [('authorization', 'Bearer {}'.format(TOKEN))]

    #create request
    req = types_pb2.CryptoSubscribeRequest(type=ASSET)
    
    trx_stream = stub.SubscribeTransactions(request=req, metadata=metadata)
    for transaction in trx_stream:
        print(transaction)


if __name__ == '__main__':
    main()
