echo "QUERYING...."
ab -n 1000 -c 4 'http://localhost:8080/query'
echo "QUERYING DONE"