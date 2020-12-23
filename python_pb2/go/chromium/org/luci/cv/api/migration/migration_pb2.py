# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: go.chromium.org/luci/cv/api/migration/migration.proto

from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from google.protobuf import reflection as _reflection
from google.protobuf import symbol_database as _symbol_database
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()


from google.protobuf import empty_pb2 as google_dot_protobuf_dot_empty__pb2
from google.protobuf import timestamp_pb2 as google_dot_protobuf_dot_timestamp__pb2
from go.chromium.org.luci.cv.api.bigquery.v1 import attempt_pb2 as go_dot_chromium_dot_org_dot_luci_dot_cv_dot_api_dot_bigquery_dot_v1_dot_attempt__pb2


DESCRIPTOR = _descriptor.FileDescriptor(
  name='go.chromium.org/luci/cv/api/migration/migration.proto',
  package='migration',
  syntax='proto3',
  serialized_options=b'Z1go.chromium.org/luci/cv/api/migration;migrationpb',
  create_key=_descriptor._internal_create_key,
  serialized_pb=b'\n5go.chromium.org/luci/cv/api/migration/migration.proto\x12\tmigration\x1a\x1bgoogle/protobuf/empty.proto\x1a\x1fgoogle/protobuf/timestamp.proto\x1a\x35go.chromium.org/luci/cv/api/bigquery/v1/attempt.proto\"1\n\x11ReportRunsRequest\x12\x1c\n\x04runs\x18\x01 \x03(\x0b\x32\x0e.migration.Run\"7\n\x18ReportFinishedRunRequest\x12\x1b\n\x03run\x18\x01 \x01(\x0b\x32\x0e.migration.Run\"C\n\x16ReportUsedNetrcRequest\x12\x13\n\x0bgerrit_host\x18\x01 \x01(\t\x12\x14\n\x0c\x61\x63\x63\x65ss_token\x18\x02 \x01(\t\".\n\x16\x46\x65tchActiveRunsRequest\x12\x14\n\x0cluci_project\x18\x01 \x01(\t\"7\n\x17\x46\x65tchActiveRunsResponse\x12\x1c\n\x04runs\x18\x01 \x03(\x0b\x32\x0e.migration.Run\"T\n\x03Run\x12\"\n\x07\x61ttempt\x18\x01 \x01(\x0b\x32\x11.bigquery.Attempt\x12\n\n\x02id\x18\x02 \x01(\t\x12\x1d\n\x03\x63ls\x18\x03 \x03(\x0b\x32\x10.migration.RunCL\"\xb1\x02\n\x05RunCL\x12\n\n\x02id\x18\x01 \x01(\x03\x12\"\n\x02gc\x18\x02 \x01(\x0b\x32\x16.bigquery.GerritChange\x12\x30\n\x0cupdated_time\x18\x03 \x01(\x0b\x32\x1a.google.protobuf.Timestamp\x12)\n\x07trigger\x18\x04 \x01(\x0b\x32\x18.migration.RunCL.Trigger\x12\"\n\x04\x64\x65ps\x18\x05 \x03(\x0b\x32\x14.migration.RunCL.Dep\x1aV\n\x07Trigger\x12(\n\x04time\x18\x01 \x01(\x0b\x32\x1a.google.protobuf.Timestamp\x12\x12\n\naccount_id\x18\x03 \x01(\x03\x12\r\n\x05\x65mail\x18\x04 \x01(\t\x1a\x1f\n\x03\x44\x65p\x12\n\n\x02id\x18\x01 \x01(\x03\x12\x0c\n\x04hard\x18\x02 \x01(\x08\x32\xc9\x02\n\tMigration\x12\x42\n\nReportRuns\x12\x1c.migration.ReportRunsRequest\x1a\x16.google.protobuf.Empty\x12P\n\x11ReportFinishedRun\x12#.migration.ReportFinishedRunRequest\x1a\x16.google.protobuf.Empty\x12L\n\x0fReportUsedNetrc\x12!.migration.ReportUsedNetrcRequest\x1a\x16.google.protobuf.Empty\x12X\n\x0f\x46\x65tchActiveRuns\x12!.migration.FetchActiveRunsRequest\x1a\".migration.FetchActiveRunsResponseB3Z1go.chromium.org/luci/cv/api/migration;migrationpbb\x06proto3'
  ,
  dependencies=[google_dot_protobuf_dot_empty__pb2.DESCRIPTOR,google_dot_protobuf_dot_timestamp__pb2.DESCRIPTOR,go_dot_chromium_dot_org_dot_luci_dot_cv_dot_api_dot_bigquery_dot_v1_dot_attempt__pb2.DESCRIPTOR,])




