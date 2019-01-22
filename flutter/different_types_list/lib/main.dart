import 'package:flutter/material.dart';

void main() {
  runApp(MaterialApp(title: 'different List', home: Scaffold(body: MyApp())));
}

class MyApp extends StatelessWidget {
  final List<Widget> items = List.generate(
    100,
    (i) => i % 3 == 0 ? Container() : Card(),
  );

  @override
  Widget build(BuildContext context) {
    return ListView.builder(
      itemCount: items.length,
      itemBuilder: (context, index) {
        if (items[index] is Container) {
          return ListTile(title: Text(items[index].toString()));
        } else if (items[index] is Card) {
          return ListTile(title: Text(items[index].toString()));
        }
      },
    );
  }
}
