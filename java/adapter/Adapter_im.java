package adapter;

public class Adapter_im implements Adapter {
	Calc adapter = new Calc();
	

	@Override
	public double div(double d, double i) {
		return adapter.div(d, i);
	}

	@Override
	public double mul(double d, double i) {
		return adapter.mul(d, i);
	}

	@Override
	public double add(double d, double i) {
		return d+i;
	}

}
