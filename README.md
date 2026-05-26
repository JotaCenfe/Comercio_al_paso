# Comercio al Paso

## Sistema de Gestión de E-commerce en Go

## Descripción

**Comercio al Paso** es un sistema básico de gestión de e-commerce desarrollado en el lenguaje de programación **Go**. El nombre del sistema hace referencia a un comercio sencillo, rápido y práctico, donde el usuario puede realizar una compra de manera ágil mediante una aplicación de consola.

El sistema permite registrar los datos de un cliente, visualizar productos disponibles, agregar productos al carrito de compras, generar un pedido y simular un proceso de pago.

Este proyecto fue desarrollado como parte de la **Etapa 1: Planeación del software** de la materia de Programación Orientada a Objetos.

---

## Modalidad de trabajo

El presente proyecto fue desarrollado de manera **individual**, cumpliendo con las indicaciones establecidas para la actividad académica.

---

## Objetivo general

Desarrollar el sistema **Comercio al Paso** utilizando el lenguaje de programación **Go**, con la finalidad de gestionar productos, usuarios, carritos de compra, pedidos y pagos dentro de una plataforma de comercio electrónico sencilla, aplicando estructuras, métodos, interfaces y funciones auxiliares para organizar correctamente la lógica del sistema.

---

## Alcance del proyecto

El sistema **Comercio al Paso** comprende el desarrollo de una aplicación de consola que permite simular el proceso básico de compra dentro de una tienda virtual.

El sistema permite:

- Registrar información básica del cliente.
- Mostrar una lista de productos disponibles.
- Seleccionar productos mediante su identificador.
- Ingresar la cantidad de productos a comprar.
- Validar la disponibilidad de stock.
- Agregar productos al carrito de compras.
- Calcular el total de la compra.
- Generar un pedido con estado inicial pendiente de pago.
- Seleccionar un método de pago.
- Simular el pago del pedido.
- Actualizar el estado del pedido a pagado.

El sistema no contempla conexión a base de datos, interfaz gráfica, autenticación de usuarios ni pagos reales, debido a que su finalidad es académica y se enfoca en la estructura funcional del sistema.

---

## Paquetes utilizados

Para el desarrollo del sistema no se utilizaron paquetes externos o de terceros. Únicamente se empleó el paquete estándar de Go:

```go
import "fmt"
