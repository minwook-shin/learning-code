package decorator;

public abstract class Deco implements Linux {
	Linux linux;
	public Deco(Linux linux) {
		this.linux = linux;
	}
	
}
