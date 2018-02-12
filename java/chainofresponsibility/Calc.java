package chainofresponsibility;

public abstract class Calc {
	
	private Calc calc;
	
	abstract protected boolean operator(Request request);

	public boolean process(Request request) {
		if (operator(request)) {
			return true;
		} else {
			if (calc != null)
				return calc.operator(request);
		}
		return false;
	}
}
