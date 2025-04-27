package com.silver.apichallenge.usecase;

import com.silver.apichallenge.adapter.repository.UserRepository;
import com.silver.apichallenge.entity.Country;
import com.silver.apichallenge.entity.User;
import org.springframework.stereotype.Service;

import java.util.Comparator;
import java.util.List;
import java.util.Map;
import java.util.stream.Collectors;

@Service
public class UserUseCase {
    private final UserRepository userRepository;
    private final int MINIMUM_SCORE = 900;
    private final int LIMIT_TOP_COUNTRIES = 5;

    public UserUseCase(UserRepository userRepository) {
        this.userRepository = userRepository;
    }

    public void saveUsers(List<User> users) {
        this.userRepository.StoreUser(users);
    }

    public List<User> getSuperUsers() {
        return this.userRepository.GetUsers()
                .stream()
                .filter(user -> user.getScore() > MINIMUM_SCORE && user.getActive())
                .toList();
    }

    public List<Country> getTopCountries() {
        return this.userRepository.GetUsers()
                .stream()
                .map(User::getCountry)
                .collect(Collectors.groupingBy(country -> country, Collectors.counting()))
                .entrySet()
                .stream()
                .sorted(Map.Entry.comparingByValue(Comparator.reverseOrder()))
                .map(entry -> new Country(entry.getKey(), entry.getValue().intValue()))
                .limit(LIMIT_TOP_COUNTRIES)
                .collect(Collectors.toList());
        /*
        return this.userRepository.GetUsers()
            .stream()
            .map(User::getCountry)
            .collect(Collectors.groupingBy(country -> country, Collectors.counting()))
            .entrySet()
            .stream()
            .sorted(Map.Entry.<String, Long>comparingByValue(Comparator.reverseOrder()))
            .map(entry -> new Country(entry.getKey(), entry.getValue().intValue()))
            .toList(); // Java 16+

        List<Country> countries = new ArrayList<>();

        this.userRepository.GetUsers()
            .stream()
            .map(User::getCountry)
            .collect(Collectors.groupingBy(country -> country, Collectors.counting()))
            .forEach((country, count) -> {
                countries.add(new Country(country, count.intValue()));
            });

        return countries;
        */
    }
}
