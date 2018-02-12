package chainofresponsibility;

public class Request {
	
	private int i,j;
	
	private String string;

	public Request(int i, int j, String string) {
		super();
		this.i = i;
		this.j = j;
		this.string = string;
	}

	public int getI() {
		return i;
	}

	public void setI(int i) {
		this.i = i;
	}

	public int getJ() {
		return j;
	}

	public void setJ(int j) {
		this.j = j;
	}

	public String getString() {
		return string;
	}

	public void setString(String string) {
		this.string = string;
	}
	
	

}
