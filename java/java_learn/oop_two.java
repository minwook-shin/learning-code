package java_learn;

import test_pack.*;

public class oop_two {
	public static void main(String[] args) {
		oop sum_sample = new oop("test",20,201700);
		test test_sample = new test();
		
		test_sample.setTestname("hello, test!");
		
		System.out.println(test_sample.getTestname());
		
		System.out.println(sum_sample.sum(10, 20));
		
		sum_sample.GetInfo();
		
		oop.wallet -= 100; 
		System.out.println(oop.getWallet());
		oop.wallet -= 100; 
		System.out.println(oop.getWallet());
		oop.wallet -= 100; 
		System.out.println(oop.getWallet());
		oop.wallet -= 100; 
		System.out.println(oop.getWallet());
		oop.wallet -= 100; 
		System.out.println(oop.getWallet());
	}
}
