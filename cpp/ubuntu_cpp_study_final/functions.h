#pragma once

#include<iostream>
#include<string>
#include<vector>
#include<fstream>
#include<map>

using namespace std;

class management
{
public:
	management();
	~management();

	int select_num = 0;

	void enter_info();
	void search_user();
	void edit_user();
	void delete_user();
	void print_user();
	void sort_user();
	int wellcome();
	int close_program();

protected:

private:
	string name;
	int score;

	vector<string>user_info = {};
	vector<int>user_score = {};
	map<string, int>sort_info = {};
	map<string, int>::iterator j;
	string find_user;
	string del_user;
	int edit_score;
	string str_score;
	string print_all_info(int score, string name, string grade);
	string convert_grade(int score);
	string result;
};





