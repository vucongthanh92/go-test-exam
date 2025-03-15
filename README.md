# go-test-exam

use Makefile to run
- make run: start project
- make wire: build dependencies
- make swagger: build documentation


URL document swagger: http://localhost:5001/swagger/index.html#


API:
- get list category:
curl --location 'http://localhost:5001/api/v1/categories' \
--data ''


- get list product by filter:
curl --location --request GET 'http://localhost:5001/api/v1/products?pageSize=100&pageIndex=1&status=Available&category_id=94d0da61-0bbe-4be8-8435-2b72f03a29ea&from_date=2024-01-02&to_date=2024-01-05' \
--header 'Content-Type: application/json' \
--data '{
    "category_name": "category 01"
}'


- generate PDF
curl --location --request GET 'http://localhost:5001/api/v1/product/gen-pdf?pageSize=100&pageIndex=1&status=Available' \
--header 'Content-Type: application/json' \
--data '{
    "category_name": "category 01"
}'


- product per category
curl --location --request GET 'http://localhost:5001/api/v1/statistics/products-per-category' \
--header 'Content-Type: application/json' \
--data '{
    "category_name": "category 01"
}'


- product per supplier
curl --location --request GET 'http://localhost:5001/api/v1/statistics/products-per-supplier' \
--header 'Content-Type: application/json' \
--data '{
    "category_name": "category 01"
}'


- calculate distance:
curl --location 'http://localhost:5001/api/v1/distance/stock_city?city=Marseille' \
--data ''