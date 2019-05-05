from decimal import Decimal
from money import Money, xrates

xrates.install('money.exchange.SimpleBackend')
xrates.base = 'USD'
xrates.setrate('USD', Decimal('1'))
xrates.setrate('KRW', Decimal('1165.95'))

a = Money(100, 'USD')
b = Money(1, 'KRW')

print(a.to('KRW'))
print(b.to('USD'))
print(a.to('KRW') + b)