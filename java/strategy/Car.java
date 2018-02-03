package com.strategy;

public class Car {
	Ifuel fuel;
	Isize size;
	
	public Car(Ifuel fuel, Isize size) {
		this.fuel = fuel;
		this.size = size;
	}
	
	public void dofuel() {
		fuel.dofuel();
	}
	public void dosize() {
		size.dosize();
	}
}
