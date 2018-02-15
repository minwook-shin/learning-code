package mediator;

import java.util.ArrayList;

public class Mediator implements Imediator {
	
	private ArrayList<Colleague> TestArray = new ArrayList<Colleague>();
	
	public Mediator(Colleague e) {
		e.setMediator(this);
        TestArray.add(e);
	}

	@Override
	public void send(String colleague, String event) {
		for (Colleague s : TestArray) {
            if (s.getString() == colleague) {
                    s.receive(colleague, event);
            }
    }
	}

}
