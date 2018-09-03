package builder

type BuildProcess interface {
	SetWheels() BuildProcess
	SetSeats() BuildProcess
	SetStructure() BuildProcess
	GetVehicle() VehicleProduct
}

type VehicleProduct struct {
	Wheels    int
	Seats     int
	Structure string
}

type ManufacturingDirector struct {
	Builder BuildProcess
}

func (f *ManufacturingDirector) Construct() {
	f.Builder.SetWheels().SetSeats().SetStructure()
}

func (f *ManufacturingDirector) SetBuilder(b BuildProcess) {
	f.Builder = b
}

type CarBuilder struct {
	Product VehicleProduct
}

func (c *CarBuilder) SetWheels() BuildProcess {
	c.Product.Wheels = 4
	return c
}

func (c *CarBuilder) SetSeats() BuildProcess {
	c.Product.Seats = 5
	return c
}

func (c *CarBuilder) SetStructure() BuildProcess {
	c.Product.Structure = "Car"
	return c
}

func (c *CarBuilder) GetVehicle() VehicleProduct {
	return c.Product
}

type BikeBuilder struct {
	Product VehicleProduct
}

func (b *BikeBuilder) SetWheels() BuildProcess {
	b.Product.Wheels = 2
	return b
}

func (b *BikeBuilder) SetSeats() BuildProcess {
	b.Product.Seats = 1
	return b
}

func (b *BikeBuilder) SetStructure() BuildProcess {
	b.Product.Structure = "Motorbike"
	return b
}

func (b *BikeBuilder) GetVehicle() VehicleProduct {
	return b.Product
}
