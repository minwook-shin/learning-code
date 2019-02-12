import 'package:flutter/material.dart';
import 'package:mvc_pattern/mvc_pattern.dart';

class Model {
  static int _counter = 0;
  static int _incrementCounter() => ++_counter;
  static int _decrementCounter() => --_counter;
}

class Controller extends ControllerMVC {
  static get counter => Model._counter;
  static void incrementController() => Model._incrementCounter();
  static void decrementController() => Model._decrementCounter();
}

void main() => runApp(MyApp());

class MyApp extends AppMVC {
  @override
  Widget build(BuildContext context) {
    return MaterialApp(
      title: 'Flutter mvc',
      home: View(),
    );
  }
}

class View extends StatefulWidget {
  @override
  State createState() => ViewState();
}

class ViewState extends StateMVC<View> {
  ViewState() : super(Controller());

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      body: Center(
        child: Column(
          mainAxisAlignment: MainAxisAlignment.center,
          children: <Widget>[
            Text(Controller.counter.toString()),
            OutlineButton(
                onPressed: () {
                  setState(() {
                    Controller.incrementController();
                  });
                },
                child: Icon(Icons.add)),
            OutlineButton(
                onPressed: () {
                  setState(() {
                    Controller.decrementController();
                  });
                },
                child: Icon(Icons.remove))
          ],
        ),
      ),
    );
  }
}
