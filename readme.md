# Laravel + GO + Angular

## Deploy

### Frontend

Run in command line:

`$ cd frontend/`

`$ npm run install`

`$ npm run start`

### Backend

Backend is deployed by a docker compose file inside backend folder

`$ cd backend/`

`$ sudo docker-compose up --build`

- #### Database:

![](./docs/docker-compose-database.png)

- #### Laravel:

![](./docs/docker-compose-laravel.png)

- #### Products Microservice:

![](./docs/docker-compose-products.png)

- #### Users Microservice:

![](./docs/docker-compose-users.png)

- #### Prometheus:

![](./docs/docker-compose-prometheus.png)

- #### redis:

![](./docs/docker-compose-redis.png)

- #### traefik:

![](./docs/docker-compose-redis.png)

### Pages

![](./docs/page-home-noAuth.png)

This is the main page where you can see all the products without being logged.

![](./docs/page-login.png)

You can registry and login with a user

![](./docs/page-auth.png)

Once the user is logged the header will be updated

![](./docs/page-auth-admin.png)

If the user has enought provileges, he will be admin and will see the "Panel Admin" screen.

![](./docs/page-product.png)

This is the product's page where all the info of it will be displayed. 

![](./docs/page-comments.png)

Also the users can publish comments to a single product