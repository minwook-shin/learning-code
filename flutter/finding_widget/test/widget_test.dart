import 'package:flutter/material.dart';
import 'package:flutter_test/flutter_test.dart';

import 'package:finding_widget/main.dart';

void main() {
  testWidgets('text', (WidgetTester tester) async {
    await tester.pumpWidget(MyApp());
    expect(find.text('0'), findsOneWidget);
    await tester.tap(find.byIcon(Icons.add));
    await tester.pump();
    expect(find.text('1'), findsOneWidget);
  });

  testWidgets('byIcon', (WidgetTester tester) async {
    await tester.pumpWidget(MyApp());
    expect(find.byIcon(Icons.publish), findsOneWidget);
    await tester.tap(find.byIcon(Icons.add));
    await tester.pump();
    expect(find.text('1'), findsOneWidget);
  });

  testWidgets('byTooltip', (WidgetTester tester) async {
    await tester.pumpWidget(MyApp());
    expect(find.byTooltip('Increment'), findsOneWidget);
    await tester.tap(find.byIcon(Icons.add));
    await tester.pump();
    expect(find.text('1'), findsOneWidget);
  });

  testWidgets('byKey', (WidgetTester tester) async {
    await tester.pumpWidget(MyApp());
    expect(find.byKey(Key("download")), findsOneWidget);
    await tester.tap(find.byIcon(Icons.add));
    await tester.pump();
    expect(find.text('1'), findsOneWidget);
  });

  testWidgets('byType', (WidgetTester tester) async {
    await tester.pumpWidget(MyApp());
    expect(find.byType(FloatingActionButton), findsOneWidget);
    await tester.tap(find.byIcon(Icons.add));
    await tester.pump();
    expect(find.text('1'), findsOneWidget);
  });

  testWidgets('byWidget', (WidgetTester tester) async {
    expect(find.byWidget(Padding(padding: EdgeInsets.zero)), findsNothing);
  });
}
