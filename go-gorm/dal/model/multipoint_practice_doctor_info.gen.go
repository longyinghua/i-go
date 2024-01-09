// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"
)

const TableNameMultipointPracticeDoctorInfo = "multipoint_practice_doctor_info"

// MultipointPracticeDoctorInfo 多点执业医生信息
type MultipointPracticeDoctorInfo struct {
	ID                   int64     `gorm:"column:id;primaryKey;autoIncrement:true;comment:id" json:"id"`                                                                             // id
	DrCode               string    `gorm:"column:dr_code;comment:医生工号" json:"dr_code"`                                                                                               // 医生工号
	DrName               string    `gorm:"column:dr_name;comment:医生姓名" json:"dr_name"`                                                                                               // 医生姓名
	DrMultiCode          string    `gorm:"column:dr_multi_code;comment:医生当前执业点医生代码" json:"dr_multi_code"`                                                                            // 医生当前执业点医生代码
	DeptMultiCode        string    `gorm:"column:dept_multi_code;comment:医生当前执业点科室代码" json:"dept_multi_code"`                                                                        // 医生当前执业点科室代码
	BaseDrID             int64     `gorm:"column:base_dr_id;comment:基础医生id" json:"base_dr_id"`                                                                                       // 基础医生id
	HosID                int64     `gorm:"column:hos_id;not null;comment:医院id" json:"hos_id"`                                                                                        // 医院id
	PriDeptID            int64     `gorm:"column:pri_dept_id;comment:一级科室id" json:"pri_dept_id"`                                                                                     // 一级科室id
	SecDeptID            int64     `gorm:"column:sec_dept_id;comment:二级科室id" json:"sec_dept_id"`                                                                                     // 二级科室id
	DiagTypes            string    `gorm:"column:diag_types;comment:医生门诊类型数组[], 0未设置 1特需门诊 2专家门诊 3专病专科门诊 4普通门诊 5高级专家门诊 6特约门诊 7疑难门诊 8名老中医 9膏方门诊 15特色门诊 703571精品门诊" json:"diag_types"` // 医生门诊类型数组[], 0未设置 1特需门诊 2专家门诊 3专病专科门诊 4普通门诊 5高级专家门诊 6特约门诊 7疑难门诊 8名老中医 9膏方门诊 15特色门诊 703571精品门诊
	IsMaster             []uint8   `gorm:"column:is_master;comment:是否主执业" json:"is_master"`                                                                                          // 是否主执业
	ScheInfoJSON         string    `gorm:"column:___sche_info_json;comment:排班信息json" json:"___sche_info_json"`                                                                       // 排班信息json
	ApptRuleJSON         string    `gorm:"column:___appt_rule_json;comment:放号规则json" json:"___appt_rule_json"`                                                                       // 放号规则json
	IsVideo              []uint8   `gorm:"column:is_video;comment:视频问诊" json:"is_video"`                                                                                             // 视频问诊
	IsChat               []uint8   `gorm:"column:is_chat;comment:图文聊天问诊" json:"is_chat"`                                                                                             // 图文聊天问诊
	ShowGoodDoctor       []uint8   `gorm:"column:show_good_doctor;comment:是否显示好医生" json:"show_good_doctor"`                                                                          // 是否显示好医生
	CurrentAppointNum    int32     `gorm:"column:current_appoint_num;comment:当月已预约数" json:"current_appoint_num"`                                                                     // 当月已预约数
	CurrentAllAppointNum int32     `gorm:"column:current_all_appoint_num;comment:当月总号源数" json:"current_all_appoint_num"`                                                             // 当月总号源数
	SourceID             int64     `gorm:"column:source_id;comment:号源id" json:"source_id"`                                                                                           // 号源id
	DrJSONInfo           string    `gorm:"column:dr_json_info;comment:医生json详情" json:"dr_json_info"`                                                                                 // 医生json详情
	CreateAt             time.Time `gorm:"column:create_at;default:CURRENT_TIMESTAMP;comment:创建时间" json:"create_at"`                                                                 // 创建时间
	UpdateAt             time.Time `gorm:"column:update_at;default:CURRENT_TIMESTAMP;comment:修改时间" json:"update_at"`                                                                 // 修改时间
	CreateUserID         int64     `gorm:"column:create_user_id;comment:创建人id" json:"create_user_id"`                                                                                // 创建人id
	UpdateUserID         int64     `gorm:"column:update_user_id;comment:修改人id" json:"update_user_id"`                                                                                // 修改人id
	OrderNum             int32     `gorm:"column:order_num;default:1;comment:序号" json:"order_num"`                                                                                   // 序号
	Version              int32     `gorm:"column:version;default:1;comment:乐观锁号" json:"version"`                                                                                     // 乐观锁号
	AppKey               string    `gorm:"column:app_key;default:edc;comment:应用key" json:"app_key"`                                                                                  // 应用key
	DeleteFlag           []uint8   `gorm:"column:delete_flag;default:b'0;comment:逻辑删除" json:"delete_flag"`                                                                           // 逻辑删除
	UUID                 string    `gorm:"column:uuid;comment:uuid" json:"uuid"`                                                                                                     // uuid
	Status               int32     `gorm:"column:status;default:1;comment:业务状态，草稿，正常，禁用" json:"status"`                                                                              // 业务状态，草稿，正常，禁用
	DraftData            string    `gorm:"column:draft_data;comment:草稿" json:"draft_data"`                                                                                           // 草稿
	ProjectID            int64     `gorm:"column:project_id;comment:项目id" json:"project_id"`                                                                                         // 项目id
	Tags                 string    `gorm:"column:tags;comment:标签" json:"tags"`                                                                                                       // 标签
	ZyID                 string    `gorm:"column:zy_id" json:"zy_id"`
}

// TableName MultipointPracticeDoctorInfo's table name
func (*MultipointPracticeDoctorInfo) TableName() string {
	return TableNameMultipointPracticeDoctorInfo
}
