## Authentication-Authorization-Golang
An organization-user JWT authentication system. Made using golang, gin, gorm and go-jwt.

#### INSTALLATION

 - Clone this repo
 - Change directory to the cloned repo
 - Ensure you have `go` installed on your machine.
 - Run `go mod download`
 - Ensure you have `postgresql` installed on your machine OR you can create a cloud `postgres` database
 - Run `cp .env.sample .env`
 - Set the environment variables as mentioned in the `.env.sample`.
 - Run `go run main.go`

#### DATABASE USED

`postgresql` has been used for this task because I have been building web-apps using postgresql from a long time. I have build many backend API services using `Django` all using `postgresql`. Unlike mongoDB, it is a RDBMS, hence defining relations between 2-3 tables becomes pretty easy.

#### FRAMEWORKS AND LIBRARIES USED

 - `gin` - It is most popular go framework that is designed for building APIs. Hence, I have used gin to complete the task.
 - `go-jwt` - It is the most popular package for implementing JWTs in Go. It has many pre-defined functions for signing and validating JWTs. Hence, used I have used it here, to sign and validate access_token and refresh_token.
 - `gorm` - It is the most popular ORM package in the Go ecosystem. Makes it really easy to create models and querying through database.

#### DATABASE STRUCTURE

<img width="949" alt="image" src="https://user-images.githubusercontent.com/62551163/227226472-5aa441da-527f-4296-af1e-24777387a7f7.png">


#### APIs BUILD

 - **/removeMember** POST - Remove the member from an organization provided the current user is an admin of the organization.
 - **/createUser** POST - Creates a new user for an organization provided that the current user is an admin of the organization.
 - **/getUsersFromOrganization/{organizationID}** GET - Get members of an organization provided the current user is also a member.
 - **/login** POST - The current user is logged in, provided the correct username and password is passed.
 - **/logout** GET - The current user is logged out.
 
 The following APIs were not supposed to be included in the task but were build to make data modification easier.
 
 - **/createOrganization** POST - Creates an organization and makes the creator of the organization an admin of the organization.
 - **/makeUserAdmin** POST - Makes an existing user an admin of an organization. [No validations here. Just for modifying data.]

#### REFERENCE VIDEO

https://user-images.githubusercontent.com/62551163/227226228-7cd1b34d-5192-4960-902a-c85eea0c6631.mp4



