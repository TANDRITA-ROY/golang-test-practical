# golang-test-practical

Step 1: 
    Ensure your MySQL database is running and configured as per the provided schema.

        1.  Connect to MySQL: mysql -u username -p

        2.  Create the database and tables: Using mysql.sql queries.

Step 2: 
    Run the application:
        go run main.go

    The application will start and listen on port 8080. You can test the endpoints using tools like Postman or curl.

    Endpoints:
        1.  GET /person/:person_id/info:

        curl -X GET http://localhost:8080/person/1/info

        2.  POST /person/create:

        curl -X POST http://localhost:8080/person/create \
        -H "Content-Type: application/json" \
        -d '{
            "name": "YOURNAME",
            "phone_number": "123-456-7890",
            "city": "Sacramento",
            "state": "CA",
            "street1": "112 Main St",
            "street2": "Apt 12",
            "zip_code": "12345"
            }'

Kindly Note: Code sets up the required REST endpoints, connects to the MySQL database, and handles the data