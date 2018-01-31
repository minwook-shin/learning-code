#include"functions.h"

int main() 
{
	ubuntu_kr_cpp_study_final::management main;
	
	while (true)
	{
		main.select_num = main.wellcome();
		switch (main.select_num)
		{
		case 1:
			main.enter_info();
			break;
		case 2:
			main.search_user();
			break;
		case 3:
			main.edit_user();
			break;
		case 4:
			main.delete_user();
			break;
		case 5:
			main.print_user();
			break;
		case 6:
			main.sort_user();
			break;
		case 7:
			main.close_program();
		default:
			exit(0);
		}
	}
	return 0;
}