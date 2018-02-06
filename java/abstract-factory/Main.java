package src;

public class Main {

	public static void main(String[] args) {
		Ifactory test = new Factory();
		
		test.getUbuntu().bash();
		
		test.getWIN().gui();
	}

}
