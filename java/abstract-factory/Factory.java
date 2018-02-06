package src;

public class Factory implements Ifactory {

	@Override
	public Ubuntu getUbuntu() {
		return new LTS();
	}

	@Override
	public Windows getWIN() {
		return new XP();
	}
	
}
