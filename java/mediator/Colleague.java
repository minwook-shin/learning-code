package mediator;

public abstract class Colleague {
	Imediator mediator;
	
	public void setMediator(Imediator mediator) {
		this.mediator = mediator;
		}
	
	public void send(String name, String event) {
        mediator.send(name, event);
        }
	
	abstract public void send(String event);
	  
    abstract public void receive(String name, String event);
    
    abstract public String getString();
}
