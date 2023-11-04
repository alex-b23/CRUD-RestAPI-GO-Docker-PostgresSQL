
# CRUD Rest API 

This CRUD Rest API was a solo project in order to showcase my ability to program an API using Go , Docker & PostgresSQL. This is a simple API that send calls to the database. The database it's based on a video game and it stores different information about the user such as the player's username password total wins and loses in game and total ammount of games played


## Installation

To Install my project simply copy the github clone and run the docker file using the commands mentioned below. It's neccesary to have docker installed and running in order to use the docker files.

```bash
  github clone https://github.com/alex-b23/CRUD-RestAPI-GO-Docker-PostgreSQL-Players.git
  docker compose up -d go_db
  docker compose build
  docker compose up go-app
```
    
## Features

- (GET) "/players" - Router in order to see all the players in the database
- (POST) "players - Router used in order to create a new player to the database
- (GET) "/players/{id}" - Router in order to see a player by a specified ID
- (PUT) "/players/{id}" - Router used in order to update a players information in the database
- (DELETE) "/players/{id}" - Router used in order to delete a player by a specified ID 
- Project runs on docker, which means there's no need to install dependicies or other software
- The project runs locally with localhost:3000


## Testing

- The project functionality was tested using Postman
- All routes performed as intended

  
## Tech Stack

- The main programming language used to create this API was Golang
- The database used for this project was PostgresSQL
- The projects utilizes Docker for the database
- The project was tested using Postman 
