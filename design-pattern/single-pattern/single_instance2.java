package com.main;

public class single_instance2 {
	public single_instance2() {
		single_class test = single_class.getinstance();
		System.out.println(test.getI());
	}
}
