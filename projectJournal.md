# T800 Project Journal

## Date: November 10-15, 2019

### Amit: User Profile

#### Features:

- Register new user in system
- Users can update profile
- Users can delete profile
- Create an authentication feature to user

#### Requirements:

- Learn basic React, Go and authentication
- Develop UI component in React for Login.
- Create Login authentication
- Implement Go API’s
  - Create new user profile
  - Authenticate user
  - View user profile
  - Update user profile
  - Delete user profile
- Design architecture for microservice
- Implement logging mechanism
- Create Unit tests

### Neil: Timeline Service

#### Features:

- Everyone can see the timeline page as the default landing page
- Fast READ request on timeline page
- Internal POST image into timeline through Image service
- Internal POST comment count and like count through comment service

#### Requirements:

- Understand React
- Create UI timeline Page using React
- Understanding GO API framework - Gorilla Mux
- Create a microservice structure for timeline service
- Implement GO API’s for this microservice:
  - GET timeline
  - Add Image details
  - Update Comment count
  - Update Like count
- Setup REDIS cluster for storage and replication
- Design scalable architecture - AFK scale cube
- Log all service calls, API calls, and user details
- Create unit tests

### Priyal: Image Service

#### Features:

- Everyone should be able to get the images and their description
- Logged-in Users should be able to post images
- Logged-in users should be able to update description

#### Requirements:

- Learn Go and React
- Understand AWS development pipeline
- UI for team details
- Create a microservice for team details management with the following APIs:
  - Post image
  - Update description for image
  - View image and description
- Setup MongoDB cluster for team details management
- Cloudwatch integration for API logs

### Shabari: Comment & Like Service

#### Features

- Allow users to comment and like posts. (There could be multiple users)
- Users can edit or delete comments on any post
- Users can dislike a previously liked post

#### Requirements

- Understand Go APIs (gorilla/mux), essentials of ReactJs, Kubernetes and AWS development pipeline.
- Setup MongoDB cluster to store grades, award details.
- Go code to perform CRUD operations on comments.
- Unit tests to check functions performing CRUD operations.
- REST APIs to access the CRUD operations.
- Web UI for the comments and like to display to users.

## Date: November 15-20, 2019

### Amit

- Created base code for user profile service.
- Created MongoDB sharded cluster and hosted in EC2 private instance.
- Added get user API.
- React code for front end.
- Setup S3 bucket to store config file.
- Changes in Go code to read S3 bucket credentials from environment variables and use them to get config file so that the code repository does not expose any credentials.

### Priyal

- Created base code for image service.
- Created MongoDB sharded cluster and hosted in EC2 private instance.
- Added get image API.
- React code for front end.
- Setup S3 bucket to store config file.
- Changes in Go code to read S3 bucket credentials from environment variables and use them to get config file so that the code repository does not expose any credentials.
- Added post image API.
- S3 integration to store image to S3 in private bucket.

### Neil

- Base GO API code for timeline microservice
- Redis Cluster on EC2 instances with master slave replication
- Timeline API (GET)
- Image API (SNS POST)
- Dockerize code and deploy to AWS Private network
- Initial React Timeline Page

### Shabari

- Connect to mongodb (mongo atlas) from golang and read simple collections
- Create, Read and Delete operations using golang for comments and likes
- SNS code to publish like and comment updated for timeline
- Deploy in private AWS cluster with Network Load Balancer
- APIs using gorilla/mux to expose the operations

## Date: November 20-23, 2019

### Amit

- Added user login API.
- Added user signup API.
- S3 integration to store image to S3 in private bucket.
- Front end changes for signup and login.
- AWS setup to deploy image service to private ec2 instance and setup network load balancer and API gateway.

### Priyal

- Front end changes to add a post.
- Setup SNS to send messages to timeline service.
- AWS setup to deploy image service to private ec2 instance and setup network load balancer, auto scaling and API gateway.
- API to get all images posted by a particular user.
- API to get all images which will be used as backup when timeline service fails.

### Neil

- UpdateCommentCount API (SNS)
- UpdateLikeCount API (SNS)
- Integrate frontend react Timeline page with all APIs
- Dockerize and deploy all services to AWS

### Shabari

- Mongodb cluster setup with replication and sharding in private AWS instances.
- Code to publish to SNS topics
- API Gateway setup to expose the REST APIs for operations
- Frontend to POST Comments, Likes and read them for each picture.
