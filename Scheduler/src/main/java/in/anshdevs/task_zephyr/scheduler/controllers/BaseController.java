package in.anshdevs.task_zephyr.scheduler.controllers;

import in.anshdevs.task_zephyr.scheduler.dto.Job;
import in.anshdevs.task_zephyr.scheduler.dto.JobStatus;
import in.anshdevs.task_zephyr.scheduler.service.SchedulerServiceImpl;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.http.HttpHeaders;
import org.springframework.http.HttpStatus;
import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.*;

import java.util.*;

@RestController
public class BaseController {
    @Autowired
    private SchedulerServiceImpl service;

    @GetMapping("/check")
    Map<Object,Object> check(){
        HashMap<Object,Object> map = new HashMap<>();
        map.put("status", "ok");
        map.put("message", "working as intended");
        map.put("data", null);
        return map;
    }

    @PostMapping("/create-job")
    public ResponseEntity<Map<Object,Object>> createJob(@RequestBody Job job){
        job.setStatus(JobStatus.REGISTERED.name());
        var resp = service.CreateJob(job);
        HashMap<Object,Object> map = new HashMap<>();
        map.put("status", "ok");
        map.put("message", "job created successfully");
        map.put("data", resp);
        HttpHeaders headers=new HttpHeaders();
        headers.add("Content-Type","application/json");
        return new ResponseEntity<>(map,headers, HttpStatus.CREATED);
    }

    @GetMapping("/get-job/{id}")

    Map<Object,Object> getJobById(@PathVariable("id") String id){
        var job = service.GetJobById(id);
        HashMap<Object,Object> map=new HashMap<>();
        map.put("status","successful");
        map.put("data",job);
        return map;
    }
}