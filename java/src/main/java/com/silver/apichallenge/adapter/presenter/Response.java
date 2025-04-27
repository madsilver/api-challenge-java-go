package com.silver.apichallenge.adapter.presenter;

public class Response<T> {
    private Integer status;
    private Body<T> body;

    public Response(Integer status, T data, long start) {
        this.status = status;
        long executionTimeMs = System.currentTimeMillis() - start;
        this.body = new Body<T>(data,  executionTimeMs);
    }

    public Integer getStatus() {
        return status;
    }

    public Body<T> getBody() {
        return body;
    }
}
