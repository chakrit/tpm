FROM alpine AS source
WORKDIR /src
RUN apk add go

ADD go.mod /src
ADD go.sum /src
RUN go mod download

ADD . /src
RUN go build -o tpm .

# ---

FROM alpine
COPY --from=source /src/views /views
COPY --from=source /src/tpm /
CMD /tpm