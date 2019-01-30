import 'package:mockito/mockito.dart';
import 'package:flutter_test/flutter_test.dart';
import '../lib/calc.dart';
import 'package:mockito_basic/helper.dart';

class MockHelper extends Mock implements Helper {}

void main() {
  group('mock', () {
    test('verify', () {
      var helper = MockHelper();
      helper.bonusClick();
      verify(helper.bonusClick());
    });

    test('verify call', () {
      var helper = MockHelper();
      helper.bonusClick();
      helper.bonusClick();
      verify(helper.bonusClick()).called(2);
    });
    test('bonusClick return', () {
      var helper = MockHelper();
      var calc = Calc();
      when(helper.bonusClick()).thenReturn(10);
      expect(calc.syncHelper(helper), 11);
    });
    test('reset', () {
      var helper = MockHelper();
      var calc = Calc();
      when(helper.bonusClick()).thenReturn(10);
      expect(calc.syncHelper(helper), 11);
      reset(helper);
      expect(helper.bonusClick(), null);
    });
  });
}
