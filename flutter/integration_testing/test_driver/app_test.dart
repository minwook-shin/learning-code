import 'package:flutter_driver/flutter_driver.dart';
import 'package:test/test.dart';

void main() {
  group('App test case', () {
    FlutterDriver driver;

    setUpAll(() async {
      driver = await FlutterDriver.connect();
    });

    test('starts at 0', () async {
      expect(await driver.getText(find.byValueKey('counter')), "0");
    });

    test('increments', () async {
      await driver.tap(find.byTooltip('Increment'));
      expect(await driver.getText(find.byValueKey('counter')), "1");
    });

    tearDownAll(() async {
      if (driver != null) {
        driver.close();
      }
    });
  });
}
