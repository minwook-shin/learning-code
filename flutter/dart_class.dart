class test{
  var a;
  var b;
  var c;
  var d;
  static const e = "";
  
  test(a,b,c){
    this.a = a;
    this.b = b;
    this.c = c;
  }
  
  // test(this.a,this.b,this.c)
  
  printValue(){
    print(this.a);
    print(this.b);
    print(this.c);
  }
  
  adder(test t){
    return t.a + t.b + t.c;
  }
  
  test.named(){
    print("identifier");
  }
  
  test.another(a,b,c):this(a,b,c);
  
}

class ImmutableTest {
  final num a, b, c;
  const ImmutableTest(this.a, this.b,this.c);
}

class superClass {
  String a;

  superClass.method(Map data) {
    print('super');
  }
}

class subClass extends superClass {
  subClass.method(Map data) : super.method(data) {
    print('sub');
  }
}

abstract class abstractClass{
  void m(){}
}

class implementClass implements abstractClass{
  @override
  m(){
    
  }
}

enum Color { red, green, blue }


void main() {
  var tmp = test(1,2,3);
  tmp.printValue();
  
  tmp.a = 5;
  tmp.printValue();
  
  print(tmp?.d);
  
  print(tmp.adder(tmp));
  
  test.named();
  
  var tmp1 = const ImmutableTest(5,5,5);
  var tmp2 = const ImmutableTest(5,5,5);
  print(tmp1 == tmp2);
  
  var tmp3 = test(5,5,5);
  var tmp4 = const ImmutableTest(5,5,5);
  print(tmp3 != tmp4);
  
  subClass.method({});
  
  var tmp5 = test.another(10,10,10);
  print(tmp5.a);
}

