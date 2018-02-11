package visitor;

public class Body implements CarAccept {

	@Override
	public void accept(Visitor visit) {
		visit.visit(this);

	}

}
