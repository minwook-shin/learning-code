import 'package:web_socket_channel/io.dart';
import 'package:flutter/material.dart';
import 'package:web_socket_channel/web_socket_channel.dart';

void main() {
  runApp(MaterialApp(title: 'WebSocket', home: Scaffold(body: MyApp())));
}

class MyApp extends StatefulWidget {
  @override
  _MyAppState createState() => _MyAppState();
}

class _MyAppState extends State<MyApp> {
  TextEditingController _controller = TextEditingController();
  final WebSocketChannel ch =
      IOWebSocketChannel.connect('ws://echo.websocket.org');

  @override
  void dispose() {
    ch.sink.close();
    super.dispose();
  }

  @override
  Widget build(BuildContext context) {
    return Column(
      children: <Widget>[
        TextFormField(
          controller: _controller,
        ),
        StreamBuilder(
          stream: ch.stream,
          builder: (context, snapshot) {
            if (snapshot.hasData) {
              return Text(snapshot.data.toString());
            }
            return Text("");
          },
        ),
        OutlineButton(
          onPressed: () {
            ch.sink.add(_controller.text);
          },
          child: Icon(Icons.send),
        )
      ],
    );
  }
}
