import 'package:flutter/material.dart';

void main() {
  runApp(MaterialApp(
    title: 'Form Focus',
    home: MyApp(),
  ));
}

class MyApp extends StatefulWidget {
  @override
  _MyAppState createState() => _MyAppState();
}

class _MyAppState extends State<MyApp> {
  FocusNode focusNode1;
  FocusNode focusNode2;

  @override
  void initState() {
    super.initState();
    focusNode1 = FocusNode();
    focusNode2 = FocusNode();

  }

  @override
  void dispose() {
    super.dispose();
    focusNode1.dispose();
    focusNode2.dispose();
  }

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      body: Column(
        children: [
          TextField(
            autofocus: true,
            focusNode: focusNode1,
          ),
          TextField(
            focusNode: focusNode2,
          ),
          OutlineButton(
            onPressed: () => FocusScope.of(context).requestFocus(focusNode1),
            child: Text("첫번째 텍스트로 포커스 이동"),
          ),
          OutlineButton(
            onPressed: () => FocusScope.of(context).requestFocus(focusNode2),
            child: Text("두번째 텍스트로 포커스 이동"),
          ),
        ],
      ),
    );
  }
}