_REPORTRUNSREQUEST = _descriptor.Descriptor(
  name='ReportRunsRequest',
  full_name='migration.ReportRunsRequest',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  create_key=_descriptor._internal_create_key,
  fields=[
    _descriptor.FieldDescriptor(
      name='runs', full_name='migration.ReportRunsRequest.runs', index=0,
      number=1, type=11, cpp_type=10, label=3,
      has_default_value=False, default_value=[],
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
  serialized_start=185,
  serialized_end=234,
)


_REPORTFINISHEDRUNREQUEST = _descriptor.Descriptor(
  name='ReportFinishedRunRequest',
  full_name='migration.ReportFinishedRunRequest',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  create_key=_descriptor._internal_create_key,
  fields=[
    _descriptor.FieldDescriptor(
      name='run', full_name='migration.ReportFinishedRunRequest.run', index=0,
      number=1, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
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
  serialized_start=236,
  serialized_end=291,
)


_REPORTUSEDNETRCREQUEST = _descriptor.Descriptor(
  name='ReportUsedNetrcRequest',
  full_name='migration.ReportUsedNetrcRequest',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  create_key=_descriptor._internal_create_key,
  fields=[
    _descriptor.FieldDescriptor(
      name='gerrit_host', full_name='migration.ReportUsedNetrcRequest.gerrit_host', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='access_token', full_name='migration.ReportUsedNetrcRequest.access_token', index=1,
      number=2, type=9, cpp_type=9, label=1,
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
  serialized_start=293,
  serialized_end=360,
)


_FETCHACTIVERUNSREQUEST = _descriptor.Descriptor(
  name='FetchActiveRunsRequest',
  full_name='migration.FetchActiveRunsRequest',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  create_key=_descriptor._internal_create_key,
  fields=[
    _descriptor.FieldDescriptor(
      name='luci_project', full_name='migration.FetchActiveRunsRequest.luci_project', index=0,
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
  serialized_start=362,
  serialized_end=408,
)


_FETCHACTIVERUNSRESPONSE = _descriptor.Descriptor(
  name='FetchActiveRunsResponse',
  full_name='migration.FetchActiveRunsResponse',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  create_key=_descriptor._internal_create_key,
  fields=[
    _descriptor.FieldDescriptor(
      name='runs', full_name='migration.FetchActiveRunsResponse.runs', index=0,
      number=1, type=11, cpp_type=10, label=3,
      has_default_value=False, default_value=[],
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
  serialized_start=410,
  serialized_end=465,
)


_RUN = _descriptor.Descriptor(
  name='Run',
  full_name='migration.Run',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  create_key=_descriptor._internal_create_key,
  fields=[
    _descriptor.FieldDescriptor(
      name='attempt', full_name='migration.Run.attempt', index=0,
      number=1, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='id', full_name='migration.Run.id', index=1,
      number=2, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='cls', full_name='migration.Run.cls', index=2,
      number=3, type=11, cpp_type=10, label=3,
      has_default_value=False, default_value=[],
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
  serialized_start=467,
  serialized_end=551,
)


_RUNCL_TRIGGER = _descriptor.Descriptor(
  name='Trigger',
  full_name='migration.RunCL.Trigger',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  create_key=_descriptor._internal_create_key,
  fields=[
    _descriptor.FieldDescriptor(
      name='time', full_name='migration.RunCL.Trigger.time', index=0,
      number=1, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='account_id', full_name='migration.RunCL.Trigger.account_id', index=1,
      number=3, type=3, cpp_type=2, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='email', full_name='migration.RunCL.Trigger.email', index=2,
      number=4, type=9, cpp_type=9, label=1,
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
  serialized_start=740,
  serialized_end=826,
)

_RUNCL_DEP = _descriptor.Descriptor(
  name='Dep',
  full_name='migration.RunCL.Dep',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  create_key=_descriptor._internal_create_key,
  fields=[
    _descriptor.FieldDescriptor(
      name='id', full_name='migration.RunCL.Dep.id', index=0,
      number=1, type=3, cpp_type=2, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='hard', full_name='migration.RunCL.Dep.hard', index=1,
      number=2, type=8, cpp_type=7, label=1,
      has_default_value=False, default_value=False,
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
  serialized_start=828,
  serialized_end=859,
)

_RUNCL = _descriptor.Descriptor(
  name='RunCL',
  full_name='migration.RunCL',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  create_key=_descriptor._internal_create_key,
  fields=[
    _descriptor.FieldDescriptor(
      name='id', full_name='migration.RunCL.id', index=0,
      number=1, type=3, cpp_type=2, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='gc', full_name='migration.RunCL.gc', index=1,
      number=2, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='updated_time', full_name='migration.RunCL.updated_time', index=2,
      number=3, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='trigger', full_name='migration.RunCL.trigger', index=3,
      number=4, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='deps', full_name='migration.RunCL.deps', index=4,
      number=5, type=11, cpp_type=10, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
  ],
  extensions=[
  ],
  nested_types=[_RUNCL_TRIGGER, _RUNCL_DEP, ],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=554,
  serialized_end=859,
)

_REPORTRUNSREQUEST.fields_by_name['runs'].message_type = _RUN
_REPORTFINISHEDRUNREQUEST.fields_by_name['run'].message_type = _RUN
_FETCHACTIVERUNSRESPONSE.fields_by_name['runs'].message_type = _RUN
_RUN.fields_by_name['attempt'].message_type = go_dot_chromium_dot_org_dot_luci_dot_cv_dot_api_dot_bigquery_dot_v1_dot_attempt__pb2._ATTEMPT
_RUN.fields_by_name['cls'].message_type = _RUNCL
_RUNCL_TRIGGER.fields_by_name['time'].message_type = google_dot_protobuf_dot_timestamp__pb2._TIMESTAMP
_RUNCL_TRIGGER.containing_type = _RUNCL
_RUNCL_DEP.containing_type = _RUNCL
_RUNCL.fields_by_name['gc'].message_type = go_dot_chromium_dot_org_dot_luci_dot_cv_dot_api_dot_bigquery_dot_v1_dot_attempt__pb2._GERRITCHANGE
_RUNCL.fields_by_name['updated_time'].message_type = google_dot_protobuf_dot_timestamp__pb2._TIMESTAMP
_RUNCL.fields_by_name['trigger'].message_type = _RUNCL_TRIGGER
_RUNCL.fields_by_name['deps'].message_type = _RUNCL_DEP
DESCRIPTOR.message_types_by_name['ReportRunsRequest'] = _REPORTRUNSREQUEST
DESCRIPTOR.message_types_by_name['ReportFinishedRunRequest'] = _REPORTFINISHEDRUNREQUEST
DESCRIPTOR.message_types_by_name['ReportUsedNetrcRequest'] = _REPORTUSEDNETRCREQUEST
DESCRIPTOR.message_types_by_name['FetchActiveRunsRequest'] = _FETCHACTIVERUNSREQUEST
DESCRIPTOR.message_types_by_name['FetchActiveRunsResponse'] = _FETCHACTIVERUNSRESPONSE
DESCRIPTOR.message_types_by_name['Run'] = _RUN
DESCRIPTOR.message_types_by_name['RunCL'] = _RUNCL
_sym_db.RegisterFileDescriptor(DESCRIPTOR)

ReportRunsRequest = _reflection.GeneratedProtocolMessageType('ReportRunsRequest', (_message.Message,), {
  'DESCRIPTOR' : _REPORTRUNSREQUEST,
  '__module__' : 'go.chromium.org.luci.cv.api.migration.migration_pb2'
  # @@protoc_insertion_point(class_scope:migration.ReportRunsRequest)
  })
_sym_db.RegisterMessage(ReportRunsRequest)

ReportFinishedRunRequest = _reflection.GeneratedProtocolMessageType('ReportFinishedRunRequest', (_message.Message,), {
  'DESCRIPTOR' : _REPORTFINISHEDRUNREQUEST,
  '__module__' : 'go.chromium.org.luci.cv.api.migration.migration_pb2'
  # @@protoc_insertion_point(class_scope:migration.ReportFinishedRunRequest)
  })
_sym_db.RegisterMessage(ReportFinishedRunRequest)

ReportUsedNetrcRequest = _reflection.GeneratedProtocolMessageType('ReportUsedNetrcRequest', (_message.Message,), {
  'DESCRIPTOR' : _REPORTUSEDNETRCREQUEST,
  '__module__' : 'go.chromium.org.luci.cv.api.migration.migration_pb2'
  # @@protoc_insertion_point(class_scope:migration.ReportUsedNetrcRequest)
  })
_sym_db.RegisterMessage(ReportUsedNetrcRequest)

FetchActiveRunsRequest = _reflection.GeneratedProtocolMessageType('FetchActiveRunsRequest', (_message.Message,), {
  'DESCRIPTOR' : _FETCHACTIVERUNSREQUEST,
  '__module__' : 'go.chromium.org.luci.cv.api.migration.migration_pb2'
  # @@protoc_insertion_point(class_scope:migration.FetchActiveRunsRequest)
  })
_sym_db.RegisterMessage(FetchActiveRunsRequest)

FetchActiveRunsResponse = _reflection.GeneratedProtocolMessageType('FetchActiveRunsResponse', (_message.Message,), {
  'DESCRIPTOR' : _FETCHACTIVERUNSRESPONSE,
  '__module__' : 'go.chromium.org.luci.cv.api.migration.migration_pb2'
  # @@protoc_insertion_point(class_scope:migration.FetchActiveRunsResponse)
  })
_sym_db.RegisterMessage(FetchActiveRunsResponse)

Run = _reflection.GeneratedProtocolMessageType('Run', (_message.Message,), {
  'DESCRIPTOR' : _RUN,
  '__module__' : 'go.chromium.org.luci.cv.api.migration.migration_pb2'
  # @@protoc_insertion_point(class_scope:migration.Run)
  })
_sym_db.RegisterMessage(Run)

RunCL = _reflection.GeneratedProtocolMessageType('RunCL', (_message.Message,), {

  'Trigger' : _reflection.GeneratedProtocolMessageType('Trigger', (_message.Message,), {
    'DESCRIPTOR' : _RUNCL_TRIGGER,
    '__module__' : 'go.chromium.org.luci.cv.api.migration.migration_pb2'
    # @@protoc_insertion_point(class_scope:migration.RunCL.Trigger)
    })
  ,

  'Dep' : _reflection.GeneratedProtocolMessageType('Dep', (_message.Message,), {
    'DESCRIPTOR' : _RUNCL_DEP,
    '__module__' : 'go.chromium.org.luci.cv.api.migration.migration_pb2'
    # @@protoc_insertion_point(class_scope:migration.RunCL.Dep)
    })
  ,
  'DESCRIPTOR' : _RUNCL,
  '__module__' : 'go.chromium.org.luci.cv.api.migration.migration_pb2'
  # @@protoc_insertion_point(class_scope:migration.RunCL)
  })
_sym_db.RegisterMessage(RunCL)
_sym_db.RegisterMessage(RunCL.Trigger)
_sym_db.RegisterMessage(RunCL.Dep)


DESCRIPTOR._options = None

_MIGRATION = _descriptor.ServiceDescriptor(
  name='Migration',
  full_name='migration.Migration',
  file=DESCRIPTOR,
  index=0,
  serialized_options=None,
  create_key=_descriptor._internal_create_key,
  serialized_start=862,
  serialized_end=1191,
  methods=[
  _descriptor.MethodDescriptor(
    name='ReportRuns',
    full_name='migration.Migration.ReportRuns',
    index=0,
    containing_service=None,
    input_type=_REPORTRUNSREQUEST,
    output_type=google_dot_protobuf_dot_empty__pb2._EMPTY,
    serialized_options=None,
    create_key=_descriptor._internal_create_key,
  ),
  _descriptor.MethodDescriptor(
    name='ReportFinishedRun',
    full_name='migration.Migration.ReportFinishedRun',
    index=1,
    containing_service=None,
    input_type=_REPORTFINISHEDRUNREQUEST,
    output_type=google_dot_protobuf_dot_empty__pb2._EMPTY,
    serialized_options=None,
    create_key=_descriptor._internal_create_key,
  ),
  _descriptor.MethodDescriptor(
    name='ReportUsedNetrc',
    full_name='migration.Migration.ReportUsedNetrc',
    index=2,
    containing_service=None,
    input_type=_REPORTUSEDNETRCREQUEST,
    output_type=google_dot_protobuf_dot_empty__pb2._EMPTY,
    serialized_options=None,
    create_key=_descriptor._internal_create_key,
  ),
  _descriptor.MethodDescriptor(
    name='FetchActiveRuns',
    full_name='migration.Migration.FetchActiveRuns',
    index=3,
    containing_service=None,
    input_type=_FETCHACTIVERUNSREQUEST,
    output_type=_FETCHACTIVERUNSRESPONSE,
    serialized_options=None,
    create_key=_descriptor._internal_create_key,
  ),
])
_sym_db.RegisterServiceDescriptor(_MIGRATION)

DESCRIPTOR.services_by_name['Migration'] = _MIGRATION

# @@protoc_insertion_point(module_scope)
