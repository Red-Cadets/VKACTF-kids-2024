package com.example.demo.models;

import java.util.UUID;

import org.springframework.data.redis.core.RedisHash;

import com.fasterxml.jackson.annotation.JsonProperty;

@RedisHash(value = "Review", timeToLive = 60L)
public class Review {
    @JsonProperty("id")
    private UUID id = UUID.randomUUID();
    @JsonProperty("rating")
    private int rating;
    @JsonProperty("email")
    private String email;
    @JsonProperty("message")
    private String message;

    public String getMessage() {
        return message;
    }

    public void setMessage(String message) {
        this.message = message;
    }

    public Review() {

    }

    public Review(final int rating, final String email, final String message) {
        this.rating = rating;
        this.email = email;
        this.message = message;
    }

    public UUID getId() {
        return id;
    }

    public String getEmail() {
        return email;
    }

    public void setEmail(String email) {
        this.email = email;
    }

    public int getRating() {
        return rating;
    }

    public void setRating(int rating) {
        this.rating = rating;
    }
}
