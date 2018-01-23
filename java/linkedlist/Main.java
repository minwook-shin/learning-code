package linkedlist;

import java.util.LinkedList;

public class Main {

	public static void main(String[] args) {
		LinkedList<Integer> linked_int = new LinkedList<Integer>();
		
		linked_int.add(1);
		linked_int.add(2);
		linked_int.add(3);
		linked_int.add(4);
		linked_int.add(5);
		
		System.out.println(linked_int.toString());
		
		linked_int.add(0, 10);
		
		System.out.println(linked_int.toString());
		
		linked_int.set(4, 50);
		
		System.out.println(linked_int.size());
		
		linked_int.remove(0);
		
		System.out.println(linked_int.toString());
		
		linked_int.clear();
		
		System.out.println(linked_int.toString());
	}

}
