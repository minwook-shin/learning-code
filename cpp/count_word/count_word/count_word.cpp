// count_word.cpp: 콘솔 응용 프로그램의 진입점을 정의합니다.
//

#include "stdafx.h"
#include <iostream>

using namespace std;

int main()
{
    char str_sample[] = "hello my name is cpp thank you";
    cout << str_sample << endl;

    int len = strlen(str_sample);

    int count=0;

    for (int i = 0; i < len + 1; i++) {
        if (str_sample[i] != ' '&& str_sample[i] != '\0') {
            count++;
        }
    }
    cout << count << endl;
    return 0;
}

