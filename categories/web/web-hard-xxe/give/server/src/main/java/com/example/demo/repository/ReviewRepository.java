package com.example.demo.repository;

import org.springframework.data.repository.CrudRepository;
import org.springframework.stereotype.Repository;

import com.example.demo.models.Review;

@Repository
public interface ReviewRepository extends CrudRepository<Review, String> {

}
