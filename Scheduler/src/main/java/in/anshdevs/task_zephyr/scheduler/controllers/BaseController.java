package in.anshdevs.task_zephyr.scheduler.controllers;

import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.RestController;

import java.util.HashMap;
import java.util.Map;

@RestController
public class BaseController {
    @GetMapping("/check")
    Map<String,String> check(){
        HashMap<String, String> map = new HashMap<>();
        map.put("status", "ok");
        map.put("message", "working as intended");
        map.put("data", null);
        return map;
    }
}