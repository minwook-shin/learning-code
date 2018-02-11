package decorator;

public class NewVersionDeco extends Deco {

	public NewVersionDeco(Linux linux) {
		super(linux);
		// TODO Auto-generated constructor stub
	}

	@Override
	public void getVersion() {
		System.out.println("NEW version");
	}

}
