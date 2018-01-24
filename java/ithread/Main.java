package ithread;

public class Main {

	public static void main(String[] args) {
		Itest test = new Itest();
		Itest test2 = new Itest();
		Thread test_thread = new Thread(test,"first");
		Thread test_thread2 = new Thread(test2,"second");
		test_thread.start();
		test_thread2.start();
		System.out.println(Thread.currentThread().getName());
	}

}
