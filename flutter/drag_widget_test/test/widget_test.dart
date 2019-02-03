import 'package:flutter/material.dart';
import 'package:flutter_test/flutter_test.dart';
import 'package:drag_widget_test/main.dart';

void main() {
  testWidgets('drag', (WidgetTester tester) async {
    await tester.pumpWidget(MyApp());
    await tester.pump();

    await tester.drag(find.byKey(Key('2')), Offset(500.0, 0.0));
    await tester.pumpAndSettle();

    expect(find.text('2'), findsNothing);
    expect(find.text('1'), findsOneWidget);
  });
}
