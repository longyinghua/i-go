// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package query

import (
	"context"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/schema"

	"gorm.io/gen"
	"gorm.io/gen/field"

	"gorm.io/plugin/dbresolver"

	"go-gorm-swagger-zap/dal/model"
)

func newUserAppointmentInfo(db *gorm.DB, opts ...gen.DOOption) userAppointmentInfo {
	_userAppointmentInfo := userAppointmentInfo{}

	_userAppointmentInfo.userAppointmentInfoDo.UseDB(db, opts...)
	_userAppointmentInfo.userAppointmentInfoDo.UseModel(&model.UserAppointmentInfo{})

	tableName := _userAppointmentInfo.userAppointmentInfoDo.TableName()
	_userAppointmentInfo.ALL = field.NewAsterisk(tableName)
	_userAppointmentInfo.ID = field.NewInt64(tableName, "id")
	_userAppointmentInfo.VisitID = field.NewInt64(tableName, "visit_id")
	_userAppointmentInfo.DrID = field.NewInt64(tableName, "dr_id")
	_userAppointmentInfo.DrSchID = field.NewInt64(tableName, "dr_sch_id")
	_userAppointmentInfo.HosID = field.NewInt64(tableName, "hos_id")
	_userAppointmentInfo.PriDeptID = field.NewInt64(tableName, "pri_dept_id")
	_userAppointmentInfo.SecDeptID = field.NewInt64(tableName, "sec_dept_id")
	_userAppointmentInfo.SchType = field.NewInt32(tableName, "sch_type")
	_userAppointmentInfo.SchDate = field.NewTime(tableName, "sch_date")
	_userAppointmentInfo.StartTime = field.NewString(tableName, "start_time")
	_userAppointmentInfo.EndTime = field.NewString(tableName, "end_time")
	_userAppointmentInfo.BookingNo = field.NewString(tableName, "booking_no")
	_userAppointmentInfo.BookingSeq = field.NewString(tableName, "booking_seq")
	_userAppointmentInfo.PayPrice = field.NewInt32(tableName, "pay_price")
	_userAppointmentInfo.AppointState = field.NewInt32(tableName, "appoint_state")
	_userAppointmentInfo.IsNotice = field.NewField(tableName, "is_notice")
	_userAppointmentInfo.NoticeState = field.NewInt32(tableName, "notice_state")
	_userAppointmentInfo.PayStatus = field.NewString(tableName, "pay_status")
	_userAppointmentInfo.PayTime = field.NewTime(tableName, "pay_time")
	_userAppointmentInfo.AdminRemark = field.NewString(tableName, "admin_remark")
	_userAppointmentInfo.SourceID = field.NewInt64(tableName, "source_id")
	_userAppointmentInfo.MedicalCardType = field.NewString(tableName, "medical_card_type")
	_userAppointmentInfo.MedicalCardNo = field.NewString(tableName, "medical_card_no")
	_userAppointmentInfo.PayType = field.NewString(tableName, "pay_type")
	_userAppointmentInfo.ChannelID = field.NewInt64(tableName, "channel_id")
	_userAppointmentInfo.UserID = field.NewInt64(tableName, "user_id")
	_userAppointmentInfo.OpenID = field.NewString(tableName, "open_id")
	_userAppointmentInfo.IsHandle = field.NewField(tableName, "is_handle")
	_userAppointmentInfo.IsConfirmed = field.NewField(tableName, "is_confirmed")
	_userAppointmentInfo.CancelType = field.NewString(tableName, "cancel_type")
	_userAppointmentInfo.CancelReason = field.NewString(tableName, "cancel_reason")
	_userAppointmentInfo.AppointCode = field.NewString(tableName, "appoint_code")
	_userAppointmentInfo.IfScheNo = field.NewString(tableName, "if_sche_no")
	_userAppointmentInfo.VisitType = field.NewInt32(tableName, "visit_type")
	_userAppointmentInfo.IfHospitalData = field.NewString(tableName, "if_hospital_data")
	_userAppointmentInfo.VisitSnapshot = field.NewString(tableName, "visit_snapshot")
	_userAppointmentInfo.DrSchInfoSnapshot = field.NewString(tableName, "dr_sch_info_snapshot")
	_userAppointmentInfo.IsIfHospital = field.NewField(tableName, "is_if_hospital")
	_userAppointmentInfo.IDCardNum = field.NewString(tableName, "id_card_num")
	_userAppointmentInfo.VisitName = field.NewString(tableName, "visit_name")
	_userAppointmentInfo.VisitMobile = field.NewString(tableName, "visit_mobile")
	_userAppointmentInfo.HospitalName = field.NewString(tableName, "hospital_name")
	_userAppointmentInfo.DeptName = field.NewString(tableName, "dept_name")
	_userAppointmentInfo.DoctorName = field.NewString(tableName, "doctor_name")
	_userAppointmentInfo.IsRescheduled = field.NewField(tableName, "is_rescheduled")
	_userAppointmentInfo.ApptExtendsJSON = field.NewString(tableName, "appt_extends_json")
	_userAppointmentInfo.ChannelCode = field.NewString(tableName, "channel_code")
	_userAppointmentInfo.ChannelOrderNo = field.NewString(tableName, "channel_order_no")
	_userAppointmentInfo.CreateAt = field.NewTime(tableName, "create_at")
	_userAppointmentInfo.UpdateAt = field.NewTime(tableName, "update_at")
	_userAppointmentInfo.CreateUserID = field.NewInt64(tableName, "create_user_id")
	_userAppointmentInfo.UpdateUserID = field.NewInt64(tableName, "update_user_id")
	_userAppointmentInfo.OrderNum = field.NewInt32(tableName, "order_num")
	_userAppointmentInfo.Version = field.NewInt32(tableName, "version")
	_userAppointmentInfo.AppKey = field.NewString(tableName, "app_key")
	_userAppointmentInfo.DeleteFlag = field.NewField(tableName, "delete_flag")
	_userAppointmentInfo.UUID = field.NewString(tableName, "uuid")
	_userAppointmentInfo.Status = field.NewInt32(tableName, "status")
	_userAppointmentInfo.DraftData = field.NewString(tableName, "draft_data")
	_userAppointmentInfo.ProjectID = field.NewInt64(tableName, "project_id")
	_userAppointmentInfo.Tags = field.NewString(tableName, "tags")
	_userAppointmentInfo.ZyID = field.NewString(tableName, "zy_id")

	_userAppointmentInfo.fillFieldMap()

	return _userAppointmentInfo
}

