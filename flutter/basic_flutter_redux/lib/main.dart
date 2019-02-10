import 'package:flutter/material.dart';
import 'package:flutter_redux/flutter_redux.dart';
import 'package:redux/redux.dart';

enum Actions { Increment, Decrement }

int counter(state, action) {
  switch (action) {
    case Actions.Increment:
      return state + 1;
      break;
    case Actions.Decrement:
      return state - 1;
      break;
    default:
      return 0;
  }
}

void main() {
  runApp(MaterialApp(title: 'flutter redux', home: FlutterReduxApp()));
}

class FlutterReduxApp extends StatelessWidget {
  final store = Store(counter, initialState: 0);

  @override
  Widget build(BuildContext context) {
    return StoreProvider(
      store: store,
      child: Scaffold(
        body: Padding(
          padding: EdgeInsets.all(30.0),
          child: Column(children: [
            StoreConnector<int, String>(
                converter: (store) => store.state.toString(),
                builder: (context, count) {
                  return Text(count);
                }),
            StoreConnector<int, VoidCallback>(
              converter: (store) {
                return () => store.dispatch(Actions.Increment);
              },
              builder: (context, callback) {
                return OutlineButton(
                    onPressed: callback, child: Icon(Icons.add));
              },
            ),
            StoreConnector<int, VoidCallback>(
              converter: (store) {
                return () => store.dispatch(Actions.Decrement);
              },
              builder: (context, callback) {
                return OutlineButton(
                    onPressed: callback, child: Icon(Icons.remove));
              },
            ),
          ]),
        ),
      ),
    );
  }
}
