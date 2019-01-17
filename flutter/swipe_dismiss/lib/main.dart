import 'package:flutter/material.dart';

void main() {
  runApp(MaterialApp(
      title: 'Swipe Dismiss',
      home: Scaffold(body:MyApp())));
}

class MyApp extends StatefulWidget {
  @override
  MyAppState createState() => MyAppState();
}

class MyAppState extends State<MyApp> {
  final items = List.generate(10, (i) => "리스트 ${i + 1}");

  @override
  Widget build(BuildContext context) {
    return  ListView.builder(
          itemCount: items.length,
          itemBuilder: (context, index) {
            return Dismissible(
              key: Key(items[index]),
              onDismissed: (direction) {
                setState(() {
                  items.removeAt(index);
                });
              },
              child: ListTile(title: Text(items[index])),
              background: Container(
                  alignment: Alignment.center,
                  color: Colors.red,
                  child: Text(
                    "삭제",
                    style: TextStyle(fontSize: 50),
                  )),
            );
          },
        );
  }
}
