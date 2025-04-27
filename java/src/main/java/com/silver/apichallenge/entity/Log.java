package com.silver.apichallenge.entity;

import com.fasterxml.jackson.annotation.JsonProperty;

public class Log {
    @JsonProperty("data")
    private String date;
    @JsonProperty("acao")
    private String action;

    public String getDate() {
        return date;
    }

    public void setDate(String date) {
        this.date = date;
    }

    public String getAction() {
        return action;
    }

    public void setAction(String action) {
        this.action = action;
    }
}
