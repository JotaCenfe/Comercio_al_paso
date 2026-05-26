package main

import "fmt"

// =====================================================
// SISTEMA: COMERCIO AL PASO
// Sistema de gestión de e-commerce desarrollado en Go
// =====================================================

// =======================
// STRUCT PRODUCTO
// =======================

type Producto struct {
	ID          int
	Nombre      string
	Descripcion string
	Precio      float64
	Stock       int
}

func (p Producto) MostrarInformacion() {
	fmt.Printf("ID: %d | Producto: %s | Descripción: %s | Precio: $%.2f | Stock: %d\n",
		p.ID, p.Nombre, p.Descripcion, p.Precio, p.Stock)
}

func (p *Producto) ReducirStock(cantidad int) bool {
	if p.Stock >= cantidad {
		p.Stock -= cantidad
		return true
	}
	return false
}

// =======================
// STRUCT USUARIO
// =======================

type Usuario struct {
	ID       int
	Nombre   string
	Correo   string
	Telefono string
}

func (u Usuario) MostrarDatos() {
	fmt.Printf("Cliente: %s | Correo: %s | Teléfono: %s\n",
		u.Nombre, u.Correo, u.Telefono)
}

// =======================
// STRUCT CARRITO
// =======================

type ItemCarrito struct {
	Producto Producto
	Cantidad int
}

type Carrito struct {
	Items []ItemCarrito
}

func (c *Carrito) AgregarProducto(producto Producto, cantidad int) {
	item := ItemCarrito{
		Producto: producto,
		Cantidad: cantidad,
	}

	c.Items = append(c.Items, item)
}

func (c Carrito) CalcularTotal() float64 {
	total := 0.0

	for _, item := range c.Items {
		total += item.Producto.Precio * float64(item.Cantidad)
	}

	return total
}

func (c Carrito) MostrarCarrito() {
	fmt.Println("\nProductos agregados al carrito:")

	if len(c.Items) == 0 {
		fmt.Println("El carrito está vacío.")
		return
	}

	for _, item := range c.Items {
		subtotal := item.Producto.Precio * float64(item.Cantidad)

		fmt.Printf("- %s | Cantidad: %d | Subtotal: $%.2f\n",
			item.Producto.Nombre, item.Cantidad, subtotal)
	}
}

// =======================
// STRUCT PEDIDO
// =======================

type Pedido struct {
	ID      int
	Usuario Usuario
	Carrito Carrito
	Estado  string
}

func (p Pedido) MostrarResumen() {
	fmt.Println("\n========== RESUMEN DEL PEDIDO ==========")
	fmt.Println("Sistema: Comercio al Paso")
	fmt.Printf("Pedido Nro: %d\n", p.ID)
	p.Usuario.MostrarDatos()
	p.Carrito.MostrarCarrito()
	fmt.Printf("Total a pagar: $%.2f\n", p.Carrito.CalcularTotal())
	fmt.Printf("Estado del pedido: %s\n", p.Estado)
	fmt.Println("========================================")
}

// =======================
// INTERFAZ DE PAGO
// =======================

type MetodoPago interface {
	ProcesarPago(monto float64) bool
}

// =======================
// PAGO CON TARJETA
// =======================

type PagoTarjeta struct {
	NumeroTarjeta string
	Titular       string
}

func (pt PagoTarjeta) ProcesarPago(monto float64) bool {
	fmt.Printf("\nProcesando pago con tarjeta por un monto de $%.2f...\n", monto)
	fmt.Println("Pago aprobado correctamente.")
	return true
}

// =======================
// PAGO POR TRANSFERENCIA
// =======================

type PagoTransferencia struct {
	Banco  string
	Cuenta string
}

func (pt PagoTransferencia) ProcesarPago(monto float64) bool {
	fmt.Printf("\nProcesando pago por transferencia por un monto de $%.2f...\n", monto)
	fmt.Println("Transferencia confirmada correctamente.")
	return true
}

// =======================
// FUNCIONES AUXILIARES
// =======================

func mostrarProductos(productos []Producto) {
	fmt.Println("\n========== PRODUCTOS DISPONIBLES ==========")
	for _, producto := range productos {
		producto.MostrarInformacion()
	}
	fmt.Println("===========================================")
}

func buscarProductoPorID(productos []Producto, id int) (int, bool) {
	for i, producto := range productos {
		if producto.ID == id {
			return i, true
		}
	}
	return -1, false
}

// =======================
// FUNCIÓN PRINCIPAL
// =======================

