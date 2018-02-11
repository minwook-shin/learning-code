package visitor;

interface Visitor {
	void visit(Car car);
	void visit(Body car);
	void visit(Engine car);
}
