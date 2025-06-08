# dpoc-service

A service built with the 🚀 Beego web framework

**Example request to the deployed service endpoint: [dopc-service.onrender.com/api/v1/delivery-order-price?venue_slug=home-assignment-venue-helsinki&cart_value=1000&user_lat=60.17094&user_lon=24.93087](https://dopc-service.onrender.com/api/v1/delivery-order-price?venue_slug=home-assignment-venue-helsinki&cart_value=1000&user_lat=60.17094&user_lon=24.93087)**

## Docker

```shell
# Build the Docker image
docker-compose build

# Run the container
docker-compose up
```

Then visit `localhost:8000/swagger/` in your browser.

```shell
$ curl "http://localhost:8000/api/v1/delivery-order-price?venue_slug=home-assignment-venue-helsinki&cart_value=1000&user_lat=60.17094&user_lon=24.93087"
> {
>    "total_price": 1190,
>    "small_order_surcharge": 0,
>    "cart_value": 1000,
>    "delivery": {
>        "fee": 190,
>        "distance": 177
>    }
> }
```

## API Documentation

The endpoints of this API have been documented with Swagger

Base URL: `<host>:8000/api/v1/`

Swagger API Docs: `<host>:8000/swagger/`