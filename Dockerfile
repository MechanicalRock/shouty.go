FROM golang:1.17

WORKDIR /app

#COPY . .
#RUN go-wrapper download   # "go get -d -v ./..."
#RUN go-wrapper install    # "go install -v ./..."


#RUN go-wrapper download -u github.com/gucumber/gucumber/cmd/gucumber

#CMD ["go-wrapper", "run"] # ["app"]

CMD ["bash"]
