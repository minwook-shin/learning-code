package flyweightpattern;

public class Java implements Iexe {

	public Java(){
		System.out.println("Java created");
	}
	
	@Override
	public void exe() {
		System.out.println("execute");
	}

}
