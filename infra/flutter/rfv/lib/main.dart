import 'dart:convert';
import 'package:flutter/material.dart';
import 'package:grpc/grpc.dart';
import 'package:rfv/infra/rpc/rfv/google/protobuf/empty.pb.dart';
import 'package:http/http.dart' as http;
import 'infra/rpc/rfv/rfv.pbgrpc.dart';

void main() => runApp(RFV());

class RFV extends StatelessWidget {
  @override
  Widget build(BuildContext context) {
    return MaterialApp(
      title: 'RFV',
      theme: ThemeData(
        brightness: Brightness.dark,
        fontFamily: 'sans-serif',
      ),
      home: IndexPage(title: 'RFCs'),
    );
  }
}

class IndexPage extends StatefulWidget {
  IndexPage({Key key, this.title}) : super(key: key);

  final String title;

  _IndexPageState createState() => _IndexPageState(HTTPFetcher('localhost', 8080));
}

class _IndexPageState extends State<IndexPage> {
  _IndexPageState(this._fetcher);

  Fetcher _fetcher;
  Future<List<RFC>> _rfcs;

  @override
  initState() {
    super.initState();
    _rfcs = _fetcher.fetchIndex();
  }

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(
        title: Text(widget.title),
      ),
      body: FutureBuilder(
        future: _rfcs,
        builder: (context, snapshot) {
          if (snapshot.hasData) {
            Iterable<Widget> rfcs = snapshot.data.map<Widget>((RFC rfc) => buildTile(context, rfc));
            return ListView(children: rfcs.toList());
          }
          if (snapshot.hasError) {
            return Center(child: Text(snapshot.error.toString()));
          }

          return Center(child: CircularProgressIndicator());
        },
      ),
    );
  }

  Widget buildTile(BuildContext context, RFC rfc) {
    return MergeSemantics(child: ListTile(
        isThreeLine: true,
        title: Text('${rfc.id} ${rfc.title}'),
        subtitle: Text(rfc.title),
        onTap: () {
          Navigator.push(
            context,
            MaterialPageRoute(builder: (BuildContext context) => SinglePage(
              title: rfc.id,
              fetcher: _fetcher,
            )),
          );
        },
      )
    );
  }
}

class SinglePage extends StatefulWidget {
  SinglePage({Key key, this.title, this.fetcher}) : super(key: key);

  final String title;
  final Fetcher fetcher;

  @override
  _SinglePageState createState() => _SinglePageState(this.fetcher);
}

class _SinglePageState extends State<SinglePage> {
  _SinglePageState(this._fetcher);

  final Fetcher _fetcher;
  Future<RFC> _rfc;

  @override
  initState() {
    super.initState();
    _rfc = _fetcher.fetch(widget.title.toLowerCase());
  }

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(
        title: Text(widget.title),
      ),
      body: FutureBuilder(
        future: _rfc,
        builder: (BuildContext context, AsyncSnapshot<RFC> snapshot) {
          if (snapshot.hasData) {
            return Padding(
              padding: const EdgeInsets.all(NavigationToolbar.kMiddleSpacing),
              child: Column(
                mainAxisSize: MainAxisSize.min,
                crossAxisAlignment: CrossAxisAlignment.start,
                children: <Widget>[
                  Text(
                    snapshot.data.title,
                    style: Theme.of(context).textTheme.title,
                  ),
                ],
              ),
            );
          }
          if (snapshot.hasError) {
            return Center(child: Text(snapshot.error.toString()));
          }

          return Center(child: CircularProgressIndicator());
        },
      ),
    );
  }
}

abstract class Fetcher {
  Future<List<RFC>> fetchIndex();
  Future<RFC> fetch(String id);
}

class MockFetcher implements Fetcher {
  final List<RFC> _rfcs = [
    RFC('RFC8672', 'TLS Server Identity Pinning with Tickets'),
    RFC('RFC8671', 'Support for Adj-RIB-Out in the BGP Monitoring Protocol (BMP)'),
    RFC('RFC8658', 'RADIUS Attributes for Softwire Mechanisms Based on Address plus Port (A+P)'),
  ];

  Future<List<RFC>> fetchIndex() async {
    return _rfcs;
  }

  Future<RFC> fetch(String id) async {
    return _rfcs.firstWhere((rfc) => rfc.id == id);
  }
}

class GRPCFetcher implements Fetcher {
  GRPCFetcher(String host, int port) {
    _client = EntryRepoClient(ClientChannel(
        host,
        port: port,
        options: ChannelOptions(credentials: ChannelCredentials.insecure()),
    ));
  }

  EntryRepoClient _client;

  Future<List<RFC>> fetchIndex() async {
    Entries entries = await _client.fetchIndex(Empty());
    return entries.entries.map((Entry entry) => RFC.fromEntry(entry));
  }

  Future<RFC> fetch(String id) async {
    Entry entry = await _client.fetch(FetchRequest()..id = id);
    return RFC.fromEntry(entry);
  }
}

class HTTPFetcher implements Fetcher {
  HTTPFetcher(this._host, this._port);

  final String _host;
  final int _port;

  Future<List<RFC>> fetchIndex() async {
    final http.Response response = await http.get(_endpoint());
    if (400 <= response.statusCode) {
      throw Exception(response.reasonPhrase);
    }
    
    final List<Map<String, dynamic>> decoded = (json.decode(response.body) as List<dynamic>).cast<Map<String, dynamic>>();
    return decoded.map((Map<String, dynamic> rfc) => RFC.fromJSON(rfc)).toList();
  }

  Future<RFC> fetch(String id) async {
    final http.Response response = await http.get(_endpoint([id]));
    if (400 <= response.statusCode) {
      throw Exception(response.reasonPhrase);
    }

    return RFC.fromJSON(json.decode(response.body));
  }

  String _endpoint([List<String> paths]) {
    final String joined = paths?.reduce((curr, next) => '$curr/$next') ?? '';
    return 'http://${_address()}/$joined';
  }

  String _address() {
    return '$_host:$_port';
  }
}

class RFC {
  RFC(this.id, this.title);

  String id;
  String title;

  factory RFC.fromEntry(Entry entry) {
    return RFC(
      entry.id,
      entry.title,
    );
  }

  factory RFC.fromJSON(Map<String, dynamic> json) {
    return RFC(
      json['id'],
      json['title'],
    );
  }
}