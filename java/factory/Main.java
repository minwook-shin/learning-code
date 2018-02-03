package factory;

public class Main {

	public static void main(String[] args) {
		Factory test1 = new Factory();
		Factory test2 = new Factory();
		
		System.out.println(test1.select("car", 700000, "compact"));
		System.out.println(test2.select("phone", 200000, "compact"));
		
	}

}
