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
		cout << "��ǰ�� �����մϴ�" << endl;
	else
		cout << "��ǰ�� �����ϴ�" << endl;

	if (5000 <= money) {
		cout << "��ǰ�� �����ϼ̽��ϴ�." << endl;
		money -= 5000;
		cout << "���� �� : "<< money << endl;
	}
	else if (1000 <= money && money < 5000) {
		cout << "�� ��üǰ�� �����ϼ̽��ϴ�." << endl;
		money -= 1000;
		cout << "���� �� : " << money << endl;
	}
	else
		cout << "���� �����մϴ�." << endl;
	return 0;
}