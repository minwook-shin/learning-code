package observer;

import java.util.Observable;
import java.util.Observer;

public class Main {

	public static void main(String[] args) {
		Button button = new Button();
		button.addObserver(new Observer() {
			@Override
			public void update(Observable o, Object arg) {
				System.out.println(o+" is Clicked");
			}
		});

		button.run();
		button.run();
		button.run();
		
	}

}
