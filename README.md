# Plantilla Base para Desarrollo en Go con Docker y VS Code

Este archivo README documenta la plantilla base oficial para cualquier desarrollo en Go dentro de contenedores Docker, diseñada para integrarse fácilmente con Visual Studio Code, proporcionar recarga automática del código fuente, y soportar depuración remota mediante Delve.

Esta configuración está pensada para reutilizarse en múltiples módulos o aplicaciones independientes en Go, especialmente dentro de arquitecturas modulares o basadas en microservicios.

---

## 🚀 Objetivo

Proveer un entorno:
- Portable y reproducible.
- Aislado mediante contenedores Docker.
- Depurable desde VS Code.
- Con recarga automática del código fuente.
- Compatible con redes Docker compartidas (`samm-network`).

---

## 🧰 Tecnologías utilizadas

| Tecnología        | Propósito                                      |
|------------------|-----------------------------------------------|
| **Go**           | Lenguaje de programación principal             |
| **Docker**       | Contenerización del entorno de desarrollo      |
| **Docker Compose** | Orquestación del contenedor y red             |
| **Dev Containers** | Entorno de desarrollo reproducible en VS Code |
| **Delve**         | Depurador para Go                              |
| **Air**           | Recarga automática al guardar archivos         |
| **Visual Studio Code** | Editor y entorno principal                   |

---

## 📁 Estructura del Proyecto

```
mi-aplicacion/
├── .devcontainer/
│   └── devcontainer.json
├── .vscode/
│   └── launch.json
├── .air.toml
├── docker-compose.yml
├── Dockerfile
├── go.mod
├── go.sum
├── main.go
└── README.md
```

---

## 🌍 Descripción de los Archivos Clave

### `Dockerfile`
Define una imagen basada en Go que incluye:
- Go 1.21+
- Delve (debugger)
- Air (hot reload para desarrollo)

El comando por defecto ejecuta Air:
```dockerfile
CMD ["air"]
```

---

### `docker-compose.yml`
- Define el servicio `go-app`.
- Monta tu código local en `/app`.
- Expone los puertos `8080` (HTTP) y `40000` (debug).
- Se conecta automáticamente a la red Docker `samm-network`.

---

### `.devcontainer/devcontainer.json`
- Permite abrir el contenedor desde VS Code como entorno de desarrollo.
- Instala extensiones necesarias y configura el entorno al abrir el contenedor.
- Ejecuta `go mod tidy` después de inicializar.

---

### `.vscode/launch.json`
- Permite hacer debugging en vivo con Delve desde VS Code.
- Requiere que el contenedor esté ejecutando Delve en el puerto `40000`.

---

### `.air.toml`
- Configura Air para recompilar el binario automáticamente al guardar cambios en archivos `.go`.

---

## ✅ Flujo de Trabajo Sugerido

### 1. Abrir el proyecto en Dev Container
```bash
Dev Containers: Reopen in Container
```

### 2. Desarrollo con recarga automática
- Edita archivos `.go` dentro del contenedor.
- Air recompilará automáticamente.

### 3. Depuración
- Coloca breakpoints.
- Usa la configuración "Attach to Go in Docker" para depurar desde VS Code.

### 4. Confirmar sincronización
- Los archivos modificados dentro del contenedor se reflejan en tu sistema de archivos local.

---

## 🧠 Notas Adicionales

- Usa el usuario `root` en el contenedor para evitar problemas de permisos.
- Si necesitas reiniciar el entorno:
```bash
docker-compose down && docker-compose up --build
```
- Puedes cambiar el nombre del servicio o puertos editando `docker-compose.yml`.

---

## ♻️ Reutilización

Puedes clonar esta plantilla para crear nuevos módulos en Go. Solo necesitas:
1. Renombrar la carpeta del proyecto.
2. Actualizar `go.mod` con el nombre correcto del módulo.
3. Cambiar los puertos si hay conflictos con otros servicios.

---

## 📚 Recursos
- [Air](https://github.com/cosmtrek/air) – recarga automática
- [Delve](https://github.com/go-delve/delve) – debugger para Go
- [Go VS Code Extension](https://marketplace.visualstudio.com/items?itemName=golang.Go)
- [Dev Containers](https://containers.dev/)

---

Esta plantilla es ahora la base oficial para todos los desarrollos en Go dentro del ecosistema de contenedores Docker + VS Code.