package factory;

public class Factory {
	public Super select(String class_name, int i, String size) {
		if("car".equals(class_name)) {
			return new Minicar(i, size);
			}
		else if("phone".equals(class_name)){
			return new Miniphone(i, size);
			}
		return null;
	}
}
