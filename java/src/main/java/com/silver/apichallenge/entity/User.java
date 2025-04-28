package com.silver.apichallenge.entity;

import com.fasterxml.jackson.annotation.JsonProperty;
import java.util.List;

public record User(
    String id,
    @JsonProperty("nome") String name,
    @JsonProperty("idade") Integer age,
    Integer score,
    @JsonProperty("ativo") Boolean active,
    @JsonProperty("pais") String country,
    @JsonProperty("equipe") Team team,
    List<Log> logs
) {}
