# Nginx Traffic Intelligence Agent

Agente local instalable por `.deb` para analizar logs estructurados de Nginx, clasificar tráfico humano vs bot, estimar sesiones y exponer métricas vía API y dashboard web.

## Estado Actual
- Fase: bootstrap de arquitectura y estándares.
- Alcance actual: documentación base, ADR inicial y convenciones de desarrollo.
- Aún no hay implementación de backend/frontend.

## Stack Decidido
- Backend daemon + API: Go (arquitectura hexagonal, puertos/adaptadores).
- Frontend: React + TypeScript (diseño basado en componentes).
- Entrega operativa: paquete Debian (`.deb`) + `systemd` + integración segura con Nginx.

## Lectura Rápida (orden recomendado)
1. ADR inicial: [`.adr/0001-initial-architecture-and-tech-stack.md`](.adr/0001-initial-architecture-and-tech-stack.md)
2. Guía de trabajo y ownership: [`AGENTS.md`](AGENTS.md)
3. Índice operativo: [`.ai/README.md`](.ai/README.md)
4. Arquitectura y seguridad:
   - [`.ai/standards/ARCHITECTURE.md`](.ai/standards/ARCHITECTURE.md)
   - [`.ai/standards/SECURITY.md`](.ai/standards/SECURITY.md)
   - [`.ai/general/THREAT_MODEL.md`](.ai/general/THREAT_MODEL.md)
5. Plan y alcance:
   - [`.ai/general/ROADMAP.md`](.ai/general/ROADMAP.md)
   - [`.ai/general/IMPLEMENTATION_PLAN.md`](.ai/general/IMPLEMENTATION_PLAN.md)
   - [`.ai/general/DEFINITIONS.md`](.ai/general/DEFINITIONS.md)

## Principios del Repositorio
- Cambios pequeños, deterministas y con rollback claro.
- Sin breaking changes sin ADR.
- Seguridad por defecto (least privilege, validación estricta, redacción de datos sensibles).
- Integración con Nginx siempre segura (`nginx -t` antes de reload).
