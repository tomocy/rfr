import 'package:flutter/material.dart';
import 'package:grpc/grpc.dart';
import 'package:rfv/infra/rpc/rfv/google/protobuf/empty.pb.dart';
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

  _IndexPageState createState() => _IndexPageState(MockFetcher());
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
        subtitle: Text(rfc.title)
    ));
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
}