package memento;

public class Main {

	public static void main(String[] args) {
		Originator originator = new Originator();
		
		originator.setState("state 1");
		originator.setState("state 2");
		originator.setState("state 3");

		originator.restore();
		System.out.println(originator.getState());
		originator.restore();
		System.out.println(originator.getState());
		originator.restore();
		System.out.println(originator.getState());
	}

}
