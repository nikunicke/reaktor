# reaktor warehouse

Here is my implementation for the *Reaktor Junior 2021 Spring* pre-assignment. Live version running on [**Heroku**](https://guarded-cliffs-12756.herokuapp.com/)

> *Your client is a clothing brand that is looking for a simple web app to use in their warehouses. To do their work efficiently, the warehouse workers need a fast and simple listing page per product category, where they can check simple product and availability information from a single UI.*

Read more about the assignment [**here**](https://www.reaktor.com/junior-dev-assignment/)

---

My solution consists of two parts. The first being a *middleman* backend that makes concurrent requests to the clients legacy APIs and processes the response data accordingly. The second part is a small [**React app**](), which renders a simple user interface on the browser.

### [Warehouse](https://github.com/nikunicke/reaktor/tree/master/warehouse) (backend)

The warehouse consists of three domain types:
* Warehouse
    * Represents the warehouse as a whole and includes product- and availability services.
    * Updates the warehouse by requesting data from the legacy APIs in a concurrent manner.
* Products
    * Represents a product. Currently supports implementations for finding all products within a category
* Availability
    * Represents the availability data. This type is split into other types in order to simplify parsing of xml fields.

##### Packages
* api
    * Implements a client that communicates with the legacy APIs provided in the assignment
* http
    * Implements a server that communicates with requests from the frontend

These packages and types are put together in cmd/warehouse/main.

The backend is currently serving a static build of the react app. To run and serve on your system, simply run
```console
    docker-compose up
```
Navigate to [localhost:5000](http://localhost:5000) or whichever $PORT you have set up in your env variables.

### [Warehouse-client](https://github.com/nikunicke/reaktor/tree/master/warehouse-client) (frontend)

The warehouse-client uses [*react-bootstrap-table2*](https://react-bootstrap-table.github.io/react-bootstrap-table2/) to render a simple table for each product category. Unfortunately I forgot to join the values in the color field, which resulted in values with multiple colors to be displayed without spaces between them.

To run separately:
```console
    cd warehouse-client
    npm start
```
