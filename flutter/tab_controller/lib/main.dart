import 'package:flutter/material.dart';

void main() {
  runApp(MaterialApp(
    home: DefaultTabController(
      child: Scaffold(
        appBar: AppBar(
          title: Text('Tabs'),
          bottom: TabBar(
            tabs: [
              Tab(icon: Icon(Icons.apps), text: "앱"),
              Tab(icon: Icon(Icons.audiotrack), text: "오디오"),
              Tab(icon: Icon(Icons.settings), text: "설정"),
            ],
          ),
        ),
        body: TabBarView(
          children: [
            Container(
              color: Colors.red,
            ),
            Container(
              color: Colors.green,
            ),
            Container(
              color: Colors.blue,
            ),
          ],
        ),
      ),
      length: 3,
    ),
  ));
}
