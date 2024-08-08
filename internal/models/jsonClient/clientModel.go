package clientModel

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

type Cliente2 struct {
	Nombre string `json:"nombre"`
	Apellido string `json:"apellido"`
	Rut string `json:"rut"`
	Telefono string `json:"telefono"`
	Correo string `json:"correo"`
	Direccion string `json:"direccion"`
	Fecha_nacimiento string `json:"fecha_nacimiento"`
	Fecha_creacion string `json:"fecha_creacion"`
	Fecha_actualizacion *string `json:"fecha_actualizacion,omitempty"`
	Instagram string `json:"instagram"`
}
// Definir las estructuras del JSON para el test de /cache
type Cliente3 struct {
	Nombre string `json:"nombre"`
	Apellido string `json:"apellido"`
	Rut string `json:"rut"`
	Telefono string `json:"telefono"`
	Correo string `json:"correo"`
	Direccion string `json:"direccion"`
	Fecha_nacimiento string `json:"fecha_nacimiento"`
	Fecha_creacion string `json:"fecha_creacion"`
	Instagram string `json:"instagram"`
	Carrito Carrito3 `json:"carrito,omitempty"`
}
type Carrito3 struct {
	Id_carrito int `json:"id_carrito"`
	Productos []Producto3 `json:"productos"`
	Promociones []Promocion3 `json:"promociones,omitempty"`
}
type Promocion3 struct {
	Id_promocion string `json:"id_promocion"`
	Off float64 `json:"off"`
	Sku string `json:"sku"`
}
type Producto3 struct {
	Id_producto string `json:"id_producto"`
	Nombre_producto string `json:"nombre_producto"`
	Precio string `json:"precio"`
	Sku string `json:"sku"`
}