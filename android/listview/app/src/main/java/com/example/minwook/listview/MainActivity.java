package com.example.minwook.listview;

import android.support.v7.app.AppCompatActivity;
import android.os.Bundle;
import android.view.View;
import android.widget.AdapterView;
import android.widget.ArrayAdapter;
import android.widget.ListAdapter;
import android.widget.ListView;
import android.widget.Toast;

public class MainActivity extends AppCompatActivity {

    @Override
    protected void onCreate(Bundle savedInstanceState) {
        super.onCreate(savedInstanceState);
        setContentView(R.layout.activity_main);
        String[] arr = {"안드로이드 폰","아이폰","윈도우 폰"};
        ListAdapter listAdapter = new ArrayAdapter<>(this,android.R.layout.simple_list_item_1,arr);
        ListView li = findViewById(R.id.list);
        li.setAdapter(listAdapter);
        li.setOnItemClickListener(new AdapterView.OnItemClickListener() {
            @Override
            public void onItemClick(AdapterView<?> adapterView, View view, int i, long l) {
                String item = String.valueOf(adapterView.getItemAtPosition(i));
                Toast.makeText(MainActivity.this,item,Toast.LENGTH_LONG).show();
            }
        });
    }

}
