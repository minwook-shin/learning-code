import 'dart:io';
import 'package:flutter/material.dart';
import 'package:path_provider/path_provider.dart';

void main() {
  runApp(
    MaterialApp(title: 'Reading Writing Files', home: Scaffold(body: MyApp())),
  );
}

class MyApp extends StatefulWidget {
  @override
  _MyAppState createState() => _MyAppState();
}

class _MyAppState extends State<MyApp> {
  int _counter;

  Future read() async {
    try {
      final directory = await getApplicationDocumentsDirectory();
      return int.parse(
          await File(directory.path + '/counter.txt').readAsString());
    } catch (e) {
      return 0;
    }
  }

  Future write(int counter) async {
    final directory = await getApplicationDocumentsDirectory();
    return File(directory.path + '/counter.txt')
        .writeAsString(counter.toString());
  }

  @override
  void initState() {
    super.initState();
    read().then((value) {
      setState(() {
        _counter = value;
      });
    });
  }

  @override
  Widget build(BuildContext context) {
    return ListView(children: [
      Text(_counter.toString(), textAlign: TextAlign.center),
      OutlineButton(
          onPressed: () {
            setState(() {
              _counter++;
            });
            return write(_counter);
          },
          child: Icon(Icons.add)),
      OutlineButton(
          onPressed: () {
            setState(() {
              _counter--;
            });
            return write(_counter);
          },
          child: Icon(Icons.remove))
    ]);
  }
}
