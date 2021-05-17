FROM scratch
WORKDIR /app
ADD config config
ADD walletTracky /app/
ENTRYPOINT ["./walletTracky"]
EXPOSE 8080