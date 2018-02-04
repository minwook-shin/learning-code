package builder;

public class Computer {
	
	private String HDD;
	private String RAM;

	private boolean GraphicsCard;
	
	public String getHDD() {
		return HDD;
	}

	public String getRAM() {
		return RAM;
	}

	public boolean isGraphicsCard() {
		return GraphicsCard;
	}

	
	Computer(Build builder) {
		this.HDD=builder.HDD;
		this.RAM=builder.RAM;
		this.GraphicsCard=builder.GraphicsCard;
	}

}
