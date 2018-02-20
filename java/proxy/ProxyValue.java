package proxy;

public class ProxyValue implements Value{
	
	private String val;
	RealValue realvalue;
	
	public ProxyValue(String val) {
		setVal(val);
	}

	@Override
	public void getValue() {
		if(realvalue == null)
			realvalue = new RealValue(val);
		realvalue.getValue();
	}
	
	public void setVal(String val) {
		this.val = val;
	}

}
