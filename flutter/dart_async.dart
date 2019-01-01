void main() {
  hello() => Future.delayed(Duration(seconds: 1),()=>"hello, world");
  
  Future func() async{
    var test = await hello();
    print(test);
  }
  func();
  print("delayed <hello,world>");
  
  var fail = 0;
  error() => Future.delayed(Duration(seconds: 1),() => fail == null);
  Future handleError() async{
    var test = await error();
    try{
      assert(test);
    }
    catch(e){
      print(e);
    }
  }
  handleError();
  print("delayed <assert>");
  
  Future<String> future() => Future.delayed(Duration(seconds: 1),()=>"future api");
  Future<void> printDelayed() {
    return future().then((str)=>print(str));
  }
  printDelayed();
  print("delayed <future api>");
}