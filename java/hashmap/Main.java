package hashmap;

import java.util.HashMap;

public class Main {

	public static void main(String[] args) {
		HashMap<Integer, String> htest = new HashMap<Integer, String>();
		
		htest.put(0, "number0");
		htest.put(1, "number1");
		htest.put(2, "number2");
		htest.put(3, "number3");
		htest.put(4, "number4");

		System.out.println(htest.toString());
		
		System.out.println(htest.get(4));
		
		htest.remove(0);
		
		System.out.println(htest.toString());
		
		htest.clear();
		
		System.out.println(htest.toString());
	}

}
