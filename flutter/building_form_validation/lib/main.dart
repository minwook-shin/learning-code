import 'package:flutter/material.dart';

void main() => runApp(MaterialApp(
      title: 'Form Validation',
      home: Scaffold(
        body: MyApp(),
      ),
    ));

class MyApp extends StatefulWidget {
  @override
  MyAppState createState() => MyAppState();
}

class MyAppState extends State<MyApp> {
  final _key = GlobalKey<FormState>();
  @override
  Widget build(BuildContext context) {
    return Form(
      key: _key,
      child: Column(
        children: [
          TextFormField(
            validator: (value) {
              if (value.isEmpty) {
                return '텍스트 필드에 작성된 문자열이 없습니다.';
              }
            },
          ),
          OutlineButton(
            onPressed: () {
              _key.currentState.validate();
            },
            child: Text("버튼"),
          ),
        ],
      ),
    );
  }
}
