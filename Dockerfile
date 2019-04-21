FROM scratch

ENV PORT 8000
EXPOSE $PORT

COPY mongo-api /
CMD ["/mongo-api"]
