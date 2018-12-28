void main() {
  if(true){
    print('hello, world!');
  }
  
  var msg = "hello, world";
  for(var i = 0;i<5;i++){
    print(msg[i]);
  }
  
  var list = [1,2,3];
  list.forEach((i)=>print(i));
  for(var i in list){
    print(i);
  }
  
  var count = 0;
  while(count<3){
    print(count);
    count++;
  }
  
  do{
    print("hello");
  }while(false);
  
  while(true){
    print("while...");
    break;
  }
  
  for(var i = 0;i<4;i++){
    if(i==0){
      continue;
    }
    print(i);
  }
  
  var command = 'dart';
  switch (command) {
    case 'c':
    case 'python':
    case 'java':
    case 'go':
      print("go");
      break;
    case 'dart':
      print("dart");
      continue defaultValue;
    defaultValue:
    default:
      print("default");
  }
  
  try{
  assert(command == null);
  }catch(e){
    print(e);
  }finally{
    print("done!");
  }
}