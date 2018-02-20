package proxy;

public class RealValue implements Value{
	
	private String val;
	
	public RealValue(String val) {
		setVal(val);
		save(val);
	}

	private void save(String val) {
		System.out.println("Load "+val);
	}

	@Override
	public void getValue() {
		System.out.println("get Real Value");
	}
	
	public void setVal(String val) {
		this.val = val;
	}
}
