#include<iostream>

using namespace std;

template<typename t>
void out() {
    cout << typeid(t).name() << endl;
}
template<typename t>
void out(t data) {
    cout << typeid(t).name() << endl;
    cout << data << endl;
}

class test
{
public:
    test();
    ~test();

private:

};

enum enum_test
{
    a,
    b,
    c
};


int main() {
    out<int>();
    out(10);
    out(3.14);
    out("a");

    out<test>();
    out<enum_test>();


    return 0;
}