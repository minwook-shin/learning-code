#include<iostream>

using namespace std;

int main() {
	int goods;
	int money;

	cout << "enter buy goods : ";
	cin >> goods;
	cout << "enter your money : ";
	cin >> money;

	if (1 <= goods)
		cout << "상품이 존재합니다" << endl;
	else
		cout << "상품이 없습니다" << endl;

	if (5000 <= money) {
		cout << "상품을 구매하셨습니다." << endl;
		money -= 5000;
		cout << "남은 돈 : "<< money << endl;
	}
	else if (1000 <= money && money < 5000) {
		cout << "싼 대체품을 구매하셨습니다." << endl;
		money -= 1000;
		cout << "남은 돈 : " << money << endl;
	}
	else
		cout << "돈이 부족합니다." << endl;
	return 0;
}