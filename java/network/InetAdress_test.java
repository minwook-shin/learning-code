package network;

import java.net.InetAddress;
import java.net.UnknownHostException;

public class InetAdress_test {

	public static void main(String[] args) {
		InetAddress test = null;
		try {
			test = InetAddress.getLocalHost();
		} catch (UnknownHostException e) {
			// TODO Auto-generated catch block
			e.printStackTrace();
		}
		System.out.println(test.getHostName() + " " + test.getHostAddress());
		InetAddress google_address = null;
		try {
			google_address = InetAddress.getByName("www.google.com");
		} catch (UnknownHostException e) {
			// TODO Auto-generated catch block
			e.printStackTrace();
		}
		System.out.println(google_address);
	}
	
}
