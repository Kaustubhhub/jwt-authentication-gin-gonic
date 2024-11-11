# Golang Backend Project with JWT authentication

This is a Golang backend project designed to run locally using MongoDB as the database.

## Prerequisites

- Go (version 1.18+ recommended)
- MongoDB (running locally or in the cloud)
- Git (for cloning the repository)

## Getting Started

### 1. Clone the Repository

First, clone the repository to your local machine:

```bash
git clone (https://github.com/Kaustubhhub/jwt-authentication-gin-gonic.git)
cd jwt-authentication-gin-gonic
```

### 2. Create a `.env` File

Create a `.env` file in the root of the project directory. Add the following environment variables:

```env
PORT=9000
MONGODB_URL=mongodb://localhost:27017/your-database-name
```

- `PORT`: The port on which the server will run (defaulted to 9000).
- `MONGODB_URL`: Connection URL for MongoDB. Replace `localhost:27017` with your MongoDB URL if it's hosted externally, and `your-database-name` with the name of your MongoDB database.

### 3. Install Dependencies

Install any required Go packages with:

```bash
go mod tidy
```

### 4. Run the Project

Run the server locally using:

```bash
go run main.go
```

This will start the backend server on `http://localhost:9000` by default.

### 5. Verify Setup

Open your browser or a tool like Postman and navigate to `http://localhost:9000/api-1` to verify that the server is running properly.

## Additional Notes

- Ensure MongoDB is running locally or that you have access to a cloud-hosted MongoDB instance.
- Customize the `.env` file as needed for other environment-specific configurations.

---
