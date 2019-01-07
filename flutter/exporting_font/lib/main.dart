import 'package:flutter/material.dart';

void main() => runApp(MaterialApp(
      title: 'Fonts',
      home: MyApp(),
    ));

class MyApp extends StatelessWidget {
  @override
  Widget build(BuildContext context) {
    return Scaffold(
        body: ListView(
      children: <Widget>[
        Text(
          '구글 폰트 사이트에서 받은 나눔 고딕입니다.',
          style: TextStyle(
            fontFamily: 'Nanum Gothic',
            fontStyle: FontStyle.normal,
          ),
        ),
        Text(
          '구글 폰트 사이트에서 받은 나눔 고딕 볼드체입니다.',
          style: TextStyle(
            fontFamily: 'Nanum Gothic bold',
            fontStyle: FontStyle.normal,
          ),
        ),
        Text(
          '기본입니다.',
        ),
      ],
    ));
  }
}
