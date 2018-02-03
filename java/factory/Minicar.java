package factory;

public class Minicar extends Super {
	
	private int price;
	private String size;
	
	Minicar(int price, String size){
		this.price = price;
		this.size = size;
	}
	
	public int getPrice() {
		return price;
	}

	public void setFuel(int price) {
		this.price = price;
	}

	public String getSize() {
		return size;
	}

	public void setSize(String size) {
		this.size = size;
	}

	@Override
	public void price() {
		// TODO Auto-generated method stub
		
	}

	@Override
	public void size() {
		// TODO Auto-generated method stub
		
	}
	
	@Override
	public String toString() {
		String strprice = String.valueOf(this.price);
		return strprice + this.size;
	}

}
