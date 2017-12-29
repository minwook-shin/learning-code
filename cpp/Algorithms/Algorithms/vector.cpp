#include<iostream>
#include<algorithm>
#include<vector>

using namespace std;

int main() {
	int array_int[] = { 4,7,3,6,8,1,2,9,5 };
	vector<int> vec(array_int,array_int+9);

	for (int i = 0; i < 9; i++) {
		cout << vec[i] << "\t";
	}

	sort(vec.begin(), vec.end());

	cout << endl;
	for (int i = 0; i < 9; i++) {
		cout << vec[i] << "\t";
	}

	reverse(vec.begin(), vec.end());

	cout << endl;
	for (int i = 0; i < 9; i++) {
		cout << vec[i] << "\t";
	}

	cout << endl;
	cout << *min_element(vec.begin(), vec.end()) << endl;
	cout << *max_element(vec.begin(), vec.end()) << endl;

	vec.push_back(10);
	vec.push_back(10);
	cout << count(vec.begin(), vec.end(), 10) << endl;

	if (find(vec.begin(), vec.end(), 1) != vec.end()) {
		cout << "found number" << endl;
	}
	cout << endl;
	for (int i = 0; i < 11; i++) {
		cout << vec[i] << "\t";
	}

	vec.erase(vec.end()-1);
	cout << endl;
	for (int i = 0; i < 10; i++) {
		cout << vec[i] << "\t";
	}

	vec.erase(vec.end() - 1);
	cout << endl;
	for (int i = 0; i < 9; i++) {
		cout << vec[i] << "\t";
	}
	
	vec.insert(vec.begin()+5,77);
	cout << endl;
	for (int i = 0; i < 10; i++) {
		cout << vec[i] << "\t";
	}
	return 0;
}