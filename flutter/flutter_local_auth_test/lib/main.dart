import 'package:flutter/material.dart';
import 'package:local_auth/local_auth.dart';

void main() => runApp(MaterialApp(home: MyApp()));

class MyApp extends StatefulWidget {
  @override
  _MyAppState createState() => _MyAppState();
}

class _MyAppState extends State<MyApp> {
  bool _isLocalAuth;
  String _authorized;

  @override
  initState() {
    super.initState();
    checkBio();
  }

  checkBio() async {
    var isLocalAuth;
    isLocalAuth = await LocalAuthentication().canCheckBiometrics;
    setState(() {
      _isLocalAuth = isLocalAuth;
    });
  }

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      body: Column(
        children: <Widget>[
          Text(_isLocalAuth.toString()),
          Text(_authorized == null ? "" : _authorized),
          OutlineButton(
              child: Text('Authenticate'),
              onPressed: () async {
                bool authenticated = false;
                authenticated = await LocalAuthentication()
                    .authenticateWithBiometrics(
                        localizedReason: '지문 인식을 진행해주십시오.');
                setState(() {
                  _authorized = authenticated ? 'Authorized' : 'Not Authorized';
                });
              })
        ],
      ),
    );
  }
}
