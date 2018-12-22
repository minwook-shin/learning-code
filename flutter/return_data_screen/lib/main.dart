import 'package:flutter/material.dart';

void main() {
  runApp(MaterialApp(
    title: 'Returning Data',
    home: HomeScreen(),
  ));
}

class HomeScreen extends StatelessWidget {
  final _scaffold = GlobalKey<ScaffoldState>();
  @override
  Widget build(BuildContext context) {
    return Scaffold(
      key: _scaffold,
      body: Center(
        child: OutlineButton(
          onPressed: () {
            _navigateAndDisplaySelection(context);
          },
          child: Text('선택하기'),
        ),
      ),
    );
  }

  _navigateAndDisplaySelection(BuildContext context) async {
    final result = await Navigator.push(
      context,
      MaterialPageRoute(builder: (context) => SelectionScreen()),
    );

    _scaffold.currentState
      ..removeCurrentSnackBar()
      ..showSnackBar(SnackBar(content: Text("$result")));
  }
}


class SelectionScreen extends StatelessWidget {
  @override
  Widget build(BuildContext context) {
    return Scaffold(
      body: Column(
        children: <Widget>[
          OutlineButton(
            onPressed: () {
              Navigator.pop(context, '예');
            },
            child: Text('예'),
          ),
          OutlineButton(
            onPressed: () {
              Navigator.pop(context, '아니오');
            },
            child: Text('아니오'),
          ),
        ],
      ),
    );
  }
}
