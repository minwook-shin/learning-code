#include<iostream>

using namespace std;

class parent {
public:
    parent() {
        cout << "calling parent" << endl;
    }
    ~parent() {
        cout << "deleting parent" << endl;
    }
    int a;
protected:
    int b;
private:
    int c;
};

class child : public parent{
public:
    child() {
        cout << "calling child" << endl;
    }
    ~child() {
        cout << "deleting child" << endl;
    }
protected:
    int d;
};

class child_two : private parent {
public:
    child_two() {
        cout << "calling 2nd child" << endl;
    }
    ~child_two() {
        cout << "deleting 2nd child" << endl;
    }
private:
    int d;
};

int main() {
    parent sample_p;
    child sample_c;
    child_two sample_two;


    return 0;
}