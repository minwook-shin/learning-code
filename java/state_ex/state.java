package state_ex;

public class state {
	
	IState State;
	
	IState test_state1 = new IState() {
		
		@Override
		public void on() {
			System.out.println("on");
			State = test_state2;
		}
		
		@Override
		public void off() {
			System.out.println("none");
		}
	};
	
	IState test_state2 = new IState() {
		
		@Override
		public void on() {
			System.out.println("none");
		}
		
		@Override
		public void off() {
			System.out.println("off");
			State = test_state1;
		}
	};
	
	public state() {
		State = test_state1;
	}
	
	public void on() {
		State.on();
	}
	
	public void off() {
		State.off();
	}
	
}

interface IState {
	
	void on();
	void off();
	
}

