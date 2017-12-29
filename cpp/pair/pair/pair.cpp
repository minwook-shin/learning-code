#include<iostream>
#include<vector>

using namespace std;

int main() 
{
	vector < pair<int, int>> vec;
	vec.push_back(make_pair(1, 2));
	cout << vec[0].first<<" "<<vec[0].second << endl;
	return 0;
}