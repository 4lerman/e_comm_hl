# E-commerce Service

## Overview

The E-commerce Service is a backend application built with Go that provides various endpoints for managing orders and products in an e-commerce platform. The service uses Gorilla Mux for routing and Swagger for API documentation. It also integrates with PostgreSQL for database management and is containerized using Docker and Docker Compose.

## Features

- **Order Management**: Create, update, delete, and fetch orders.
- **Product Management**: Create, update, delete, and fetch products.
- **Search Functionality**: Search for orders by status or user.
- **Swagger Documentation**: Interactive API documentation.
- **Dockerized Deployment**: Easy setup and deployment using Docker and Docker Compose.

## Prerequisites

- Go 1.22 or later
- Docker
- Docker Compose

## Installation

1. Clone the repository:

   ```sh
   git clone <repository-url>
   cd e_comm_hl
   ```

2. Build and run the Docker container:

   ```sh
   docker-compose up --build
   ```

## Usage

1. Access Swagger Documentation: Open http://localhost:8080/swagger/ to view and interact with the API documentation.
