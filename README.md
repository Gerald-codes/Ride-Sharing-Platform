# Ride-Sharing-Platform
(ETI) Assignment 1 - Semester 3.2

---
## Design consideration of your microservices

For Console FontEnd, it is Monolith frontend as everything is implemented in that file. 

For WEB FrontEnd, it will be Micro frontend as each individual service will be splited to different pages and it will also import other libraries to use.

For BackEnd, It can cater to both mobile and desktop as both devices can connect and access the functions through API calls.

---
## Architecture diagram

![image](https://user-images.githubusercontent.com/77374003/145786170-c4d719d5-a1e5-461c-b4b8-08f49acd1071.png)


FrontEnd will send and retrieve data through API calls as BackEnd listens through respectives ports (Driver - 1000, Passenger - 1001 & Trip - 1002).

Based on the API calls from FrontEnd, BackEnd will run functions that will either retreive, insert or update information in the database. 

For Database, its supposed to be in separate Database connections, but for this assignment, it will be in the same database but different table.

Note: BackEnd wont have connection with another Backend, Some functions will have return values hence there is 2 different arrows.

---
## Instructions for setting up and running your microservices

#### SQL (MySQL)
- Run the codes in MySQL-DB-Query.csv to create the respective tables

#### Servers (Passenger, Driver and Trip services)
- cd to src/Driver_MS/Backend & go run main.go
- cd to src/Passenger_MS/Backend & go run main.go
- cd to src/Trip_MS/Backend & go run main.go

#### User Interface (Console)
- cd to src/UI & go run main.go
---
