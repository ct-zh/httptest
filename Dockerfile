FROM alpine
ADD httptest /httptest
ENTRYPOINT ["/httptest"]
EXPOSE 11234