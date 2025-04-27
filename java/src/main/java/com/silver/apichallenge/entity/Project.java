package com.silver.apichallenge.entity;

import com.fasterxml.jackson.annotation.JsonProperty;

public class Project {
    @JsonProperty("nome")
    private String name;
    @JsonProperty("concluido")
    private Boolean completed;

    public String getName() {
        return name;
    }

    public void setName(String name) {
        this.name = name;
    }

    public Boolean getCompleted() {
        return completed;
    }

    public void setCompleted(Boolean completed) {
        this.completed = completed;
    }
}