// userAppointmentInfo 用户预约表
type userAppointmentInfo struct {
	userAppointmentInfoDo userAppointmentInfoDo

	ALL               field.Asterisk
	ID                field.Int64  // id
	VisitID           field.Int64  // 就诊人id
	DrID              field.Int64  // 医生id
	DrSchID           field.Int64  // 医生排班id
	HosID             field.Int64  // 医院id
	PriDeptID         field.Int64  // 主科室id
	SecDeptID         field.Int64  // 科室id
	SchType           field.Int32  // 排班类型
	SchDate           field.Time   // 就诊日期
	StartTime         field.String // 开始时间
	EndTime           field.String // 结束时间
	BookingNo         field.String // 接口医院预约编号
	BookingSeq        field.String // 接口医院顺序号
	PayPrice          field.Int32  // 价格(分)
	AppointState      field.Int32  // 接口返回状态 0待返回 1成功收到返回 2补发确认预约成功 3补发确认预约失败
	IsNotice          field.Field  // 是否通知
	NoticeState       field.Int32  // 通知状态
	PayStatus         field.String // 支付状态：已支付-PAYED，未支付-UNPAY，无需支付-NOPAY
	PayTime           field.Time   // 支付时间
	AdminRemark       field.String // 后台手工备注
	SourceID          field.Int64  // 号源id
	MedicalCardType   field.String // 就诊卡类别
	MedicalCardNo     field.String // 就诊卡号
	PayType           field.String // 支付类型
	ChannelID         field.Int64  // 渠道id
	UserID            field.Int64  // 用户id
	OpenID            field.String // 第三方OpenId
	IsHandle          field.Field  // 是否处理（非接口给医院）
	IsConfirmed       field.Field  // 第三方是否已确认此订单
	CancelType        field.String // 取消类型
	CancelReason      field.String // 取消原因
	AppointCode       field.String // 预约码
	IfScheNo          field.String // 接口医院排班编号
	VisitType         field.Int32  // 初复诊（初诊1，复诊2）
	IfHospitalData    field.String // 接口医院相关数据（科室，医生的医院编码）
	VisitSnapshot     field.String // 就诊人快照
	DrSchInfoSnapshot field.String // 医生排班快照
	IsIfHospital      field.Field  // 是否接口医院（冗）
	IDCardNum         field.String // 身份证号（冗）
	VisitName         field.String // 就诊人姓名(冗)
	VisitMobile       field.String // 就诊人手机号(冗)
	HospitalName      field.String // 医院名称(冗)
	DeptName          field.String // 科室名称(冗)
	DoctorName        field.String // 医生姓名(冗)
	IsRescheduled     field.Field  // 是否改期（默认否）
	ApptExtendsJSON   field.String // 预约扩展信息json
	ChannelCode       field.String // 渠道编码
	ChannelOrderNo    field.String // 渠道订单号
	CreateAt          field.Time   // 创建时间
	UpdateAt          field.Time   // 修改时间
	CreateUserID      field.Int64  // 创建人id
	UpdateUserID      field.Int64  // 修改人id
	OrderNum          field.Int32  // 序号
	Version           field.Int32  // 乐观锁号
	AppKey            field.String // 应用key
	DeleteFlag        field.Field  // 逻辑删除
	UUID              field.String // uuid
	Status            field.Int32  // 业务状态，草稿，正常，禁用
	DraftData         field.String // 草稿
	ProjectID         field.Int64  // 项目id
	Tags              field.String // 标签
	ZyID              field.String // 助医网id

	fieldMap map[string]field.Expr
}

