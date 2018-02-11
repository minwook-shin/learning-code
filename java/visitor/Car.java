package visitor;

public class Car implements CarAccept{
	
	CarAccept[] list;
	
	public Car() {
		this.list = new CarAccept[] {
	            new Body(), new Engine()
	        };

	}

	@Override
	public void accept(Visitor visit) {
		for (CarAccept v : list) {
            v.accept(visit);;
        }
        visit.visit(this);
	}

}
