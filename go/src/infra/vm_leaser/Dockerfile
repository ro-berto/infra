# Copyright 2022 The LUCI Authors. All rights reserved.
# Use of this source code is governed under the Apache License, Version 2.0
# that can be found in the LICENSE file.

# Use an empty base image
# https://github.com/GoogleContainerTools/distroless
FROM gcr.io/distroless/static-debian11:latest

# The binary is built on the host. Copy it into the image.
COPY vm-leaser-service /

EXPOSE 8080

ENTRYPOINT ["/vm-leaser-service"]
