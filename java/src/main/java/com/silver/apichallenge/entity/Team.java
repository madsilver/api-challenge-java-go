package com.silver.apichallenge.entity;

import com.fasterxml.jackson.annotation.JsonProperty;

import java.util.List;

public class Team {
    @JsonProperty("nome")
    private String name;
    @JsonProperty("lider")
    private Boolean leader;
    @JsonProperty("projetos")
    private List<Project> projects;

    public String getName() {
        return name;
    }

    public void setName(String name) {
        this.name = name;
    }

    public Boolean getLeader() {
        return leader;
    }

    public void setLeader(Boolean leader) {
        this.leader = leader;
    }

    public List<Project> getProjects() {
        return projects;
    }

    public void setProjects(List<Project> projects) {
        this.projects = projects;
    }
}
