import unittest

class TestMethods(unittest.TestCase):

    def setUp(self):
        print("set up!")

    def tearDown(self):
        print("Cleaning job")

    def test_upper(self):
        self.assertEqual('hello'.upper(), 'HELLO')

    def test_isupper(self):
        self.assertTrue('HELLO'.isupper())
        self.assertFalse('hello'.isupper())

    def test_split(self):
        s = 'hello, world!'
        self.assertEqual(s.split(), ['hello,', 'world!'])


    def test_even(self):
        for i in range(0, 6):
            with self.subTest(i=i):
                self.assertEqual(i % 2, 0)

    @unittest.skip("skipping")
    def test_nothing(self):
        self.assertTrue(True)


def suite():
    suite = unittest.TestSuite()
    suite.addTest(TestMethods('test_upper'))
    suite.addTest(TestMethods('test_isupper'))
    return suite

    
if __name__ == '__main__':
    runner = unittest.TextTestRunner()
    runner.run(suite())


    def testSomething():
        assert 0 is not None

    testcase = unittest.FunctionTestCase(testSomething)
    testcase.run()

    unittest.main()


    


