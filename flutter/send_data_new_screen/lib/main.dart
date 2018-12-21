import 'package:flutter/material.dart';

class Thread {
  Thread(this.title, this.body);

  final title;
  final body;
}

void main() {
  runApp(MaterialApp(
    title: 'Passing Data',
    home: MainScreen(
      array: List.generate(
        10,
        (i) => Thread(
              '$i',
              '$i 번째 리스트입니다.',
            ),
      ),
    ),
  ));
}

class MainScreen extends StatelessWidget {
  final List<Thread> array;

  MainScreen({Key key, @required this.array}) : super(key: key);

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      body: ListView.builder(
        itemCount: array.length,
        itemBuilder: (context, index) {
          return ListTile(
            title: Text(array[index].title),
            onLongPress: () {
              Navigator.push(
                context,
                MaterialPageRoute(
                  builder: (context) => DetailScreen(t: array[index]),
                ),
              );
            },
          );
        },
      ),
    );
  }
}

class DetailScreen extends StatelessWidget {
  final t;

  DetailScreen({Key key, @required this.t}) : super(key: key);

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(
        title: Text("${t.title}"),
      ),
      body: Text('${t.body}'),
    );
  }
}
