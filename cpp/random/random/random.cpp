#include<iostream>
#include<time.h>

using namespace std;

int main() {
	srand((unsigned int)time(0));

	cout << "random : " << (rand() % 100 + 1) << endl;

	double percent;
	percent = rand() % 10000 / 100.f;
	cout << "random : " << percent << endl;

	if(percent <= 80.f)
		cout << "in 80%" << endl;
	if (percent <= 60.f)
		cout << "in 60%" << endl;
	if (percent <= 40.f)
		cout << "in 40%" << endl;
	if (percent <= 20.f)
		cout << "in 20%" << endl;
	if (percent <= 1.f)
		cout << "in 1%" << endl;
	return 0;
}