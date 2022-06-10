# JUMIA MDS CHALLENGE

Platform that allows management of products for ecommerce platform.   Products should have an unique SKU and could be commercialized in multiple countries.

## Explicit Features/Tasks

- [x]  Provide an api that can get a product by SKU
- [x]  Provide API to sell products with validation of availability of stock
- [x]  Provide API for bulk update of orders from CSV
- [x]  For each CSV line, the stock update could be positive or negative
- [x]  If a product doesn’t exist, it should be created.
- [x]  When a stock goes below a configurable threshold, an alert should be emitted so that we could order more of that product and avoid an out of stock scenario.

## Application requirements stack used

1. Postgres database as the core database
2. Gofiber main Golang web framework used
3. Gorm as ORM for the database
4. Docker to build the application images
5. Docker-compose to run the application in an isolated environment 

### Application setup

1. Retrieve the source code from git 

```bash
git clone https://github.com/ryananyangu/jumia.git
```

1. Get into the directory

```bash
cd jumia
```

1. Install application dependancies

```bash
go mod tidy
```

1. Build application docker image

```bash
make build
```

1. Update the Environment variables on docker compose file **docker-compose.yml .** You can also leave the default configurations for testing purposes.

```yaml
# Postgres admin default login credentials update to your preffered
PGADMIN_DEFAULT_EMAIL: admin@admin.com
PGADMIN_DEFAULT_PASSWORD: root

# Main application environment variables make sure to match with DB
ENV: PROD
DB_HOST: db
DB_PORT: 5432
DB_USER: postgres
DB_PASSWORD: postgres
DB_NAME: postgres

# Default postgres credentials
POSTGRES_USER: postgres
POSTGRES_PASSWORD: postgres
```

1. Starting the application in an isolated environment using docker compose

```bash
make run-env
```

1. Application has started at 
    1. [App monitoring link](http://localhost:8080/Dashboard)
    2. [Postgres GUI link](http://localhost:5050/) 
    3. [App base API link](http://localhost:8080/api/v1)

## API’s Specifications

1. [Create product API link](http://127.0.0.1:8080/api/v1/product/update)
    
    Sample request
    
    ```bash
    curl --location --request POST 'http://127.0.0.1:8080/api/v1/product/update' \
    --header 'Content-Type: application/json' \
    --data-raw '{
        "name":"Samsung Phone",
        "sku" : "UYUT-879847564793-PO",
        "stock":2,
        "alertThold":10,
        "country":"ke"
    }'
    ```
    
    Sample response
    
    ```json
    {
        "ID": 1,
        "CreatedAt": "2022-06-09T14:50:06.893242148+03:00",
        "UpdatedAt": "2022-06-09T14:50:36.320068757+03:00",
        "DeletedAt": null,
        "name": "Samsung Phone",
        "sku": "UYUT-879847564793-PO",
        "stock": 14,
        "country": "ke",
        "alertThold": 10
    }
    ```
    
2. [Get product by SKU and country code](http://127.0.0.1:8080/api/v1/product/get?sku=UYUT-879847564793-PO&country=ke) link
    
    Sample request
    
    ```bash
    curl --location --request GET 'http://127.0.0.1:8080/api/v1/product/get?sku=UYUT-879847564793-PO&country=ke'
    ```
    
    Sample response 
    
    ```json
    {
        "ID": 1,
        "CreatedAt": "2022-06-09T14:50:06.893242148+03:00",
        "UpdatedAt": "2022-06-09T14:55:23.263848332+03:00",
        "DeletedAt": null,
        "name": "Samsung Phone",
        "sku": "UYUT-879847564793-PO",
        "stock": 14,
        "country": "ke",
        "alertThold": 10
    }
    ```
    
3. [Selling a product link](http://127.0.0.1:8080/api/v1/product/sell)
    
    Sample request
    
    ```bash
    curl --location --request POST 'http://127.0.0.1:8080/api/v1/product/sell' \
    --header 'Content-Type: application/json' \
    --data-raw '{
        "sku": "UYUT-879847564793-PO",
        "stock": -1,
        "country": "ke"
    }'
    ```
    
    Sample response 
    
    ```json
    {
        "ID": 1,
        "CreatedAt": "2022-06-09T15:11:48.678974033+03:00",
        "UpdatedAt": "2022-06-09T15:18:42.179705072+03:00",
        "DeletedAt": null,
        "name": "Samsung Phone",
        "sku": "UYUT-879847564793-PO",
        "stock": 1,
        "country": "ke",
        "alertThold": 10
    }
    ```
    
4. [Stock bulk update link](http://127.0.0.1:8080/api/v1/product/bulk/update)
    
    Sample request 
    
    ```bash
    curl --location --request POST 'http://127.0.0.1:8080/api/v1/product/bulk/update' \
    --data-raw '[
        {
            "name": "Samsung Phone",
            "sku": "UYUT-879847564793-PO",
            "stock": 14,
            "country": "ke",
            "alertThold": 10
        },
        {
            "name": "Samsung TAB",
            "sku": "UYUT-879847564793-TAB",
            "stock": 200,
            "country": "ke",
            "alertThold": 10
        },
        {
            "name": "Samsung TV",
            "sku": "UYUT-879847564793-TV",
            "stock": 5,
            "country": "ke",
            "alertThold": 10
        },
        {
            "name": "Samsung Phone",
            "sku": "UYUT-879847564793-PO",
            "stock": -6,
            "country": "ke",
            "alertThold": 10
        }
    ]'
    ```
    
    Sample response 
    
    ```json
    {
        "message": "Request recieved successfull\n"
    }
    ```
    
5. [Product alert threshold update link](http://127.0.0.1:8080/api/v1/product/threshold/update)
    
    Sample request
    
    ```bash
    curl --location --request PUT 'http://127.0.0.1:8080/api/v1/product/threshold/update' \
    --header 'Content-Type: application/json' \
    --data-raw '{
        "ID" : 1,
        "alertThold":20
    }'
    ```
    
    Sample response 
    
    ```json
    {
        "message": "Product ID 1 updated successfull"
    }
    ```
    
6. [Product update using batch file](http://127.0.0.1:8080/api/v1/product/stock/update/batchfile)
    
    Sample request
    
    ```bash
    curl --location --request POST 'http://127.0.0.1:8080/api/v1/product/stock/update/batchfile' \
    --form 'fileUpload=@"/{path}/file_2.csv"'
    ```
    
    Sample response
    
    ```json
    {
        "message": "Request recieved successfull\n"
    }
    ```
    

## NOTES:

- Runtime of the batch file is relatively high still after the following steps to optimize were done
    - Processing the records concurrently using go routines
    - Making sure the number of spawned go routines do not exceed the database max connections setting. used a channel to block go routines until previous are done.
    - Create index on the database for **sku** and **country** given that they are the most commonly queried against columns. **This reduced the slow queries**
- Proposal for handling the long processing time
    - Usage of queues that the main web application is the producer and consumers are the control and insert into the database.