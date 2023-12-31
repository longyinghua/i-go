// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package query

import (
	"context"
	"database/sql"

	"gorm.io/gorm"

	"gorm.io/gen"

	"gorm.io/plugin/dbresolver"
)

var (
	Q                                 = new(Query)
	Book                              *book
	DoctorScheduleRuleInfo            *doctorScheduleRuleInfo
	HospitalDeptInfo                  *hospitalDeptInfo
	MedicalExaminationAppointmentInfo *medicalExaminationAppointmentInfo
	MultipointPracticeDoctorInfo      *multipointPracticeDoctorInfo
	SqlTest                           *sqlTest
	Student                           *student
	UserAppointmentInfo               *userAppointmentInfo
)

func SetDefault(db *gorm.DB, opts ...gen.DOOption) {
	*Q = *Use(db, opts...)
	Book = &Q.Book
	DoctorScheduleRuleInfo = &Q.DoctorScheduleRuleInfo
	HospitalDeptInfo = &Q.HospitalDeptInfo
	MedicalExaminationAppointmentInfo = &Q.MedicalExaminationAppointmentInfo
	MultipointPracticeDoctorInfo = &Q.MultipointPracticeDoctorInfo
	SqlTest = &Q.SqlTest
	Student = &Q.Student
	UserAppointmentInfo = &Q.UserAppointmentInfo
}

func Use(db *gorm.DB, opts ...gen.DOOption) *Query {
	return &Query{
		db:                                db,
		Book:                              newBook(db, opts...),
		DoctorScheduleRuleInfo:            newDoctorScheduleRuleInfo(db, opts...),
		HospitalDeptInfo:                  newHospitalDeptInfo(db, opts...),
		MedicalExaminationAppointmentInfo: newMedicalExaminationAppointmentInfo(db, opts...),
		MultipointPracticeDoctorInfo:      newMultipointPracticeDoctorInfo(db, opts...),
		SqlTest:                           newSqlTest(db, opts...),
		Student:                           newStudent(db, opts...),
		UserAppointmentInfo:               newUserAppointmentInfo(db, opts...),
	}
}

type Query struct {
	db *gorm.DB

	Book                              book
	DoctorScheduleRuleInfo            doctorScheduleRuleInfo
	HospitalDeptInfo                  hospitalDeptInfo
	MedicalExaminationAppointmentInfo medicalExaminationAppointmentInfo
	MultipointPracticeDoctorInfo      multipointPracticeDoctorInfo
	SqlTest                           sqlTest
	Student                           student
	UserAppointmentInfo               userAppointmentInfo
}

func (q *Query) Available() bool { return q.db != nil }

func (q *Query) clone(db *gorm.DB) *Query {
	return &Query{
		db:                                db,
		Book:                              q.Book.clone(db),
		DoctorScheduleRuleInfo:            q.DoctorScheduleRuleInfo.clone(db),
		HospitalDeptInfo:                  q.HospitalDeptInfo.clone(db),
		MedicalExaminationAppointmentInfo: q.MedicalExaminationAppointmentInfo.clone(db),
		MultipointPracticeDoctorInfo:      q.MultipointPracticeDoctorInfo.clone(db),
		SqlTest:                           q.SqlTest.clone(db),
		Student:                           q.Student.clone(db),
		UserAppointmentInfo:               q.UserAppointmentInfo.clone(db),
	}
}

func (q *Query) ReadDB() *Query {
	return q.ReplaceDB(q.db.Clauses(dbresolver.Read))
}

func (q *Query) WriteDB() *Query {
	return q.ReplaceDB(q.db.Clauses(dbresolver.Write))
}

func (q *Query) ReplaceDB(db *gorm.DB) *Query {
	return &Query{
		db:                                db,
		Book:                              q.Book.replaceDB(db),
		DoctorScheduleRuleInfo:            q.DoctorScheduleRuleInfo.replaceDB(db),
		HospitalDeptInfo:                  q.HospitalDeptInfo.replaceDB(db),
		MedicalExaminationAppointmentInfo: q.MedicalExaminationAppointmentInfo.replaceDB(db),
		MultipointPracticeDoctorInfo:      q.MultipointPracticeDoctorInfo.replaceDB(db),
		SqlTest:                           q.SqlTest.replaceDB(db),
		Student:                           q.Student.replaceDB(db),
		UserAppointmentInfo:               q.UserAppointmentInfo.replaceDB(db),
	}
}

type queryCtx struct {
	Book                              IBookDo
	DoctorScheduleRuleInfo            IDoctorScheduleRuleInfoDo
	HospitalDeptInfo                  IHospitalDeptInfoDo
	MedicalExaminationAppointmentInfo IMedicalExaminationAppointmentInfoDo
	MultipointPracticeDoctorInfo      IMultipointPracticeDoctorInfoDo
	SqlTest                           ISqlTestDo
	Student                           IStudentDo
	UserAppointmentInfo               IUserAppointmentInfoDo
}

func (q *Query) WithContext(ctx context.Context) *queryCtx {
	return &queryCtx{
		Book:                              q.Book.WithContext(ctx),
		DoctorScheduleRuleInfo:            q.DoctorScheduleRuleInfo.WithContext(ctx),
		HospitalDeptInfo:                  q.HospitalDeptInfo.WithContext(ctx),
		MedicalExaminationAppointmentInfo: q.MedicalExaminationAppointmentInfo.WithContext(ctx),
		MultipointPracticeDoctorInfo:      q.MultipointPracticeDoctorInfo.WithContext(ctx),
		SqlTest:                           q.SqlTest.WithContext(ctx),
		Student:                           q.Student.WithContext(ctx),
		UserAppointmentInfo:               q.UserAppointmentInfo.WithContext(ctx),
	}
}

func (q *Query) Transaction(fc func(tx *Query) error, opts ...*sql.TxOptions) error {
	return q.db.Transaction(func(tx *gorm.DB) error { return fc(q.clone(tx)) }, opts...)
}

func (q *Query) Begin(opts ...*sql.TxOptions) *QueryTx {
	tx := q.db.Begin(opts...)
	return &QueryTx{Query: q.clone(tx), Error: tx.Error}
}

type QueryTx struct {
	*Query
	Error error
}

func (q *QueryTx) Commit() error {
	return q.db.Commit().Error
}

func (q *QueryTx) Rollback() error {
	return q.db.Rollback().Error
}

func (q *QueryTx) SavePoint(name string) error {
	return q.db.SavePoint(name).Error
}

func (q *QueryTx) RollbackTo(name string) error {
	return q.db.RollbackTo(name).Error
}
