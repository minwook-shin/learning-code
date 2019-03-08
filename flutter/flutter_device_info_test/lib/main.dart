import 'package:flutter/material.dart';
import 'package:device_info/device_info.dart';
import 'dart:io';

void main() {
  runApp(MyApp());
}

class MyApp extends StatelessWidget {
  @override
  Widget build(BuildContext context) {
    return MaterialApp(
      title: 'Flutter device_info',
      home: MyHomePage(),
    );
  }
}

class MyHomePage extends StatefulWidget {
  @override
  _MyHomePageState createState() => _MyHomePageState();
}

class _MyHomePageState extends State<MyHomePage> {
  List _data = <dynamic>[];
  static final DeviceInfoPlugin plugin = DeviceInfoPlugin();

  @override
  void initState() {
    super.initState();
    initPlatform();
  }

  Future<void> initPlatform() async {
    if (Platform.isAndroid) {
      setState(() async {
        _data = getAndroidDevice(await plugin.androidInfo);
      });
    }
    if (Platform.isIOS) {
      setState(() async {
        _data = getIosDevice(await plugin.iosInfo);
      });
    }
  }

  getAndroidDevice(AndroidDeviceInfo device) {
    return [
      device.version.securityPatch,
      device.version.sdkInt,
      device.version.release,
      device.version.codename,
      device.board,
      device.bootloader,
      device.brand,
      device.device,
      device.display,
      device.fingerprint,
      device.hardware,
      device.host,
      device.id,
      device.manufacturer,
      device.model,
      device.product,
      device.androidId,
    ];
  }

  getIosDevice(IosDeviceInfo device) {
    return [
      device.name,
      device.systemName,
      device.systemVersion,
      device.model,
      device.localizedModel,
      device.identifierForVendor,
    ];
  }

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      body: Center(
        child: Text(
          _data.toString(),
        ),
      ),
    );
  }
}
