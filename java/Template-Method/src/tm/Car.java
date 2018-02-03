package tm;

public abstract class Car {

	public void engine() {
		System.out.println("preparing...");
		fuel();
		consum();
		System.out.println("running car!");
	}
	
	protected abstract void fuel();
	protected abstract void consum();
}
