package com.silver.apichallenge.adapter.presenter;

public record Response<T>(Integer status, Body<T> body) {
    public Response(Integer status, T data, long start) {
        this(status, new Body<T>("", System.currentTimeMillis() - start, data));
    }
}
