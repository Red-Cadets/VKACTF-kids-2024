package com.example.demo.controllers;

import java.io.IOException;

import org.apache.coyote.BadRequestException;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.http.MediaType;
import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.CrossOrigin;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.PostMapping;
import org.springframework.web.bind.annotation.RequestBody;
import org.springframework.web.bind.annotation.RestController;

import com.example.demo.models.Issue;
import com.example.demo.models.Review;
import com.example.demo.repository.IssueRepository;
import com.example.demo.repository.ReviewRepository;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.JsonMappingException;
import com.fasterxml.jackson.dataformat.xml.XmlMapper;

@RestController
@CrossOrigin(origins = "*")
public class IssueController {

    @Autowired
    private IssueRepository issueRepository;
    @Autowired
    private ReviewRepository reviewRepository;
    String xmlRes;

    @GetMapping("/review")
    public ResponseEntity<?> getAll() throws JsonMappingException, JsonProcessingException {
        try {
            Iterable<Review> reviews = reviewRepository.findAll();
            xmlRes = "<?xml version=\"1.0\" encoding=\"UTF-8\"?><Reviews>";
            reviews.forEach(review -> {
                if (review != null) {
                    String message = review.getMessage() != null ? review.getMessage() : "";

                    xmlRes += "<Review><id>" + review.getId().toString() + "</id><rating>"
                            + String.valueOf(review.getRating()) + "</rating><message>" + message + "</message><email>"
                            + review.getEmail() + "</email></Review>";
                }
            });
            xmlRes += "</Reviews>";
            return ResponseEntity.ok().contentType(MediaType.APPLICATION_XML).body(xmlRes);
        } catch (Exception e) {
            return ResponseEntity.badRequest().contentType(MediaType.APPLICATION_XML).body(e.toString()
                    + "<?xml version=\"1.0\" encoding=\"UTF-8\"?><message>Не удалось обработать запрос</message>");
        }

    }

    @PostMapping("/review")
    public ResponseEntity<?> createReview(@RequestBody String str)
            throws JsonMappingException, JsonProcessingException, BadRequestException, IOException {
        try {
            if (str.toLowerCase().contains("http://")
                    || str.toLowerCase().contains("https://")
                    || str.toLowerCase().contains("file://")
                    || str.toLowerCase().contains("//")) {
                return ResponseEntity.badRequest().contentType(MediaType.APPLICATION_XML).body(
                        "<?xml version=\"1.0\" encoding=\"UTF-8\"?><message>Ссылки не допустимы.</message>");
            }
            XmlMapper xmlMapper = new XmlMapper();
            Review value = xmlMapper.readValue(str, Review.class);
            if (!value.getEmail().matches("^[a-zA-Z0-9_!#$%&'*+/=?`{|}~^.-]+@[a-zA-Z0-9.-]+$")) {
                return ResponseEntity.badRequest().contentType(MediaType.APPLICATION_XML).body(
                        "<?xml version=\"1.0\" encoding=\"UTF-8\"?><message>Введите валидный email</message>");
            }
            if (!value.getMessage().matches("^[a-z A-Z 0-9 , ; . ! ?]+$")) {
                throw new BadRequestException(
                        "<?xml version=\"1.0\" encoding=\"UTF-8\"?><message>Введите валидный текст без специальных символов.</message>");
            }
            reviewRepository.save(value);
            return ResponseEntity.ok().contentType(MediaType.APPLICATION_XML).body(
                    "<?xml version=\"1.0\" encoding=\"UTF-8\"?><message>Благодарим за Ваш отзыв!</message>");
        } catch (JsonMappingException e) {
            return ResponseEntity.badRequest().contentType(MediaType.APPLICATION_XML).body(
                    "<?xml version=\"1.0\" encoding=\"UTF-8\"?><message>Не удалось обработать запрос</message>");
        }

    }

    @PostMapping("/issue")
    public ResponseEntity<?> createIssue(@RequestBody String str)
            throws JsonMappingException, JsonProcessingException, IOException {
        try {
            XmlMapper xmlMapper = new XmlMapper();
            Issue value = xmlMapper.readValue(str, Issue.class);
            if (!value.getEmail().matches("^[a-zA-Z0-9_!#$%&'*+/=?`{|}~^.-]+@[a-zA-Z0-9.-]+$")) {
                return ResponseEntity.badRequest().contentType(MediaType.APPLICATION_XML).body(
                        "<?xml version=\"1.0\" encoding=\"UTF-8\"?><message>Введите валидный email</message>");
            }
            if (!value.getMessage().matches("^[a-z A-Z 0-9 , ; . ! ?]+$")) {
                throw new BadRequestException(
                        "<?xml version=\"1.0\" encoding=\"UTF-8\"?><message>Введите валидный текст без специальных символов.</message>");
            }
            issueRepository.save(value);
            System.out.println(value.getId());
            return ResponseEntity.ok().contentType(MediaType.APPLICATION_XML).body(
                    "<?xml version=\"1.0\" encoding=\"UTF-8\"?><message>Благодарим за Вашу обратную связь. Мы ответим Вам в ближайшее время.</message>");
        } catch (JsonMappingException e) {
            return ResponseEntity.badRequest().contentType(MediaType.APPLICATION_XML).body(
                    "<?xml version=\"1.0\" encoding=\"UTF-8\"?><message>Не удалось обработать запрос</message>");
        }

    }

}
