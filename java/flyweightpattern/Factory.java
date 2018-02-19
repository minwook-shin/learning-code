package flyweightpattern;

import java.util.HashMap;

public final class Factory {
	
	private HashMap<String, Java>map = new HashMap<String, Java>();
	
	public Java getInstance(String t){
		Java instance = map.get(t);
		if(instance == null && t == "JAVA"){
			instance = new Java(); 	   
			}
		map.put(t, instance);
		return instance;
	}
}
