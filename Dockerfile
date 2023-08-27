FROM public.ecr.aws/docker/library/golang:1.20.0-alpine3.16 AS build-env
RUN apk --no-cache add g++ ca-certificates tzdata build-base git

COPY . /project
WORKDIR /project
RUN go install

FROM public.ecr.aws/docker/library/alpine:3.16
RUN apk --no-cache add g++ ca-certificates tzdata git
WORKDIR /workspace
COPY --from=build-env /go/bin/backend-challenge-mid-level /go/bin/backend-challenge-mid-level
ENTRYPOINT ["/go/bin/backend-challenge-mid-level"]
CMD [ "server" ]
