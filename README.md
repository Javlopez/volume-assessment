# volume-assessment

This is the readme for api calculation

We have some commands in order to run the app:

- make test
  - Will run the tests
- make api
  - Will run the api on port 8080
- make docs
  - Will run swagger on different port each time

### Make test
```bash
$ make test 
echo "running tests"
running tests
go test ./... -v
=== RUN   TestApi
=== RUN   TestApi/Should_raise_error_if_path_is_wrong
=== RUN   TestApi/Should_raise_error_if_missing_body_data
=== RUN   TestApi/Should_raise_error_if_the_json_is_malformed
=== RUN   TestApi/Should_raise_error_if_the_route_is_malformed
=== RUN   TestApi/Should_return_1_flight_paths
=== RUN   TestApi/Should_return_2_flight_paths
=== RUN   TestApi/Should_return_4_flight_paths
--- PASS: TestApi (0.00s)
    --- PASS: TestApi/Should_raise_error_if_path_is_wrong (0.00s)
    --- PASS: TestApi/Should_raise_error_if_missing_body_data (0.00s)
    --- PASS: TestApi/Should_raise_error_if_the_json_is_malformed (0.00s)
    --- PASS: TestApi/Should_raise_error_if_the_route_is_malformed (0.00s)
    --- PASS: TestApi/Should_return_1_flight_paths (0.00s)
    --- PASS: TestApi/Should_return_2_flight_paths (0.00s)
    --- PASS: TestApi/Should_return_4_flight_paths (0.00s)
PASS
ok  	github.com/Javlopez/volume-assessment	0.002s
=== RUN   TestFlightsCalculation
=== RUN   TestFlightsCalculation/Should_return_1_flight
=== RUN   TestFlightsCalculation/Should_return_2_flight
=== RUN   TestFlightsCalculation/Should_return_4_flight
--- PASS: TestFlightsCalculation (0.00s)
    --- PASS: TestFlightsCalculation/Should_return_1_flight (0.00s)
    --- PASS: TestFlightsCalculation/Should_return_2_flight (0.00s)
    --- PASS: TestFlightsCalculation/Should_return_4_flight (0.00s)
=== RUN   TestQueue
=== RUN   TestQueue/Should_return_empty_queue
=== RUN   TestQueue/Should_be_able_to_push_data_into_the_queue
=== RUN   TestQueue/Should_be_able_to_shift_element_from_the_queue
--- PASS: TestQueue (0.00s)
    --- PASS: TestQueue/Should_return_empty_queue (0.00s)
    --- PASS: TestQueue/Should_be_able_to_push_data_into_the_queue (0.00s)
    --- PASS: TestQueue/Should_be_able_to_shift_element_from_the_queue (0.00s)
PASS
ok  	github.com/Javlopez/volume-assessment/calculate	(cached)
?   	github.com/Javlopez/volume-assessment/docs	[no test files]
```

### Make Docs
```bash
$ make docs
echo "starting swagger documentation"
starting swagger documentation
swagger serve -F=swagger swagger.yaml
2022/09/24 01:18:21 serving docs at http://localhost:46129/docs
```

### Make Api
```bash
$ make api      
echo "running api"
running api
go run main.go
2022/09/24 01:19:08 Api is running.....

```