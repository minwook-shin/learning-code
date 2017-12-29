#include<iostream>
#include<algorithm>

using namespace std;

int main() {
	int array_int[10] = { 9,3,7,4,5,6,2,8,1,10 };

	sort(array_int, array_int + 10);

	for (int i = 0; i < 10; i++) {
		cout << array_int[i];
	}
	return 0;
}