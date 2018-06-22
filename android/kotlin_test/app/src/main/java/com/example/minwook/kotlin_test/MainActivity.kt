package com.example.minwook.kotlin_test

import android.support.v7.app.AppCompatActivity
import android.os.Bundle

import kotlinx.android.synthetic.main.activity_main.*

class MainActivity : AppCompatActivity() {

    override fun onCreate(savedInstanceState: Bundle?) {
        super.onCreate(savedInstanceState)
        setContentView(R.layout.activity_main)
        var arr = arrayListOf<Int>()

        result.text = "0"
        saveB.text ="saved number"

        number_text.setText(result.text)

        add_button.setOnClickListener {
            var r:Int = number_text.text.toString().toInt()
            r+=1
            result.text = r.toString()
            number_text.setText(result.text)
        }
        remove_button.setOnClickListener{
            var r:Int = number_text.text.toString().toInt()
            r-=1
            result.text = r.toString()
            number_text.setText(result.text)
        }

        saveB.setOnClickListener {
            arr.add(result.text.toString().toInt())
            printarr.text = arr.toString()
        }
        r.setOnClickListener {
            arr.clear()
            printarr.text ="clear!"
        }
    }
}
