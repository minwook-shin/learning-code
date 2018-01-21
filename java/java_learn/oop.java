package java_learn;

public class oop {
	private int age;
	private String name;
	private int num;
	public static int wallet = 1000;
	
	public static int getWallet() {
		return wallet;
	}

	public static void setWallet(int wallet) {
		oop.wallet = wallet;
	}

	protected int sum(int i,int j) {
		int result;
		result = i+j;
		return result;
	}
	
	public void GetInfo() {
		System.out.println(name+age+num);
	}
	
	public oop() {
		System.out.println("생성자 호출!");
	}
	
	public oop(String name,int age,int num) {
		this.name= name;
		this.age = age;
		this.num = num;
		System.out.println("생성자 호출!!");
	}

	public int getAge() {
		return age;
	}

	public void setAge(int age) {
		this.age = age;
	}

	public String getName() {
		return name;
	}

	public void setName(String name) {
		this.name = name;
	}

	public int getNum() {
		return num;
	}

	public void setNum(int num) {
		this.num = num;
	}
}