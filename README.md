# users-and-books
This is my final homework for OneLab. The system is a library management API. It is divided into two microservices. 

Books Microservice is the main and manages:
* user registration and authentication using JWT
* book addition
* book borrowing
* runs on port 8585

Transactions Microservice manages:
* Financial transactions for each book borrowing
* runs on port 8787

Microservices communicate via HTTP. The books microservice receives the request to borrow a book, registers it in its own database, and sends the request to the transactions microservice to register the financial transaction of book borrowing.
The application is packed in one Docker container with all dependencies which you need to run using `docker-compose up -d`. The gorm's auto migration is enabled, so the application is ready to work after the previous command.

For your convenience, there is a main_test.http file that could be used to test the endpoints using GoLand and Swagger documentation that you could access via `/api/v1/swagger/`.

