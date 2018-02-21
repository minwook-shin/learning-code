package command;

import java.util.ArrayList;

public class Main {
	public static void main(String[] args) {
		ArrayList<Command> test = new ArrayList<Command>();
		
		test.add(new Command() {
			
			@Override
			public String execute() {
				return "COM 1";
			}
		});
		test.add(new Command() {
			
			@Override
			public String execute() {
				return "COM 2";
			}
		});
		test.add(new Command() {
			
			@Override
			public String execute() {
				return "COM 3";
			}
		});
		
	for (Command command : test) {
		System.out.println(command.execute());
	}
	}
}
