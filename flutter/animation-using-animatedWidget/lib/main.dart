import 'package:flutter/material.dart';

void main() => runApp(MyApp());

class MyApp extends StatefulWidget {
  @override
  _LogoAppState createState() => _LogoAppState();
}

class _LogoAppState extends State<MyApp> with SingleTickerProviderStateMixin {
  AnimationController controller;
  Animation animation;

  @override
  void initState() {
    super.initState();
    controller =
        AnimationController(duration: Duration(seconds: 5), vsync: this);
    controller.forward();
  }

  @override
  void dispose() {
    super.dispose();
    controller.dispose();
  }

  @override
  Widget build(BuildContext context) {
    animation = CurvedAnimation(parent: controller, curve: Curves.bounceInOut);

    animation.addStatusListener((status) {
      if (status == AnimationStatus.completed) {
        controller.reverse();
      } else if (status == AnimationStatus.dismissed) {
        controller.forward();
      }
    });

    return MaterialApp(
      home: Scaffold(body: Animated(animation: animation)),
    );
  }
}

class Animated extends AnimatedWidget {
  Animated({Key key, @required Animation<double> animation})
      : super(key: key, listenable: animation);

  Widget build(BuildContext context) {
    final Animation<double> animation = listenable;
    return Center(
      child: Container(
          height: Tween(begin: 0.0, end: 200.0).evaluate(animation),
          width: Tween(begin: 0.0, end: 200.0).evaluate(animation),
          child: CircularProgressIndicator()),
    );
  }
}
