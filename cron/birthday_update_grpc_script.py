import grpc
import logging
from proto.birthday_pb2 import BirthdayUpdateRequest
from proto.birthday_pb2_grpc import BirthdayUpdaterStub

def run():
  channel = grpc.insecure_channel('api:7000')
  stub = BirthdayUpdaterStub(channel)
  response = stub.UpdateInfos(BirthdayUpdateRequest())
  logging.info("Client received: " + response.status)
  
logging.basicConfig(format='%(asctime)s - %(message)s', datefmt='%d-%b-%y %H:%M:%S', level=logging.INFO)
run()