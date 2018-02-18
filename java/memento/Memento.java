package memento;


public class Memento {
	private String state;

	public String getState() {
		return state;
	}
	
	public void setState(String state) {
		this.state = state;
	}
	
	public Memento(String state) {
		setState(state);
	}
}
