package list.data.structure;

public class Main {

	public static void main(String[] args) {
		Arraylist test = new Arraylist();
		
		test.add_last(10);
		test.add_last(20);
		test.add_last(30);
		test.add_last(50);
		
		test.add(3,40);
		
		test.add_first(0);
		
		System.out.println(test.toString());
		
		test.remove(4);
		
		System.out.println(test.toString());
		
		test.remove_first();
		
		System.out.println(test.toString());
		
		test.remove_last();
		
		System.out.println(test.toString());
		
		System.out.println(test.get(0));
		System.out.println(test.get(1));
		System.out.println(test.get(2));
		
		System.out.println(test.size());
		
		System.out.println(test.indexOf(30));
		
		Arraylist.Iterator it = test.iterator();
		
		System.out.println(it.hasnext());
		System.out.println(it.next());
		System.out.println(it.hasnext());
		System.out.println(it.next());
		System.out.println(it.hasnext());
		System.out.println(it.next());
		System.out.println(it.hasnext());
		System.out.println(it.next());
		
		System.out.println(it.hasprev());
		System.out.println(it.prev());
		System.out.println(it.hasprev());
		System.out.println(it.prev());
		System.out.println(it.hasprev());
		System.out.println(it.prev());
		System.out.println(it.hasprev());
		System.out.println(it.prev());
		System.out.println(it.hasprev());
		
		it.add(100);
		System.out.println(test.toString());
		
		it.remove();
		System.out.println(test.toString());
		
	}

}
