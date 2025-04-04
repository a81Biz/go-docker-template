FROM golang:1.23

WORKDIR /app

# Copiar el código fuente
COPY . .

# Instalar Air (live-reload) y Delve (debugger)
RUN go install github.com/air-verse/air@latest && \
    go install github.com/go-delve/delve/cmd/dlv@latest

# Inicializar módulo Go si no existe
RUN [ ! -f go.mod ] && go mod init api-manager || true

# Descargar dependencias
RUN go mod tidy

# Compilar con flags para debugging
RUN go build -gcflags="all=-N -l" -o /app/main .

# Puerto para la app (8080) y para debugging (40000)
EXPOSE 8080 40000

# Comando para iniciar ambos servicios (Air + Delve)
CMD ["sh", "-c", "air -c .air.toml & dlv debug --headless --listen=:40000 --api-version=2 --accept-multiclient --continue --output=/tmp/__debug_bin ./main.go"]