package com.tutorial.spring;

import org.springframework.stereotype.Controller;
import org.springframework.ui.Model;
import org.springframework.web.bind.annotation.RequestMapping;

@Controller
public class PageController {
    @RequestMapping(value = "/home")
    public String index(Model model){
        model.addAttribute("name", "testing");
        return "index";
    }
}
