package catgo

import "gorm.io/gorm"

type Model struct {
	db    *gorm.DB
	child interface{}
}

func (c *Model) Init(child interface{}) *Model {
	c.child = child
	return c
}

func (c *Model) InitDb(params ...interface{}) *Model {
	c.db = Gorm()
	return c
}

func (c *Model) Find(params ...interface{}) *Model {

	c.db.Debug().Find(c.child)
	Dump(c.child)
	return c
}
