# Whaletrace b2b client - python

# Client example

1.  Make sure you have Python 2.7 or Python 3.4 or higher
2.  Ensure you have `pip` version 9.0.1 or higher 
	`python -m pip install --upgrade pip`
3.  Install gRPC and gRPC tools: 
	```
    python -m pip install grpcio
    python -m pip install grpcio-tools
	```
4.  Generate types `python -m grpc_tools.protoc -I../ --python_out=. --grpc_python_out=. ../types.proto`  
	
	Steps 3 is optional, because this directory already contains prepared `types_pb2.py` and `types_pb2_grpc.py` files 

4.  Invoke connection to the server  
	```python
    channel = grpc.insecure_channel(SERVER)
	``` 
5.  Initialize the TransactionServer Stub(client) service.  
	```python
    stub = types_pb2_grpc.TransactionServerStub(channel)
	```
6.  Add your TOKEN into header  
	```python
    metadata = [('authorization', 'Bearer {}'.format(TOKEN))]
	```
7.  Subscribe to required stream and listen to incoming data  
	```python
    req = types_pb2.CryptoSubscribeRequest(type=ASSET)
    
    trx_stream = stub.SubscribeTransactions(request=req, metadata=metadata)
    for tweet in tweet_stream:
        print(tweet)
	```
    
For full example, see `client.py` file in any of the sub-directory.
For additional resources, see the [grpc library](https://grpc.io/docs/tutorials/basic/python/)
