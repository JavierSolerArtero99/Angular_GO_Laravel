FROM golang:1.15

# Argumento que nos permite establecer el nombre del módulo a crear
ARG MODULE_NAME

# Establecemos el directorio de trabajo en función del MODULE_NAME
WORKDIR /go/src/${MODULE_NAME}

# Creamos el módulo de go con el nombre MODULE_NAME e instalamos fresh para hot reloading
RUN rm -f go.mod || true && \
    rm -f go.sum || true && \
    go mod init ${MODULE_NAME} && \
    go mod tidy && \
    go get github.com/pilu/fresh 

# Exponemos el puerto 3000 del contenedor
EXPOSE 8080

# Ejecutamos el comando fresh para iniciar la aplicación
CMD ["fresh"]