package composite;

import java.util.ArrayList;
import java.util.List;

public class Folder implements Feature {
	private List<Feature> private_folder = new ArrayList<Feature>();

	@Override
	public void print() {
		for(Feature i : private_folder) {
			i.print();
		}
	}
	
	public void add(Feature t) {
		private_folder.add(t);
	}
}
