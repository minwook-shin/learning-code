import 'package:flutter/material.dart';
import 'package:flutter/foundation.dart' show SynchronousFuture;
import 'package:flutter_localizations/flutter_localizations.dart';

class LocalizationsClass {
  final Locale locale;
  LocalizationsClass(this.locale);

  static Map _local = {
    'en': {'title': 'Hello'},
    'ko': {'title': '안녕'}
  };

  String get title {
    return _local[locale.languageCode]['title'];
  }
}

class LocalizationsDelegateClass
    extends LocalizationsDelegate<LocalizationsClass> {
  @override
  bool isSupported(Locale locale) => ['en', 'ko'].contains(locale.languageCode);

  @override
  Future<LocalizationsClass> load(Locale locale) {
    return SynchronousFuture<LocalizationsClass>(LocalizationsClass(locale));
  }

  @override
  bool shouldReload(LocalizationsDelegateClass old) => false;
}

void main() {
  runApp(MaterialApp(
      localizationsDelegates: [
        LocalizationsDelegateClass(),
        GlobalMaterialLocalizations.delegate,
        GlobalWidgetsLocalizations.delegate
      ],
      supportedLocales: [Locale('en', ''), Locale('ko', '')],
      home: MyAppPage()));
}

class MyAppPage extends StatefulWidget {
  @override
  _MyHomePageState createState() => _MyHomePageState();
}

class _MyHomePageState extends State<MyAppPage> {
  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(
          title: Text(
              Localizations.of<LocalizationsClass>(context, LocalizationsClass)
                  .title)),
    );
  }
}
