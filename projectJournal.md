# T800 Project Journal

### Neil: Contest Management

#### Features:
* Everyone can see the contest page as the default landing page
* Admins can create a contest for a competition or project
* Admins can update a contest’s description, etc. if required
* Admins can delete/deactivate/close contests
* If student clicks on any of the contest, it will redirect to team management page

#### Requirements:
* Understand React 
* Create UI Contest Page using React
* Understanding GO API framework - Gorilla Mux
* Create a microservice structure for Contest Management
* Implement GO API’s for this microservice:
  * Create Contest
  * Update Contest
  * Delete Contest
  * View Contests
* Setup MongoDB cluster for storage and replication
* Setup Redis elastic cache for reading contests
* Design scalable architecture - AFK scale cube
* Log all service calls, API calls, and user details
* Create unit tests 

### Shabari: Grading and Awards
#### Features
* Allow the creator of the contest to assign grades. (There could be multiple graders)
* Creator will also have the option to provide awards for the contest.
* Once the final grades are out anyone can view the leaderboard.
* The Creator will be able to view scores given by each grader.

#### Requirements
* Understand Go APIs (gorilla/mux), essentials of ReactJs, Kubernetes and AWS development pipeline.
* Setup MongoDB cluster to store grades, award details.
* Setup Redis to dynamically keep track of leaderboard.
* Go code to perform CRUD operations on grades.
* Unit tests to check functions performing CRUD operations.
* REST APIs to access the CRUD operations.
* Web UI for graders and the general public.
