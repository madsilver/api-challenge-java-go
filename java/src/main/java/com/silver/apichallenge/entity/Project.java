package com.silver.apichallenge.entity;

import com.fasterxml.jackson.annotation.JsonProperty;

public record Project (
    @JsonProperty("nome") String name,
    @JsonProperty("concluido") Boolean completed
) {}