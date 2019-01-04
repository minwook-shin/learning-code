import 'package:flutter/material.dart';

void main() => runApp(MaterialApp(
      title: 'fade in/out',
      home: MyApp(),
    ));

class MyApp extends StatefulWidget {
  @override
  _MyAppState createState() => _MyAppState();
}

class _MyAppState extends State<MyApp> {
  bool _visible = true;

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      body: ListView(
        children: [
          AnimatedOpacity(
            opacity: _visible ? 1.0 : 0.1,
            duration: Duration(milliseconds: 500),
            child: Container(
              height: 500,
              color: Colors.blue,
            ),
          ),
          OutlineButton(
            child: Text("fade in/out"),
            onPressed: () {
              setState(() {
                _visible = !_visible;
              });
            },
          )
        ],
      ),
    );
  }
}
