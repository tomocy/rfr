///
//  Generated code. Do not modify.
//  source: rfv.proto
//
// @dart = 2.3
// ignore_for_file: camel_case_types,non_constant_identifier_names,library_prefixes,unused_import,unused_shown_name,return_of_invalid_type

import 'dart:core' as $core;

import 'package:protobuf/protobuf.dart' as $pb;

class FindRequest extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo('FindRequest', package: const $pb.PackageName('rfv'), createEmptyInstance: create)
    ..a<$core.int>(1, 'id', $pb.PbFieldType.O3)
    ..hasRequiredFields = false
  ;

  FindRequest._() : super();
  factory FindRequest() => create();
  factory FindRequest.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory FindRequest.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  FindRequest clone() => FindRequest()..mergeFromMessage(this);
  FindRequest copyWith(void Function(FindRequest) updates) => super.copyWith((message) => updates(message as FindRequest));
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static FindRequest create() => FindRequest._();
  FindRequest createEmptyInstance() => create();
  static $pb.PbList<FindRequest> createRepeated() => $pb.PbList<FindRequest>();
  @$core.pragma('dart2js:noInline')
  static FindRequest getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<FindRequest>(create);
  static FindRequest _defaultInstance;

  @$pb.TagNumber(1)
  $core.int get id => $_getIZ(0);
  @$pb.TagNumber(1)
  set id($core.int v) { $_setSignedInt32(0, v); }
  @$pb.TagNumber(1)
  $core.bool hasId() => $_has(0);
  @$pb.TagNumber(1)
  void clearId() => clearField(1);
}

class RFCs extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo('RFCs', package: const $pb.PackageName('rfv'), createEmptyInstance: create)
    ..pc<RFC>(1, 'rfcs', $pb.PbFieldType.PM, subBuilder: RFC.create)
    ..hasRequiredFields = false
  ;

  RFCs._() : super();
  factory RFCs() => create();
  factory RFCs.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory RFCs.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  RFCs clone() => RFCs()..mergeFromMessage(this);
  RFCs copyWith(void Function(RFCs) updates) => super.copyWith((message) => updates(message as RFCs));
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static RFCs create() => RFCs._();
  RFCs createEmptyInstance() => create();
  static $pb.PbList<RFCs> createRepeated() => $pb.PbList<RFCs>();
  @$core.pragma('dart2js:noInline')
  static RFCs getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<RFCs>(create);
  static RFCs _defaultInstance;

  @$pb.TagNumber(1)
  $core.List<RFC> get rfcs => $_getList(0);
}

class RFC extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo('RFC', package: const $pb.PackageName('rfv'), createEmptyInstance: create)
    ..a<$core.int>(1, 'id', $pb.PbFieldType.O3)
    ..aOS(2, 'title')
    ..hasRequiredFields = false
  ;

  RFC._() : super();
  factory RFC() => create();
  factory RFC.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory RFC.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  RFC clone() => RFC()..mergeFromMessage(this);
  RFC copyWith(void Function(RFC) updates) => super.copyWith((message) => updates(message as RFC));
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static RFC create() => RFC._();
  RFC createEmptyInstance() => create();
  static $pb.PbList<RFC> createRepeated() => $pb.PbList<RFC>();
  @$core.pragma('dart2js:noInline')
  static RFC getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<RFC>(create);
  static RFC _defaultInstance;

  @$pb.TagNumber(1)
  $core.int get id => $_getIZ(0);
  @$pb.TagNumber(1)
  set id($core.int v) { $_setSignedInt32(0, v); }
  @$pb.TagNumber(1)
  $core.bool hasId() => $_has(0);
  @$pb.TagNumber(1)
  void clearId() => clearField(1);

  @$pb.TagNumber(2)
  $core.String get title => $_getSZ(1);
  @$pb.TagNumber(2)
  set title($core.String v) { $_setString(1, v); }
  @$pb.TagNumber(2)
  $core.bool hasTitle() => $_has(1);
  @$pb.TagNumber(2)
  void clearTitle() => clearField(2);
}

