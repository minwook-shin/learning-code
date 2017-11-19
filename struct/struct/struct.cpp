#include<iostream>

using namespace std;
#define size 64
struct student
{
    char name[size];
    int num;
    int kor;
    int eng;
    int math;
};

int main() {
    student korean_student = {};

    strcpy_s(korean_student.name, "minsoo");
    strcat_s(korean_student.name, " kim");
    korean_student.num = 20171600;
    korean_student.kor = 100;
    korean_student.eng = 100;
    korean_student.math = 50;


    if (strcmp(korean_student.name, "minsoo kim") == 0) {
        cout << "found student :  " << korean_student.name << endl;
    }


    cout << "name : " << korean_student.name << endl;
    cout << "number : " << korean_student.num<< endl;
    cout << "korean : " << korean_student.kor << endl;
    cout << "english : " << korean_student.eng << endl;
    cout << "math : " << korean_student.math << endl;


    return 0;
}