FROM scratch AS prod

WORKDIR /app
ADD . /app/

CMD ["./myblog"]