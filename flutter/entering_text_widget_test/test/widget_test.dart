import 'package:flutter/material.dart';
import 'package:flutter_test/flutter_test.dart';
import 'package:entering_text_widget_test/main.dart';

void main() {
  testWidgets('entering test', (WidgetTester tester) async {
    await tester.pumpWidget(MyApp());
    await tester.enterText(find.byKey(Key('textForm')), "hello");

    await tester.tap(find.byKey(Key('button')));
    await tester.pump();

    expect(find.byKey(Key('text')), findsOneWidget);
    expect(find.text("world! hello,"), findsNothing);
  });
}
