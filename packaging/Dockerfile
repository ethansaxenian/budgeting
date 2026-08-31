FROM debian:bookworm-slim AS package

ARG VERSION=0.1.0

WORKDIR /package

COPY bin/budgeting usr/bin/budgeting
COPY packaging/budgeting.service usr/lib/systemd/system/budgeting.service

RUN mkdir -p DEBIAN /out

COPY packaging/postinst DEBIAN/postinst

RUN chmod 0755 DEBIAN/postinst

RUN printf '%s\n' \
  'Package: budgeting' \
  "Version: ${VERSION}" \
  'Section: web' \
  'Priority: optional' \
  'Architecture: arm64' \
  'Maintainer: Ethan Saxenian' \
  'Description: Monthly expense tracking web application' \
  > DEBIAN/control

RUN dpkg-deb --build --root-owner-group . "/out/budgeting_${VERSION}_arm64.deb"

FROM scratch

COPY --from=package /out/ /
