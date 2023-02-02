FROM golang:1.19

RUN mkdir /tcgraf

ADD . /tcgraf

WORKDIR /tcgraf

RUN go build -o tcg

CMD ["/tcgraf/tcg"]