#include<iostream>

using namespace std;

int main() {
    FILE *pfile = NULL;

    fopen_s(&pfile, "test.txt", "wt");

    if (pfile != NULL) {
        char *ptext = "hello";
        fwrite(ptext,1,5,pfile);
        fclose(pfile);
        cout << "making file" << endl;
    }
    fopen_s(&pfile, "test.txt", "rt");
    if (pfile) {
        char str[6] = {};
        fread(str, 1, 5, pfile);
        cout << "reading file : "<< str << endl;
    }

    return 0;
}