package com.silver.apichallenge.adapter.controller;

import com.fasterxml.jackson.core.type.TypeReference;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.silver.apichallenge.adapter.presenter.Response;
import com.silver.apichallenge.adapter.presenter.UserCreate;
import com.silver.apichallenge.entity.Country;
import com.silver.apichallenge.entity.TeamInfo;
import com.silver.apichallenge.entity.User;
import com.silver.apichallenge.usecase.UserUseCase;
import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.*;
import org.springframework.web.multipart.MultipartFile;

import java.io.IOException;
import java.util.List;

@RestController
public class UserController {
    private UserUseCase userUseCase;

    public UserController(UserUseCase userUseCase) {
        this.userUseCase = userUseCase;
    }

    @PostMapping("/users")
    public ResponseEntity<Response<UserCreate>> saveUsers(@RequestParam("file") MultipartFile file) {
        long start = System.currentTimeMillis();
        ObjectMapper mapper = new ObjectMapper();
        TypeReference<List<User>> typeReference = new TypeReference<List<User>>(){};
        List<User> users;
        try {
            users = mapper.readValue(file.getInputStream(), typeReference);
            this.userUseCase.saveUsers(users);
        } catch (IOException e) {
            throw new RuntimeException(e);
        }
        UserCreate userCreate = new UserCreate("File received successfully", users.size());
        return ResponseEntity.ok(new Response<>(200, userCreate, start));
    }

    @GetMapping("/superusers")
    public ResponseEntity<Response<List<User>>> getUsers() {
        long start = System.currentTimeMillis();
        List<User> users = this.userUseCase.getSuperUsers();
        return ResponseEntity.ok(new Response<>(200, users, start));
    }

    @GetMapping("/top-countries")
    public ResponseEntity<Response<List<Country>>> getTopCountries() {
        long start = System.currentTimeMillis();
        List<Country> countries = this.userUseCase.getTopCountries();
        return ResponseEntity.ok(new Response<>(200, countries, start));
    }

    @GetMapping("/team-insights")
    public ResponseEntity<Response<List<TeamInfo>>> getTeamInsights() {
        long start = System.currentTimeMillis();
        List<TeamInfo> teamInfo = this.userUseCase.getTeamInsights();
        return ResponseEntity.ok(new Response<>(200, teamInfo, start));
    }
}
