import 'package:flutter/material.dart';

void main() {
  runApp(MaterialApp(title: 'Grid List', home: Scaffold(body: MyApp())));
}

class MyApp extends StatelessWidget {
  final items = List.generate(
      100,
      (i) => Text(
            i.toString(),
          ));

  @override
  Widget build(BuildContext context) {
    return GridView.count(
      crossAxisCount: 5,
      children: items,
      mainAxisSpacing: 30.0,
      crossAxisSpacing: 30.0,
      reverse: false,
    );
  }
}
