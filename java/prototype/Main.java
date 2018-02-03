package prototype;

public class Main {

	public static void main(String[] args) throws CloneNotSupportedException {
		Website test = new Website(1234);
		Website test_copy = (Website) test.copy();
		
		System.out.println(test.getId()+"\n"+test_copy.getId() + "\n" + test.getPassword()+ "\n" + test_copy.getPassword());
	}

}
