# Personalities Management System

This is a simple web application that provides a RESTful API for managing a database of fictional personalities. The application is built with Go and connects to a PostgreSQL database that is run on Docker. The API can be tested using Postman or a similar tool. Additionally, a simple frontend is provided using React.

## Dependencies

### Backend

- Go: A programming language that provides a set of features for building efficient and scalable applications.
- database/sql: A Go package that provides a generic interface to SQL databases.
- encoding/json: A Go package that provides functionality for encoding and decoding JSON data.
- github.com/gorilla/handlers: A Go package that provides middleware functionality for the Gorilla web toolkit.
- github.com/gorilla/mux: A Go package that provides a request router and dispatcher.
- net/http: A Go package that provides HTTP client and server implementations.
- os: A Go package that provides a platform-independent interface to operating system functionality.
- postgresql: A powerful open-source relational database management system.

### Frontend

- React: A JavaScript library for building user interfaces.
- Node.js: A JavaScript runtime built on Chrome's V8 JavaScript engine.
- NPM: A package manager for Node.js.

## Getting Started

To run the application, follow these steps:

1. Clone the repository to your local machine.
2. Open a terminal window and navigate to the project directory.
3. Run the command `docker-compose up` to start the database.
4. Open a web browser and go to http://localhost:54321 to view PgAdmin. Use the following credentials to log in:
- Connection:
- Email: `admin@admin.com`
- Password: `admin`
   - General:
     - Name: `YOUR_SERVER_NAME`
   - Connection:
   - For the host, type in ``docker-compose exec postgres sh, then `hostname -i`, and copy the adress which will be pasted on PgAdmin`s setup as shown below:
     - Host: `YOUR_HOST_IP`
     - Port: `5432`
     - Maintenance database: `root`
     - Username: `root`
     - Password: `root`
5. Save the server settings and connect to the database.
6. Navigate to the frontend directory by running `cd frontend` in the terminal window.
7. Run the command `npm install` to install the required dependencies for the frontend.
8. Run the command `npm start` to start the frontend server.
9. Open another terminal window and navigate to the project directory.
10. Run the command `go run main.go` to start the backend server.
11. Open a web browser and go to `http://localhost:8000` to use the application.

## Usage

The application provides the following API endpoints:

- GET `/api/personalidades`: Retrieves a list of all personalities in the database.
- GET `/api/personalidades/{id}`: Retrieves a specific personality from the database by ID.
- POST `/api/personalidades`: Creates a new personality in the database.
- DELETE `/api/personalidades/{id}`: Deletes a specific personality from the database by ID.
- PUT `/api/personalidades/{id}`: Updates an existing personality in the database by ID.

The API returns data in JSON format.

In addition to the API, a simple frontend is provided using React. The frontend only allows visualization, as it is been used with the purpose of learning how to integrate it with the backend correctly.

## Conclusion

This application demonstrates how to use Go, PostgreSQL, and React to create a simple web application for managing a database of fictional personalities. By following the steps outlined above, you should be able to get the application up and running on your local machine in no time.
