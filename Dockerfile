FROM python:3.12-slim AS builder

WORKDIR /app

RUN pip install --no-cache-dir pytest

COPY compute.py test_compute.py .

RUN python -m pytest -v

FROM python:3.12-slim

COPY --from=builder /app/compute.py /bin/compute

RUN chmod +x /bin/compute
