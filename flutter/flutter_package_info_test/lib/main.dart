import 'package:flutter/material.dart';
import 'package:package_info/package_info.dart';

void main() => runApp(MyApp());

class MyApp extends StatelessWidget {
  @override
  Widget build(BuildContext context) {
    return MaterialApp(
      title: 'Flutter package_info',
      home: MyHomePage(),
    );
  }
}

class MyHomePage extends StatefulWidget {
  @override
  _MyHomePageState createState() => _MyHomePageState();
}

class _MyHomePageState extends State<MyHomePage> {
  var _packageInfo = PackageInfo();

  @override
  initState() {
    super.initState();
    getPackageInfo();
  }

  getPackageInfo() async {
    var packageInfo = await PackageInfo.fromPlatform();
    setState(() {
      _packageInfo = packageInfo;
    });
  }

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      body: Column(
        children: <Widget>[
          Text(_packageInfo.appName ?? 'Not set'),
          Text(_packageInfo.packageName ?? 'Not set'),
          Text(_packageInfo.version ?? 'Not set'),
          Text(_packageInfo.buildNumber ?? 'Not set'),
        ],
      ),
    );
  }
}
