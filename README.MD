# FindMyRecipe

Find your recipes from available ingredients!
Share your recipes!

## File Structure
1. Frontend
```
- public -> Contains all assets that will be exposed in production
- src
    - assets -> assets that are not exposed (must be imported in svelte files)
    - lib -> svelte components
```
2. Backend
```
- server
    - config -> .env configurations
    - database -> connection to database, and database operations
    - handler -> route handlers; contains functions for each route
    - model -> schema for database tables
    - router -> router for API calls
    - util -> utility functions
- dataset
main.go
```

## Initial Development
- Ensure that you have Node and Go installed in your machine
- **Ensure that you are using Node v18 or above, or else the package dependencies will break. Use `nvm` for node manager**
- Run the following commands
```
node --version 
// ensure that the version is v18 or above
npm install
go mod tidy
```
- Ensure that you have Postgres installed in your machine (preferrably Postgres 14)
- Open postgres through `psql`
- Run the following command
```
CREATE ROLE masterchef LOGIN SUPERUSER PASSWORD 'allrecipes';
CREATE DATABASE findmyrecipe;
```
- If the postgres is running on a different port other than 5432, run the following SQL code in postgres to determine active port
```
SELECT * FROM pg_settings WHERE name = 'port';
```
## For Milestone 1:
- Run `./demo/milestone1.sh`. This executable will create the tables and populate the tables with the sample data. It will
will also define functions for the features implemented in Milestone 1.
- To run the implemented features, run `./demo/featureN.sh` where N is the number of feature. Currently, there are 4 implemented
features in the application.
- To reset or drop all the tables, run `./demo/drop.sh`.

## Running the Application in DEV mode
```
npm run dev
go run .
```
**localhost:3000/** -> Port 3000 is reserved for Vite HMR (hot module replacement; automatically reloads the page with new changes)

## Running the Application in PROD mode
```
npm run build
go run .
```
**localhost:8080/** -> After building the frontend, all bundles should be able to be directly accessed from port 8080. 

## For Milestone 2:
- More in [here](https://github.com/chic27/FindMyRecipe/blob/main/demo/milestone2/milestone.md)
