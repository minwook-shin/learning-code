import 'package:flutter/material.dart';
import 'package:shared_preferences/shared_preferences.dart';

void main() {
  runApp(
      MaterialApp(title: 'Shared preferences', home: Scaffold(body: MyApp())));
}

class MyApp extends StatefulWidget {
  @override
  MyAppState createState() => MyAppState();
}

class MyAppState extends State<MyApp> {
  int _counter;

  @override
  void initState() {
    super.initState();
    (() async {
      SharedPreferences prefs = await SharedPreferences.getInstance();
      setState(() {
        _counter = prefs.getInt('counter');
      });
    })();
  }

  @override
  Widget build(BuildContext context) {
    return ListView(
      children: [
        Text(_counter.toString(), textAlign: TextAlign.center),
        OutlineButton(
            onPressed: () async {
              SharedPreferences prefs = await SharedPreferences.getInstance();
              setState(() {
                _counter++;
                prefs.setInt('counter', _counter);
              });
            },
            child: Icon(Icons.add)),
        OutlineButton(
            onPressed: () async {
              SharedPreferences prefs = await SharedPreferences.getInstance();
              setState(() {
                _counter--;
                prefs.setInt('counter', _counter);
              });
            },
            child: Icon(Icons.remove)),
      ],
    );
  }
}
