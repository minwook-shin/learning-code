package list.linked;

public class Main {

	public static void main(String[] args) {
		Linkedlist test = new Linkedlist();
		test.add_fisrt(0);
		test.add_fisrt(10);
		test.add_fisrt(20);
		test.add_fisrt(30);
		
		test.add_last(10);
		test.add_last(20);
		test.add_last(30);

		test.add(2, 0);
		
		System.out.println(test.toString());
		
		test.remove_first();
		
		System.out.println(test.toString());
		
		test.remove(1);
		
		System.out.println(test.toString());
		
		test.remove_last();
		
		System.out.println(test.toString());
		
		System.out.println(test.size());
		
		System.out.println(test.get(3));
		
		System.out.println(test.indexOf(0));
		
		Linkedlist.List_it it = test.iterator();
		
		it.next();
		it.next();
		it.next();
		it.next();
	
		System.out.println(it.hasnext());
		System.out.println(test.toString());
	}

}
