import 'package:flutter/material.dart';
import 'package:connectivity/connectivity.dart';

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
  String networkStatus = 'null';
  @override
  void initState() {
    super.initState();
    initConnect();
  }

  initConnect() async {
    ConnectivityResult connectResult;
    connectResult = await Connectivity().checkConnectivity();
    updateStatus(connectResult);
  }

  updateStatus(ConnectivityResult result) async {
    if (result == ConnectivityResult.wifi) {
      String wifiName, wifiIP;
      wifiName = (await Connectivity().getWifiName()).toString();
      wifiIP = (await Connectivity().getWifiIP()).toString();

      setState(() {
        networkStatus = result.toString() + "/" + wifiName + "/" + wifiIP;
      });
    } else if (result == ConnectivityResult.mobile) {
      setState(() {
        networkStatus = result.toString();
      });
    } else if (result == ConnectivityResult.none) {
      setState(() => networkStatus = result.toString());
    }
  }

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      body: Center(child: Text(networkStatus)),
    );
  }
}
