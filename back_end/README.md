# E-Commerce Application

This is an E-Commerce application built using Go (Golang), Next.js, GORM, GIN, and SQLite. The project aims to create a fully functional e-commerce website where users can browse products, add items to their cart, place orders, and more.

## Table of Contents

- [Features](#features)
- [Technologies Used](#technologies-used)
- [Prerequisites](#prerequisites)
- [Getting Started](#getting-started)
- [Installation](#installation)
- [Usage](#usage)
- [Contributing](#contributing)
- [License](#license)

## Features

The E-Commerce application will include the following features:

- User registration and authentication
- Product browsing with categories and filters
- Product search functionality
- Product details page with images, descriptions, and specifications
- Shopping cart management
- Checkout process with order summary
- Order history for users
- Admin panel to manage products, categories, and orders

## Technologies Used

The application is built using the following technologies:

- Go (Golang) - Backend programming language
- Next.js - Frontend framework for server-rendered React applications
- GORM - Object-relational mapping (ORM) library for Go
- GIN - Web framework for building APIs in Go
- SQLite - Lightweight and serverless relational database

## Prerequisites

Before running the application, make sure you have the following installed:

- Go (Golang): [Official Installation Guide](https://golang.org/doc/install)
- Node.js and npm: [Official Installation Guide](https://docs.npmjs.com/downloading-and-installing-node-js-and-npm)
- SQLite: [Official Website](https://www.sqlite.org/download.html)

## Getting Started

To get started with the E-Commerce application, follow the steps below:

1. Clone this repository to your local machine.
2. Set up the backend using Go, GORM, and GIN.
3. Set up the frontend using Next.js.
4. Connect the frontend and backend to make API calls.
5. Set up the SQLite database and seed it with sample data.

## Installation

Follow these steps to install the required dependencies:

### Backend

1. Open a terminal and navigate to the `backend` directory.
2. Run `go mod download` to download the Go modules.
3. Run `go run main.go` to start the backend server.

### Frontend

1. Open another terminal and navigate to the `frontend` directory.
2. Run `npm install` to install the required npm packages.
3. Run `npm run dev` to start the frontend development server.

## Usage

After successfully installing and running both the backend and frontend, open your web browser and visit `http://localhost:3000` to access the E-Commerce application.

## Contributing

Contributions to this E-Commerce application are welcome. If you find any issues or want to add new features, feel free to fork the repository and submit a pull request.

## License

This project is licensed under the [MIT License](LICENSE). Feel free to use, modify, and distribute the code as per the terms of this license.
