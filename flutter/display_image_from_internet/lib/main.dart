import 'package:flutter/material.dart';

void main() {
  runApp(MaterialApp(
      title: 'Images from internet', 
      home: Scaffold(body: MyApp())));
}

class MyApp extends StatelessWidget {
  @override
  Widget build(BuildContext context) {
    return ListView(
      children: <Widget>[
        Image.network(
          'https://picsum.photos/250?image=9',
          fit: BoxFit.fill,
        ),
        Image.network(
          'https://github.com/flutter/plugins/raw/master/packages/video_player/doc/demo_ipod.gif?raw=true',
          fit: BoxFit.fill,
        ),
      ],
    );
  }
}
