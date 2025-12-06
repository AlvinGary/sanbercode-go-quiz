# sanbercode-go-quiz

## Usage

run the command `go run main.go` in terminal to start the local server.

## Purpose

To do the Quiz assignment from Sanbercode Go mini bootcamp by learning more about how to create rest api using Go programming language, with GIN framework, and PostgresSQL database.<br>
This program have categories feature where user can view, create, update, and delete category data locally.

## Path List

1. `localhost:8000/api/categories`<br>
   Using GET method.<br>
   This API used for fetch all categories data.
2. `localhost:8000/api/categories`<br>
   Using POST method.<br>
   This API used for create new categories data.
3. `localhost:8000/api/categories/:id`<br>
   Using GET method.<br>
   This API used for fetch specific categories data.<br>
   Using PUT method for update specific categories data.<br>
   Using DELETE method for delete specific categories data.
