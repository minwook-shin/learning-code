import persistent
import ZODB.FileStorage
import transaction


class TestScore(persistent.Persistent):

    def __init__(self):
        self.score = 0

    def visit(self, i):
        self.score += i


if __name__ == '__main__':
    storage = ZODB.FileStorage.FileStorage('mydata.fs')

    db = ZODB.DB(storage)

    connection = db.open()

    root = connection.root

    root.storage = TestScore()

    root.storage.visit(1)
    transaction.commit()
    print(root.storage.score)

    root.storage.visit(1)
    transaction.abort()
    print(root.storage.score)

    save = transaction.savepoint()
    root.storage.visit(1)
    root.storage.visit(1)
    root.storage.visit(1)
    print(root.storage.score)
    save.rollback()
    print(root.storage.score)

    root.storage.visit(1)
    transaction.doom()
    print(root.storage.score)

    # root.storage.visit(1)
    # transaction.commit()
    # print(root.storage.score)
