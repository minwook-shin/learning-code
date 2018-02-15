package mediator;

public class Main {

	public static void main(String[] args) {
		Com_1 test1 = new Com_1();
		Com_2 test2 = new Com_2();
		
		Mediator test_M1 = new Mediator(test1);
		Mediator test_M2 = new Mediator(test2);
		
		test_M1.send("COM 1","hello world!");
		test_M2.send("COM 2","hello world!");
		
		test1.send("hello world!");
		test2.send("hello world!");
	}

}
