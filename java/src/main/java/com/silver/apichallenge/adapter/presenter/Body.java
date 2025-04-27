package com.silver.apichallenge.adapter.presenter;

import com.fasterxml.jackson.annotation.JsonProperty;
import java.time.LocalDateTime;
import java.time.format.DateTimeFormatter;

public class Body<T> {
    private final String timestamp;
    private final Long ExecutionTimeMs;
    private final T data;

    public Body(T data, Long executionTimeMs) {
        this.ExecutionTimeMs = executionTimeMs;
        this.data = data;
        this.timestamp = LocalDateTime.now()
                .format(DateTimeFormatter.ofPattern("yyyy-MM-dd HH:mm:ss"));
    }

    public String getTimestamp() {
        return timestamp;
    }

    @JsonProperty("execution_time_ms")
    public Long getExecutionTimeMs() {
        return ExecutionTimeMs;
    }

    public T getData() {
        return data;
    }
}
