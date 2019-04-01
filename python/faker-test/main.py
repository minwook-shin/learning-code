from faker import Faker

fake = Faker('ko_KR')

print("random_int",fake.random_int(min=0, max=9999))
print("random_element",fake.random_element(elements=('a', 'b', 'c')))
print("random_elements",fake.random_elements(elements=('a', 'b', 'c'), length=None, unique=False))
print("random_choices",fake.random_choices(elements=('a', 'b', 'c'), length=None))
print("random_letter",fake.random_letter())
print("random_digit",fake.random_digit())

print("numerify",fake.numerify(text="###"))
print("lexify",fake.lexify(text="????", letters="abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"))
print("bothify",fake.bothify(text="## ??", letters="abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"))
print("hexify",fake.hexify(text="^^^^", upper=False))

print("random_digit_not_null_or_empty",fake.random_digit_not_null_or_empty())
print("random_digit_or_empty",fake.random_digit_or_empty())
print("random_digit_not_null",fake.random_digit_not_null())

print("randomize_nb_elements",fake.randomize_nb_elements(number=10, le=False, ge=False, min=None, max=None))
print("random_letters",fake.random_letters(length=16))
print("random_number",fake.random_number(digits=None, fix_len=False))
print("random_uppercase_letter",fake.random_uppercase_letter())
print("random_lowercase_letter",fake.random_lowercase_letter())
print("random_sample",fake.random_sample(elements=('a', 'b', 'c'), length=None))

print("address",fake.address())
print("street_name",fake.street_name())
print("city",fake.city())
print("country",fake.country())
print("postcode",fake.postcode())
print("country_code",fake.country_code(representation="alpha-2"))
print("street_address",fake.street_address())

print("color_name",fake.color_name())
print("rgb_css_color",fake.rgb_css_color())

print("company",fake.company())

print("credit_card_full",fake.credit_card_full(card_type=None))

print("date",fake.date(pattern="%Y-%m-%d", end_datetime=None))
print("timezone",fake.timezone())

print("file_name",fake.file_name(category=None, extension=None))
print("file_path",fake.file_path(depth=1, category=None, extension=None))
print("mime_type",fake.mime_type(category=None))
print("file_extension",fake.file_extension(category=None))

print("latlng",fake.latlng())

print("uri_extension",fake.uri_extension())
print("ascii_safe_email",fake.ascii_safe_email())
print("uri",fake.uri())
print("ipv4",fake.ipv4(network=False, address_class=None, private=None))

print("job",fake.job())

print("sentences",fake.sentences(nb=3, ext_word_list=None))
print("text",fake.text(max_nb_chars=200, ext_word_list=None))
print("paragraphs",fake.paragraphs(nb=3, ext_word_list=None))
print("word",fake.word(ext_word_list=None))

print("locale",fake.locale())
print("language_code",fake.language_code())

print("profile",fake.profile(fields=None, sex=None))
print("name_female",fake.name_female())
print("name_male",fake.name_male())
print("phone_number",fake.phone_number())

print("pystr",fake.pystr(min_chars=None, max_chars=20))
print("pyfloat",fake.pyfloat(left_digits=None, right_digits=None, positive=False))
print("pybool",fake.pybool())

print("user_agent",fake.user_agent())
print("chrome",fake.chrome(version_from=13, version_to=63, build_from=800, build_to=899))
print("firefox",fake.firefox())
print("safari",fake.safari())
print("internet_explorer",fake.internet_explorer())
