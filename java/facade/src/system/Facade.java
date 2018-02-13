package system;

public class Facade {
	
	public Facade() {
		System.out.println("calling facade");
	}
	COM_1 com1 = new COM_1();
	COM_2 com2 = new COM_2();
	COM_3 com3 = new COM_3();
		
	public void boot() {
		com1.RaidCOM1();
		com2.RaidCOM2();
		com3.RaidCOM3();
	}

}
