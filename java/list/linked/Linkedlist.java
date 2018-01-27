package list.linked;

public class Linkedlist {
	private Node head;
	private Node tail;
	private int size = 0;
	
	private class Node {
		private Object data;
		private Node next;
		
		public Node(Object e) {
			this.data = e;
			this.next = null;
		}

	}

	public boolean add_fisrt(Object e) {
		Node new_node = new Node(e);
		new_node.next = head;
		head = new_node;
		size++;
		if(head.next == null)
			tail = head;
		return true;
	}
	
	public boolean add_last(Object e) {
		Node new_node = new Node(e);
		if(size == 0)
			add_fisrt(e);
		else {
			tail.next = new_node;
			tail = new_node;
			size++;
		}
		
		return true;
	}
	
	protected Node node(int index) {
		Node n = head;
		for(int i = 0; i<index;i++)
			n = n.next;
		return n;
	}
	 
	 public boolean add(int index, Object e) {
		 if(index == 0)
			 add_fisrt(e);
		 else {
			Node temp = node(index-1);
			Node temp_next = node(index);
			Node new_node = new Node(e);
			temp.next = new_node;
			new_node.next = temp_next;
			size++;
			if(new_node.next == null)
				tail = new_node;
		 }
		 
		 return true;
	 }
	 
	 @Override
	 public String toString() {
		 if(head == null)
			 return "[]";
		 else {
			 	Node temp = head;
			 	String str = "[";
			 	while(temp.next != null) {
			 		str += temp.data + ",";
			 		temp = temp.next;
			 		if(temp == tail)
			 			str+= temp.data;
			 	}
			 	
			 	return str + "]";
		 }
		}
	 
	 public boolean remove_first() {
		 Node temp = head;
		 head = head.next;
		 temp = null;
		 size--;
		 
		 return true;
	 }
	 
	 public boolean remove(int index) {
		 if(index == 0)
			 return remove_first();
		 else {
			 Node temp = node(index-1);
			 Node del = temp.next;
			 temp.next = temp.next.next;
			 if(del == tail)
				 tail = temp;
			 del = null;
			 size--;
		 }
		 
		 return true;
	 }
	 
	 public boolean remove_last() {
		 return remove(size-1);
	 }
	 
	 public int size() {
		 return size;
	 }
	 
	 public Object get(int i) {
		 Node temp = node(i);
		 
		 return temp.data;
	 }
	 public int indexOf(Object e) {
		 Node temp = head;
		 int index = 0;
		 while(temp.data != e) {
			 temp = temp.next;
			 index ++;
			 if(temp == null)
				 return -1;
		 }
		 return index;
	 }
	 
	 public List_it iterator() {
		 return new List_it();
	 }
	 
	 public class List_it{
		 private Node next;
		 private Node current;
		 private int next_i;
		 List_it(){
			next = head; 
		 }
		 public Object next() {
			 current = next;
			 next = next.next;
			 next_i++;
			 
			 return current.data;
		 }
		 
		 public boolean hasnext() {
			 return next_i < size;
		 }
	 }
	 
}
