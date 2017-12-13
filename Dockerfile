FROM golang:1.8

WORKDIR /app

RUN go get -u github.com/lsegal/cmd/gucumber

# COPY . .
# RUN go-wrapper download   # "go get -d -v ./..."
# RUN go-wrapper install    # "go install -v ./..."

# CMD ["go-wrapper", "run"] # ["app"]

CMD ["bash"]