func (u userAppointmentInfo) Table(newTableName string) *userAppointmentInfo {
	u.userAppointmentInfoDo.UseTable(newTableName)
	return u.updateTableName(newTableName)
}

func (u userAppointmentInfo) As(alias string) *userAppointmentInfo {
	u.userAppointmentInfoDo.DO = *(u.userAppointmentInfoDo.As(alias).(*gen.DO))
	return u.updateTableName(alias)
}

func (u *userAppointmentInfo) updateTableName(table string) *userAppointmentInfo {
	u.ALL = field.NewAsterisk(table)
	u.ID = field.NewInt64(table, "id")
	u.VisitID = field.NewInt64(table, "visit_id")
	u.DrID = field.NewInt64(table, "dr_id")
	u.DrSchID = field.NewInt64(table, "dr_sch_id")
	u.HosID = field.NewInt64(table, "hos_id")
	u.PriDeptID = field.NewInt64(table, "pri_dept_id")
	u.SecDeptID = field.NewInt64(table, "sec_dept_id")
	u.SchType = field.NewInt32(table, "sch_type")
	u.SchDate = field.NewTime(table, "sch_date")
	u.StartTime = field.NewString(table, "start_time")
	u.EndTime = field.NewString(table, "end_time")
	u.BookingNo = field.NewString(table, "booking_no")
	u.BookingSeq = field.NewString(table, "booking_seq")
	u.PayPrice = field.NewInt32(table, "pay_price")
	u.AppointState = field.NewInt32(table, "appoint_state")
	u.IsNotice = field.NewField(table, "is_notice")
	u.NoticeState = field.NewInt32(table, "notice_state")
	u.PayStatus = field.NewString(table, "pay_status")
	u.PayTime = field.NewTime(table, "pay_time")
	u.AdminRemark = field.NewString(table, "admin_remark")
	u.SourceID = field.NewInt64(table, "source_id")
	u.MedicalCardType = field.NewString(table, "medical_card_type")
	u.MedicalCardNo = field.NewString(table, "medical_card_no")
	u.PayType = field.NewString(table, "pay_type")
	u.ChannelID = field.NewInt64(table, "channel_id")
	u.UserID = field.NewInt64(table, "user_id")
	u.OpenID = field.NewString(table, "open_id")
	u.IsHandle = field.NewField(table, "is_handle")
	u.IsConfirmed = field.NewField(table, "is_confirmed")
	u.CancelType = field.NewString(table, "cancel_type")
	u.CancelReason = field.NewString(table, "cancel_reason")
	u.AppointCode = field.NewString(table, "appoint_code")
	u.IfScheNo = field.NewString(table, "if_sche_no")
	u.VisitType = field.NewInt32(table, "visit_type")
	u.IfHospitalData = field.NewString(table, "if_hospital_data")
	u.VisitSnapshot = field.NewString(table, "visit_snapshot")
	u.DrSchInfoSnapshot = field.NewString(table, "dr_sch_info_snapshot")
	u.IsIfHospital = field.NewField(table, "is_if_hospital")
	u.IDCardNum = field.NewString(table, "id_card_num")
	u.VisitName = field.NewString(table, "visit_name")
	u.VisitMobile = field.NewString(table, "visit_mobile")
	u.HospitalName = field.NewString(table, "hospital_name")
	u.DeptName = field.NewString(table, "dept_name")
	u.DoctorName = field.NewString(table, "doctor_name")
	u.IsRescheduled = field.NewField(table, "is_rescheduled")
	u.ApptExtendsJSON = field.NewString(table, "appt_extends_json")
	u.ChannelCode = field.NewString(table, "channel_code")
	u.ChannelOrderNo = field.NewString(table, "channel_order_no")
	u.CreateAt = field.NewTime(table, "create_at")
	u.UpdateAt = field.NewTime(table, "update_at")
	u.CreateUserID = field.NewInt64(table, "create_user_id")
	u.UpdateUserID = field.NewInt64(table, "update_user_id")
	u.OrderNum = field.NewInt32(table, "order_num")
	u.Version = field.NewInt32(table, "version")
	u.AppKey = field.NewString(table, "app_key")
	u.DeleteFlag = field.NewField(table, "delete_flag")
	u.UUID = field.NewString(table, "uuid")
	u.Status = field.NewInt32(table, "status")
	u.DraftData = field.NewString(table, "draft_data")
	u.ProjectID = field.NewInt64(table, "project_id")
	u.Tags = field.NewString(table, "tags")
	u.ZyID = field.NewString(table, "zy_id")

	u.fillFieldMap()

	return u
}

