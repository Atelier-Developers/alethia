# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: birthday.proto
"""Generated protocol buffer code."""
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from google.protobuf import reflection as _reflection
from google.protobuf import symbol_database as _symbol_database
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()




DESCRIPTOR = _descriptor.FileDescriptor(
  name='birthday.proto',
  package='helloworld',
  syntax='proto3',
  serialized_options=b'Z+github.com/Atelier-Developers/alethia/proto',
  create_key=_descriptor._internal_create_key,
  serialized_pb=b'\n\x0e\x62irthday.proto\x12\nhelloworld\"\x17\n\x15\x42irthdayUpdateRequest\"(\n\x16\x42irthdayUpdateResponse\x12\x0e\n\x06status\x18\x01 \x01(\t2i\n\x0f\x42irthdayUpdater\x12V\n\x0bUpdateInfos\x12!.helloworld.BirthdayUpdateRequest\x1a\".helloworld.BirthdayUpdateResponse\"\x00\x42-Z+github.com/Atelier-Developers/alethia/protob\x06proto3'
)




_BIRTHDAYUPDATEREQUEST = _descriptor.Descriptor(
  name='BirthdayUpdateRequest',
  full_name='helloworld.BirthdayUpdateRequest',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  create_key=_descriptor._internal_create_key,
  fields=[
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=30,
  serialized_end=53,
)


_BIRTHDAYUPDATERESPONSE = _descriptor.Descriptor(
  name='BirthdayUpdateResponse',
  full_name='helloworld.BirthdayUpdateResponse',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  create_key=_descriptor._internal_create_key,
  fields=[
    _descriptor.FieldDescriptor(
      name='status', full_name='helloworld.BirthdayUpdateResponse.status', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=55,
  serialized_end=95,
)

DESCRIPTOR.message_types_by_name['BirthdayUpdateRequest'] = _BIRTHDAYUPDATEREQUEST
DESCRIPTOR.message_types_by_name['BirthdayUpdateResponse'] = _BIRTHDAYUPDATERESPONSE
_sym_db.RegisterFileDescriptor(DESCRIPTOR)

BirthdayUpdateRequest = _reflection.GeneratedProtocolMessageType('BirthdayUpdateRequest', (_message.Message,), {
  'DESCRIPTOR' : _BIRTHDAYUPDATEREQUEST,
  '__module__' : 'birthday_pb2'
  # @@protoc_insertion_point(class_scope:helloworld.BirthdayUpdateRequest)
  })
_sym_db.RegisterMessage(BirthdayUpdateRequest)

BirthdayUpdateResponse = _reflection.GeneratedProtocolMessageType('BirthdayUpdateResponse', (_message.Message,), {
  'DESCRIPTOR' : _BIRTHDAYUPDATERESPONSE,
  '__module__' : 'birthday_pb2'
  # @@protoc_insertion_point(class_scope:helloworld.BirthdayUpdateResponse)
  })
_sym_db.RegisterMessage(BirthdayUpdateResponse)


DESCRIPTOR._options = None

_BIRTHDAYUPDATER = _descriptor.ServiceDescriptor(
  name='BirthdayUpdater',
  full_name='helloworld.BirthdayUpdater',
  file=DESCRIPTOR,
  index=0,
  serialized_options=None,
  create_key=_descriptor._internal_create_key,
  serialized_start=97,
  serialized_end=202,
  methods=[
  _descriptor.MethodDescriptor(
    name='UpdateInfos',
    full_name='helloworld.BirthdayUpdater.UpdateInfos',
    index=0,
    containing_service=None,
    input_type=_BIRTHDAYUPDATEREQUEST,
    output_type=_BIRTHDAYUPDATERESPONSE,
    serialized_options=None,
    create_key=_descriptor._internal_create_key,
  ),
])
_sym_db.RegisterServiceDescriptor(_BIRTHDAYUPDATER)

DESCRIPTOR.services_by_name['BirthdayUpdater'] = _BIRTHDAYUPDATER

# @@protoc_insertion_point(module_scope)
