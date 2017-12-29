#include<iostream>

using namespace std;

int main() {
    double a;
	cout << "enter num(a) : ";
	cin >> a;

    a += 5;
    cout << "a + 5 = "<< a << endl;

    a -= 5;
    cout << "a - 5 = " << a << endl;

	a /= 2;
	cout << "a / 2 = " << a << endl;

	cout << "a == 20 = " << (a == 20) << endl;

	cout << "a != 20 = " << (a != 20) << endl;

	cout << "(a >= 10 && a <= 20)" << (a >= 10 && a <= 20)<< endl;

	cout << "(a >= 10 || a <= 20)" << (a >= 10 || a <= 20) << endl;

	return 0;
}