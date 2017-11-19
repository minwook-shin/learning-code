#include<iostream>

using namespace std;

int main() {
    for (int i = 0; i < 10; i++) {
        cout <<"hello * 10" <<endl;
    }

    cout << "-----------" << endl;

    for (int i = 0; i < 10; i++) {
        cout << "2 *" << i << " = " << 2 * i << endl;
    }

    cout << "-----------" << endl;

    for (int i = 1; i < 10; i++) {
        for (int j = 1; j < 10; j++) {
            cout << i << " *" << j << " = " << i * j << endl;
        }
    }

    cout << "-----------" << endl;

    for (int i = 1; i < 10; i++) {
        for (int j = 1; j < 10; j++) {
            if (i * j % 3 == 0 && i *j % 7 == 0) {
                cout << i << " *" << j << " = " << i * j << endl;
            }
        }
    }
    cout << "-----------" << endl;
    for (int i = 0; i < 5; i++) {
        for (int j = 1; j <= i + 1; j++) {
            cout << '*';
        }
        cout << endl;
    }

	cout << "-----------" << endl;
    for (int i = 0; i <= 5; i++) {
        for (int j = 1; j <= 5 - i; j++) {
            cout << '*';
        }
        cout << endl;
    }

	cout << "-----------" << endl;
    for (int i = 0; i < 5; i++) {
        for (int j = 0; j < 4 - i; j++) {
            cout << ' ';
        }
        for (int j = 0; j < i * 2 + 1; j++) {
            cout << '*';
        }
        cout << endl;
    }
    return 0;
}