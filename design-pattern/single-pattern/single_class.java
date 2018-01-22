package com.main;

public class single_class {
	private static single_class instance = new single_class();
	public int i = 100;
	
	private single_class() {
		
	}
	
	public static single_class getinstance() {
		if(instance ==null)
			instance = new single_class();
		return instance;
	}

	public int getI() {
		return i;
	}

	public void setI(int i) {
		this.i = i;
	}
}
