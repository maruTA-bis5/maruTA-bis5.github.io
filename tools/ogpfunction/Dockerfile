FROM golang:1.19-alpine AS builder
WORKDIR /src/github.com/maruTA-bis5/maruTA-bis5.github.io/tools/ogpfunction/
ADD . /src/github.com/maruTA-bis5/maruTA-bis5.github.io/tools/ogpfunction/
RUN go mod tidy
# make portable binary
# https://okzk.hatenablog.com/entry/2016/08/03/234738
RUN CGO_ENABLED=0 go build -o bin/ogp main.go

FROM mcr.microsoft.com/azure-functions/dotnet:4-appservice 
ENV AzureWebJobsScriptRoot=/home/site/wwwroot \
    AzureFunctionsJobHost__Logging__Console__IsEnabled=true

COPY ./*.json /home/site/wwwroot/
COPY ./link-card-api/ /home/site/wwwroot/link-card-api/
COPY --from=builder /src/github.com/maruTA-bis5/maruTA-bis5.github.io/tools/ogpfunction/bin/ogp /home/site/wwwroot/bin/
