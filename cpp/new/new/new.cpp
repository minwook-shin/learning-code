#include<iostream>

using namespace std;

int main() {
    int *pnum = new int;
    *pnum = 100;
    cout << "*pnum : " << *pnum << endl;;
    delete pnum;
//  cout << "*pnum" << *pnum << endl;;

    int *parray = new int[10];
    parray[1] = 200;
    cout << "parray[1] : " << parray[1] << endl;;
    delete[] parray;
//  cout << "parray[1]" << parray[1] << endl;;

    return 0;
}