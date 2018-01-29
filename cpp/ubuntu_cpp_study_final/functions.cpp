#include"functions.h"

management::management()
{
	cout << "������ ȣ��" << endl;
}

management::~management()
{
	cout << "�Ҹ��� ȣ��" << endl;
}

int management::wellcome()
{
	int num;
	cout << "�Է�(1) / �˻�(2) / ����(3) / ����(4) / ���(5) / ����(6) / �����ϰ� ����(7) :";
	cin >> num;
	return num;
}

string management::print_all_info(int score, string name, string grade)
{
	str_score = to_string(score);
	string blank = " ";
	string full = name + blank + str_score + blank + grade;
	return full;
}

void management::enter_info()
{
	cout << "�Է� ���" << endl;
	cout << "�̸� : ";
	cin >> name;
	cout << "���� : ";
	cin >> score;
	user_info.push_back(name);
	user_score.push_back(score);
}

void management::search_user()
{
	cout << "������ ã���� ���ΰ���? : " << endl;
	cin >> find_user;
	for (int i = 0; i < user_info.size(); i++)
	{
		if (find_user.find(user_info[i]) != string::npos)
		{
			cout << "ã�ҽ��ϴ�!" << endl;
			cout << print_all_info((int)user_score[i], (string)user_info[i], convert_grade((int)user_score[i])) << endl;
		}
	}
}

void management::edit_user()
{
	cout << "������ �����Ͻ� ���ΰ���? : " << endl;
	cin >> del_user;
	for (int i = 0; i < user_info.size(); i++)
	{
		if (del_user.find(user_info[i]) != string::npos)
		{
			cout << "������ ����� ã�ҽ��ϴ�!" << endl;
			user_score.erase(user_score.begin() + i);
			cout << "��� �����Ͻ� �����ΰ���?" << endl;
			cin >> edit_score;
			user_score.insert(user_score.begin() + i, (int)edit_score);
			break;
		}
	}
}

void management::delete_user()
{
	cout << "������ �����Ͻ� ���ΰ���? : " << endl;
	cin >> del_user;
	for (int i = 0; i < user_info.size(); i++)
	{
		if (del_user.find(user_info[i]) != string::npos)
		{
			cout << "������ ����� ã�ҽ��ϴ�!" << endl;
			user_info.erase(user_info.begin() + i);
			user_score.erase(user_score.begin() + i);
			break;
		}
	}
}

void management::print_user()
{
	cout << "��ü ����մϴ�." << endl;
	for (int i = 0; i < user_info.size(); i++)
	{
		cout << print_all_info((int)user_score[i], (string)user_info[i], convert_grade((int)user_score[i])) << endl;
	}
}

void management::sort_user()
{
	cout << "�����մϴ�." << endl;
	for (int i = 0; i < user_info.size(); i++)
	{
		sort_info.insert(map<string, int>::value_type((string)user_info[i], (int)user_score[i]));
	}
	for (j = sort_info.begin();j != sort_info.end(); j++) {
		cout << "[" << j->first << "] " << j->second << endl;
	}
}

int management::close_program()
{
	ofstream files("save_user.txt");
	cout << "�����ϱ� ���� ���Ͽ� �����ϴ� ���Դϴ�." << endl;
	for (int i = 0; i <user_info.size(); i++) {
		files << print_all_info((int)user_score[i], (string)user_info[i], convert_grade((int)user_score[i])) << "\n";
	}
	files.close();
	exit(0);
	return 0;
}

string management::convert_grade(int score)
{
	if (score >= 90)
	{
		result = "A+";
	}
	else if (90 >score&& score >= 50)
	{
		result = "A";
	}
	else if (50 >score&&  score >= 20)
	{
		result = "B";
	}
	else 
	{
		result = "F";
	}
	return result;
}