# API Manager – Proyecto SAMM

Este documento detalla la construcción, ejecución y depuración del **API Manager** del sistema SAMM. El enfoque actual está en levantar correctamente el microservicio usando Docker y depurarlo desde Visual Studio Code.

---

## 🧠 ¿Qué es el API Manager?

Es un microservicio escrito en **Go (Golang)** que actúa como puerta de entrada para redirigir solicitudes HTTP hacia los distintos módulos (como `users-module`) del sistema SAMM. Por ahora, el objetivo principal es ponerlo en funcionamiento y habilitar su depuración.

---

## 🧰 Tecnologías utilizadas

| Tecnología         | Propósito                                         |
|--------------------|--------------------------------------------------|
| **Go 1.21**         | Lenguaje base del API Manager                    |
| **Gin**             | Framework HTTP para crear el servidor            |
| **Docker**          | Empaquetado y ejecución del servicio             |
| **Delve**           | Depurador oficial de programas Go                |
| **VS Code**         | Editor desde donde se conecta el depurador       |

---

## 📁 Estructura del Proyecto

```
api-manager/
├── Dockerfile         # Contenedor del microservicio
├── go.mod             # Módulo de Go con dependencias
├── main.go            # Código principal del servicio
└── .vscode/
    └── launch.json    # Configuración de debugging remoto en VS Code
```

---

## 🐳 Dockerfile explicado

```dockerfile
FROM golang:1.21

WORKDIR /app

RUN go install github.com/go-delve/delve/cmd/dlv@latest

COPY . .

RUN go mod tidy
RUN go build -o api-manager .

EXPOSE 8080 40000

CMD ["dlv", "debug", "--headless", "--listen=0.0.0.0:40000", "--api-version=2", "--accept-multiclient"]
```

### ¿Qué hace?
- Usa la imagen oficial de Go.
- Instala Delve para debugging.
- Copia el proyecto y sus dependencias.
- Compila el binario.
- Expone el puerto 8080 (HTTP) y 40000 (debug).
- Ejecuta el binario en modo depuración.

---

## 🏃‍♂️ Cómo levantar el contenedor

### 1. Construir la imagen

```bash
docker build -t api-manager-debug .
```

### 2. Ejecutar el contenedor

```bash
docker run --rm \
  -p 8080:8080 \
  -p 40000:40000 \
  --network samm-network \
  api-manager-debug
```

> Esto lo conecta a la red de SAMM y habilita debugging desde el host.

---

## 🐞 Depuración desde VS Code

### Requisitos
- Extensión oficial de **Go** para VS Code.

### Archivo `.vscode/launch.json`

```json
{
  "version": "0.2.0",
  "configurations": [
    {
      "name": "Attach to Go in Docker",
      "type": "go",
      "request": "attach",
      "mode": "remote",
      "remotePath": "/app",
      "port": 40000,
      "host": "localhost",
      "program": "${workspaceFolder}",
      "showLog": true
    }
  ]
}
```

### Pasos para debug
1. Abrir `main.go` y colocar un breakpoint.
2. Ir al panel de depuración.
3. Seleccionar **Attach to Go in Docker**.
4. Presionar ▶️ Iniciar depuración.

---

## 🔍 Estado actual

- ✅ API Manager corriendo en Go dentro de Docker.
- ✅ Integrado con la red de SAMM.
- ✅ Puede recibir solicitudes (con token `Bearer`).
- ✅ Depurable directamente desde VS Code con breakpoints.

---

## 📌 Siguiente(s) pasos sugeridos

- Implementar validación real de JWT.
- Añadir lógica de enrutamiento dinámico desde base de datos.
- Centralizar logs o métricas.
- Integrar un sistema de configuración externo (archivos `.env` o servicio de config).

---