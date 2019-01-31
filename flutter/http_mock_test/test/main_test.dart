import 'package:http/http.dart' as http;
import 'package:mockito/mockito.dart';
import 'package:flutter_test/flutter_test.dart';
import 'package:http_mock_test/main.dart';

class MockClient extends Mock implements http.Client {}

main() {
  group('mockito http test', () {
    test('return title string', () async {
      final client = MockClient();
      when(client
              .get('https://postman-echo.com/response-headers?title=blog_test'))
          .thenAnswer(
              (_) async => http.Response('{"title": "blog_test"}', 200));
      expect(await fetchPost(client), isInstanceOf<TextClass>());
    });
  });
}
