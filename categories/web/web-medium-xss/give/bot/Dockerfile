FROM selenium/standalone-chrome:114.0

USER root

RUN apt-get update && apt-get install -y --no-install-recommends \
    python3-pip \
    socat \
    && rm -rf /var/lib/apt/lists/*

WORKDIR /app

COPY . /app

RUN pip3 install --no-cache-dir -r requirements.txt

CMD socat TCP-LISTEN:7776,fork EXEC:"python3 /app/bot.py"