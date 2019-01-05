import 'package:flutter/material.dart';

void main() => runApp(MaterialApp(
      title: 'SnackBar',
      home: Scaffold(
        body: MyApp(),
      ),
    ));

class MyApp extends StatelessWidget {
  final snackBar = SnackBar(
    content: Text('SnackBar!'),
    action: SnackBarAction(
      label: 'Undo',
      onPressed: () {
      },
    ),

  );
  @override
  Widget build(BuildContext context) {
    return 
      OutlineButton(
        onPressed: () {
          Scaffold.of(context).showSnackBar(snackBar);
        },
        child: Text('SnackBar'),
      );
  }
}
