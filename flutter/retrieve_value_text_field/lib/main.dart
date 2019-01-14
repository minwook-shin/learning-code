import 'package:flutter/material.dart';

void main() {
  runApp(MaterialApp(
    title: 'Retrieve Text',
    home: MyApp(),
  ));
}

class MyApp extends StatefulWidget {
  @override
  _MyAppState createState() => _MyAppState();
}

class _MyAppState extends State<MyApp> {
  TextEditingController _textController;

  @override
  void initState() {
    super.initState();
    _textController = TextEditingController();
  }

  @override
  void dispose() {
    super.dispose();
    _textController.dispose();
  }

  @override
  Widget build(BuildContext context) {
    return Scaffold(
        body: Column(
      children: <Widget>[
        TextField(
          controller: _textController,
        ),
        OutlineButton(
          onPressed: () {
            print(_textController.text);
          },
        )
      ],
    ));
  }
}
