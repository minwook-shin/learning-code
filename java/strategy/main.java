package com.strategy;

public class main {

	public static void main(String[] args) {
		Car minicar = new Car(new C_oil(), new C_size());
		minicar.dofuel();
		minicar.dosize();
	}

}
