#include<iostream>

using namespace std;

int g_num = 0;

int sum(int a, int b) {

    return a + b;
}

void output() {
    cout << "text func" << endl;
}

void output_num(int a) {
    cout << "number func"<< a << endl;
}

void change_num(int a) {
    g_num = 10000;
}

void change_num(int *pnum) {
    *pnum = 20000;
}

int default_sum(int a, int b = 10) {

    return a + b;
}


int main() {

    cout << sum(10, 20) << endl;
    cout << sum(40, 10) << endl;

    output();
    output_num(10);

    cout << g_num << endl;
    change_num(g_num);
    cout << g_num << endl;

    int num = 0;
    cout << num << endl;
    change_num(&num);
    cout << num << endl;

    cout << default_sum(10) << endl;

    return 0;
}