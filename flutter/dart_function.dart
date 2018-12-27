void main() {
  void test() {
  	print("hello, world!");
	}
  test();
  
  test2() {
  	print("hello, world!");
	}
  test2();
  
  void test3() => print("hello, world!");
  test3();
  
  void test4(bool trig,bool option){
    if(trig){
      print(option);
    }
  }
  test4(true, false);
  
  void test5({bool trig, bool option}) {
    if(trig){
      print(option);
    }
  }
  test5(trig: true, option: false);
  
  String test6(bool trig, [String str]){
    if(trig){
      return str;
    }
    return "";
  }
  test6(false);
  
  void test7({String str = "hello,world!"}){
    print(str);
  }
  test7();
  
  void test8({List<int> list = const [1,2,3]}){
    print(list);
  }
  test8();
  
  void printArr(int element) {
  print(element);
  }
  var list = [1, 2, 3];
  list.forEach(printArr);
  
  var testList = ['1', '2', '3'];
  list.forEach((item) {
    print(testList[item-1]);
  });
  
  Function adder(int n){
    return (int i)=> n + i;
  }
  var tmp = adder(5);
  print(tmp(5));
  
  foo() {}
  print(foo() == null);
}