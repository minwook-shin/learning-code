import 'package:flutter_driver/flutter_driver.dart';
import 'package:test/test.dart';

void main() {
  group('Scrolling test code', () {
    FlutterDriver driver;

    setUpAll(() async {
      driver = await FlutterDriver.connect();
    });

    test('verifies item', () async {
      await driver.scrollUntilVisible(
          find.byValueKey('listview'), find.byValueKey('30'),
          dyScroll: -100);
      expect(await driver.getText(find.byValueKey('30')), '30');
    });

    tearDownAll(() async {
      if (driver != null) {
        await driver.close();
      }
    });
  });
}
