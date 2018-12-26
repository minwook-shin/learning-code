void main() {
  var hello = 'hello,world';
  print(hello);
  
  dynamic variable;
  print(variable);
  variable = "10";
  print(variable);
  
  int year = 10;
  print(year);
  String name = "string";
  print(name);
  
  int lineCount;
  print(lineCount == null);
  
  final test = 'test';
  final String testStr = 'test';
  print(test);
  print(testStr);
 
  const test2 = 'test';
  const String testStr2 = 'test';
  print(test2);
  print(testStr2);
  
  double d = 1;
  print(d == 1.0);

  var ip = int.parse('1');
  print(ip == 1);

  var dp = double.parse('1.1');
  print(dp == 1.1);

  String ts = 1.toString();
  print(ts == '1');

  String tsf = 3.14159.toStringAsFixed(2);
  print(tsf == '3.14');
  
  var empty = '';
  print(empty.isEmpty);

  var nan = 0 / 0;
  print(nan.isNaN);
  
  var list = [1, 2, 3];
  print(list);
  
  var literals = {
  'first': '1',
  'second': '2',
  'third': '3'
	};
  print(literals);
  
  var constructor = Map();
  constructor['first'] = '1';
  constructor['second'] = '2';
  constructor['third'] = '3';
  print(constructor);
  
  #symbol;
}