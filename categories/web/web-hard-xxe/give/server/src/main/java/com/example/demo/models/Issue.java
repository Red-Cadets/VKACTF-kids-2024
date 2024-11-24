package com.example.demo.models;

import java.util.UUID;

import org.springframework.data.redis.core.RedisHash;

@RedisHash(value = "Issue", timeToLive = 60L)
public class Issue {

    private UUID id = UUID.randomUUID();
    private String message;
    private String email;

    public String getEmail() {
        return email;
    }

    public void setEmail(String email) {
        this.email = email;
    }

    public Issue() {
    }

    public Issue(final String message, final String email) {
        this.id = UUID.randomUUID();
        this.message = message;
        this.email = email;
    }

    public UUID getId() {
        return id;
    }

    public String getMessage() {
        return message;
    }

    public void setMessage(String message) {
        this.message = message;
    }

}
