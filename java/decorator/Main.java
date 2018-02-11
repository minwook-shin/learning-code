package decorator;

public class Main {

	public static void main(String[] args) {
		Linux test1 = new Ubuntu();
		test1.getVersion();
		
		Linux test2 = new NewVersionDeco(new Ubuntu());
		test2.getVersion();
		
		
	}

}
