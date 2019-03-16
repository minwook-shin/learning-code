import 'package:flutter/material.dart';
import 'package:path_provider/path_provider.dart';
import 'dart:io';

void main() => runApp(MyApp());

class MyApp extends StatelessWidget {
  @override
  Widget build(BuildContext context) {
    return MaterialApp(
      title: 'Flutter path_provider',
      home: MyHomePage(),
    );
  }
}

class MyHomePage extends StatefulWidget {
  @override
  _MyHomePageState createState() => _MyHomePageState();
}

class _MyHomePageState extends State<MyHomePage> {
  Future<Directory> _tempDir;
  Future<Directory> _appDir;
  Future<Directory> _externalDir;

  @override
  void initState() {
    super.initState();
    _requestTempDir();
    _requestAppDir();
    _requestExternalDir();
  }

  void _requestTempDir() {
    setState(() {
      _tempDir = getTemporaryDirectory();
    });
  }

  void _requestAppDir() {
    setState(() {
      _appDir = getApplicationDocumentsDirectory();
    });
  }

  void _requestExternalDir() {
    setState(() {
      _externalDir = getExternalStorageDirectory();
    });
  }

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      body: Column(
        mainAxisAlignment: MainAxisAlignment.center,
        children: <Widget>[
          FutureBuilder<Directory>(
            future: _tempDir,
            builder: (BuildContext context, AsyncSnapshot<Directory> snapshot) {
              if (snapshot.connectionState == ConnectionState.done) {
                if (snapshot.hasData) {
                  return Text('path: ${snapshot.data.path}');
                } else {
                  return const Text('path unavailable');
                }
              }
            },
          ),
          FutureBuilder<Directory>(
            future: _appDir,
            builder: (BuildContext context, AsyncSnapshot<Directory> snapshot) {
              if (snapshot.connectionState == ConnectionState.done) {
                if (snapshot.hasData) {
                  return Text('path: ${snapshot.data.path}');
                } else {
                  return const Text('path unavailable');
                }
              }
            },
          ),
          FutureBuilder<Directory>(
            future: _externalDir,
            builder: (BuildContext context, AsyncSnapshot<Directory> snapshot) {
              if (snapshot.connectionState == ConnectionState.done) {
                if (snapshot.hasData) {
                  return Text('path: ${snapshot.data.path}');
                } else {
                  return const Text('path unavailable');
                }
              }
            },
          ),
        ],
      ),
    );
  }
}
