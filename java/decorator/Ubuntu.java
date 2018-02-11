package decorator;

public class Ubuntu implements Linux {

	@Override
	public void getVersion() {
		System.out.println("Default version");
		
	}

}
