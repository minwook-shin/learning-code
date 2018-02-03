package tm;

public class super_car extends Car {

	@Override
	protected void fuel() {
		System.out.println("oil");
	}

	@Override
	protected void consum() {
		System.out.println("high");
	}

}
