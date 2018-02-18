package memento;

import java.util.Stack;

public class Originator {
	private Stack<Memento> list = new Stack<>();
	private String state;
	
	public Memento setState(String state) {
		this.state = state;
		return list.push(new Memento(state));
	}
	
	public Memento getList() {
		return list.pop();
	}
	
	public String getState() {
		return state;
	}
	
	public void restore() {
		Memento val = getList();
		this.state = val.getState();
	}

}
