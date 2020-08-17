package com.tutorial.spring;

import org.springframework.web.bind.annotation.*;

import java.util.HashMap;
import java.util.Map;

@RestController
public class TestController {

    @RequestMapping(value = "/api", method = RequestMethod.GET)
    public Map<String, Object> getTest() {
        Map<String, Object> result = new HashMap<>();
        result.put("application_name", "testing");
        result.put("update_date", "2020.01.01");
        return result;
    }

}


