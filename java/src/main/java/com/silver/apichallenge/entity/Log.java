package com.silver.apichallenge.entity;

import com.fasterxml.jackson.annotation.JsonProperty;

public record Log (
    @JsonProperty("data") String date,
    @JsonProperty("acao") String action
) {}