func (u *userAppointmentInfo) WithContext(ctx context.Context) IUserAppointmentInfoDo {
	return u.userAppointmentInfoDo.WithContext(ctx)
}

func (u userAppointmentInfo) TableName() string { return u.userAppointmentInfoDo.TableName() }

func (u userAppointmentInfo) Alias() string { return u.userAppointmentInfoDo.Alias() }

func (u userAppointmentInfo) Columns(cols ...field.Expr) gen.Columns {
	return u.userAppointmentInfoDo.Columns(cols...)
}

func (u *userAppointmentInfo) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := u.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (u *userAppointmentInfo) fillFieldMap() {
	u.fieldMap = make(map[string]field.Expr, 62)
	u.fieldMap["id"] = u.ID
	u.fieldMap["visit_id"] = u.VisitID
	u.fieldMap["dr_id"] = u.DrID
	u.fieldMap["dr_sch_id"] = u.DrSchID
	u.fieldMap["hos_id"] = u.HosID
	u.fieldMap["pri_dept_id"] = u.PriDeptID
	u.fieldMap["sec_dept_id"] = u.SecDeptID
	u.fieldMap["sch_type"] = u.SchType
	u.fieldMap["sch_date"] = u.SchDate
	u.fieldMap["start_time"] = u.StartTime
	u.fieldMap["end_time"] = u.EndTime
	u.fieldMap["booking_no"] = u.BookingNo
	u.fieldMap["booking_seq"] = u.BookingSeq
	u.fieldMap["pay_price"] = u.PayPrice
	u.fieldMap["appoint_state"] = u.AppointState
	u.fieldMap["is_notice"] = u.IsNotice
	u.fieldMap["notice_state"] = u.NoticeState
	u.fieldMap["pay_status"] = u.PayStatus
	u.fieldMap["pay_time"] = u.PayTime
	u.fieldMap["admin_remark"] = u.AdminRemark
	u.fieldMap["source_id"] = u.SourceID
	u.fieldMap["medical_card_type"] = u.MedicalCardType
	u.fieldMap["medical_card_no"] = u.MedicalCardNo
	u.fieldMap["pay_type"] = u.PayType
	u.fieldMap["channel_id"] = u.ChannelID
	u.fieldMap["user_id"] = u.UserID
	u.fieldMap["open_id"] = u.OpenID
	u.fieldMap["is_handle"] = u.IsHandle
	u.fieldMap["is_confirmed"] = u.IsConfirmed
	u.fieldMap["cancel_type"] = u.CancelType
	u.fieldMap["cancel_reason"] = u.CancelReason
	u.fieldMap["appoint_code"] = u.AppointCode
	u.fieldMap["if_sche_no"] = u.IfScheNo
	u.fieldMap["visit_type"] = u.VisitType
	u.fieldMap["if_hospital_data"] = u.IfHospitalData
	u.fieldMap["visit_snapshot"] = u.VisitSnapshot
	u.fieldMap["dr_sch_info_snapshot"] = u.DrSchInfoSnapshot
	u.fieldMap["is_if_hospital"] = u.IsIfHospital
	u.fieldMap["id_card_num"] = u.IDCardNum
	u.fieldMap["visit_name"] = u.VisitName
	u.fieldMap["visit_mobile"] = u.VisitMobile
	u.fieldMap["hospital_name"] = u.HospitalName
	u.fieldMap["dept_name"] = u.DeptName
	u.fieldMap["doctor_name"] = u.DoctorName
	u.fieldMap["is_rescheduled"] = u.IsRescheduled
	u.fieldMap["appt_extends_json"] = u.ApptExtendsJSON
	u.fieldMap["channel_code"] = u.ChannelCode
	u.fieldMap["channel_order_no"] = u.ChannelOrderNo
	u.fieldMap["create_at"] = u.CreateAt
	u.fieldMap["update_at"] = u.UpdateAt
	u.fieldMap["create_user_id"] = u.CreateUserID
	u.fieldMap["update_user_id"] = u.UpdateUserID
	u.fieldMap["order_num"] = u.OrderNum
	u.fieldMap["version"] = u.Version
	u.fieldMap["app_key"] = u.AppKey
	u.fieldMap["delete_flag"] = u.DeleteFlag
	u.fieldMap["uuid"] = u.UUID
	u.fieldMap["status"] = u.Status
	u.fieldMap["draft_data"] = u.DraftData
	u.fieldMap["project_id"] = u.ProjectID
	u.fieldMap["tags"] = u.Tags
	u.fieldMap["zy_id"] = u.ZyID
}

