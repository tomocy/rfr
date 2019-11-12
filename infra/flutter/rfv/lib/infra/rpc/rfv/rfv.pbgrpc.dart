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

class EntryRepoClient extends $grpc.Client {
  static final _$fetchIndex = $grpc.ClientMethod<$0.Empty, $1.Entries>(
      '/rfv.EntryRepo/FetchIndex',
      ($0.Empty value) => value.writeToBuffer(),
      ($core.List<$core.int> value) => $1.Entries.fromBuffer(value));
  static final _$fetch = $grpc.ClientMethod<$1.FetchRequest, $1.Entry>(
      '/rfv.EntryRepo/Fetch',
      ($1.FetchRequest value) => value.writeToBuffer(),
      ($core.List<$core.int> value) => $1.Entry.fromBuffer(value));

  EntryRepoClient($grpc.ClientChannel channel, {$grpc.CallOptions options})
      : super(channel, options: options);

  $grpc.ResponseFuture<$1.Entries> fetchIndex($0.Empty request,
      {$grpc.CallOptions options}) {
    final call = $createCall(
        _$fetchIndex, $async.Stream.fromIterable([request]),
        options: options);
    return $grpc.ResponseFuture(call);
  }

  $grpc.ResponseFuture<$1.Entry> fetch($1.FetchRequest request,
      {$grpc.CallOptions options}) {
    final call = $createCall(_$fetch, $async.Stream.fromIterable([request]),
        options: options);
    return $grpc.ResponseFuture(call);
  }
}

abstract class EntryRepoServiceBase extends $grpc.Service {
  $core.String get $name => 'rfv.EntryRepo';

  EntryRepoServiceBase() {
    $addMethod($grpc.ServiceMethod<$0.Empty, $1.Entries>(
        'FetchIndex',
        fetchIndex_Pre,
        false,
        false,
        ($core.List<$core.int> value) => $0.Empty.fromBuffer(value),
        ($1.Entries value) => value.writeToBuffer()));
    $addMethod($grpc.ServiceMethod<$1.FetchRequest, $1.Entry>(
        'Fetch',
        fetch_Pre,
        false,
        false,
        ($core.List<$core.int> value) => $1.FetchRequest.fromBuffer(value),
        ($1.Entry value) => value.writeToBuffer()));
  }

  $async.Future<$1.Entries> fetchIndex_Pre(
      $grpc.ServiceCall call, $async.Future<$0.Empty> request) async {
    return fetchIndex(call, await request);
  }

  $async.Future<$1.Entry> fetch_Pre(
      $grpc.ServiceCall call, $async.Future<$1.FetchRequest> request) async {
    return fetch(call, await request);
  }

  $async.Future<$1.Entries> fetchIndex(
      $grpc.ServiceCall call, $0.Empty request);
  $async.Future<$1.Entry> fetch(
      $grpc.ServiceCall call, $1.FetchRequest request);
}
