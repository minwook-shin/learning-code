import 'package:flutter/material.dart';

void main() => runApp(MaterialApp(
      title: 'Drawer',
      home: MyApp(),
    ));

class MyApp extends StatelessWidget {
  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(),
      drawer: Drawer(
        child: ListView(
          padding: EdgeInsets.zero,
          children: [
            DrawerHeader(
              child: Text('Header'),
              decoration: BoxDecoration(
                
                color: Colors.blue,
              ),
            ),
            ListTile(
              title: Text('Item'),
            ),
          ],
        ),
      ),
    );
  }
}
