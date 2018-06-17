package com.example.minwook.calc;

import android.support.v7.app.AppCompatActivity;
import android.os.Bundle;
import android.view.View;
import android.widget.EditText;
import android.widget.TextView;

public class MainActivity extends AppCompatActivity {

    @Override
    protected void onCreate(Bundle savedInstanceState) {
        super.onCreate(savedInstanceState);
        setContentView(R.layout.activity_main);
    }

    public void add_click(View v){
        EditText a = findViewById(R.id.number1);
        EditText b = findViewById(R.id.number2);

        TextView result = findViewById(R.id.result);

        int n1 = Integer.parseInt(a.getText().toString());
        int n2 = Integer.parseInt(b.getText().toString());
        result.setText(Integer.toString(n1+n2));
    }

    public void sub_click(View v){
        EditText a = findViewById(R.id.number1);
        EditText b = findViewById(R.id.number2);

        TextView result = findViewById(R.id.result);

        int n1 = Integer.parseInt(a.getText().toString());
        int n2 = Integer.parseInt(b.getText().toString());
        result.setText(Integer.toString(n1-n2));
    }
    public void mul_click(View v){
        EditText a = findViewById(R.id.number1);
        EditText b = findViewById(R.id.number2);

        TextView result = findViewById(R.id.result);

        int n1 = Integer.parseInt(a.getText().toString());
        int n2 = Integer.parseInt(b.getText().toString());
        result.setText(Integer.toString(n1*n2));
    }
    public void div_click(View v){
        EditText a = findViewById(R.id.number1);
        EditText b = findViewById(R.id.number2);

        TextView result = findViewById(R.id.result);

        int n1 = Integer.parseInt(a.getText().toString());
        int n2 = Integer.parseInt(b.getText().toString());
        result.setText(Integer.toString(n1/n2));
    }
}
