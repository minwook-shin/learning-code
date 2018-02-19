package flyweightpattern;

public class Main {

	public static void main(String[] args) {

		Factory factory = new Factory();
		
		Java test = factory.getInstance("JAVA");
		test.exe();
		test = factory.getInstance("JAVA");
		test.exe();
	}

}
