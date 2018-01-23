package arraylist;

import java.util.ArrayList;

public class Main {

	public static void main(String[] args) {
		ArrayList<Integer>int_arr = new ArrayList<Integer>();

		int_arr.add(1);
		int_arr.add(2);
		int_arr.add(3);
		int_arr.add(4);
		int_arr.add(5);
		
		System.out.println(int_arr.get(0));
		
		int_arr.set(0, 10);
		
		System.out.println(int_arr.size());
		
		int_arr.remove(4);
		
		System.out.println(int_arr.toString());
		
		int_arr.clear();
		
		System.out.println(int_arr.toString());
	}

}
