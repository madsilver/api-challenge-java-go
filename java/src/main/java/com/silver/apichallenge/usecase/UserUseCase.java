package com.silver.apichallenge.usecase;

import com.silver.apichallenge.adapter.repository.UserRepository;
import com.silver.apichallenge.entity.*;
import org.springframework.stereotype.Service;

import java.util.*;
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
                .filter(user -> user.score() > MINIMUM_SCORE && user.active())
                .toList();
    }

    public List<Country> getTopCountries() {
        return this.userRepository.GetUsers()
                .stream()
                .map(User::country)
                .collect(Collectors.groupingBy(country -> country, Collectors.counting()))
                .entrySet()
                .stream()
                .sorted(Map.Entry.comparingByValue(Comparator.reverseOrder()))
                .map(entry -> new Country(entry.getKey(), entry.getValue().intValue()))
                .limit(LIMIT_TOP_COUNTRIES)
                .collect(Collectors.toList());
    }

    public List<TeamInfo> getTeamInsights() {
        return this.userRepository.GetUsers()
                .stream()
                .filter(user -> user.team() != null)
                .collect(Collectors.groupingBy(user -> user.team().name(), Collectors.toList()))
                .entrySet()
                .stream()
                .map(entry -> {
                    String teamName = entry.getKey();
                    List<User> users = entry.getValue();
                    TeamInfo teamInfo = new TeamInfo(teamName);
                    teamInfo.setTotalMembers(users.size());
                    teamInfo.setActivePercentage((float) users.stream().filter(User::active).count() / users.size() * 100);
                    teamInfo.setLeaders(users.stream().filter(user -> user.team().leader()).count());
                    teamInfo.setCompletedProjects(users.stream().flatMap(user -> user.team().projects().stream()).filter(Project::completed).count());
                    return teamInfo;
                })
                .toList();
    }

    public List<TeamInfo> getTeamInsights2() {
        Map<String, TeamInfo> map = new HashMap<>();
        List<User> users = this.userRepository.GetUsers();
        for (User user : users) {
            if (user.team() == null) {
                continue;
            }
            String teamName = user.team().name();
            if (!map.containsKey(teamName)) {
                map.put(teamName, new TeamInfo(teamName));
            }
            TeamInfo teamInfo = map.get(teamName);
            teamInfo.incTotalMembers();
            if (user.team().leader()) {
                teamInfo.incLeaders();
            }
            if (user.active()) {
                teamInfo.incActivePercentage();
            }
            if (user.team().projects() == null) {
                continue;
            }
            for (Project project : user.team().projects()) {
                if (project.completed()) {
                    teamInfo.incCompletedProjects();
                }
            }
        }
        List<TeamInfo> teams = new ArrayList<>();
        for (Map.Entry<String, TeamInfo> entry : map.entrySet()) {
            TeamInfo teamInfo = entry.getValue();
            teamInfo.setActivePercentage(teamInfo.getActivePercentage() / teamInfo.getTotalMembers() * 100);
            teams.add(teamInfo);
        }
        return teams;
    }
}