func (u userAppointmentInfo) clone(db *gorm.DB) userAppointmentInfo {
	u.userAppointmentInfoDo.ReplaceConnPool(db.Statement.ConnPool)
	return u
}

func (u userAppointmentInfo) replaceDB(db *gorm.DB) userAppointmentInfo {
	u.userAppointmentInfoDo.ReplaceDB(db)
	return u
}

type userAppointmentInfoDo struct{ gen.DO }

type IUserAppointmentInfoDo interface {
	gen.SubQuery
	Debug() IUserAppointmentInfoDo
	WithContext(ctx context.Context) IUserAppointmentInfoDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IUserAppointmentInfoDo
	WriteDB() IUserAppointmentInfoDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IUserAppointmentInfoDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IUserAppointmentInfoDo
	Not(conds ...gen.Condition) IUserAppointmentInfoDo
	Or(conds ...gen.Condition) IUserAppointmentInfoDo
	Select(conds ...field.Expr) IUserAppointmentInfoDo
	Where(conds ...gen.Condition) IUserAppointmentInfoDo
	Order(conds ...field.Expr) IUserAppointmentInfoDo
	Distinct(cols ...field.Expr) IUserAppointmentInfoDo
	Omit(cols ...field.Expr) IUserAppointmentInfoDo
	Join(table schema.Tabler, on ...field.Expr) IUserAppointmentInfoDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IUserAppointmentInfoDo
	RightJoin(table schema.Tabler, on ...field.Expr) IUserAppointmentInfoDo
	Group(cols ...field.Expr) IUserAppointmentInfoDo
	Having(conds ...gen.Condition) IUserAppointmentInfoDo
	Limit(limit int) IUserAppointmentInfoDo
	Offset(offset int) IUserAppointmentInfoDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IUserAppointmentInfoDo
	Unscoped() IUserAppointmentInfoDo
	Create(values ...*model.UserAppointmentInfo) error
	CreateInBatches(values []*model.UserAppointmentInfo, batchSize int) error
	Save(values ...*model.UserAppointmentInfo) error
	First() (*model.UserAppointmentInfo, error)
	Take() (*model.UserAppointmentInfo, error)
	Last() (*model.UserAppointmentInfo, error)
	Find() ([]*model.UserAppointmentInfo, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.UserAppointmentInfo, err error)
	FindInBatches(result *[]*model.UserAppointmentInfo, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.UserAppointmentInfo) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IUserAppointmentInfoDo
	Assign(attrs ...field.AssignExpr) IUserAppointmentInfoDo
	Joins(fields ...field.RelationField) IUserAppointmentInfoDo
	Preload(fields ...field.RelationField) IUserAppointmentInfoDo
	FirstOrInit() (*model.UserAppointmentInfo, error)
	FirstOrCreate() (*model.UserAppointmentInfo, error)
	FindByPage(offset int, limit int) (result []*model.UserAppointmentInfo, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IUserAppointmentInfoDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (u userAppointmentInfoDo) Debug() IUserAppointmentInfoDo {
	return u.withDO(u.DO.Debug())
}

func (u userAppointmentInfoDo) WithContext(ctx context.Context) IUserAppointmentInfoDo {
	return u.withDO(u.DO.WithContext(ctx))
}

func (u userAppointmentInfoDo) ReadDB() IUserAppointmentInfoDo {
	return u.Clauses(dbresolver.Read)
}

func (u userAppointmentInfoDo) WriteDB() IUserAppointmentInfoDo {
	return u.Clauses(dbresolver.Write)
}

func (u userAppointmentInfoDo) Session(config *gorm.Session) IUserAppointmentInfoDo {
	return u.withDO(u.DO.Session(config))
}

func (u userAppointmentInfoDo) Clauses(conds ...clause.Expression) IUserAppointmentInfoDo {
	return u.withDO(u.DO.Clauses(conds...))
}

func (u userAppointmentInfoDo) Returning(value interface{}, columns ...string) IUserAppointmentInfoDo {
	return u.withDO(u.DO.Returning(value, columns...))
}

func (u userAppointmentInfoDo) Not(conds ...gen.Condition) IUserAppointmentInfoDo {
	return u.withDO(u.DO.Not(conds...))
}

func (u userAppointmentInfoDo) Or(conds ...gen.Condition) IUserAppointmentInfoDo {
	return u.withDO(u.DO.Or(conds...))
}

func (u userAppointmentInfoDo) Select(conds ...field.Expr) IUserAppointmentInfoDo {
	return u.withDO(u.DO.Select(conds...))
}

func (u userAppointmentInfoDo) Where(conds ...gen.Condition) IUserAppointmentInfoDo {
	return u.withDO(u.DO.Where(conds...))
}

func (u userAppointmentInfoDo) Order(conds ...field.Expr) IUserAppointmentInfoDo {
	return u.withDO(u.DO.Order(conds...))
}

func (u userAppointmentInfoDo) Distinct(cols ...field.Expr) IUserAppointmentInfoDo {
	return u.withDO(u.DO.Distinct(cols...))
}

func (u userAppointmentInfoDo) Omit(cols ...field.Expr) IUserAppointmentInfoDo {
	return u.withDO(u.DO.Omit(cols...))
}

func (u userAppointmentInfoDo) Join(table schema.Tabler, on ...field.Expr) IUserAppointmentInfoDo {
	return u.withDO(u.DO.Join(table, on...))
}

func (u userAppointmentInfoDo) LeftJoin(table schema.Tabler, on ...field.Expr) IUserAppointmentInfoDo {
	return u.withDO(u.DO.LeftJoin(table, on...))
}

func (u userAppointmentInfoDo) RightJoin(table schema.Tabler, on ...field.Expr) IUserAppointmentInfoDo {
	return u.withDO(u.DO.RightJoin(table, on...))
}

func (u userAppointmentInfoDo) Group(cols ...field.Expr) IUserAppointmentInfoDo {
	return u.withDO(u.DO.Group(cols...))
}

func (u userAppointmentInfoDo) Having(conds ...gen.Condition) IUserAppointmentInfoDo {
	return u.withDO(u.DO.Having(conds...))
}

func (u userAppointmentInfoDo) Limit(limit int) IUserAppointmentInfoDo {
	return u.withDO(u.DO.Limit(limit))
}

func (u userAppointmentInfoDo) Offset(offset int) IUserAppointmentInfoDo {
	return u.withDO(u.DO.Offset(offset))
}

func (u userAppointmentInfoDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IUserAppointmentInfoDo {
	return u.withDO(u.DO.Scopes(funcs...))
}

func (u userAppointmentInfoDo) Unscoped() IUserAppointmentInfoDo {
	return u.withDO(u.DO.Unscoped())
}

func (u userAppointmentInfoDo) Create(values ...*model.UserAppointmentInfo) error {
	if len(values) == 0 {
		return nil
	}
	return u.DO.Create(values)
}

func (u userAppointmentInfoDo) CreateInBatches(values []*model.UserAppointmentInfo, batchSize int) error {
	return u.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (u userAppointmentInfoDo) Save(values ...*model.UserAppointmentInfo) error {
	if len(values) == 0 {
		return nil
	}
	return u.DO.Save(values)
}

func (u userAppointmentInfoDo) First() (*model.UserAppointmentInfo, error) {
	if result, err := u.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.UserAppointmentInfo), nil
	}
}

func (u userAppointmentInfoDo) Take() (*model.UserAppointmentInfo, error) {
	if result, err := u.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.UserAppointmentInfo), nil
	}
}