func main() {
	fmt.Println("=======================================")
	fmt.Println("        SISTEMA COMERCIO AL PASO       ")
	fmt.Println("=======================================")
	fmt.Println("Sistema de gestión de e-commerce rápido y sencillo")

	// Registro de usuario
	var nombre string
	var correo string
	var telefono string

	fmt.Println("\n========== REGISTRO DE CLIENTE ==========")
	fmt.Print("Ingrese su nombre: ")
	fmt.Scanln(&nombre)

	fmt.Print("Ingrese su correo: ")
	fmt.Scanln(&correo)

	fmt.Print("Ingrese su teléfono: ")
	fmt.Scanln(&telefono)

	usuario := Usuario{
		ID:       1,
		Nombre:   nombre,
		Correo:   correo,
		Telefono: telefono,
	}

	// Lista de productos disponibles
	productos := []Producto{
		{
			ID:          1,
			Nombre:      "Laptop Lenovo",
			Descripcion: "Laptop para trabajo y estudio",
			Precio:      850.00,
			Stock:       10,
		},
		{
			ID:          2,
			Nombre:      "Mouse inalámbrico",
			Descripcion: "Mouse ergonómico USB",
			Precio:      25.50,
			Stock:       20,
		},
		{
			ID:          3,
			Nombre:      "Teclado mecánico",
			Descripcion: "Teclado mecánico retroiluminado",
			Precio:      60.00,
			Stock:       15,
		},
		{
			ID:          4,
			Nombre:      "Audífonos Bluetooth",
			Descripcion: "Audífonos inalámbricos recargables",
			Precio:      35.00,
			Stock:       12,
		},
		{
			ID:          5,
			Nombre:      "Memoria USB 64GB",
			Descripcion: "Unidad de almacenamiento portátil",
			Precio:      12.00,
			Stock:       30,
		},
	}

	carrito := Carrito{}
	opcion := 0

	for opcion != 3 {
		fmt.Println("\n========== MENÚ PRINCIPAL ==========")
		fmt.Println("1. Ver productos disponibles")
		fmt.Println("2. Agregar producto al carrito")
		fmt.Println("3. Finalizar compra")
		fmt.Print("Seleccione una opción: ")
		fmt.Scanln(&opcion)

		switch opcion {
		case 1:
			mostrarProductos(productos)

		case 2:
			var idProducto int
			var cantidad int

			mostrarProductos(productos)

			fmt.Print("\nIngrese el ID del producto que desea comprar: ")
			fmt.Scanln(&idProducto)

			indice, encontrado := buscarProductoPorID(productos, idProducto)

			if !encontrado {
				fmt.Println("Producto no encontrado. Intente nuevamente.")
				continue
			}

			fmt.Print("Ingrese la cantidad que desea comprar: ")
			fmt.Scanln(&cantidad)

			if cantidad <= 0 {
				fmt.Println("La cantidad debe ser mayor a cero.")
				continue
			}

			if productos[indice].ReducirStock(cantidad) {
				carrito.AgregarProducto(productos[indice], cantidad)
				fmt.Println("Producto agregado correctamente al carrito.")
			} else {
				fmt.Println("Stock insuficiente para realizar la compra.")
			}

		case 3:
			fmt.Println("\nFinalizando compra...")

		default:
			fmt.Println("Opción no válida. Intente nuevamente.")
		}
	}

	if len(carrito.Items) == 0 {
		fmt.Println("\nNo se agregaron productos al carrito. Compra cancelada.")
		return
	}

	pedido := Pedido{
		ID:      1001,
		Usuario: usuario,
		Carrito: carrito,
		Estado:  "Pendiente de pago",
	}

	pedido.MostrarResumen()

	// Selección del método de pago
	var opcionPago int
	var metodoPago MetodoPago

	fmt.Println("\n========== MÉTODO DE PAGO ==========")
	fmt.Println("1. Pago con tarjeta")
	fmt.Println("2. Pago por transferencia")
	fmt.Print("Seleccione una opción de pago: ")
	fmt.Scanln(&opcionPago)

	switch opcionPago {
	case 1:
		metodoPago = PagoTarjeta{
			NumeroTarjeta: "1234-5678-9012-3456",
			Titular:       usuario.Nombre,
		}

	case 2:
		metodoPago = PagoTransferencia{
			Banco:  "Banco Pichincha",
			Cuenta: "0000000001",
		}

	default:
		fmt.Println("Método de pago no válido. Se seleccionará pago con tarjeta por defecto.")
		metodoPago = PagoTarjeta{
			NumeroTarjeta: "1234-5678-9012-3456",
			Titular:       usuario.Nombre,
		}
	}

	pagoExitoso := metodoPago.ProcesarPago(pedido.Carrito.CalcularTotal())

	if pagoExitoso {
		pedido.Estado = "Pagado"
		fmt.Println("\nEl pedido ha sido pagado correctamente en Comercio al Paso.")
	} else {
		pedido.Estado = "Pago rechazado"
		fmt.Println("\nNo se pudo completar el pago.")
	}

	pedido.MostrarResumen()
}
