package list.data.structure;

public class Arraylist {
	private Object e_data[] = new Object[100];
	private int size = 0;
	
	public boolean add_last(Object e){
		e_data[size] = e;
		size ++;
		
		return true;
	}

	public boolean add(int index, Object e) {
		for(int i = size-1; i >= index;i--) {
			e_data[i+1] = e_data[i];
		}
		e_data[index] = e;
		size++;
		
		return true;
	}
	
	public boolean add_first(Object e) {
		return add(0, e);
	}
	
	@Override
	public String toString() {
		String text = "[";
		for (int i = 0; i < size; i++) {
			text += e_data[i];
			if(i< size-1)
				text += ",";
		}
		return text + "]";
	}

	public boolean remove(int index) {
		for(int i = index + 1;i<=size-1;i++ ) {
			e_data[i-1] = e_data[i];
		}
		size--;
		e_data[size] = null;
		return true;
	}

	public boolean remove_first() {
		return remove(0);
	}
	public boolean remove_last() {
		return remove(size -1 );
	}

	public Object get(int index) {
		return e_data[index];
	}
	
	public int size() {
		return size;
	}

	public int indexOf(Object e) {
		for(int i =0; i<size; i++) {
			if(e.equals(e_data[i])){
				return i;
			}
		}
		return -1;
	}

	public Iterator iterator() {
		
		return new Iterator();
	}
	class Iterator{
		private int next_i =0;
		public Object next() {
			Object return_val = e_data[next_i];
			next_i++;
			return return_val;
		}
		public boolean hasnext() {
			return next_i < size;
		}
		
		public Object prev() {
			next_i--;
			Object return_val = e_data[next_i];
			return return_val;
		}
		public boolean hasprev() {
			return next_i > 0;
		}
		public boolean add(Object e) {
			Arraylist.this.add(next_i, e);
			next_i++;
			return true;
		}
		public boolean remove() {
			Arraylist.this.remove(next_i-1);
			next_i--;
			return true;
		}
	}
}
