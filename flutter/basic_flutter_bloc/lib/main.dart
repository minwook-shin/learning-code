import 'package:flutter/material.dart';
import 'package:bloc/bloc.dart';
import 'package:flutter_bloc/flutter_bloc.dart';

enum CounterEvent { increment, decrement }

class BlocDelegateClass extends BlocDelegate {
  @override
  void onTransition(Transition transition) {}
}

void main() {
  BlocSupervisor().delegate = BlocDelegateClass();
  runApp(MyApp());
}

class MyApp extends StatefulWidget {
  @override
  State<StatefulWidget> createState() => _MyAppState();
}

class CounterBloc extends Bloc<CounterEvent, int> {
  @override
  get initialState => 0;

  @override
  Stream<int> mapEventToState(currentState, event) async* {
    switch (event) {
      case CounterEvent.increment:
        yield currentState + 1;
        break;
      case CounterEvent.decrement:
        yield currentState - 1;
        break;
    }
  }
}

class _MyAppState extends State<MyApp> {
  final _counterBloc = CounterBloc();

  @override
  Widget build(BuildContext context) {
    return MaterialApp(
      title: 'bloc code',
      home: BlocProvider(
        bloc: _counterBloc,
        child: MyappPage(),
      ),
    );
  }
}

class MyappPage extends StatelessWidget {
  @override
  Widget build(BuildContext context) {
    final _counterBloc = BlocProvider.of<CounterBloc>(context);

    return Scaffold(
      body: BlocBuilder(
        bloc: _counterBloc,
        builder: (context, count) {
          return Center(child: Text(count.toString()));
        },
      ),
      floatingActionButton: Column(
        mainAxisAlignment: MainAxisAlignment.end,
        children: [
          FloatingActionButton(
            child: Icon(Icons.add),
            onPressed: () {
              _counterBloc.dispatch(CounterEvent.increment);
            },
          ),
          FloatingActionButton(
            child: Icon(Icons.remove),
            onPressed: () {
              _counterBloc.dispatch(CounterEvent.decrement);
            },
          ),
        ],
      ),
    );
  }
}
