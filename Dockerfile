FROM scratch
ADD ./stripper /stripper
ENTRYPOINT ["/stripper"]
