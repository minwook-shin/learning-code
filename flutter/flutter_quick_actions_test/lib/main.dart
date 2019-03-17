import 'package:flutter/material.dart';
import 'package:quick_actions/quick_actions.dart';

void main() => runApp(MyApp());

class MyApp extends StatelessWidget {
  @override
  Widget build(BuildContext context) {
    return MaterialApp(
      title: 'Flutter quick_actions',
      home: MyHomePage(),
    );
  }
}

class MyHomePage extends StatefulWidget {
  @override
  _MyHomePageState createState() => _MyHomePageState();
}

class _MyHomePageState extends State<MyHomePage> {
  @override
  void initState() {
    super.initState();
    final QuickActions quickActions = const QuickActions();
    quickActions.initialize((String shortcutType) {
      if (shortcutType == 'action_main') {
        print("main");
      }
      if (shortcutType == 'action_help') {
        print("help");
      }
    });

    quickActions.setShortcutItems(<ShortcutItem>[
      ShortcutItem(
          type: 'action_main', localizedTitle: '메인 스크린', icon: 'icon_main'),
      ShortcutItem(
          type: 'action_help', localizedTitle: '도움말', icon: 'icon_help')
    ]);
  }

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      body: Center(
        child: Text(""),
      ),
    );
  }
}
