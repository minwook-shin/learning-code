#include<iostream>

using namespace std;

struct node
{
    int data;
    node *next;
};
class list {
private:
    node *head, *tail;
public:
    list() {
        head = NULL;
        tail = NULL;
        cout << "preparing linked list..." << endl;
    }
    void display() {
        node *temp = new node;
        temp = head;
        while (temp!=NULL)
        {
            cout <<"|"<< temp->data << "|" ;
            temp = temp->next;
        }
        cout << endl;
    }
    void create(int value) {
        node *temp = new node;
        temp->data = value;
        temp->next = NULL;
        if (head == NULL) {
            head = temp;
            tail = temp;
            temp = NULL;
        }
        else {
            tail->next = temp;
            tail = temp;
        }
        cout << "creating " <<value<< endl;
    }
    void insert(int value) {
        node *temp = new node;
        temp->data = value;
        temp->next = head;
        head = temp;
        cout << "inserting " << value << endl;
    }
    void insert_any(int pos, int value) {
        node *prev = new node;
        node *curr = new node;
        node *temp = new node;
        curr = head;
        for (int i = 1; i < pos; i++) {
            prev = curr;
            curr = curr->next;
        }
        temp->data = value;
        prev->next = temp;
        temp->next = curr;
        cout << "inserting " << value<< "\t position:"<< pos << endl;
    }
    void del_first() {
        node *temp = new node;
        temp = head;
        head = head->next;
        delete temp;
        cout << "deleting first number" << endl;
    }
    void del_end() {
        node *prev = new node;
        node *curr = new node;
        curr = head;
        while (curr->next != NULL) {
            prev = curr;
            curr = curr->next;
        }
        tail = prev;
        prev->next = NULL;
        delete curr;
        cout << "deleting end number" << endl;
    }
    void del_any(int pos) {
        node *curr = new node;
        node *prev = new node;
        curr = head;
        for (int i = 1; i < pos; i++) {
            prev = curr;
            curr = curr->next;
        }
        prev->next = curr->next;
        cout << "deleting " << pos << "\t position:" << pos << endl;
    }
};
int main() {
    list a;
    a.create(5);
    a.create(10);
    a.create(15);
    a.create(20);
    a.insert(1);
    a.insert_any(3,100);
    a.display();
    a.del_first();
    a.display();
    a.del_end();
    a.display();
    a.del_any(2);
    a.display();

    return 0;
}