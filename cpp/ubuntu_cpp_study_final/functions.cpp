#include"functions.h"

management::management()
{
	cout << "생성자 호출" << endl;
}

management::~management()
{
	cout << "소멸자 호출" << endl;
}

int management::wellcome()
{
	int num;
	cout << "입력(1) / 검색(2) / 수정(3) / 삭제(4) / 출력(5) / 정렬(6) / 저장하고 종료(7) :";
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
	cout << "입력 기능" << endl;
	cout << "이름 : ";
	cin >> name;
	cout << "점수 : ";
	cin >> score;
	user_info.push_back(name);
	user_score.push_back(score);
}

void management::search_user()
{
	cout << "누구를 찾으실 것인가요? : " << endl;
	cin >> find_user;
	for (int i = 0; i < user_info.size(); i++)
	{
		if (find_user.find(user_info[i]) != string::npos)
		{
			cout << "찾았습니다!" << endl;
			cout << print_all_info((int)user_score[i], (string)user_info[i], convert_grade((int)user_score[i])) << endl;
		}
	}
}

void management::edit_user()
{
	cout << "누구를 수정하실 것인가요? : " << endl;
	cin >> del_user;
	for (int i = 0; i < user_info.size(); i++)
	{
		if (del_user.find(user_info[i]) != string::npos)
		{
			cout << "수정할 사람을 찾았습니다!" << endl;
			user_score.erase(user_score.begin() + i);
			cout << "어떻게 수정하실 예정인가요?" << endl;
			cin >> edit_score;
			user_score.insert(user_score.begin() + i, (int)edit_score);
			break;
		}
	}
}

void management::delete_user()
{
	cout << "누구를 삭제하실 것인가요? : " << endl;
	cin >> del_user;
	for (int i = 0; i < user_info.size(); i++)
	{
		if (del_user.find(user_info[i]) != string::npos)
		{
			cout << "삭제할 사람을 찾았습니다!" << endl;
			user_info.erase(user_info.begin() + i);
			user_score.erase(user_score.begin() + i);
			break;
		}
	}
}

void management::print_user()
{
	cout << "전체 출력합니다." << endl;
	for (int i = 0; i < user_info.size(); i++)
	{
		cout << print_all_info((int)user_score[i], (string)user_info[i], convert_grade((int)user_score[i])) << endl;
	}
}

void management::sort_user()
{
	cout << "정렬합니다." << endl;
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
	cout << "종료하기 전에 파일에 저장하는 중입니다." << endl;
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