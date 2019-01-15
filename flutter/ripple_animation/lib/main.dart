import 'package:flutter/material.dart';

void main() {
  runApp(MaterialApp(
    title: 'InkWell Demo',
    home: Scaffold(
      body: Center(child: InkWellButton()),
    ),
  ));
}

class InkWellButton extends StatelessWidget {
  @override
  Widget build(BuildContext context) {
    return InkWell(
      onTap: () {
        print("탭");
      },
      child: Container(
        padding: EdgeInsets.all(50.0),
        child: Text('Flat Button'),
      ),
    );
  }
}
