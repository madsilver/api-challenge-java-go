package com.silver.apichallenge.entity;

import com.fasterxml.jackson.annotation.JsonProperty;

import java.util.List;

public record Team (
    @JsonProperty("nome") String name,
    @JsonProperty("lider") Boolean leader,
    @JsonProperty("projetos") List<Project> projects
) {}
