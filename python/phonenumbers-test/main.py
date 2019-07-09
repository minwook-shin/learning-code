import phonenumbers


phone1 = phonenumbers.parse("+821000000000", None)
print(phone1)
type(phone1)

phone2 = phonenumbers.parse("010 0000 0000", "KR")
print(phone2)

print(phone1 == phone2)

print(phonenumbers.is_valid_number(phone2))
print(phonenumbers.is_possible_number(phone2))


try:
  x = phonenumbers.parse("01000000000", None)
except phonenumbers.phonenumberutil.NumberParseException as e:
  print(e)

try:
  x = phonenumbers.parse("hello,world", None)
except phonenumbers.phonenumberutil.NumberParseException as e:
  print(e)

print(phonenumbers.format_number(phone1, phonenumbers.PhoneNumberFormat.NATIONAL))
print(phonenumbers.format_number(phone1, phonenumbers.PhoneNumberFormat.INTERNATIONAL))
print(phonenumbers.format_number(phone1, phonenumbers.PhoneNumberFormat.E164))

formatter = phonenumbers.AsYouTypeFormatter("KR")
print(formatter.input_digit("1"))
print(formatter.input_digit("0"))
print(formatter.input_digit("0"))
print(formatter.input_digit("0"))
print(formatter.input_digit("0"))
print(formatter.input_digit("0"))
print(formatter.input_digit("0"))
print(formatter.input_digit("0"))


text = "Call me at 010-0000-0000 after 10 am."
for match in phonenumbers.PhoneNumberMatcher(text, "KR"):
  print(match)

for match in phonenumbers.PhoneNumberMatcher(text, "KR"):
  print(phonenumbers.format_number(match.number, phonenumbers.PhoneNumberFormat.E164))


from phonenumbers import geocoder

kr_number = phonenumbers.parse("01000000000", "KR")
print(geocoder.description_for_number(kr_number, "ko"))


from phonenumbers import timezone
kr_number = phonenumbers.parse("01000000000", "KR")
print(timezone.time_zones_for_number(kr_number))
