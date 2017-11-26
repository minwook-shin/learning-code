#include<iostream>

using namespace std;

class calc {
public:
    calc() {
        cout << "calling calc()" << endl;
    }
    calc(char* pname) {
        strcpy_s(public_str, pname);
        cout << "calling calc() " << "+ "<< public_str << endl;
    }
    ~calc() {
        cout << "deleting calc()" << endl;
    }
private:
    int private_num;

public:
    char public_str[32];
    void display() {
        cout << "printing " << public_str << endl;
    }
};

int main() {
    calc sample("hello world");

    sample.display();

    return 0;
}