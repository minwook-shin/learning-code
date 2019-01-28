import 'package:flutter/material.dart';

void main() {
  runApp(MaterialApp(title: 'List test code', home: MyApp()));
}

class MyApp extends StatelessWidget {
  final List<String> items = List<String>.generate(100, (i) => i.toString());

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      body: ListView.builder(
        key: Key('listview'),
        itemCount: items.length,
        itemBuilder: (context, index) {
          return ListTile(
            onTap: () {
              print("click test");
            },
            title: Text(items[index].toString(), key: Key(index.toString())),
          );
        },
      ),
    );
  }
}
