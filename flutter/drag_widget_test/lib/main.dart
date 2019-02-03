import 'package:flutter/material.dart';

void main() => runApp(MyApp());

class MyApp extends StatelessWidget {
  @override
  Widget build(BuildContext context) {
    return MaterialApp(
        title: 'Flutter Demo', 
        home: Scaffold(body: MyAppStateful())
        );
  }
}

class MyAppStateful extends StatefulWidget {
  @override
  _MyAppState createState() => _MyAppState();
}

class _MyAppState extends State<MyAppStateful> {
  var list = List.generate(100, (i) => i.toString());
  @override
  Widget build(BuildContext context) {
    return ListView.builder(
      itemBuilder: (context, i) {
        return Dismissible(
            onDismissed: (d) {
              setState(() {
                list.removeAt(i);
              });
            },
            child: ListTile(title: Center(child: Text(list[i].toString()))),
            key: Key(list[i]),
            background: Container(color: Colors.amber));
      },
      itemCount: list.length,
    );
  }
}
