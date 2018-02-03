package com.main;

public class single_instance {
	public single_instance() {
		single_class test = single_class.getinstance();
		System.out.println(test.getI());
		test.setI(500);
		System.out.println(test.getI());
	}
}
