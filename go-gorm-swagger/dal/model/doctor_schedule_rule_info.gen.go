// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"
)

const TableNameDoctorScheduleRuleInfo = "doctor_schedule_rule_info"

// DoctorScheduleRuleInfo 医生排班规则
type DoctorScheduleRuleInfo struct {
	ID                  int64     `gorm:"column:id;primaryKey;autoIncrement:true;comment:id" json:"id"`                                                                             // id
	BaseDrID            int64     `gorm:"column:base_dr_id;comment:基础医生id" json:"base_dr_id"`                                                                                       // 基础医生id
	DrID                int64     `gorm:"column:dr_id;comment:医生id" json:"dr_id"`                                                                                                   // 医生id
	HosID               int64     `gorm:"column:hos_id;comment:医院id" json:"hos_id"`                                                                                                 // 医院id
	DiagType            int32     `gorm:"column:diag_type;comment:0未设置 1特需门诊 2专家门诊 3专病专科门诊 4普通门诊 5高级专家门诊 6特约门诊 7疑难门诊 8名老中医 9膏方门诊 15特色门诊 703571精品门诊 5270普外科换药便民门诊" json:"diag_type"` // 0未设置 1特需门诊 2专家门诊 3专病专科门诊 4普通门诊 5高级专家门诊 6特约门诊 7疑难门诊 8名老中医 9膏方门诊 15特色门诊 703571精品门诊 5270普外科换药便民门诊
	DiagPeriod          int32     `gorm:"column:diag_period;comment:门诊周期（1每周/2每月单周/3每月双周/4每月第一周/5每月第二周/6每月第三周/7每月第四周/8全年单周/9全年双周）" json:"diag_period"`                              // 门诊周期（1每周/2每月单周/3每月双周/4每月第一周/5每月第二周/6每月第三周/7每月第四周/8全年单周/9全年双周）
	WeekDay             int32     `gorm:"column:week_day;comment:星期x（1-7）" json:"week_day"`                                                                                         // 星期x（1-7）
	SchBeginTime        time.Time `gorm:"column:sch_begin_time;comment:排班开始时间" json:"sch_begin_time"`                                                                               // 排班开始时间
	SchEndTime          time.Time `gorm:"column:sch_end_time;comment:排班结束时间" json:"sch_end_time"`                                                                                   // 排班结束时间
	OnlyFirst           []uint8   `gorm:"column:only_first;comment:只能初诊" json:"only_first"`                                                                                         // 只能初诊
	OnlySecond          []uint8   `gorm:"column:only_second;comment:只能复诊" json:"only_second"`                                                                                       // 只能复诊
	GetResID            int64     `gorm:"column:get_res_id;comment:取号说明id" json:"get_res_id"`                                                                                       // 取号说明id
	FirstPrice          int32     `gorm:"column:first_price;comment:初诊价格(分)" json:"first_price"`                                                                                    // 初诊价格(分)
	SecPrice            int32     `gorm:"column:sec_price;comment:复诊价格(分)" json:"sec_price"`                                                                                        // 复诊价格(分)
	IsDetailTime        []uint8   `gorm:"column:is_detail_time;comment:是否选时排班" json:"is_detail_time"`                                                                               // 是否选时排班
	DetailTimeJSON      string    `gorm:"column:detail_time_json;comment:选时json[]" json:"detail_time_json"`                                                                         // 选时json[]
	SchType             int32     `gorm:"column:sch_type;comment:排班类别（0夜晚，1上午，2下午，3白天）" json:"sch_type"`                                                                            // 排班类别（0夜晚，1上午，2下午，3白天）
	AvailableAppointNum int32     `gorm:"column:available_appoint_num;comment:号源数量" json:"available_appoint_num"`                                                                   // 号源数量
	SchDayPeriod        int32     `gorm:"column:sch_day_period;comment:排班生成周期(天)" json:"sch_day_period"`                                                                            // 排班生成周期(天)
	DrName              string    `gorm:"column:dr_name;not null;comment:医生姓名(冗)" json:"dr_name"`                                                                                   // 医生姓名(冗)
	HosName             string    `gorm:"column:hos_name;comment:医院名称(冗)" json:"hos_name"`                                                                                          // 医院名称(冗)
	DeptID              int64     `gorm:"column:dept_id;comment:（二级）科室ID" json:"dept_id"`                                                                                           // （二级）科室ID
	DeptName            string    `gorm:"column:dept_name;comment:（二级）科室名称" json:"dept_name"`                                                                                       // （二级）科室名称
	CreateAt            time.Time `gorm:"column:create_at;default:CURRENT_TIMESTAMP;comment:创建时间" json:"create_at"`                                                                 // 创建时间
	UpdateAt            time.Time `gorm:"column:update_at;default:CURRENT_TIMESTAMP;comment:修改时间" json:"update_at"`                                                                 // 修改时间
	CreateUserID        int64     `gorm:"column:create_user_id;comment:创建人id" json:"create_user_id"`                                                                                // 创建人id
	UpdateUserID        int64     `gorm:"column:update_user_id;comment:修改人id" json:"update_user_id"`                                                                                // 修改人id
	OrderNum            int32     `gorm:"column:order_num;default:1;comment:序号" json:"order_num"`                                                                                   // 序号
	Version             int32     `gorm:"column:version;default:1;comment:乐观锁号" json:"version"`                                                                                     // 乐观锁号
	AppKey              string    `gorm:"column:app_key;default:edc;comment:应用key" json:"app_key"`                                                                                  // 应用key
	DeleteFlag          []uint8   `gorm:"column:delete_flag;default:b'0;comment:逻辑删除" json:"delete_flag"`                                                                           // 逻辑删除
	UUID                string    `gorm:"column:uuid;comment:uuid" json:"uuid"`                                                                                                     // uuid
	Status              int32     `gorm:"column:status;default:1;comment:业务状态，草稿，正常，禁用" json:"status"`                                                                              // 业务状态，草稿，正常，禁用
	DraftData           string    `gorm:"column:draft_data;comment:草稿" json:"draft_data"`                                                                                           // 草稿
	ProjectID           int64     `gorm:"column:project_id;comment:项目id" json:"project_id"`                                                                                         // 项目id
	Tags                string    `gorm:"column:tags;comment:标签" json:"tags"`                                                                                                       // 标签
	ZyID                string    `gorm:"column:zy_id;comment:助医网id" json:"zy_id"`                                                                                                  // 助医网id
}

// TableName DoctorScheduleRuleInfo's table name
func (*DoctorScheduleRuleInfo) TableName() string {
	return TableNameDoctorScheduleRuleInfo
}