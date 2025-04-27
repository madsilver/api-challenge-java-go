package com.silver.apichallenge.entity;

public class Country {
    private String country;
    private Integer total;

    public Country(String country, Integer total) {
        this.country = country;
        this.total = total;
    }

    public String getCountry() {
        return country;
    }

    public void setCountry(String country) {
        this.country = country;
    }

    public Integer getTotal() {
        return total;
    }

    public void setTotal(Integer total) {
        this.total = total;
    }
}