# 🛒 E-Commerce- E-commerce System

![Project Status](https://img.shields.io/badge/Status-Completed-success)
![Go Version](https://img.shields.io/badge/Go-1.21+-00ADD8?logo=go)
![Next.js Version](https://img.shields.io/badge/Next.js-14-black?logo=next.js)

**Poom Store** is a modern, high-performance e-commerce platform designed to demonstrate a robust full-stack architecture. It features a responsive frontend built with **Next.js** and a high-concurrency backend powered by **Go (Gin Framework)**.

> **Note:** This project focuses on performance optimization and manual data control by utilizing **Native SQL (`database/sql`)** instead of ORMs for the database layer.

---

## 📸 Screenshots

*(Place your screenshot here)*

---

## 🚀 Tech Stack

### **Frontend**
- **Framework:** Next.js (App Router)
- **Language:** TypeScript
- **Styling:** Tailwind CSS
- **State Management:** Zustand (Handling complex cart logic)
- **Icons:** Lucide React

### **Backend**
- **Language:** Go (Golang)
- **Framework:** Gin Web Framework
- **Database Driver:** `go-sql-driver/mysql` (Native SQL)
- **Architecture:** RESTful API, Layered Architecture (Controller-Service-Model)

### **Database**
- **System:** MySQL
- **Design:** Relational Schema (Products, Orders)

---

## ✨ Key Features

- **🛍️ Product Catalog:** Dynamic product listing with real-time **Search** and **Category Filtering**.
- **🛒 Smart Shopping Cart:** - Client-side state management using **Zustand**.
  - Handles unique item grouping and quantity adjustments instantly.
- **💳 Checkout System:** Seamless order processing flow connecting Frontend state to Backend database.
- **⚡ High Performance:** Optimized backend queries using raw SQL to ensure precise data handling and speed.
- **🔒 Security:** Implemented Custom CORS Middleware for secure cross-origin resource sharing.

---

## 🛠️ Technical Highlights (Why Native SQL?)

One of the core challenges in this project was handling complex data types and ensuring maximum query performance. 

Instead of using an ORM like GORM, I decided to migrate to **Native SQL (`database/sql`)**.
- **Reason:** To have 100% control over the generated SQL queries and resolve `uint` vs `int` type mapping issues between Go and the Database.
- **Result:** Improved query transparency and easier debugging for complex transactions.

---

## ⚙️ Installation & Setup

Follow these steps to run the project locally.

### Prerequisites
- Go (Golang) installed
- Node.js & npm installed
- MySQL Database running

### 1. Database Setup
Run the following SQL commands in your MySQL interface (e.g., phpMyAdmin):

```sql
CREATE DATABASE ecommerce_db;

USE ecommerce_db;

CREATE TABLE `products` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(255) NOT NULL,
  `price` decimal(10,2) NOT NULL,
  `image_url` text,
  `category` varchar(100) DEFAULT NULL,
  `rating` decimal(2,1) DEFAULT 5.0,
  `created_at` datetime DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
);

CREATE TABLE `orders` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `user_id` int(11) DEFAULT 1,
  `total_amount` decimal(10,2) NOT NULL,
  `status` varchar(50) DEFAULT 'pending',
  `created_at` datetime DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
);
