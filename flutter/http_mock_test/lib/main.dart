import 'dart:convert';
import 'package:flutter/material.dart';
import 'package:http/http.dart' as http;

class TextClass {
  String title;
  TextClass({this.title});
}

void main() {
  runApp(MaterialApp(title: 'Fetch Data', home: Scaffold(body: MyApp())));
}

Future<TextClass> fetchPost(http.Client client) async {
  final res = await http.get('https://postman-echo.com/response-headers?title=blog_post');
  return TextClass(title: json.decode(res.body)['title']);
}

final mainClient = http.Client();

class MyApp extends StatelessWidget {
  @override
  Widget build(BuildContext context) {
    return Center(
      child: FutureBuilder<TextClass>(
        future: fetchPost(mainClient),
        builder: (context, snapshot) {
          if (snapshot.hasData) {
            return Text(snapshot.data.title);
          } else if (snapshot.hasError) {
            return Text(snapshot.error.toString());
          }
          return LinearProgressIndicator();
        },
      ),
    );
  }
}
