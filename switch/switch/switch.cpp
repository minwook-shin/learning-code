#include<iostream>

using namespace std;

#define num_4 3
enum e {
	num_1,
	num_2,
	num_3
};

int main() {
	int a =0;
	cout << "enter number(1-3) : ";
	cin >> a;

	e num_5 = (e)4;

	switch(a){
	case num_1:
		cout << "press 0" << endl;
		break;
	case num_2:
		cout << "press 1" << endl;
		break;
	case num_3:
		cout << "press 2" << endl;
		break;
	case num_4:
		cout << "press 3" << endl;
		break;
	default:
		cout << "other number" << endl;
		break;
	}

	cout << "and num_5 is... " << num_5 << endl;
	
	return 0;
}