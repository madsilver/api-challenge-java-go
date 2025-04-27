package com.silver.apichallenge.adapter.presenter;

public class UserCreate {
    private String message;
    private int userCount;

    public UserCreate(String message, int userCount) {
        this.message = message;
        this.userCount = userCount;
    }

    public String getMessage() {
        return message;
    }

    public int getUserCount() {
        return userCount;
    }
}
