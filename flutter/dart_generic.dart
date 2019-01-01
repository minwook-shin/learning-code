abstract class StringTest{
  String getTest(String i);
  void setTest(String i);
}
abstract class Test<T>{
  T getTest(T i);
  void setTest(T i);
}

class superClass{}
class sub1 <T extends superClass>{
  toString() => "$T";
}

void main() {
  var list1 = <int>[1,2,3];
  var list2 = List<int>();
  var map1 = <int,int>{
    1 : 1,
    2 : 2,
    3 : 3
  };
  var map2 = Map<int,int>();
  print(list1.toString()+list2.toString()+map1.toString()+map2.toString());
  
  print(sub1<superClass>());
  
  T functionTest<T>(List<T> i){
    T tmp = i[0];
    return tmp;
  }
  var f = functionTest<int>(list1);
  print(f);
}