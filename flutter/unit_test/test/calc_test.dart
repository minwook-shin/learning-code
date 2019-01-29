import 'package:test/test.dart';
import 'package:unit_test/calc.dart';

void main() {
  group('Adder function', () {
    test('start at 0', () {
      expect(Calc().value, 0);
    });

    test('incremented', () {
      final calc = Calc();

      calc.increment();

      expect(calc.value, 1);
    });
  });
}
