package java_learn;

import java.util.Scanner;

public class java_learn {

	public static void main(String[] args) {
		System.out.println("hello world!");
		System.out.print("입력 : ");
		Scanner scanner = new Scanner(System.in);
		int i = scanner.nextInt();
		
		System.out.print("입력 : ");
		System.out.println(i);
		
		if(i<=1) {
			System.out.println("음수입니다.");
		}else if(i>=i) {
			System.out.println("양수입니다.");
		}
		else {
			System.out.println("0입니다.");
		}
		
		switch (i) {
		case 0:
			System.out.println("hello, zero!");
			break;

		default:
			break;
		}
		
		for(int j = 0;j<10;j++) {
			System.out.println(j);
		}
		
		int array[] = {10,20,30,40};
		
		System.out.println(array[0]);
		
		int array_two[] = new int[5];
		
		int array_three[][] = new int[5][5];
		
		for(int k=0;k<array_two.length;k++) {
			array_two[k] = k;
			System.out.println(array_two[k]);
		}
		
		for (int l = 0; l < array_three.length; l++) {
			for (int n = 0; n < array_three.length; n++) {
				array_three[l][n]=l*n;
			}
		}
		for (int m = 0; m < array_three.length; m++) {
			for (int o = 0; o < array_three.length; o++) {
				System.out.print(array_three[m][o]);
			}
			System.out.println("");
		}
		
		
		
	}

}
