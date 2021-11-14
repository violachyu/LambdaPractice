0. Tutorial: https://www.youtube.com/watch?v=x_yCX4kSchY
1. create directory: /LambdaPractice
2. init go module: $ go mod init <myapp>
3. import lambda, echo: 
    $ go get github.com/labstack/echo/v4
    $ go get github.com/aws/aws-lambda-go/lambda
4. compile into .zip: $ GOOS=linux go build -o main
5. Create lambda function & set config
