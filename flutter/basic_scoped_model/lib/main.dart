import 'dart:async';

import 'package:flutter/material.dart';
import 'package:scoped_model/scoped_model.dart';

class CounterModel extends Model {
  var value = 0;

  increment() {
    value += 1;
    notifyListeners();
  }

  decrement() {
    value -= 1;
    notifyListeners();
  }

  static CounterModel of(BuildContext context) =>
      ScopedModel.of<CounterModel>(context);
}

void main() {
  var counter = CounterModel();
  runApp(ScopedModel<CounterModel>(model: counter, child: MyApp()));
  Timer.periodic(Duration(seconds: 5), (timer) => counter.increment());
}

class MyApp extends StatelessWidget {
  @override
  Widget build(BuildContext context) {
    return MaterialApp(title: 'Demo', home: MyHomePage());
  }
}

class MyHomePage extends StatelessWidget {
  @override
  Widget build(BuildContext context) {
    return Scaffold(
        body: Center(
            child: ScopedModelDescendant<CounterModel>(
                builder: (context, child, counter) =>
                    Text(counter.value.toString()))),
        floatingActionButton: FloatingActionButton(
            onPressed: () => CounterModel.of(context).decrement(),
            tooltip: 'Increment',
            child: Icon(Icons.remove)));
  }
}
