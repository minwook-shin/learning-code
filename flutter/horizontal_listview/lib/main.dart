import 'package:flutter/material.dart';

void main() {
  runApp(MaterialApp(title: 'Horizontal List', home: Scaffold(body: MyApp())));
}

class MyApp extends StatelessWidget {
  @override
  Widget build(BuildContext context) {
    return ListView(
      scrollDirection: Axis.horizontal,
      children: [
        Card(
          child: Text("#FF0000"),
          margin: EdgeInsets.all(30),
          color: Colors.red,
        ),
        Card(
          child: Text("#FFA500"),
          margin: EdgeInsets.all(30),
          color: Colors.orange,
        ),
        Card(
          child: Text("#FFFF00"),
          margin: EdgeInsets.all(30),
          color: Colors.yellow,
        ),
        Card(
          child: Text("#008000"),
          margin: EdgeInsets.all(30),
          color: Colors.green,
        ),
        Card(
          child: Text("#0000FF"),
          margin: EdgeInsets.all(30),
          color: Colors.blue,
        ),
        Card(
          child: Text("#4B0082"),
          margin: EdgeInsets.all(30),
          color: Colors.indigo,
        ),
        Card(
          child: Text("#800080"),
          margin: EdgeInsets.all(30),
          color: Colors.purple,
        ),
      ],
    );
  }
}
