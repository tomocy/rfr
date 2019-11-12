///
//  Generated code. Do not modify.
//  source: rfv.proto
//
// @dart = 2.3
// ignore_for_file: camel_case_types,non_constant_identifier_names,library_prefixes,unused_import,unused_shown_name,return_of_invalid_type

import 'dart:core' as $core;

import 'package:protobuf/protobuf.dart' as $pb;

class FetchRequest extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo('FetchRequest', package: const $pb.PackageName('rfv'), createEmptyInstance: create)
    ..aOS(1, 'id')
    ..hasRequiredFields = false
  ;

  FetchRequest._() : super();
  factory FetchRequest() => create();
  factory FetchRequest.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory FetchRequest.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  FetchRequest clone() => FetchRequest()..mergeFromMessage(this);
  FetchRequest copyWith(void Function(FetchRequest) updates) => super.copyWith((message) => updates(message as FetchRequest));
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static FetchRequest create() => FetchRequest._();
  FetchRequest createEmptyInstance() => create();
  static $pb.PbList<FetchRequest> createRepeated() => $pb.PbList<FetchRequest>();
  @$core.pragma('dart2js:noInline')
  static FetchRequest getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<FetchRequest>(create);
  static FetchRequest _defaultInstance;

  @$pb.TagNumber(1)
  $core.String get id => $_getSZ(0);
  @$pb.TagNumber(1)
  set id($core.String v) { $_setString(0, v); }
  @$pb.TagNumber(1)
  $core.bool hasId() => $_has(0);
  @$pb.TagNumber(1)
  void clearId() => clearField(1);
}

class Entries extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo('Entries', package: const $pb.PackageName('rfv'), createEmptyInstance: create)
    ..pc<Entry>(1, 'entries', $pb.PbFieldType.PM, subBuilder: Entry.create)
    ..hasRequiredFields = false
  ;

  Entries._() : super();
  factory Entries() => create();
  factory Entries.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory Entries.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  Entries clone() => Entries()..mergeFromMessage(this);
  Entries copyWith(void Function(Entries) updates) => super.copyWith((message) => updates(message as Entries));
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static Entries create() => Entries._();
  Entries createEmptyInstance() => create();
  static $pb.PbList<Entries> createRepeated() => $pb.PbList<Entries>();
  @$core.pragma('dart2js:noInline')
  static Entries getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<Entries>(create);
  static Entries _defaultInstance;

  @$pb.TagNumber(1)
  $core.List<Entry> get entries => $_getList(0);
}

class Entry extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo('Entry', package: const $pb.PackageName('rfv'), createEmptyInstance: create)
    ..aOS(1, 'id')
    ..aOS(2, 'title')
    ..hasRequiredFields = false
  ;

  Entry._() : super();
  factory Entry() => create();
  factory Entry.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory Entry.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  Entry clone() => Entry()..mergeFromMessage(this);
  Entry copyWith(void Function(Entry) updates) => super.copyWith((message) => updates(message as Entry));
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static Entry create() => Entry._();
  Entry createEmptyInstance() => create();
  static $pb.PbList<Entry> createRepeated() => $pb.PbList<Entry>();
  @$core.pragma('dart2js:noInline')
  static Entry getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<Entry>(create);
  static Entry _defaultInstance;

  @$pb.TagNumber(1)
  $core.String get id => $_getSZ(0);
  @$pb.TagNumber(1)
  set id($core.String v) { $_setString(0, v); }
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

