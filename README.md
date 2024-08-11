## Building Management System
Create a REST API to manage buildings and apartments using Go-Fiber for the web
framework and SQLBoiler for ORM. The database will consist of two tables:
&quot;building&quot; and &quot;apartment&quot;.

### Tools and Technologies:
* Language: Go (Golang)
* Framework: Go-Fiber
* ORM: SQLBoiler
* Database: PostgreSQL (recommended but MySQL or Sqlite3 are fine)

### Database Schema:
#### building
* id: Primary key, integer, auto-increment
* name: String, unique
* address: Text

#### Apartment
* id: Primary key, integer, auto-increment
* building_id: Foreign key, integer (references building.id)
* number: String
* floor: Integer
* sq_meters: Integer

### API Endpoints:
#### Buildings
* GET /buildings: List all buildings (with or without the apartments)
* GET /buildings/{id}: Get a single building by ID
* POST /buildings: Create a new building (update if already exist)
* DELETE /buildings/{id}: Delete a building by ID

#### Apartments
* GET /apartments: List all apartments
* GET /apartments/{id}: Get a single apartment by ID
* GET /apartments/building/{buildingId}: Get all apartments in a specific building
* POST /apartments: Create a new apartment (update if already exist)
* DELETE /apartments/{id}: Delete an apartment by ID