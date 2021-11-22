package Model

// Orders 对应表结构
//type Orders struct {
//	ID           int
//	TalId        int
//	HashPhone    string
//	OrderId      string
//	OrderStatus  int
//	IsDelete     int
//	CreateUserId int
//	UpdateUserId int
//	CreatedAt    time.Time `gorm:"column:create_time"` // 这句要重写
//	UpdatedAt    time.Time `gorm:"column:update_time"` // 这句要重写
//}
//
//func (t Orders) initDb() *gorm.DB {
//	connection := "db_hotspot" //对应ini的section名称
//	db, err := model.GetDb(connection)
//	if err != nil {
//		fmt.Println(err)
//	}
//	return db
//}
//
////以上为固定结构
////以下为查询实现
//
//func (t Orders) GetById(id int) *Orders {
//	db := t.initDb()
//	db.Find(&t, id)
//	return &t
//}
