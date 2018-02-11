package visitor;

public class Implements_visit implements Visitor {

	@Override
	public void visit(Car car) {
		System.out.println("visit car");
		
	}

	@Override
	public void visit(Body car) {
		System.out.println("visit body");
		
	}

	@Override
	public void visit(Engine car) {
		System.out.println("visit engine");
		
	}

}
