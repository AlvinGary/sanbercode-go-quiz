package controllers

import (
	"database/sql"
	"log"
	"net/http"
	"sanbercode-go-quiz/structs"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

// Endpoint POST kategori
func CreateKategori(c *gin.Context, db *sql.DB) {
	BasicAuth()(c)
	if c.IsAborted() {
		return
	}
	var newKategori structs.Kategori
	if err := c.ShouldBindJSON(&newKategori); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON data"})
		return
	}
	if newKategori.Name == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Name tidak boleh kosong"})
		return
	}
	now := time.Now()
	user := c.GetString("user")
	newKategori.CreatedAt = now
	newKategori.CreatedBy = user
	newKategori.ModifiedAt = now
	newKategori.ModifiedBy = user
	query := `INSERT INTO "Kategori" (name, created_at, created_by, modified_at, modified_by) VALUES ($1, $2, $3, $4, $5) RETURNING id`
	err := db.QueryRow(query, newKategori.Name, newKategori.CreatedAt, newKategori.CreatedBy, newKategori.ModifiedAt, newKategori.ModifiedBy).Scan(&newKategori.Id)
	if err != nil {
		log.Println("Error inserting new Kategori:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal menambahkan Kategori"})
		return
	}
	c.JSON(http.StatusCreated, newKategori)
}

// Endpoint GET Kategori
func GetKategori(c *gin.Context, db *sql.DB) {
	BasicAuth()(c)
	if c.IsAborted() {
		return
	}
	rows, err := db.Query(`SELECT id, name, created_at, created_by, modified_at, modified_by FROM "Kategori"`)
	if err != nil {
		log.Println("Error fetching kategori:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal mengambil data Kategori"})
		return
	}
	defer rows.Close()

	var kategori []structs.Kategori
	for rows.Next() {
		var b structs.Kategori
		if err := rows.Scan(&b.Id, &b.Name, &b.CreatedAt, &b.CreatedBy, &b.ModifiedAt, &b.ModifiedBy); err != nil {
			log.Println("Error scanning Kategori:", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal memproses data Kategori"})
			return
		}
		kategori = append(kategori, b)
	}
	c.JSON(http.StatusOK, kategori)
}

// Endpoint GET by ID Kategori
func GetKategoriByID(c *gin.Context, db *sql.DB) {
	BasicAuth()(c)
	if c.IsAborted() {
		return
	}
	id := c.Param("id")
	var b structs.Kategori
	query := `SELECT id, name, created_at, created_by, modified_at, modified_by FROM "Kategori" WHERE "id" = $1`
	err := db.QueryRow(query, id).Scan(&b.Id, &b.Name, &b.CreatedAt, &b.CreatedBy, &b.ModifiedAt, &b.ModifiedBy)
	if err != nil {
		if err == sql.ErrNoRows {
			c.JSON(http.StatusNotFound, gin.H{"error": "Kategori tidak ditemukan"})
		} else {
			log.Println("Error fetching Kategori by ID:", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal mengambil data Kategori"})
		}
		return
	}
	c.JSON(http.StatusOK, b)
}

// Endpoint PUT by ID kategori
func UpdateKategori(c *gin.Context, db *sql.DB) {
	BasicAuth()(c)
	if c.IsAborted() {
		return
	}
	id := c.Param("id")

	// Fetch existing record to preserve created_at and created_by
	var existingKategori structs.Kategori
	querySelect := `SELECT id, name, created_at, created_by, modified_at, modified_by FROM "Kategori" WHERE id = $1`
	err := db.QueryRow(querySelect, id).Scan(&existingKategori.Id, &existingKategori.Name, &existingKategori.CreatedAt, &existingKategori.CreatedBy, &existingKategori.ModifiedAt, &existingKategori.ModifiedBy)
	if err != nil {
		if err == sql.ErrNoRows {
			c.JSON(http.StatusNotFound, gin.H{"error": "Kategori tidak ditemukan"})
		} else {
			log.Println("Error fetching existing Kategori:", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal mengambil data Kategori"})
		}
		return
	}

	var input struct {
		Name string `json:"name"`
	}
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON data"})
		return
	}
	if input.Name == ""{
		c.JSON(http.StatusBadRequest, gin.H{"error": "Name tidak boleh kosong"})
		return
	}
	now := time.Now()
	user := c.GetString("user")
	updatedKategori := structs.Kategori{
		Id:         existingKategori.Id,
		Name:       input.Name,
		CreatedAt:  existingKategori.CreatedAt,  // Preserve original
		CreatedBy:  existingKategori.CreatedBy,  // Preserve original
		ModifiedAt: now,
		ModifiedBy: user,
	}
	query := `UPDATE "Kategori" SET name = $1, modified_at = $2, modified_by = $3 WHERE id = $4`
	result, err := db.Exec(query, updatedKategori.Name, updatedKategori.ModifiedAt, updatedKategori.ModifiedBy, id)
	if err != nil {
		log.Println("Error updating Kategori:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal memperbarui Kategori"})
		return
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		log.Println("Error getting rows affected:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal memperbarui Kategori"})
		return
	}
	if rowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Kategori tidak ditemukan"})
		return
	}
	updatedKategori.Id, _ = strconv.Atoi(id)
	c.JSON(http.StatusOK, updatedKategori)
}

// Endpoint DELETE by ID kategori
func DeleteKategori(c *gin.Context, db *sql.DB) {
	id := c.Param("id")
	query := `DELETE FROM "Kategori" WHERE id = $1`
	result, err := db.Exec(query, id)
	if err != nil {
		log.Println("Error deleting Kategori:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal menghapus Kategori"})
		return
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		log.Println("Error getting rows affected:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal menghapus Kategori"})
		return
	}
	if rowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Kategori tidak ditemukan"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Kategori berhasil dihapus"})
}


// basic auth middleware
func BasicAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		user, password, hasAuth := c.Request.BasicAuth()
		if hasAuth && user == "admin" && password == "root" {
			c.Set("user", user)
			c.Next()
			return
		}
		c.Header("WWW-Authenticate", `Basic realm="Restricted"`)
		c.AbortWithStatus(http.StatusUnauthorized)
	}
}