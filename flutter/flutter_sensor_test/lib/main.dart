import 'package:flutter/material.dart';
import 'package:sensors/sensors.dart';

void main() => runApp(MyApp());

class MyApp extends StatelessWidget {
  @override
  Widget build(BuildContext context) {
    return MaterialApp(
      title: 'Flutter sensors',
      home: MyHomePage(),
    );
  }
}

class MyHomePage extends StatefulWidget {
  @override
  _MyHomePageState createState() => _MyHomePageState();
}

class _MyHomePageState extends State<MyHomePage> {
  List accelerometer;
  List userAccelerometer;
  List gyroscope;

  @override
  void initState() {
    super.initState();
    accelerometerEvents.listen((AccelerometerEvent e) {
      setState(() {
        accelerometer = <double>[e.x, e.y, e.z];
      });
    });
    gyroscopeEvents.listen((GyroscopeEvent e) {
      setState(() {
        gyroscope = <double>[e.x, e.y, e.z];
      });
    });
    userAccelerometerEvents.listen((UserAccelerometerEvent e) {
      setState(() {
        userAccelerometer = <double>[e.x, e.y, e.z];
      });
    });
  }

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      body: Center(
        child: Column(
          mainAxisAlignment: MainAxisAlignment.center,
          children: <Widget>[
            Text(
              accelerometer.toString(),
            ),
            Text(
              gyroscope.toString(),
            ),
            Text(
              userAccelerometer.toString(),
            ),
          ],
        ),
      ),
    );
  }
}
