# Finance Tracker API Documentation

## Overview

The **Finance Tracker API** is a robust, scalable, and modern RESTful API designed to empower users to manage their financial lives effectively. Built with Go, Fiber, and Turso (LibSQL), this project provides a comprehensive backend solution for tracking income, expenses, assets, liabilities, budgets, and financial goals. Whether you're a developer looking to contribute to a cutting-edge project, a recruiter seeking talent with real-world experience, or a client in need of a reliable financial management system, this API offers a professional-grade foundation.

This documentation highlights the system's architecture, features, and extensibility, showcasing its value to technical and non-technical stakeholders alike.

---

## Table of Contents

1. [Project Highlights](#project-highlights)
   - [Why This Project Stands Out](#why-this-project-stands-out)
2. [Features](#features)
   - [Implemented Functionalities](#implemented-functionalities)
   - [Future Roadmap](#future-roadmap)
3. [Technical Architecture](#technical-architecture)
   - [Technology Stack](#technology-stack)
   - [Directory Structure](#directory-structure)
   - [Data Models](#data-models)
4. [API Endpoints](#api-endpoints)
   - [User Management](#user-management)
   - [Transaction Management](#transaction-management)
   - [Category Management](#category-management)
   - [Asset Management](#asset-management)
   - [Liability Management](#liability-management)
   - [Budget Management](#budget-management)
   - [Goal Management](#goal-management)
5. [Setup and Installation](#setup-and-installation)
   - [Prerequisites](#prerequisites)
   - [Installation Steps](#installation-steps)
6. [Usage Examples](#usage-examples)
   - [Creating a User](#creating-a-user)
   - [Tracking a Transaction](#tracking-a-transaction)
7. [Contributing](#contributing)
   - [How to Get Involved](#how-to-get-involved)
   - [Code Standards](#code-standards)
8. [For Clients](#for-clients)
   - [Business Value](#business-value)
   - [Customization Options](#customization-options)
9. [Contact and Collaboration](#contact-and-collaboration)

---

## Project Highlights

### Why This Project Stands Out

- **Modern Tech Stack**: Built with Go and Fiber for high performance, paired with Turso (LibSQL) for lightweight, scalable database management.
- **Clean Architecture**: Follows RESTful principles with a modular design, making it easy to maintain, extend, and scale.
- **Real-World Relevance**: Addresses everyday financial management needs, from budgeting to goal tracking, appealing to both end-users and businesses.
- **Developer-Friendly**: Well-documented codebase with SQLC-generated queries, reducing boilerplate and ensuring type safety.
- **Recruiter Appeal**: Demonstrates proficiency in backend development, database design, and API engineering—skills in high demand.

---

## Features

### Implemented Functionalities

- **User Management**: CRUD operations for user accounts with Clerk integration for authentication-ready design.
- **Transaction Management**: Track income and expenses with categorization and detailed metadata.
- **Category Management**: Organize transactions with customizable categories.
- **Asset & Liability Tracking**: Monitor financial health by managing assets and liabilities.
- **Budget Management**: Set and track budgets tied to categories with time-bound constraints.
- **Goal Management**: Define and monitor financial goals with user-specific progress tracking.

### Future Roadmap

- **Reporting & Analytics**: Generate insights like income vs. expenses, category spending, and net worth trends.
- **Authentication**: Integrate full JWT or OAuth2-based security.
- **Pagination & Filtering**: Enhance API usability with advanced query options.
- **Goal Progress Tracking**: Add detailed progress endpoints for richer goal management.

---

## Technical Architecture

### Technology Stack

| Component          | Technology            | Purpose                              |
|--------------------|-----------------------|--------------------------------------|
| Language           | Go 1.22.4            | High-performance backend development |
| Web Framework      | Fiber v2             | Fast, lightweight HTTP routing       |
| Database           | Turso (LibSQL)       | Scalable SQLite-based storage        |
| Query Generation   | SQLC v1.26.0         | Type-safe SQL queries                |
| Environment Config | godotenv             | Manage environment variables         |

### Directory Structure

```
controller/         # API endpoint handlers
database/           # Database initialization and schema
  queries/          # SQL query definitions
  schemas/          # Database schema (DDL)
  sqlc/             # Auto-generated Go code from SQLC
router/             # Route definitions
.env.example        # Environment variable template
.gitignore         # Git ignore rules
go.mod             # Go module dependencies
main.go            # Application entry point
readME.md          # Project overview
sqlc.yaml          # SQLC configuration
```

### Data Models

The API leverages a relational database with the following core entities:

- **Users**: `id`, `clerk_id`, `is_active`, `created_at`, `updated_at`
- **Transactions**: `id`, `user_id`, `category_id`, `name`, `amount`, `description`, `type`, `date`, `created_at`, `updated_at`
- **Categories**: `id`, `name`, `description`
- **Budgets**: `id`, `category_id`, `amount`, `start_date`, `end_date`
- **Assets**: `id`, `user_id`, `name`, `amount`
- **Liabilities**: `id`, `user_id`, `name`, `amount`
- **Goals**: `id`, `user_id`, `type`, `amount`, `start_date`, `end_date`, `status`

---

## API Endpoints

### User Management

| Method | Endpoint         | Description             |
|--------|------------------|-------------------------|
| GET    | `/users`         | List all users          |
| GET    | `/users/:id`     | Get a specific user     |
| POST   | `/users`         | Create a new user       |
| PUT    | `/users/:id`     | Update a user           |
| DELETE | `/users/:id`     | Delete a user           |

### Transaction Management

| Method | Endpoint             | Description                 |
|--------|----------------------|-----------------------------|
| GET    | `/transactions`      | List all transactions       |
| GET    | `/transactions/:id`  | Get a specific transaction  |
| POST   | `/transactions`      | Create a new transaction    |
| PUT    | `/transactions/:id`  | Update a transaction        |
| DELETE | `/transactions/:id`  | Delete a transaction        |

### Category Management

| Method | Endpoint           | Description               |
|--------|--------------------|---------------------------|
| GET    | `/categories`      | List all categories       |
| GET    | `/categories/:id`  | Get a specific category   |
| POST   | `/categories`      | Create a new category     |
| PUT    | `/categories/:id`  | Update a category         |
| DELETE | `/categories/:id`  | Delete a category         |

### Asset Management

| Method | Endpoint         | Description             |
|--------|------------------|-------------------------|
| GET    | `/assets`        | List all assets         |
| GET    | `/assets/:id`    | Get a specific asset    |
| POST   | `/assets`        | Create a new asset      |
| PUT    | `/assets/:id`    | Update an asset         |
| DELETE | `/assets/:id`    | Delete an asset         |

### Liability Management

| Method | Endpoint            | Description                |
|--------|---------------------|----------------------------|
| GET    | `/liabilities`      | List all liabilities       |
| GET    | `/liabilities/:id`  | Get a specific liability   |
| POST   | `/liabilities`      | Create a new liability     |
| PUT    | `/liabilities/:id`  | Update a liability         |
| DELETE | `/liabilities/:id`  | Delete a liability         |

### Budget Management

| Method | Endpoint         | Description             |
|--------|------------------|-------------------------|
| GET    | `/budgets`       | List all budgets        |
| GET    | `/budgets/:id`   | Get a specific budget   |
| POST   | `/budgets`       | Create a new budget     |
| PUT    | `/budgets/:id`   | Update a budget         |
| DELETE | `/budgets/:id`   | Delete a budget         |

### Goal Management

| Method | Endpoint            | Description                |
|--------|---------------------|----------------------------|
| GET    | `/goals`            | List all goals             |
| GET    | `/goals/:id`        | Get a specific goal        |
| POST   | `/goals`            | Create a new goal          |
| PUT    | `/goals/:id`        | Update a goal              |
| DELETE | `/goals/:id`        | Delete a goal              |
| GET    | `/goals/users/:id`  | List goals by user ID      |

---

## Setup and Installation

### Prerequisites

- Go 1.22.4 or higher
- Turso (LibSQL) database access
- Git for version control
- Optional: SQLC for query generation

### Installation Steps

1. **Clone the Repository**
   ```bash
   git clone <repository-url>
   cd finance-tracker-api
   ```

2. **Set Up Environment Variables**
   Copy `.env.example` to `.env` and update `DB_URL` with your Turso database connection string:
   ```bash
   cp .env.example .env
   ```

3. **Install Dependencies**
   ```bash
   go mod tidy
   ```

4. **Run the Application**
   ```bash
   go run main.go
   ```

5. **Verify**
   Access `http://localhost:8080/` to see the "Server active" message.

---

## Usage Examples

### Creating a User

**Request**
```bash
curl -X POST http://localhost:8080/users \
-H "Content-Type: application/json" \
-d '{"clerk_id": "user_123", "is_active": true}'
```

**Response**
```json
{
  "id": 1,
  "clerk_id": "user_123",
  "is_active": true,
  "created_at": "2025-03-17T10:00:00Z",
  "updated_at": "2025-03-17T10:00:00Z"
}
```

### Tracking a Transaction

**Request**
```bash
curl -X POST http://localhost:8080/transactions \
-H "Content-Type: application/json" \
-d '{"user_id": 1, "category_id": 1, "name": "Grocery Shopping", "amount": 50.75, "type": "expense", "date": "2025-03-17T12:00:00Z"}'
```

**Response**
```json
{
  "id": 1,
  "user_id": 1,
  "category_id": 1,
  "name": "Grocery Shopping",
  "amount": 50.75,
  "type": "expense",
  "date": "2025-03-17T12:00:00Z",
  "created_at": "2025-03-17T12:00:00Z",
  "updated_at": "2025-03-17T12:00:00Z"
}
```

---

## Contributing

### How to Get Involved

Developers are welcome to contribute! Fork the repository, create a feature branch, and submit a pull request. Focus areas include:
- Adding reporting endpoints
- Implementing authentication
- Enhancing test coverage

### Code Standards

- Follow Go conventions (e.g., `gofmt`, `golint`).
- Use SQLC for database queries.
- Write clear, concise commit messages.

---

## For Clients

### Business Value

- **User Empowerment**: Equip your customers with tools to manage their finances seamlessly.
- **Scalability**: Built to handle growing user bases with minimal overhead.
- **Customizable**: Easily adapt the API to fit specific branding or feature needs.

### Customization Options

- Add branded reporting dashboards.
- Integrate with existing authentication systems.
- Extend data models for additional financial metrics.

---

## Contact and Collaboration

Interested in hiring a skilled developer, contributing to the project, or deploying this API for your business? Reach out via the project's repository or contacts. Let’s build something extraordinary together!
