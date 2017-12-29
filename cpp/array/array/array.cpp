#include<iostream>
using namespace std;

int main() {
    // int i[10] = {0,0,0,0,0,0,0,0,0,0};
    int i[10] = {};
    i[1] = 77;

    cout << i[5] << endl;
    cout << "----------" << endl;

    for (int index = 0; index < 10; index++) {
        cout << " | ";
        cout << i[index] ;
        cout << " | ";
    }
    cout << endl;
    cout << "----------" << endl;

    int i2[10][10] = { {1,1,1},{2,2,2} };

    for (int index = 0; index < 10; index++) {
        for (int index2 = 0; index2 < 10; index2++) {
            cout << " | ";
            cout << i2[index][index2];
            cout << " | ";
        }
        cout << endl;
    }


    return 0;
}