package com.silver.apichallenge.adapter.presenter;

import com.fasterxml.jackson.annotation.JsonProperty;
import java.time.LocalDateTime;
import java.time.format.DateTimeFormatter;

public record Body<T>(
        String timestamp,
        @JsonProperty("execution_time_ms") Long ExecutionTimeMs,
        T data) {
    public String getTimestamp()  {
        return LocalDateTime.now()
                .format(DateTimeFormatter.ofPattern("yyyy-MM-dd HH:mm:ss"));
    }
}
