package chainofresponsibility;

public class PlusCalc extends Calc {

	@Override
	protected boolean operator(Request request) {
		if(request.getString().equals("+")){
			int a = request.getI();
			int b = request.getJ();
			int result = a + b; 
			System.out.println(result);
		}
		return false;
	}

}
