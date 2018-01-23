package hashset;

import java.util.HashSet;

public class Main {

	public static void main(String[] args) {
		HashSet<Integer> hint = new HashSet<Integer>();
		
		hint.add(0);
		hint.add(4);
		hint.add(2);
		hint.add(5);
		hint.add(2);
		hint.add(5);
		hint.add(1);
		hint.add(0);
		
		System.out.println(hint.toString());
		
		hint.remove(0);
		
		System.out.println(hint.toString());
		
		System.out.println(hint.size());

	}

}
