package com.example.ewf.service;

import com.example.ewf.exception.ResourceNotFoundException;
import com.example.ewf.model.Product;
import com.example.ewf.repository.ProductRepository;
import org.junit.jupiter.api.BeforeEach;
import org.junit.jupiter.api.Test;
import org.junit.jupiter.api.extension.ExtendWith;
import org.mockito.InjectMocks;
import org.mockito.Mock;
import org.mockito.junit.jupiter.MockitoExtension;

import java.math.BigDecimal;
import java.util.Arrays;
import java.util.List;
import java.util.Optional;

import static org.assertj.core.api.Assertions.*;
import static org.mockito.ArgumentMatchers.any;
import static org.mockito.Mockito.*;

@ExtendWith(MockitoExtension.class)
class ProductServiceTest {

    @Mock
    private ProductRepository productRepository;

    @InjectMocks
    private ProductService productService;

    private Product product;

    @BeforeEach
    void setUp() {
        product = new Product("Test Product", "Test Description", new BigDecimal("9.99"), 10);
        product.setId(1L);
    }

    @Test
    void findAll_returnsAllProducts() {
        when(productRepository.findAll()).thenReturn(Arrays.asList(product));
        List<Product> result = productService.findAll();
        assertThat(result).hasSize(1);
        assertThat(result.get(0).getName()).isEqualTo("Test Product");
    }

    @Test
    void findById_existingId_returnsProduct() {
        when(productRepository.findById(1L)).thenReturn(Optional.of(product));
        Product result = productService.findById(1L);
        assertThat(result).isNotNull();
        assertThat(result.getId()).isEqualTo(1L);
    }

    @Test
    void findById_nonExistingId_throwsException() {
        when(productRepository.findById(99L)).thenReturn(Optional.empty());
        assertThatThrownBy(() -> productService.findById(99L))
                .isInstanceOf(ResourceNotFoundException.class)
                .hasMessageContaining("99");
    }

    @Test
    void create_validProduct_returnsSavedProduct() {
        when(productRepository.save(any(Product.class))).thenReturn(product);
        Product result = productService.create(product);
        assertThat(result).isNotNull();
        verify(productRepository, times(1)).save(product);
    }

    @Test
    void update_existingId_returnsUpdatedProduct() {
        Product updated = new Product("Updated", "Updated Desc", new BigDecimal("19.99"), 5);
        when(productRepository.findById(1L)).thenReturn(Optional.of(product));
        when(productRepository.save(any(Product.class))).thenAnswer(inv -> inv.getArgument(0));
        Product result = productService.update(1L, updated);
        assertThat(result.getName()).isEqualTo("Updated");
        assertThat(result.getPrice()).isEqualByComparingTo("19.99");
    }

    @Test
    void delete_existingId_deletesProduct() {
        when(productRepository.findById(1L)).thenReturn(Optional.of(product));
        productService.delete(1L);
        verify(productRepository, times(1)).delete(product);
    }

    @Test
    void delete_nonExistingId_throwsException() {
        when(productRepository.findById(99L)).thenReturn(Optional.empty());
        assertThatThrownBy(() -> productService.delete(99L))
                .isInstanceOf(ResourceNotFoundException.class);
    }
}
