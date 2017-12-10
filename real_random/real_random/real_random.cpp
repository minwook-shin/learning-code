#include "stdafx.h"
#include <iostream>
#include <random>
using namespace std;

void main(void)
{
    int min = 0;
    int max = 10;
    random_device r;
    mt19937_64 rnd(r());

    uniform_int_distribution<int> range(min, max);

    
    cout << range(rnd) << endl;

}