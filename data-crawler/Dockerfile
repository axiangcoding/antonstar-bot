FROM python:3.10-alpine
RUN #sed -i 's/dl-cdn.alpinelinux.org/mirrors.aliyun.com/g' /etc/apk/repositories
RUN apk update && apk upgrade && apk add --no-cache gcc libressl-dev musl-dev libffi-dev chromium chromium-chromedriver
ENV PATH="/usr/bin/chromedriver:${PATH}"
RUN pip install -i https://pypi.tuna.tsinghua.edu.cn/simple pip -U
RUN pip config set global.index-url https://pypi.tuna.tsinghua.edu.cn/simple

WORKDIR /app
COPY . /app/
RUN pip install -r requirements.txt
ENTRYPOINT ["python", "-u","main.py"]