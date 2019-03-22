import 'package:flutter/material.dart';
import 'package:firebase_admob/firebase_admob.dart';

void main() {
  runApp(MyApp());
  FirebaseAdMob.instance.initialize(appId: FirebaseAdMob.testAppId);
}

class MyApp extends StatelessWidget {
  @override
  Widget build(BuildContext context) {
    return MaterialApp(
      title: 'Flutter admob',
      home: MyHomePage(),
    );
  }
}

class MyHomePage extends StatefulWidget {
  @override
  _MyHomePageState createState() => _MyHomePageState();
}

MobileAdTargetingInfo targetingInfo = MobileAdTargetingInfo(
  keywords: <String>['flutterio', 'beautiful apps'],
  contentUrl: 'https://flutter.dev',
  childDirected: false,
  testDevices: <String>[],
);

class _MyHomePageState extends State<MyHomePage> {
  BannerAd myBanner = BannerAd(
    adUnitId: BannerAd.testAdUnitId,
    size: AdSize.smartBanner,
    targetingInfo: targetingInfo,
    listener: (MobileAdEvent event) {
      print("$event");
    },
  );

  @override
  void initState() {
    super.initState();
    myBanner
      ..load()
      ..show(
        anchorType: AnchorType.bottom,
      );

    RewardedVideoAd.instance.load(
        adUnitId: RewardedVideoAd.testAdUnitId, targetingInfo: targetingInfo);

    RewardedVideoAd.instance.listener =
        (RewardedVideoAdEvent event, {String rewardType, int rewardAmount}) {
      if (event == RewardedVideoAdEvent.rewarded) {
        setState(() {
          print(rewardAmount);
        });
      }
    };
  }

  @override
  Widget build(BuildContext context) {
    return Scaffold(
        body: OutlineButton(
      onPressed: () {
        RewardedVideoAd.instance.show();
      },
      child: Text("보상형 광고"),
    ));
  }
}
