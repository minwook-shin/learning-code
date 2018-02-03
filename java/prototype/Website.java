package prototype;

public class Website extends Clone{
	
	private int password;

	public Website(int password) {
		super();
		this.password = password;
	}

	public void setPassword(int password) {
		this.password = password;
	}
	
	public int getPassword() {
		return password;
	}
	
	Clone copy() throws CloneNotSupportedException {
		Clone Website = (Clone) clone();
		return Website;
	}

}
