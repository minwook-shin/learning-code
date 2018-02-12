package chainofresponsibility;

public class Main {

	public static void main(String[] args) {
		Calc plus = new PlusCalc();

		Request request1 = new Request(1,2,"+");
		Request request2 = new Request(10,2,"-");

		plus.process(request1);
		plus.process(request2);
		}
}