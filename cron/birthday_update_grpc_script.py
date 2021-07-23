import grpc
from proto.birthday_pb2 import BirthdayUpdateRequest
from proto.birthday_pb2_grpc import BirthdayUpdaterStub

def run():
  channel = grpc.insecure_channel('api:7000')
  stub = BirthdayUpdaterStub(channel)
  response = stub.UpdateInfos(BirthdayUpdateRequest())
  print("Greeter client received: " + response.status)
  
run()