# Bright-Life
This project is a backend server for life care.
 
# Contributor 
Randy Cummings

# Features  

- CRUD
- Microservice Architecture

    2 microservices

- CI/CD

    Unit Test + Github action

    Continuous integration and deploying with CircleCI

- JWT

    Key generation with openSSL

    Verify with token

- Serverless Architecture

# Tech Stack  

**Server:** Gin Framework(GoLang), AWS

**DB:** MySQL, MongoDB, HarperDB

# Functions  
1. Management users(Auth)
    1. Register
    2. Sign In / Sign Up
2. REST API
    1. HTTPS connection
3. Analysis DB

    1. Analyze the sport events and get the result information from the database.


# Run Locally  

Clone the project  

~~~bash  
  git clone https://github.com/speedspeedo/bright-life.git
~~~

Install dependencies  

~~~bash  
npm install
~~~

Start the server  

~~~bash  
nodemon --exec go run . --signal SIGTERM
~~~

# Environment Variables  

To run this project, you will need to add the following environment variables to your .env file  
`API_KEY`  

`ANOTHER_API_KEY` 

## Feedback  

If you have any feedback, please send an email to speedodeveloper1004@gmail.com
