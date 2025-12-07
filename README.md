# sanbercode-go-quiz

## Usage

Just type the path list using **Postman** or other api testing tools to start request the api.<br>
Note if you want to start using the local server:<br>
Run the command `go run main.go` in terminal. And don't forget to test the api path using `localhost:8000` instead of railway path for local server.<br>

## Purpose

To do the Quiz assignment from Sanbercode Go Mini Bootcamp by learning more about how to create rest api using **Go** programming language, with **GIN** framework, and **PostgresSQL** database. Along with **Railway** deployment.<br>
This program have `Categories` and `Books` feature where user can view, create, update, and delete Categories and Books data.

## Path List

### Categories API

1. `https://sanbercode-go-quiz-production.up.railway.app/api/categories`<br>
   Using **GET** method for fetch all categories data.<br>
   Using **POST** method for create new categories data.
2. `https://sanbercode-go-quiz-production.up.railway.app/api/categories/:id`<br>
   Using **GET** method for fetch specific categories data.<br>
   Using **PUT** method for update specific categories data.<br>
   Using **DELETE** method for delete specific categories data.
3. `https://sanbercode-go-quiz-production.up.railway.app/api/categories/:id/books`<br>
   Using **GET** method for fetching books data based on categories id.<br>

### Books API

1. `https://sanbercode-go-quiz-production.up.railway.app/api/books`<br>
   Using **GET** method for fetch all books data.
   Using **POST** method for create new books data.
2. `https://sanbercode-go-quiz-production.up.railway.app/api/books/:id`<br>
   Using **GET** method for fetch specific books data.<br>
   Using **PUT** method for update specific books data.<br>
   Using **DELETE** method for delete specific books data.
