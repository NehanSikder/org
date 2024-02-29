go build cmd/org/org.go
mkdir integration_tests
mv org integration_tests/org
cd integration_tests
touch "test1.log"
touch "test.txt"
touch "test1"
# ./org --dir=/Users/arhamsikder/Desktop/go/org/integration_tests