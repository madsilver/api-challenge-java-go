package com.silver.apichallenge.entity;

public class TeamInfo {
    private String team;
    private long totalMembers;
    private long leaders;
    private long completedProjects;
    private float activePercentage;



    public TeamInfo(String team) {
        this.team = team;
    }

    public String getTeam() {
        return team;
    }

    public void setTeam(String team) {
        this.team = team;
    }

    public long getTotalMembers() {
        return totalMembers;
    }

    public void incTotalMembers() {
        this.totalMembers++;
    }

    public long getLeaders() {
        return leaders;
    }

    public void incLeaders() {
        this.leaders++;
    }

    public long getCompletedProjects() {
        return completedProjects;
    }

    public void incCompletedProjects() {
        this.completedProjects++;
    }

    public float getActivePercentage() {
        return activePercentage;
    }

    public void incActivePercentage() {
        this.activePercentage++;
    }

    public void setActivePercentage(float activePercentage) {
        this.activePercentage = activePercentage;
    }

    public void setTotalMembers(long totalMembers) {
        this.totalMembers = totalMembers;
    }

    public void setLeaders(long leaders) {
        this.leaders = leaders;
    }

    public void setCompletedProjects(long completedProjects) {
        this.completedProjects = completedProjects;
    }
}
