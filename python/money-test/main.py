from money import Money
m = Money(2000, currency='KRW')
print(m.currency)
print(m.amount)
print(m)

m = Money(2000, currency='KRW')
print(m/2)
print(m*2)
print(m+1000)
print(m-1000)

m = Money(2000, currency='KRW')
m2 = Money(4000, currency='KRW')

print(m*2 == m2)


class KRW(Money):
    def __init__(self, amount='0'):
        super().__init__(amount=amount, currency='KRW')


price = KRW('10000')
print(price)