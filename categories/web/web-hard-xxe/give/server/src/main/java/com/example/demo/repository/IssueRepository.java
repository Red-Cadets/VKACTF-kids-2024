package com.example.demo.repository;

import org.springframework.data.repository.CrudRepository;
import org.springframework.stereotype.Repository;

import com.example.demo.models.Issue;

@Repository
public interface IssueRepository extends CrudRepository<Issue, String> {

}
