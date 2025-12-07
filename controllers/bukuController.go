package controllers

import (
	"database/sql"
	"log"
	"net/http"
	"sanbercode-go-quiz/middleware"
	"sanbercode-go-quiz/structs"
	"time"

	"github.com/gin-gonic/gin"
)

// Endpoint POST Buku
func CreateBuku(c *gin.Context, db *sql.DB) {
	middleware.BasicAuth()(c)
	if c.IsAborted() {
		return
	}
	var newBuku structs.Buku
	if err := c.ShouldBindJSON(&newBuku); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON data"})
		return
	}
	if newBuku.Title == "" || newBuku.Description == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Title dan Description tidak boleh kosong"})
		return
	}

	if newBuku.ReleaseYear == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Release Year tidak boleh kosong"})
		return
	}
	if newBuku.ReleaseYear < 1980 || newBuku.ReleaseYear > 2024 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Release Year tidak boleh kurang dari tahun 1980 atau lebih dari tahun 2024"})
		return
	}

	if newBuku.TotalPage < 100 {
		newBuku.Thickness = "Tipis"
	} else {
		newBuku.Thickness = "Tebal"
	}

	// check category id if it exists
	var exists bool
    err := db.QueryRow(`SELECT EXISTS(SELECT 1 FROM "Kategori" WHERE id = $1)`,
        newBuku.CategoryId).Scan(&exists)
    if err != nil {
        log.Println("Error checking category:", err)
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal memeriksa Category ID"})
        return
    }
    if !exists {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Category ID tidak valid"})
        return
    }

	now := time.Now()
	user := c.GetString("user")
	newBuku.CreatedAt = now
	newBuku.CreatedBy = user
	newBuku.ModifiedAt = now
	newBuku.ModifiedBy = user
	query := `INSERT INTO "Buku" (title, description, image_url, release_year, price, total_page, thickness, category_id, created_at, created_by, modified_at, modified_by) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12) RETURNING id`
	err = db.QueryRow(query, newBuku.Title, newBuku.Description, newBuku.ImageUrl, newBuku.ReleaseYear, newBuku.Price, newBuku.TotalPage, newBuku.Thickness, newBuku.CategoryId, newBuku.CreatedAt, newBuku.CreatedBy, newBuku.ModifiedAt, newBuku.ModifiedBy).Scan(&newBuku.Id)
	if err != nil {
		log.Println("Error inserting new Buku:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal menambahkan Buku"})
		return
	}
	c.JSON(http.StatusCreated, newBuku)
}

// Endpoint GET Buku
func GetBuku(c *gin.Context, db *sql.DB) {
	middleware.BasicAuth()(c)
	if c.IsAborted() {
		return
	}
	rows, err := db.Query(`SELECT id, title, description, image_url, release_year, price, total_page, thickness, category_id, created_at, created_by, modified_at, modified_by FROM "Buku"`)
	if err != nil {
		log.Println("Error fetching Buku:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal mengambil data Buku"})
		return
	}
	defer rows.Close()

	var buku []structs.Buku
	for rows.Next() {
		var b structs.Buku
		if err := rows.Scan(&b.Id, &b.Title, &b.Description, &b.ImageUrl, &b.ReleaseYear, &b.Price, &b.TotalPage, &b.Thickness, &b.CategoryId, &b.CreatedAt, &b.CreatedBy, &b.ModifiedAt, &b.ModifiedBy); err != nil {
			log.Println("Error scanning Buku:", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal memproses data Buku"})
			return
		}
		buku = append(buku, b)
	}
	c.JSON(http.StatusOK, buku)
}

// Endpoint GET by ID Buku
func GetBukuByID(c *gin.Context, db *sql.DB) {
	middleware.BasicAuth()(c)
	if c.IsAborted() {
		return
	}
	id := c.Param("id")
	var b structs.Buku
	query := `SELECT id, title, description, image_url, release_year, price, total_page, thickness, category_id, created_at, created_by, modified_at, modified_by FROM "Kategori" WHERE "id" = $1`
	err := db.QueryRow(query, id).Scan(&b.Id, &b.Title, &b.Description, &b.ImageUrl, &b.ReleaseYear, &b.Price, &b.TotalPage, &b.Thickness, &b.CategoryId, &b.CreatedAt, &b.CreatedBy, &b.ModifiedAt, &b.ModifiedBy)
	if err != nil {
		if err == sql.ErrNoRows {
			c.JSON(http.StatusNotFound, gin.H{"error": "Buku tidak ditemukan"})
		} else {
			log.Println("Error fetching Buku by ID:", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal mengambil data Buku"})
		}
		return
	}
	c.JSON(http.StatusOK, b)
}

// Endpoint DELETE by ID Buku
func DeleteBuku(c *gin.Context, db *sql.DB) {
middleware.BasicAuth()(c)
	if c.IsAborted() {
		return
	}
	id := c.Param("id")
	query := `DELETE FROM "Buku" WHERE id = $1`
	result, err := db.Exec(query, id)
	if err != nil {
		log.Println("Error deleting Buku:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal menghapus Buku"})
		return
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		log.Println("Error getting rows affected:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal menghapus Buku"})
		return
	}
	if rowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Buku tidak ditemukan"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Buku berhasil dihapus"})
}