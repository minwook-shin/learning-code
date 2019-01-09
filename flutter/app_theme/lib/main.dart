import 'package:flutter/material.dart';

void main() {
  runApp(MaterialApp(
    title: "app theme",
    home: MyApp(),
    theme: ThemeData(
      brightness: Brightness.light,
      primaryColor: Colors.blue,
      accentColor: Colors.black12,
      textTheme: TextTheme(
          headline: TextStyle(fontSize: 72.0, fontWeight: FontWeight.bold),
          title: TextStyle(fontSize: 36.0, fontStyle: FontStyle.italic),
          body1: TextStyle(fontSize: 14.0),
          display1: TextStyle(fontSize: 50.0)),
    ),
  ));
}

class MyApp extends StatelessWidget {
  @override
  Widget build(BuildContext context) {
    return Scaffold(
      body: Center(
          child: ListView(
        children: [
          Container(
            color: Theme.of(context).accentColor,
            child: Text(
              '헤드라인',
              style: Theme.of(context).textTheme.headline,
            ),
          ),
          Container(
            child: Text(
              '타이틀',
              style: Theme.of(context).textTheme.title,
            ),
          ),
          Container(
            child: Text(
              '바디',
              style: Theme.of(context).textTheme.body1,
            ),
          ),
          Container(
            child: Text(
              '디스플레이',
              style: Theme.of(context).textTheme.display1,
            ),
          ),
        ],
      )),
    );
  }
}
