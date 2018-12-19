import 'package:flutter/material.dart';

void main() => runApp(MyApp());

class FavoriteWidget extends StatefulWidget {
  @override
  _FavoriteWidgetState createState() => _FavoriteWidgetState();
}

class _FavoriteWidgetState extends State<FavoriteWidget> {
  bool _isFavorited = true;
  int _favoriteCount = 1;

  void _toggleFavorite() {
    setState(() {
      if (_isFavorited) {
        _favoriteCount -= 1;
        _isFavorited = false;
      } else {
        _favoriteCount += 1;
        _isFavorited = true;
      }
    });
  }

  @override
  Widget build(BuildContext context) {
    return Row(
      children: [
        Container(
          child: IconButton(
            icon: (_isFavorited ? Icon(Icons.star) : Icon(Icons.star_border)),
            color: Colors.red,
            onPressed: _toggleFavorite,
          ),
        ),
        SizedBox(
          width: 18.0,
          child: Container(
            child: Text('$_favoriteCount'),
          ),
        ),
      ],
    );
  }
}

class MyApp extends StatelessWidget {
  @override
  Widget build(BuildContext context) {
    Widget titleSection = Container(
      padding: EdgeInsets.all(30.0),
      child: Row(
        children: [
          Expanded(
            child: Column(
              crossAxisAlignment: CrossAxisAlignment.start,
              children: [
                Container(
                  child: Text(
                    '제목',
                    style: TextStyle(
                      fontWeight: FontWeight.bold,
                    ),
                  ),
                ),
                Text(
                  '부제목',
                ),
              ],
            ),
          ),
          FavoriteWidget(),
        ],
      ),
    );

    Widget textSection = Container(
      padding: EdgeInsets.all(30.0),
      child: Text(
        '''
설명입니다.설명입니다.설명입니다.설명입니다.설명입니다.설명입니다.설명입니다.설명입니다.설명입니다.설명입니다.설명입니다.설명입니다.설명입니다.설명입니다.
        ''',
        softWrap: true,
      ),
    );

    return MaterialApp(
      title: 'Flutter Demo',
      home: Scaffold(
        appBar: AppBar(
          title: Text('Building UI'),
        ),
        body: ListView(
          children: [
            Image.asset(
              'lake.jpg',
              width: 600.0,
              height: 240.0,
              fit: BoxFit.cover,
            ),
            titleSection,
            textSection,
          ],
        ),
      ),
    );
  }
}
