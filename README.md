# Car Objects API

## Functionality

    1. Create a new car object: Send a POST request to http://localhost:3000 with JSON payload:
    ```
    {
      "mark" : "ford",
      "model" : "focus",
      "number" : "am999r",
      "action" : "test",
      "distance" : 10000,
      "year" : 2016
    }
    ```

    2. Get data of a car object by ID: Send a GET request to http://localhost:3000?type=one&id=7 (with parameter type=one)

    3. Get data of all car objects in the table: Send a GET request to http://localhost:3000?type=all (with parameter type=all)

## Deployment

To deploy the project, run the following command:

`docker-compose up -d`