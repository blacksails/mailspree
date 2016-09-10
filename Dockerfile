FROM golang:1.7

CMD ["/go/bin/mailspree"]
EXPOSE 8080

# Setup env
ENV APP_DIR /go/src/github.com/blacksails/mailspree

# Copy project files
RUN mkdir -p $APP_DIR
COPY . $APP_DIR

# Install server command
RUN go-wrapper install github.com/blacksails/mailspree/cmd/mailspree
