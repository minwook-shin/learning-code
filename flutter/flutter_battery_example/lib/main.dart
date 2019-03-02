import 'package:flutter/material.dart';
import 'package:battery/battery.dart';

void main() => runApp(MyApp());

class MyApp extends StatelessWidget {
  @override
  Widget build(BuildContext context) {
    return MaterialApp(
      title: 'Flutter Demo',
      home: MyHomePage(),
    );
  }
}

class MyHomePage extends StatefulWidget {
  @override
  _MyHomePageState createState() => _MyHomePageState();
}

class _MyHomePageState extends State<MyHomePage> {
  var battery = Battery();
  BatteryState _batteryState;

  @override
  void initState() {
    super.initState();
    battery.onBatteryStateChanged.listen((BatteryState state) {
      setState(() {
        _batteryState = state;
      });
    });
  }

  getBatteryLevel() async {
    final int batteryLevel = await battery.batteryLevel;
    return batteryLevel;
  }

  @override
  Widget build(BuildContext context) {
    final key = new GlobalKey<ScaffoldState>();

    return Scaffold(
      key: key,
      body: Center(
        child: Column(
          mainAxisAlignment: MainAxisAlignment.center,
          children: [
            Text(_batteryState.toString()),
            OutlineButton(
              child: Icon(Icons.battery_charging_full),
              onPressed: () async {
                final int batteryLevel = await battery.batteryLevel;
                key.currentState.showSnackBar(
                    SnackBar(content: Text(batteryLevel.toString())));
              },
            )
          ],
        ),
      ),
    );
  }
}
