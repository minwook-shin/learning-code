from rx import Observable, Observer


class PrintObserver(Observer):

    def on_next(self, value):
        print(value)

    def on_completed(self):
        print("Done")

    def on_error(self, error):
        print(error)


source = Observable.from_(["python", "java", "kotlin", "go"])

source.subscribe(PrintObserver())
