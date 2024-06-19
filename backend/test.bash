#!/bin/bash

SERVER="http://127.0.0.1:8080"

# Test Product endpoints
printf "====================================================================================================\n"
echo "Testing Product endpoints..."
echo "Getting all products..."
curl -s -X GET ${SERVER}/products | jq
echo "Getting product with ID 1..."
curl -s -X GET ${SERVER}/products/1 | jq
echo "Getting products in category with ID 1..."
curl -s -X GET ${SERVER}/products/category/1 | jq
echo "Getting products with name 'iPhone'..."
curl -s -X GET ${SERVER}/products/name/T-shirt | jq
echo "Getting products with price 999.99..."
curl -s -X GET ${SERVER}/products/price/999.99 | jq
echo "Creating a new product..."
curl -s -X POST -H "Content-Type: application/json" -d '{"name":"New Product","category_id":1,"description":"New product description","price":199.99}' ${SERVER}/products | jq
echo "Getting all products..."
curl -s -X GET ${SERVER}/products | jq
echo "Updating product with ID 4..."
curl -s -X PUT -H "Content-Type: application/json" -d '{"name":"Updated Product","category_id":1,"description":"Updated product description","price":299.99}' ${SERVER}/products/4 | jq
echo "Getting all products..."
curl -s -X GET ${SERVER}/products | jq
echo "Deleting product with ID 4..."
curl -s -X DELETE ${SERVER}/products/4 | jq
echo "Getting all products..."
curl -s -X GET ${SERVER}/products | jq
printf "====================================================================================================\n"

# Test Category endpoints
printf "====================================================================================================\n"
echo "Testing Category endpoints..."
echo "Getting all categories..."
curl -s -X GET ${SERVER}/categories | jq
echo "Getting category with ID 1..."
curl -s -X GET ${SERVER}/categories/1 | jq
echo "Getting category with name 'Electronics'..."
curl -s -X GET ${SERVER}/categories/name/Electronics | jq
echo "Creating a new category..."
curl -s -X POST -H "Content-Type: application/json" -d '{"name":"New Category"}' ${SERVER}/categories | jq
echo "Getting all categories..."
curl -s -X GET ${SERVER}/categories | jq
echo "Updating category with ID 4..."
curl -s -X PUT -H "Content-Type: application/json" -d '{"name":"Updated Category"}' ${SERVER}/categories/4 | jq
echo "Getting all categories..."
curl -s -X GET ${SERVER}/categories | jq
echo "Deleting category with ID 4..."
curl -s -X DELETE ${SERVER}/categories/4 | jq
echo "Getting all categories..."
curl -s -X GET ${SERVER}/categories | jq
printf "====================================================================================================\n"

# Test Cart endpoints
printf "====================================================================================================\n"
echo "Testing Cart endpoints..."
echo "Getting cart..."
curl -s -X GET ${SERVER}/cart | jq
echo "Adding product with ID 1 to cart..."
curl -s -X POST -H "Content-Type: application/json" -d '{"product_id":1,"quantity":1}' ${SERVER}/cart | jq
echo "Getting cart..."
curl -s -X GET ${SERVER}/cart | jq
echo "Updating quantity of product with ID 1 in cart..."
curl -s -X PUT -H "Content-Type: application/json" -d '{"product_id":1,"quantity":2}' ${SERVER}/cart/product/1 | jq
echo "Getting cart..."
curl -s -X GET ${SERVER}/cart | jq
echo "Deleting product with ID 1 from cart..."
curl -s -X DELETE ${SERVER}/cart/product/1 | jq
echo "Getting cart..."
curl -s -X GET ${SERVER}/cart | jq
echo "Adding product with ID 2 to cart..."
curl -s -X POST -H "Content-Type: application/json" -d '{"product_id":2,"quantity":1}' ${SERVER}/cart | jq
echo "Adding product with ID 3 to cart..."
curl -s -X POST -H "Content-Type: application/json" -d '{"product_id":3,"quantity":2}' ${SERVER}/cart | jq
echo "Getting product with ID 2 from cart..."
curl -s -X GET ${SERVER}/cart/product/2 | jq
echo "Getting total price of cart..."
curl -s -X GET ${SERVER}/cart/total | jq
echo "Clearing cart..."
curl -s -X DELETE ${SERVER}/cart | jq
echo "Getting cart..."
curl -s -X GET ${SERVER}/cart | jq
printf "====================================================================================================\n"