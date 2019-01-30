import 'helper.dart';

class Calc {
  int value = 0;
  var help= Helper();
  void increment() => value+=1;

  int syncHelper(help) {
    this.value = help.setup();
    this.value = help.bonusClick();
    increment();
    return this.value;
  }
}
