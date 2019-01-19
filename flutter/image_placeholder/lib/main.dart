import 'package:flutter/material.dart';
import 'package:cached_network_image/cached_network_image.dart';

void main() {
  runApp(MaterialApp(title: 'Image loading', home: Scaffold(body: MyApp())));
}

class MyApp extends StatelessWidget {
  @override
  Widget build(BuildContext context) {
    return Container(
      alignment: Alignment.center,
      child: CachedNetworkImage(
        placeholder: CircularProgressIndicator(),
        imageUrl:
            'https://github.com/flutter/plugins/raw/master/packages/video_player/doc/demo_ipod.gif?raw=true',
        errorWidget: Icon(Icons.error),
        fadeInCurve: Curves.easeIn,
        fadeInDuration: Duration(seconds: 3),
      ),
    );
  }
}
