# 基础镜像
FROM golang:1.20.2

# 设置工作目录
WORKDIR /app

# 将项目文件复制到容器中
COPY . .

# 构建Go项目
RUN go build -o app

# 暴露应用程序的端口
EXPOSE 3000

# 运行应用程序
CMD ["./app"]
