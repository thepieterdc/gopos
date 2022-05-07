# Base image.
FROM golang:1.18.1

# Install dependencies.
RUN apt-get update && apt-get install -y autoconf curl libtool pkg-config

# Install libpostal.
RUN git clone https://github.com/openvenues/libpostal /tmp/libpostal
WORKDIR /tmp/libpostal
RUN ./bootstrap.sh && ./configure --datadir=/var/libpostal/
RUN make -j2 && make install && ldconfig

# Create a working directory for the project.
WORKDIR /go/src/github.com/thepieterdc/gopos

# Copy and install the dependencies.
COPY go.mod go.sum ./
RUN go mod download

# Copy and build the project.
ADD . /go/src/github.com/thepieterdc/gopos
RUN go build .

# Start the project.
CMD ["./gopos"]