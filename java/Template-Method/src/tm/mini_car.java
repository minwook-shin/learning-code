package tm;

public class mini_car extends Car {

	@Override
	protected void fuel() {
		System.out.println("oil");
	}

	@Override
	protected void consum() {
		System.out.println("low");
	}

}
