import datetime

from forex_python.converter import CurrencyRates

from forex_python.converter import CurrencyCodes


c = CurrencyRates()
usd = c.get_rates('USD')
print(usd)
krw = c.get_rates('KRW')
print(krw)

c = CurrencyRates()
usd_krw = c.get_rates('USD')['KRW']
print(usd_krw)
krw_usd = c.get_rates('KRW')['USD']
print(krw_usd)

date = datetime.datetime(2018, 5, 5, 18, 0, 0, 151012)
c = CurrencyRates()
print(c.get_rates('USD', date)['KRW'])

c = CurrencyRates()
usd_to_krw = c.get_rate('USD', 'KRW')
print(usd_to_krw)

date = datetime.datetime(2018, 5, 5, 18, 0, 0, 151012)
c = CurrencyRates()
print(c.get_rate('USD','KRW', date))

c = CurrencyRates()
con = c.convert('USD', 'KRW', 100)
print(con)

date = datetime.datetime(2018, 5, 5, 18, 0, 0, 151012)
c = CurrencyRates()
print(c.convert('USD','KRW',100, date))

c = CurrencyCodes()
sb = c.get_symbol('KRW')
print(sb)