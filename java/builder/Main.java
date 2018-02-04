package builder;

public class Main {

	public static void main(String[] args) {
		Computer test = new Build("hdd 1G", "ram 4G").setGraphicsCard(true).build();
		
		System.out.println(test.isGraphicsCard());
	}

}
