package jsonModel

// Definir las estructuras del JSON
type Producto struct {
	IDProducto      string `json:"id_producto" validate:"required"`
	NombreProducto  string `json:"nombre_producto" validate:"required"`
	Valor           int    `json:"Valor" validate:"required,gt=0"`
}

type Carrito struct {
	IDCarrito string     `json:"id_carrito" validate:"required"`
	Productos []Producto `json:"productos" validate:"required,dive"`
}

type Cliente struct {
	IDTransaction  int     `json:"id_transaction" validate:"required"`
	NombreCliente  string  `json:"nombre_cliente" validate:"required"`
	IDCliente      string  `json:"id_cliente" validate:"required,len=10"`
	CorreoCliente  string  `json:"correo_cliente" validate:"required,email"`
	CarritoCliente Carrito `json:"carrito_cliente" validate:"required"`
}