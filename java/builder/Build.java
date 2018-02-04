package builder;

public class Build {
			String HDD;
			String RAM;

			boolean GraphicsCard;
			
			public Build(String hdd, String ram){
				this.HDD=hdd;
				this.RAM=ram;
			}

			public Build setGraphicsCard(boolean GraphicsCard) {
				this.GraphicsCard = GraphicsCard;
				return this;
			}
			
			public Computer build(){
				return new Computer(this);
			}
}
