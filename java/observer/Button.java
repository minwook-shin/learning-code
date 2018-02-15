package observer;

import java.util.Observable;

public class Button extends Observable implements Runnable{

	@Override
	public void run() {
		setChanged();
		notifyObservers();
	}
}
