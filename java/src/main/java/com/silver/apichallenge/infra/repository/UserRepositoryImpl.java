package com.silver.apichallenge.infra.repository;

import com.silver.apichallenge.adapter.repository.UserRepository;
import com.silver.apichallenge.entity.User;
import org.springframework.stereotype.Repository;

import java.util.ArrayList;
import java.util.List;

@Repository
public class UserRepositoryImpl implements UserRepository {
    private List<User> users;

    public UserRepositoryImpl() {
        this.users = new ArrayList<>();
    }

    @Override
    public void StoreUser(List<User> users) {
        this.users = users;
    }

    @Override
    public List<User> GetUsers() {
        return this.users;
    }
}
