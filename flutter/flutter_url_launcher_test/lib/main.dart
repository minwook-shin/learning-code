import 'package:flutter/material.dart';
import 'package:url_launcher/url_launcher.dart';

void main() => runApp(MyApp());

class MyApp extends StatelessWidget {
  @override
  Widget build(BuildContext context) {
    return MaterialApp(
      title: 'Flutter url_launcher',
      home: MyHomePage(),
    );
  }
}

class MyHomePage extends StatefulWidget {
  @override
  _MyHomePageState createState() => _MyHomePageState();
}

class _MyHomePageState extends State<MyHomePage> {
  launchBrowser(String url) async {
    if (await canLaunch(url)) {
      await launch(url, forceSafariVC: false, forceWebView: false);
    }
  }

  launchWebView(String url) async {
    if (await canLaunch(url)) {
      await launch(url, forceSafariVC: true, forceWebView: true);
    }
  }

  launchUniversalForIos(String url) async {
    if (await canLaunch(url)) {
      final bool succeeded =
          await launch(url, forceSafariVC: false, universalLinksOnly: true);
      if (!succeeded) {
        await launch(url, forceSafariVC: true);
      }
    }
  }

  makeCall(String url) async {
    if (await canLaunch(url)) {
      await launch(url);
    }
  }

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      body: Center(
        child: Column(
          mainAxisAlignment: MainAxisAlignment.center,
          children: <Widget>[
            OutlineButton(
              onPressed: () {
                setState(() {
                  launchBrowser("https://www.google.com");
                });
              },
              child: Column(
                children: <Widget>[Icon(Icons.link), Text("launchBrowser")],
              ),
            ),
            OutlineButton(
              onPressed: () {
                setState(() {
                  launchWebView("http://www.google.com");
                });
              },
              child: Column(
                children: <Widget>[Icon(Icons.link), Text("launchWebView")],
              ),
            ),
            OutlineButton(
              onPressed: () {
                launchUniversalForIos("http://youtube.com");
              },
              child: Column(
                children: <Widget>[
                  Icon(Icons.link),
                  Text("launchUniversalForIos")
                ],
              ),
            ),
            OutlineButton(
              onPressed: () {
                setState(() {
                  makeCall("0000");
                });
              },
              child: Icon(Icons.call),
            ),
          ],
        ),
      ),
    );
  }
}