func (u userAppointmentInfoDo) Last() (*model.UserAppointmentInfo, error) {
	if result, err := u.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.UserAppointmentInfo), nil
	}
}

func (u userAppointmentInfoDo) Find() ([]*model.UserAppointmentInfo, error) {
	result, err := u.DO.Find()
	return result.([]*model.UserAppointmentInfo), err
}

func (u userAppointmentInfoDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.UserAppointmentInfo, err error) {
	buf := make([]*model.UserAppointmentInfo, 0, batchSize)
	err = u.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (u userAppointmentInfoDo) FindInBatches(result *[]*model.UserAppointmentInfo, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return u.DO.FindInBatches(result, batchSize, fc)
}

func (u userAppointmentInfoDo) Attrs(attrs ...field.AssignExpr) IUserAppointmentInfoDo {
	return u.withDO(u.DO.Attrs(attrs...))
}

func (u userAppointmentInfoDo) Assign(attrs ...field.AssignExpr) IUserAppointmentInfoDo {
	return u.withDO(u.DO.Assign(attrs...))
}

func (u userAppointmentInfoDo) Joins(fields ...field.RelationField) IUserAppointmentInfoDo {
	for _, _f := range fields {
		u = *u.withDO(u.DO.Joins(_f))
	}
	return &u
}

func (u userAppointmentInfoDo) Preload(fields ...field.RelationField) IUserAppointmentInfoDo {
	for _, _f := range fields {
		u = *u.withDO(u.DO.Preload(_f))
	}
	return &u
}

func (u userAppointmentInfoDo) FirstOrInit() (*model.UserAppointmentInfo, error) {
	if result, err := u.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.UserAppointmentInfo), nil
	}
}

func (u userAppointmentInfoDo) FirstOrCreate() (*model.UserAppointmentInfo, error) {
	if result, err := u.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.UserAppointmentInfo), nil
	}
}

func (u userAppointmentInfoDo) FindByPage(offset int, limit int) (result []*model.UserAppointmentInfo, count int64, err error) {
	result, err = u.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = u.Offset(-1).Limit(-1).Count()
	return
}

func (u userAppointmentInfoDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = u.Count()
	if err != nil {
		return
	}

	err = u.Offset(offset).Limit(limit).Scan(result)
	return
}

func (u userAppointmentInfoDo) Scan(result interface{}) (err error) {
	return u.DO.Scan(result)
}

func (u userAppointmentInfoDo) Delete(models ...*model.UserAppointmentInfo) (result gen.ResultInfo, err error) {
	return u.DO.Delete(models)
}

func (u *userAppointmentInfoDo) withDO(do gen.Dao) *userAppointmentInfoDo {
	u.DO = *do.(*gen.DO)
	return u
}
