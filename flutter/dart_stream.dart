Future<int> adder(Stream<int> i) async {
  var sum = 0;
  await for (var v in i) {
    sum += v;
  }
  return sum;
}

Stream<int> count(int i) async* {
  for (int j = 1; j <= i; j++) {
    yield j;
  }
}


main() async {
  var sum = await adder(count(10));
  print(sum);
}