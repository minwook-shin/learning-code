package cthread;

public class Main {

	public static void main(String[] args) {
		ctest test = new ctest();
		
		test.start();
		System.out.println(Thread.currentThread().getName());
	}

}
