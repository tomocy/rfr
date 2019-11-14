///
//  Generated code. Do not modify.
//  source: rfv.proto
//
// @dart = 2.3
// ignore_for_file: camel_case_types,non_constant_identifier_names,library_prefixes,unused_import,unused_shown_name,return_of_invalid_type

import 'dart:async' as $async;

import 'dart:core' as $core;

import 'package:grpc/service_api.dart' as $grpc;
import 'google/protobuf/empty.pb.dart' as $0;
import 'rfv.pb.dart' as $1;
export 'rfv.pb.dart';

class RFCRepoClient extends $grpc.Client {
  static final _$get = $grpc.ClientMethod<$0.Empty, $1.RFCs>(
      '/rfv.RFCRepo/Get',
      ($0.Empty value) => value.writeToBuffer(),
      ($core.List<$core.int> value) => $1.RFCs.fromBuffer(value));
  static final _$find = $grpc.ClientMethod<$1.FindRequest, $1.RFC>(
      '/rfv.RFCRepo/Find',
      ($1.FindRequest value) => value.writeToBuffer(),
      ($core.List<$core.int> value) => $1.RFC.fromBuffer(value));

  RFCRepoClient($grpc.ClientChannel channel, {$grpc.CallOptions options})
      : super(channel, options: options);

  $grpc.ResponseFuture<$1.RFCs> get($0.Empty request,
      {$grpc.CallOptions options}) {
    final call = $createCall(_$get, $async.Stream.fromIterable([request]),
        options: options);
    return $grpc.ResponseFuture(call);
  }

  $grpc.ResponseFuture<$1.RFC> find($1.FindRequest request,
      {$grpc.CallOptions options}) {
    final call = $createCall(_$find, $async.Stream.fromIterable([request]),
        options: options);
    return $grpc.ResponseFuture(call);
  }
}

abstract class RFCRepoServiceBase extends $grpc.Service {
  $core.String get $name => 'rfv.RFCRepo';

  RFCRepoServiceBase() {
    $addMethod($grpc.ServiceMethod<$0.Empty, $1.RFCs>(
        'Get',
        get_Pre,
        false,
        false,
        ($core.List<$core.int> value) => $0.Empty.fromBuffer(value),
        ($1.RFCs value) => value.writeToBuffer()));
    $addMethod($grpc.ServiceMethod<$1.FindRequest, $1.RFC>(
        'Find',
        find_Pre,
        false,
        false,
        ($core.List<$core.int> value) => $1.FindRequest.fromBuffer(value),
        ($1.RFC value) => value.writeToBuffer()));
  }

  $async.Future<$1.RFCs> get_Pre(
      $grpc.ServiceCall call, $async.Future<$0.Empty> request) async {
    return get(call, await request);
  }

  $async.Future<$1.RFC> find_Pre(
      $grpc.ServiceCall call, $async.Future<$1.FindRequest> request) async {
    return find(call, await request);
  }

  $async.Future<$1.RFCs> get($grpc.ServiceCall call, $0.Empty request);
  $async.Future<$1.RFC> find($grpc.ServiceCall call, $1.FindRequest request);
}
