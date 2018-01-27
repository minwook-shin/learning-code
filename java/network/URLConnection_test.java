package network;

import java.io.IOException;
import java.io.InputStream;
import java.net.MalformedURLException;
import java.net.URL;
import java.net.URLConnection;
import java.util.Scanner;

public class URLConnection_test {

	public static void main(String[] args) {
		URL url;
		InputStream in_s;
		Scanner in_scan;
		String code;
		try {
			url = new URL("http://www.google.com");
			URLConnection con = url.openConnection();
			in_s = con.getInputStream();
			in_scan = new Scanner(in_s);
			
			while(in_scan.hasNext()) {
				code = in_scan.next();
				System.out.println(code);
				}
			in_scan.close();
		} catch (MalformedURLException e) {
			// TODO Auto-generated catch block
			e.printStackTrace();
		} catch (IOException e) {
			// TODO Auto-generated catch block
			e.printStackTrace();
		}
		
	}
}