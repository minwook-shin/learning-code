class car(object):
    def __init__(self,engine,fuel):
        self.engine = engine
        self.fuel = fuel
        self.__price = "10000"

    @property
    def price(self):
        return self.__price

    @price.setter
    def price(self,n):
        self.__price = n

    def __str__(self):
        toString = self.engine + " " + self.fuel + " " + "Car"
        return toString
    def run(self):
        print("running car")

class second_car(car):
    def run(self):
        print("running second car")



minicar = car("super","oil")
print(minicar)
print(minicar.price)
minicar2 = second_car("super","oil")
print(minicar2)
minicar2.run()
print(minicar2.price)