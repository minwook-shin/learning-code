#include<iostream>

using namespace std;

struct student
{
    int score;
};

int main() {
    int i = 0;
    int * pointer= &i;

    *pointer += 100;

    cout <<"i value : "<< i << endl;
    cout <<"i address : " << &i << endl;
    cout << "pointer value : " << pointer << endl;
    cout << "pointer address : " << &pointer << endl;

    int a[10] = { 1,2,3,4,5,6,7,8,9,10 };

    cout << "array address : "<<a << endl;
    cout << "array address : " << &a << endl;
    cout << "array[0] address : " << &a[0] << endl; 
    int *array_pointer = &a[0];
    cout << "array pointer[3] value : " << array_pointer[3] << endl;

    cout << "array pointer address : " << &array_pointer << endl;
    cout << "array pointer address : " << &array_pointer + 2<< endl;

    char *text = "text string";

    cout << "char pointer value : " << text << endl;
    cout << "char pointer address : " << (int*)text << endl;

    text = "text text string";

    cout << "char pointer value : " << text << endl;
    cout << "char pointer address : " << (int*)text << endl;

    text = "text string";

    cout << "char pointer value : " << text << endl;
    cout << "char pointer address : " << (int*)text << endl;

    student korean_student = {};
    korean_student.score = 100;

    student *pointer_student = &korean_student;

    (*pointer_student).score = 70;
    pointer_student->score = 80;
    cout << "korean_student's value : " << korean_student.score << endl;
    cout << "korean_student's address : " << &korean_student.score << endl;

    void *v_pointer = &i;
    int *convert = (int *)v_pointer;
    *convert += 1000;
    cout << "i value : " << i << endl;
    cout << "void pointer i : " << v_pointer<< endl;
    v_pointer = &korean_student.score;
    cout << "void pointer score : " << v_pointer << endl;

    int myint = 1000;
    int *pint = &myint;
    int **ppint = &pint;

    cout << "myint : " << myint << endl;
    cout << "&myint : " << &myint << endl;

    cout << "pint : " << pint << endl;
    cout << "*pint : " << *pint << endl;
    cout << "&pint : " << &pint << endl;

    cout << "ppint : " << ppint << endl;
    cout << "*ppint : " << *ppint << endl;
    cout << "**ppint : " << **ppint << endl;
    cout << "&ppint : " << &ppint << endl;

    return 0;
}