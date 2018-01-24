package cthread;

public class ctest extends Thread{
	@Override
	public void run() {
		System.out.println(Thread.currentThread().getName());
		System.out.println("test_thread");
		
		for (int i = 0; i < 10; i++) {
			System.out.println(++i);
			try {
				Thread.sleep(500);
			} catch (InterruptedException e) {
				// TODO Auto-generated catch block
				e.printStackTrace();
			}
		}
		}
	}
