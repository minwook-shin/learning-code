package composite;

public class Main {

	public static void main(String[] args) {
		 Folder test = new Folder();
		 
		 File t = new File();
		 
		 test.add(t);
		 test.add(t);
		 test.add(t);
		 
		 test.print();
	}

}
