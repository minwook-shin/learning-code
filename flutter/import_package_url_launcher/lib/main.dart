import 'package:flutter/material.dart';
import 'package:url_launcher/url_launcher.dart';

void main() {
  runApp(MaterialApp(
    home: MainPage(),
  ));
}

class MainPage extends StatelessWidget {
  @override
  Widget build(BuildContext context) {
    return Scaffold(
      body: OutlineButton(
          onPressed: () {
            launch('https://www.google.com');
          },
          child: Text('홈페이지 링크'),
        ),
    );
  }
}
