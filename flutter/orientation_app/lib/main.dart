import 'package:flutter/material.dart';

void main() {
  runApp(MaterialApp(
    title: "orientation app",
    home: MyApp(),
  ));
}

class MyApp extends StatelessWidget {
  @override
  Widget build(BuildContext context) {
    return Scaffold(body: Center(
      child: OrientationBuilder(builder: (context, orientation) {
        return Text(orientation == Orientation.portrait ? "세로" : "가로",
            style: TextStyle(fontSize: 100));
      }),
    ));
  }
}
