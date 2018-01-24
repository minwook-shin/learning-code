package ithread;

public class Itest implements Runnable {
	
	int sum = 0;

	@Override
	public void run() {
		System.out.println(Thread.currentThread().getName());
		System.out.println("test_thread");
		
		for (int i = 0; i < 10; i++) {
			if(Thread.currentThread().getName()=="first")
				sum++;
			System.out.println(i+" "+sum);
			try {
				Thread.sleep(500);
			} catch (InterruptedException e) {
				// TODO Auto-generated catch block
				e.printStackTrace();
			}
		}
	}

}
