package com.silver.apichallenge.adapter.repository;

import com.silver.apichallenge.entity.User;

import java.util.List;

public interface UserRepository {
    void StoreUser(List<User> users);
    List<User> GetUsers();
}
