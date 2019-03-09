import 'package:flutter/material.dart';
import 'package:google_maps_flutter/google_maps_flutter.dart';
import 'dart:async';

void main() => runApp(MyApp());

class MyApp extends StatelessWidget {
  @override
  Widget build(BuildContext context) {
    return MaterialApp(
      title: 'Flutter google maps',
      home: MyHomePage(),
    );
  }
}

class MyHomePage extends StatefulWidget {
  @override
  _MyHomePageState createState() => _MyHomePageState();
}

class _MyHomePageState extends State<MyHomePage> {
  Completer<GoogleMapController> _controller = Completer();

  static final gwanghwamun = CameraPosition(
    target: LatLng(37.575929, 126.976849),
    zoom: 3.0,
  );

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      body: Center(
        child: GoogleMap(
          mapType: MapType.hybrid,
          initialCameraPosition: gwanghwamun,
          onMapCreated: (GoogleMapController controller) {
            _controller.complete(controller);
          },
          compassEnabled: true,
          zoomGesturesEnabled: true,
          rotateGesturesEnabled: true,
          scrollGesturesEnabled: true,
          tiltGesturesEnabled: true,
        ),
      ),
    );
  }
}
