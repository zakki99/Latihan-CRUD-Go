package main

import (
	"fmt"
)

type Product struct {
	ID    int
	Code  string
	Price int
}

// Simpan data produk dalam slice
var products []Product

// Fungsi untuk menambahkan produk baru
func createProduct(code string, price int) {
	// Hitung ID berikutnya
	nextID := len(products) + 1
	// Tambahkan produk baru ke slice
	products = append(products, Product{ID: nextID, Code: code, Price: price})
}

// Fungsi untuk membaca produk berdasarkan ID
func readProduct(id int) *Product {
	for _, p := range products {
		if p.ID == id {
			return &p
		}
	}
	return nil
}

// Fungsi untuk memperbarui produk berdasarkan ID
func updateProduct(id int, code string, price int) {
	for i, p := range products {
		if p.ID == id {
			// Update data produk
			products[i].Code = code
			products[i].Price = price
			return
		}
	}
}

// Fungsi untuk menghapus produk berdasarkan ID
func deleteProduct(id int) {
	for i, p := range products {
		if p.ID == id {
			// Hapus produk dari slice
			products = append(products[:i], products[i+1:]...)
			return
		}
	}
}

func main() {
	// Menambahkan beberapa produk
	createProduct("Laptop", 1000)
	createProduct("Smartphone", 800)

	// Mencetak produk
	fmt.Println("List of Products:")
	for _, p := range products {
		fmt.Printf("ID: %d, Code: %s, Price: $%d\n", p.ID, p.Code, p.Price)
	}

	// Mencari produk berdasarkan ID
	productID := 1
	fmt.Printf("\nFinding Product with ID %d:\n", productID)
	foundProduct := readProduct(productID)
	if foundProduct != nil {
		fmt.Printf("ID: %d, Code: %s, Price: $%d\n", foundProduct.ID, foundProduct.Code, foundProduct.Price)
	} else {
		fmt.Println("Product not found.")
	}

	// Memperbarui produk
	updateProduct(productID, "Updated Laptop", 1200)
	fmt.Printf("\nProduct with ID %d updated.\n", productID)

	// Menghapus produk
	deleteProduct(productID)
	fmt.Printf("\nProduct with ID %d deleted.\n", productID)
}
