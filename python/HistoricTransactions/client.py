import sys; sys.path.append('../') # for correct types inclusion,

import grpc

import types_pb2
import types_pb2_grpc
from google.protobuf import timestamp_pb2

import time

SERVER = "SERVER_PATH"
TOKEN = 'YOUR_TOKEN'
ASSET = "BTC"
COUNT = 10

def main():

    # connect the the server
    channel = grpc.insecure_channel(SERVER)
    stub = types_pb2_grpc.TransactionServerStub(channel)

    #add your token into the header
    metadata = [('authorization', 'Bearer {}'.format(TOKEN))]

    # create timeframe 
    now = time.time()
    seconds = int(now)
    to_time = timestamp_pb2.Timestamp(seconds=seconds)
    from_time = timestamp_pb2.Timestamp(seconds=to_time.seconds - int(3600)) # last 1 hour

    # in our case we have to use kwarg because `from` is
    # is recognized as python keyword so there would syntax be error
    # if you want get value you have to use getattr()
    historic_request_kwargs = { 'from': from_time, 'to': to_time, 
                                'type': ASSET, 'count': COUNT}
    req = types_pb2.CryptoTypeRequest(**historic_request_kwargs)
    
    trx_stream = stub.HistoricTransactions(request=req, metadata=metadata)
    for transaction in trx_stream:
        print(transaction)


if __name__ == '__main__':
    main()
