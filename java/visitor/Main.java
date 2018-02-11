package visitor;

public class Main {

	public static void main(String[] args) {
		Car test = new Car();
		
		test.accept(new Implements_visit());
	}

}
