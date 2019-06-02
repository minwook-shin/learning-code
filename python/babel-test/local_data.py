from babel import Locale

locale = Locale('en', 'US')
print(locale.territories['US'])

locale = Locale('ko', 'KR')
print(locale.territories['US'])


lang = Locale.parse('ko_KR')
print(lang.display_name)
print(lang.english_name)
print(lang.language)
print(lang.language_name)
locale = Locale('ko')

month_names = locale.months['format']['wide'].items()
for _, name in sorted(month_names):
    print(name)

month_names = locale.days['format']['wide'].items()
for _, name in sorted(month_names):
    print(name)
