import 'package:flutter/material.dart';

void main() {
  runApp(MaterialApp(
    title: 'Gesture',
    home: Scaffold(
      body: Center(child:  GestureDetectorButton()),
    ),
  ));
}

class  GestureDetectorButton extends StatelessWidget {
  @override
  Widget build(BuildContext context) {
    return GestureDetector(
      onTap: () {
        print("탭");
      },
      child: Container(
        padding: EdgeInsets.all(100.0),
        child: Text('버튼'),
        decoration: BoxDecoration(
          color: Colors.amber,
          borderRadius: BorderRadius.circular(100.0),
        ),
      ),
    );
  }
}
