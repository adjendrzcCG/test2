package com.example.ewf.controller;

import com.example.ewf.model.Product;
import com.example.ewf.repository.ProductRepository;
import com.fasterxml.jackson.databind.ObjectMapper;
import org.junit.jupiter.api.BeforeEach;
import org.junit.jupiter.api.Test;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.boot.test.autoconfigure.web.servlet.AutoConfigureMockMvc;
import org.springframework.boot.test.context.SpringBootTest;
import org.springframework.http.MediaType;
import org.springframework.test.web.servlet.MockMvc;

import java.math.BigDecimal;

import static org.hamcrest.Matchers.*;
import static org.springframework.test.web.servlet.request.MockMvcRequestBuilders.*;
import static org.springframework.test.web.servlet.result.MockMvcResultMatchers.*;

@SpringBootTest
@AutoConfigureMockMvc
class ProductControllerIntegrationTest {

    @Autowired
    private MockMvc mockMvc;

    @Autowired
    private ObjectMapper objectMapper;

    @Autowired
    private ProductRepository productRepository;

    @BeforeEach
    void setUp() {
        productRepository.deleteAll();
    }

    @Test
    void getAll_returnsEmptyList() throws Exception {
        mockMvc.perform(get("/api/products"))
                .andExpect(status().isOk())
                .andExpect(jsonPath("$", hasSize(0)));
    }

    @Test
    void create_validProduct_returns201() throws Exception {
        Product product = new Product("Widget", "A nice widget", new BigDecimal("4.99"), 100);
        mockMvc.perform(post("/api/products")
                        .contentType(MediaType.APPLICATION_JSON)
                        .content(objectMapper.writeValueAsString(product)))
                .andExpect(status().isCreated())
                .andExpect(jsonPath("$.id", notNullValue()))
                .andExpect(jsonPath("$.name", is("Widget")));
    }

    @Test
    void getById_existingProduct_returnsProduct() throws Exception {
        Product saved = productRepository.save(
                new Product("Gadget", "Cool gadget", new BigDecimal("29.99"), 50));
        mockMvc.perform(get("/api/products/" + saved.getId()))
                .andExpect(status().isOk())
                .andExpect(jsonPath("$.name", is("Gadget")));
    }

    @Test
    void getById_nonExisting_returns404() throws Exception {
        mockMvc.perform(get("/api/products/9999"))
                .andExpect(status().isNotFound());
    }

    @Test
    void update_existingProduct_returnsUpdated() throws Exception {
        Product saved = productRepository.save(
                new Product("Old Name", "Old Desc", new BigDecimal("1.00"), 1));
        Product updated = new Product("New Name", "New Desc", new BigDecimal("2.00"), 2);
        mockMvc.perform(put("/api/products/" + saved.getId())
                        .contentType(MediaType.APPLICATION_JSON)
                        .content(objectMapper.writeValueAsString(updated)))
                .andExpect(status().isOk())
                .andExpect(jsonPath("$.name", is("New Name")));
    }

    @Test
    void delete_existingProduct_returns204() throws Exception {
        Product saved = productRepository.save(
                new Product("ToDelete", "Will be deleted", new BigDecimal("0.99"), 1));
        mockMvc.perform(delete("/api/products/" + saved.getId()))
                .andExpect(status().isNoContent());
    }

    @Test
    void create_invalidProduct_returns400() throws Exception {
        Product invalid = new Product("", null, new BigDecimal("-1.00"), -1);
        mockMvc.perform(post("/api/products")
                        .contentType(MediaType.APPLICATION_JSON)
                        .content(objectMapper.writeValueAsString(invalid)))
                .andExpect(status().isBadRequest());
    }
}
