package visitor;

public class Engine implements CarAccept {

	@Override
	public void accept(Visitor visit) {
		visit.visit(this);
		
	}

}
