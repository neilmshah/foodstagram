# T800 Project Journal

Date: November 10, 2019

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

### Priyal: Team Management

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

### Shabari: Grading and Awards

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

