package model

import (
	context "context"
	fmt "fmt"
	gorm1 "github.com/infobloxopen/atlas-app-toolkit/gorm"
	errors "github.com/infobloxopen/protoc-gen-gorm/errors"
	gorm "github.com/jinzhu/gorm"
	field_mask "google.golang.org/genproto/protobuf/field_mask"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	strings "strings"
	time "time"
)

type AddressORM struct {
	Address   string
	City      string
	Country   string
	CreatedAt *time.Time
	Id        int32 `gorm:"type:integer;primary_key"`
	State     string
	Type      string
	UpdatedAt *time.Time
	User      *UserORM `gorm:"foreignkey:UserId;association_foreignkey:Id"`
	UserId    *string
	Zip       string
}

// TableName overrides the default tablename generated by GORM
func (AddressORM) TableName() string {
	return "addresses"
}

// ToORM runs the BeforeToORM hook if present, converts the fields of this
// object to ORM format, runs the AfterToORM hook, then returns the ORM object
func (m *Address) ToORM(ctx context.Context) (AddressORM, error) {
	to := AddressORM{}
	var err error
	if prehook, ok := interface{}(m).(AddressWithBeforeToORM); ok {
		if err = prehook.BeforeToORM(ctx, &to); err != nil {
			return to, err
		}
	}
	to.Id = m.Id
	to.Address = m.Address
	to.City = m.City
	to.State = m.State
	to.Zip = m.Zip
	to.Country = m.Country
	to.Type = m.Type
	if m.User != nil {
		tempUser, err := m.User.ToORM(ctx)
		if err != nil {
			return to, err
		}
		to.User = &tempUser
	}
	if m.CreatedAt != nil {
		t := m.CreatedAt.AsTime()
		to.CreatedAt = &t
	}
	if m.UpdatedAt != nil {
		t := m.UpdatedAt.AsTime()
		to.UpdatedAt = &t
	}
	if posthook, ok := interface{}(m).(AddressWithAfterToORM); ok {
		err = posthook.AfterToORM(ctx, &to)
	}
	return to, err
}

// ToPB runs the BeforeToPB hook if present, converts the fields of this
// object to PB format, runs the AfterToPB hook, then returns the PB object
func (m *AddressORM) ToPB(ctx context.Context) (Address, error) {
	to := Address{}
	var err error
	if prehook, ok := interface{}(m).(AddressWithBeforeToPB); ok {
		if err = prehook.BeforeToPB(ctx, &to); err != nil {
			return to, err
		}
	}
	to.Id = m.Id
	to.Address = m.Address
	to.City = m.City
	to.State = m.State
	to.Zip = m.Zip
	to.Country = m.Country
	to.Type = m.Type
	if m.User != nil {
		tempUser, err := m.User.ToPB(ctx)
		if err != nil {
			return to, err
		}
		to.User = &tempUser
	}
	if m.CreatedAt != nil {
		to.CreatedAt = timestamppb.New(*m.CreatedAt)
	}
	if m.UpdatedAt != nil {
		to.UpdatedAt = timestamppb.New(*m.UpdatedAt)
	}
	if posthook, ok := interface{}(m).(AddressWithAfterToPB); ok {
		err = posthook.AfterToPB(ctx, &to)
	}
	return to, err
}

// The following are interfaces you can implement for special behavior during ORM/PB conversions
// of type Address the arg will be the target, the caller the one being converted from

// AddressBeforeToORM called before default ToORM code
type AddressWithBeforeToORM interface {
	BeforeToORM(context.Context, *AddressORM) error
}

// AddressAfterToORM called after default ToORM code
type AddressWithAfterToORM interface {
	AfterToORM(context.Context, *AddressORM) error
}

// AddressBeforeToPB called before default ToPB code
type AddressWithBeforeToPB interface {
	BeforeToPB(context.Context, *Address) error
}

// AddressAfterToPB called after default ToPB code
type AddressWithAfterToPB interface {
	AfterToPB(context.Context, *Address) error
}

type UserPermissionsORM struct {
	CreatedAt  *time.Time
	Permission string
	UpdatedAt  *time.Time
	User       *UserORM `gorm:"foreignkey:UserId;association_foreignkey:Id"`
	UserId     *string
}

// TableName overrides the default tablename generated by GORM
func (UserPermissionsORM) TableName() string {
	return "user_permissions"
}

// ToORM runs the BeforeToORM hook if present, converts the fields of this
// object to ORM format, runs the AfterToORM hook, then returns the ORM object
func (m *UserPermissions) ToORM(ctx context.Context) (UserPermissionsORM, error) {
	to := UserPermissionsORM{}
	var err error
	if prehook, ok := interface{}(m).(UserPermissionsWithBeforeToORM); ok {
		if err = prehook.BeforeToORM(ctx, &to); err != nil {
			return to, err
		}
	}
	if m.User != nil {
		tempUser, err := m.User.ToORM(ctx)
		if err != nil {
			return to, err
		}
		to.User = &tempUser
	}
	to.Permission = m.Permission
	if m.CreatedAt != nil {
		t := m.CreatedAt.AsTime()
		to.CreatedAt = &t
	}
	if m.UpdatedAt != nil {
		t := m.UpdatedAt.AsTime()
		to.UpdatedAt = &t
	}
	if posthook, ok := interface{}(m).(UserPermissionsWithAfterToORM); ok {
		err = posthook.AfterToORM(ctx, &to)
	}
	return to, err
}

// ToPB runs the BeforeToPB hook if present, converts the fields of this
// object to PB format, runs the AfterToPB hook, then returns the PB object
func (m *UserPermissionsORM) ToPB(ctx context.Context) (UserPermissions, error) {
	to := UserPermissions{}
	var err error
	if prehook, ok := interface{}(m).(UserPermissionsWithBeforeToPB); ok {
		if err = prehook.BeforeToPB(ctx, &to); err != nil {
			return to, err
		}
	}
	if m.User != nil {
		tempUser, err := m.User.ToPB(ctx)
		if err != nil {
			return to, err
		}
		to.User = &tempUser
	}
	to.Permission = m.Permission
	if m.CreatedAt != nil {
		to.CreatedAt = timestamppb.New(*m.CreatedAt)
	}
	if m.UpdatedAt != nil {
		to.UpdatedAt = timestamppb.New(*m.UpdatedAt)
	}
	if posthook, ok := interface{}(m).(UserPermissionsWithAfterToPB); ok {
		err = posthook.AfterToPB(ctx, &to)
	}
	return to, err
}

// The following are interfaces you can implement for special behavior during ORM/PB conversions
// of type UserPermissions the arg will be the target, the caller the one being converted from

// UserPermissionsBeforeToORM called before default ToORM code
type UserPermissionsWithBeforeToORM interface {
	BeforeToORM(context.Context, *UserPermissionsORM) error
}

// UserPermissionsAfterToORM called after default ToORM code
type UserPermissionsWithAfterToORM interface {
	AfterToORM(context.Context, *UserPermissionsORM) error
}

// UserPermissionsBeforeToPB called before default ToPB code
type UserPermissionsWithBeforeToPB interface {
	BeforeToPB(context.Context, *UserPermissions) error
}

// UserPermissionsAfterToPB called after default ToPB code
type UserPermissionsWithAfterToPB interface {
	AfterToPB(context.Context, *UserPermissions) error
}

type UserORM struct {
	CreatedAt *time.Time
	Email     string
	Firstname string
	Id        string `gorm:"type:uuid;primary_key"`
	ImageUrl  string
	Lastname  string
	Password  string
	Role      string
	Telephone string
	Token     string
	UpdatedAt *time.Time
	Username  string
}

// TableName overrides the default tablename generated by GORM
func (UserORM) TableName() string {
	return "users"
}

// ToORM runs the BeforeToORM hook if present, converts the fields of this
// object to ORM format, runs the AfterToORM hook, then returns the ORM object
func (m *User) ToORM(ctx context.Context) (UserORM, error) {
	to := UserORM{}
	var err error
	if prehook, ok := interface{}(m).(UserWithBeforeToORM); ok {
		if err = prehook.BeforeToORM(ctx, &to); err != nil {
			return to, err
		}
	}
	to.Id = m.Id
	to.Email = m.Email
	to.Username = m.Username
	to.Password = m.Password
	to.Firstname = m.Firstname
	to.Lastname = m.Lastname
	to.Telephone = m.Telephone
	to.Token = m.Token
	to.Role = m.Role
	to.ImageUrl = m.ImageUrl
	if m.CreatedAt != nil {
		t := m.CreatedAt.AsTime()
		to.CreatedAt = &t
	}
	if m.UpdatedAt != nil {
		t := m.UpdatedAt.AsTime()
		to.UpdatedAt = &t
	}
	if posthook, ok := interface{}(m).(UserWithAfterToORM); ok {
		err = posthook.AfterToORM(ctx, &to)
	}
	return to, err
}

// ToPB runs the BeforeToPB hook if present, converts the fields of this
// object to PB format, runs the AfterToPB hook, then returns the PB object
func (m *UserORM) ToPB(ctx context.Context) (User, error) {
	to := User{}
	var err error
	if prehook, ok := interface{}(m).(UserWithBeforeToPB); ok {
		if err = prehook.BeforeToPB(ctx, &to); err != nil {
			return to, err
		}
	}
	to.Id = m.Id
	to.Email = m.Email
	to.Username = m.Username
	to.Password = m.Password
	to.Firstname = m.Firstname
	to.Lastname = m.Lastname
	to.Telephone = m.Telephone
	to.Token = m.Token
	to.Role = m.Role
	to.ImageUrl = m.ImageUrl
	if m.CreatedAt != nil {
		to.CreatedAt = timestamppb.New(*m.CreatedAt)
	}
	if m.UpdatedAt != nil {
		to.UpdatedAt = timestamppb.New(*m.UpdatedAt)
	}
	if posthook, ok := interface{}(m).(UserWithAfterToPB); ok {
		err = posthook.AfterToPB(ctx, &to)
	}
	return to, err
}

// The following are interfaces you can implement for special behavior during ORM/PB conversions
// of type User the arg will be the target, the caller the one being converted from

// UserBeforeToORM called before default ToORM code
type UserWithBeforeToORM interface {
	BeforeToORM(context.Context, *UserORM) error
}

// UserAfterToORM called after default ToORM code
type UserWithAfterToORM interface {
	AfterToORM(context.Context, *UserORM) error
}

// UserBeforeToPB called before default ToPB code
type UserWithBeforeToPB interface {
	BeforeToPB(context.Context, *User) error
}

// UserAfterToPB called after default ToPB code
type UserWithAfterToPB interface {
	AfterToPB(context.Context, *User) error
}

// DefaultCreateAddress executes a basic gorm create call
func DefaultCreateAddress(ctx context.Context, in *Address, db *gorm.DB) (*Address, error) {
	if in == nil {
		return nil, errors.NilArgumentError
	}
	ormObj, err := in.ToORM(ctx)
	if err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&ormObj).(AddressORMWithBeforeCreate_); ok {
		if db, err = hook.BeforeCreate_(ctx, db); err != nil {
			return nil, err
		}
	}
	if err = db.Create(&ormObj).Error; err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&ormObj).(AddressORMWithAfterCreate_); ok {
		if err = hook.AfterCreate_(ctx, db); err != nil {
			return nil, err
		}
	}
	pbResponse, err := ormObj.ToPB(ctx)
	return &pbResponse, err
}

type AddressORMWithBeforeCreate_ interface {
	BeforeCreate_(context.Context, *gorm.DB) (*gorm.DB, error)
}
type AddressORMWithAfterCreate_ interface {
	AfterCreate_(context.Context, *gorm.DB) error
}

func DefaultReadAddress(ctx context.Context, in *Address, db *gorm.DB) (*Address, error) {
	if in == nil {
		return nil, errors.NilArgumentError
	}
	ormObj, err := in.ToORM(ctx)
	if err != nil {
		return nil, err
	}
	if ormObj.Id == 0 {
		return nil, errors.EmptyIdError
	}
	if hook, ok := interface{}(&ormObj).(AddressORMWithBeforeReadApplyQuery); ok {
		if db, err = hook.BeforeReadApplyQuery(ctx, db); err != nil {
			return nil, err
		}
	}
	if db, err = gorm1.ApplyFieldSelection(ctx, db, nil, &AddressORM{}); err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&ormObj).(AddressORMWithBeforeReadFind); ok {
		if db, err = hook.BeforeReadFind(ctx, db); err != nil {
			return nil, err
		}
	}
	ormResponse := AddressORM{}
	if err = db.Where(&ormObj).First(&ormResponse).Error; err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&ormResponse).(AddressORMWithAfterReadFind); ok {
		if err = hook.AfterReadFind(ctx, db); err != nil {
			return nil, err
		}
	}
	pbResponse, err := ormResponse.ToPB(ctx)
	return &pbResponse, err
}

type AddressORMWithBeforeReadApplyQuery interface {
	BeforeReadApplyQuery(context.Context, *gorm.DB) (*gorm.DB, error)
}
type AddressORMWithBeforeReadFind interface {
	BeforeReadFind(context.Context, *gorm.DB) (*gorm.DB, error)
}
type AddressORMWithAfterReadFind interface {
	AfterReadFind(context.Context, *gorm.DB) error
}

func DefaultDeleteAddress(ctx context.Context, in *Address, db *gorm.DB) error {
	if in == nil {
		return errors.NilArgumentError
	}
	ormObj, err := in.ToORM(ctx)
	if err != nil {
		return err
	}
	if ormObj.Id == 0 {
		return errors.EmptyIdError
	}
	if hook, ok := interface{}(&ormObj).(AddressORMWithBeforeDelete_); ok {
		if db, err = hook.BeforeDelete_(ctx, db); err != nil {
			return err
		}
	}
	err = db.Where(&ormObj).Delete(&AddressORM{}).Error
	if err != nil {
		return err
	}
	if hook, ok := interface{}(&ormObj).(AddressORMWithAfterDelete_); ok {
		err = hook.AfterDelete_(ctx, db)
	}
	return err
}

type AddressORMWithBeforeDelete_ interface {
	BeforeDelete_(context.Context, *gorm.DB) (*gorm.DB, error)
}
type AddressORMWithAfterDelete_ interface {
	AfterDelete_(context.Context, *gorm.DB) error
}

func DefaultDeleteAddressSet(ctx context.Context, in []*Address, db *gorm.DB) error {
	if in == nil {
		return errors.NilArgumentError
	}
	var err error
	keys := []int32{}
	for _, obj := range in {
		ormObj, err := obj.ToORM(ctx)
		if err != nil {
			return err
		}
		if ormObj.Id == 0 {
			return errors.EmptyIdError
		}
		keys = append(keys, ormObj.Id)
	}
	if hook, ok := (interface{}(&AddressORM{})).(AddressORMWithBeforeDeleteSet); ok {
		if db, err = hook.BeforeDeleteSet(ctx, in, db); err != nil {
			return err
		}
	}
	err = db.Where("id in (?)", keys).Delete(&AddressORM{}).Error
	if err != nil {
		return err
	}
	if hook, ok := (interface{}(&AddressORM{})).(AddressORMWithAfterDeleteSet); ok {
		err = hook.AfterDeleteSet(ctx, in, db)
	}
	return err
}

type AddressORMWithBeforeDeleteSet interface {
	BeforeDeleteSet(context.Context, []*Address, *gorm.DB) (*gorm.DB, error)
}
type AddressORMWithAfterDeleteSet interface {
	AfterDeleteSet(context.Context, []*Address, *gorm.DB) error
}

// DefaultStrictUpdateAddress clears / replaces / appends first level 1:many children and then executes a gorm update call
func DefaultStrictUpdateAddress(ctx context.Context, in *Address, db *gorm.DB) (*Address, error) {
	if in == nil {
		return nil, fmt.Errorf("Nil argument to DefaultStrictUpdateAddress")
	}
	ormObj, err := in.ToORM(ctx)
	if err != nil {
		return nil, err
	}
	lockedRow := &AddressORM{}
	db.Model(&ormObj).Set("gorm:query_option", "FOR UPDATE").Where("id=?", ormObj.Id).First(lockedRow)
	if hook, ok := interface{}(&ormObj).(AddressORMWithBeforeStrictUpdateCleanup); ok {
		if db, err = hook.BeforeStrictUpdateCleanup(ctx, db); err != nil {
			return nil, err
		}
	}
	if hook, ok := interface{}(&ormObj).(AddressORMWithBeforeStrictUpdateSave); ok {
		if db, err = hook.BeforeStrictUpdateSave(ctx, db); err != nil {
			return nil, err
		}
	}
	if err = db.Save(&ormObj).Error; err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&ormObj).(AddressORMWithAfterStrictUpdateSave); ok {
		if err = hook.AfterStrictUpdateSave(ctx, db); err != nil {
			return nil, err
		}
	}
	pbResponse, err := ormObj.ToPB(ctx)
	if err != nil {
		return nil, err
	}
	return &pbResponse, err
}

type AddressORMWithBeforeStrictUpdateCleanup interface {
	BeforeStrictUpdateCleanup(context.Context, *gorm.DB) (*gorm.DB, error)
}
type AddressORMWithBeforeStrictUpdateSave interface {
	BeforeStrictUpdateSave(context.Context, *gorm.DB) (*gorm.DB, error)
}
type AddressORMWithAfterStrictUpdateSave interface {
	AfterStrictUpdateSave(context.Context, *gorm.DB) error
}

// DefaultPatchAddress executes a basic gorm update call with patch behavior
func DefaultPatchAddress(ctx context.Context, in *Address, updateMask *field_mask.FieldMask, db *gorm.DB) (*Address, error) {
	if in == nil {
		return nil, errors.NilArgumentError
	}
	var pbObj Address
	var err error
	if hook, ok := interface{}(&pbObj).(AddressWithBeforePatchRead); ok {
		if db, err = hook.BeforePatchRead(ctx, in, updateMask, db); err != nil {
			return nil, err
		}
	}
	pbReadRes, err := DefaultReadAddress(ctx, &Address{Id: in.GetId()}, db)
	if err != nil {
		return nil, err
	}
	pbObj = *pbReadRes
	if hook, ok := interface{}(&pbObj).(AddressWithBeforePatchApplyFieldMask); ok {
		if db, err = hook.BeforePatchApplyFieldMask(ctx, in, updateMask, db); err != nil {
			return nil, err
		}
	}
	if _, err := DefaultApplyFieldMaskAddress(ctx, &pbObj, in, updateMask, "", db); err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&pbObj).(AddressWithBeforePatchSave); ok {
		if db, err = hook.BeforePatchSave(ctx, in, updateMask, db); err != nil {
			return nil, err
		}
	}
	pbResponse, err := DefaultStrictUpdateAddress(ctx, &pbObj, db)
	if err != nil {
		return nil, err
	}
	if hook, ok := interface{}(pbResponse).(AddressWithAfterPatchSave); ok {
		if err = hook.AfterPatchSave(ctx, in, updateMask, db); err != nil {
			return nil, err
		}
	}
	return pbResponse, nil
}

type AddressWithBeforePatchRead interface {
	BeforePatchRead(context.Context, *Address, *field_mask.FieldMask, *gorm.DB) (*gorm.DB, error)
}
type AddressWithBeforePatchApplyFieldMask interface {
	BeforePatchApplyFieldMask(context.Context, *Address, *field_mask.FieldMask, *gorm.DB) (*gorm.DB, error)
}
type AddressWithBeforePatchSave interface {
	BeforePatchSave(context.Context, *Address, *field_mask.FieldMask, *gorm.DB) (*gorm.DB, error)
}
type AddressWithAfterPatchSave interface {
	AfterPatchSave(context.Context, *Address, *field_mask.FieldMask, *gorm.DB) error
}

// DefaultPatchSetAddress executes a bulk gorm update call with patch behavior
func DefaultPatchSetAddress(ctx context.Context, objects []*Address, updateMasks []*field_mask.FieldMask, db *gorm.DB) ([]*Address, error) {
	if len(objects) != len(updateMasks) {
		return nil, fmt.Errorf(errors.BadRepeatedFieldMaskTpl, len(updateMasks), len(objects))
	}

	results := make([]*Address, 0, len(objects))
	for i, patcher := range objects {
		pbResponse, err := DefaultPatchAddress(ctx, patcher, updateMasks[i], db)
		if err != nil {
			return nil, err
		}

		results = append(results, pbResponse)
	}

	return results, nil
}

// DefaultApplyFieldMaskAddress patches an pbObject with patcher according to a field mask.
func DefaultApplyFieldMaskAddress(ctx context.Context, patchee *Address, patcher *Address, updateMask *field_mask.FieldMask, prefix string, db *gorm.DB) (*Address, error) {
	if patcher == nil {
		return nil, nil
	} else if patchee == nil {
		return nil, errors.NilArgumentError
	}
	var err error
	var updatedUser bool
	var updatedCreatedAt bool
	var updatedUpdatedAt bool
	for i, f := range updateMask.Paths {
		if f == prefix+"Id" {
			patchee.Id = patcher.Id
			continue
		}
		if f == prefix+"Address" {
			patchee.Address = patcher.Address
			continue
		}
		if f == prefix+"City" {
			patchee.City = patcher.City
			continue
		}
		if f == prefix+"State" {
			patchee.State = patcher.State
			continue
		}
		if f == prefix+"Zip" {
			patchee.Zip = patcher.Zip
			continue
		}
		if f == prefix+"Country" {
			patchee.Country = patcher.Country
			continue
		}
		if f == prefix+"Type" {
			patchee.Type = patcher.Type
			continue
		}
		if !updatedUser && strings.HasPrefix(f, prefix+"User.") {
			updatedUser = true
			if patcher.User == nil {
				patchee.User = nil
				continue
			}
			if patchee.User == nil {
				patchee.User = &User{}
			}
			if o, err := DefaultApplyFieldMaskUser(ctx, patchee.User, patcher.User, &field_mask.FieldMask{Paths: updateMask.Paths[i:]}, prefix+"User.", db); err != nil {
				return nil, err
			} else {
				patchee.User = o
			}
			continue
		}
		if f == prefix+"User" {
			updatedUser = true
			patchee.User = patcher.User
			continue
		}
		if !updatedCreatedAt && strings.HasPrefix(f, prefix+"CreatedAt.") {
			if patcher.CreatedAt == nil {
				patchee.CreatedAt = nil
				continue
			}
			if patchee.CreatedAt == nil {
				patchee.CreatedAt = &timestamppb.Timestamp{}
			}
			childMask := &field_mask.FieldMask{}
			for j := i; j < len(updateMask.Paths); j++ {
				if trimPath := strings.TrimPrefix(updateMask.Paths[j], prefix+"CreatedAt."); trimPath != updateMask.Paths[j] {
					childMask.Paths = append(childMask.Paths, trimPath)
				}
			}
			if err := gorm1.MergeWithMask(patcher.CreatedAt, patchee.CreatedAt, childMask); err != nil {
				return nil, nil
			}
		}
		if f == prefix+"CreatedAt" {
			updatedCreatedAt = true
			patchee.CreatedAt = patcher.CreatedAt
			continue
		}
		if !updatedUpdatedAt && strings.HasPrefix(f, prefix+"UpdatedAt.") {
			if patcher.UpdatedAt == nil {
				patchee.UpdatedAt = nil
				continue
			}
			if patchee.UpdatedAt == nil {
				patchee.UpdatedAt = &timestamppb.Timestamp{}
			}
			childMask := &field_mask.FieldMask{}
			for j := i; j < len(updateMask.Paths); j++ {
				if trimPath := strings.TrimPrefix(updateMask.Paths[j], prefix+"UpdatedAt."); trimPath != updateMask.Paths[j] {
					childMask.Paths = append(childMask.Paths, trimPath)
				}
			}
			if err := gorm1.MergeWithMask(patcher.UpdatedAt, patchee.UpdatedAt, childMask); err != nil {
				return nil, nil
			}
		}
		if f == prefix+"UpdatedAt" {
			updatedUpdatedAt = true
			patchee.UpdatedAt = patcher.UpdatedAt
			continue
		}
	}
	if err != nil {
		return nil, err
	}
	return patchee, nil
}

// DefaultListAddress executes a gorm list call
func DefaultListAddress(ctx context.Context, db *gorm.DB) ([]*Address, error) {
	in := Address{}
	ormObj, err := in.ToORM(ctx)
	if err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&ormObj).(AddressORMWithBeforeListApplyQuery); ok {
		if db, err = hook.BeforeListApplyQuery(ctx, db); err != nil {
			return nil, err
		}
	}
	db, err = gorm1.ApplyCollectionOperators(ctx, db, &AddressORM{}, &Address{}, nil, nil, nil, nil)
	if err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&ormObj).(AddressORMWithBeforeListFind); ok {
		if db, err = hook.BeforeListFind(ctx, db); err != nil {
			return nil, err
		}
	}
	db = db.Where(&ormObj)
	db = db.Order("id")
	ormResponse := []AddressORM{}
	if err := db.Find(&ormResponse).Error; err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&ormObj).(AddressORMWithAfterListFind); ok {
		if err = hook.AfterListFind(ctx, db, &ormResponse); err != nil {
			return nil, err
		}
	}
	pbResponse := []*Address{}
	for _, responseEntry := range ormResponse {
		temp, err := responseEntry.ToPB(ctx)
		if err != nil {
			return nil, err
		}
		pbResponse = append(pbResponse, &temp)
	}
	return pbResponse, nil
}

type AddressORMWithBeforeListApplyQuery interface {
	BeforeListApplyQuery(context.Context, *gorm.DB) (*gorm.DB, error)
}
type AddressORMWithBeforeListFind interface {
	BeforeListFind(context.Context, *gorm.DB) (*gorm.DB, error)
}
type AddressORMWithAfterListFind interface {
	AfterListFind(context.Context, *gorm.DB, *[]AddressORM) error
}

// DefaultCreateUserPermissions executes a basic gorm create call
func DefaultCreateUserPermissions(ctx context.Context, in *UserPermissions, db *gorm.DB) (*UserPermissions, error) {
	if in == nil {
		return nil, errors.NilArgumentError
	}
	ormObj, err := in.ToORM(ctx)
	if err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&ormObj).(UserPermissionsORMWithBeforeCreate_); ok {
		if db, err = hook.BeforeCreate_(ctx, db); err != nil {
			return nil, err
		}
	}
	if err = db.Create(&ormObj).Error; err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&ormObj).(UserPermissionsORMWithAfterCreate_); ok {
		if err = hook.AfterCreate_(ctx, db); err != nil {
			return nil, err
		}
	}
	pbResponse, err := ormObj.ToPB(ctx)
	return &pbResponse, err
}

type UserPermissionsORMWithBeforeCreate_ interface {
	BeforeCreate_(context.Context, *gorm.DB) (*gorm.DB, error)
}
type UserPermissionsORMWithAfterCreate_ interface {
	AfterCreate_(context.Context, *gorm.DB) error
}

// DefaultApplyFieldMaskUserPermissions patches an pbObject with patcher according to a field mask.
func DefaultApplyFieldMaskUserPermissions(ctx context.Context, patchee *UserPermissions, patcher *UserPermissions, updateMask *field_mask.FieldMask, prefix string, db *gorm.DB) (*UserPermissions, error) {
	if patcher == nil {
		return nil, nil
	} else if patchee == nil {
		return nil, errors.NilArgumentError
	}
	var err error
	var updatedUser bool
	var updatedCreatedAt bool
	var updatedUpdatedAt bool
	for i, f := range updateMask.Paths {
		if !updatedUser && strings.HasPrefix(f, prefix+"User.") {
			updatedUser = true
			if patcher.User == nil {
				patchee.User = nil
				continue
			}
			if patchee.User == nil {
				patchee.User = &User{}
			}
			if o, err := DefaultApplyFieldMaskUser(ctx, patchee.User, patcher.User, &field_mask.FieldMask{Paths: updateMask.Paths[i:]}, prefix+"User.", db); err != nil {
				return nil, err
			} else {
				patchee.User = o
			}
			continue
		}
		if f == prefix+"User" {
			updatedUser = true
			patchee.User = patcher.User
			continue
		}
		if f == prefix+"Permission" {
			patchee.Permission = patcher.Permission
			continue
		}
		if !updatedCreatedAt && strings.HasPrefix(f, prefix+"CreatedAt.") {
			if patcher.CreatedAt == nil {
				patchee.CreatedAt = nil
				continue
			}
			if patchee.CreatedAt == nil {
				patchee.CreatedAt = &timestamppb.Timestamp{}
			}
			childMask := &field_mask.FieldMask{}
			for j := i; j < len(updateMask.Paths); j++ {
				if trimPath := strings.TrimPrefix(updateMask.Paths[j], prefix+"CreatedAt."); trimPath != updateMask.Paths[j] {
					childMask.Paths = append(childMask.Paths, trimPath)
				}
			}
			if err := gorm1.MergeWithMask(patcher.CreatedAt, patchee.CreatedAt, childMask); err != nil {
				return nil, nil
			}
		}
		if f == prefix+"CreatedAt" {
			updatedCreatedAt = true
			patchee.CreatedAt = patcher.CreatedAt
			continue
		}
		if !updatedUpdatedAt && strings.HasPrefix(f, prefix+"UpdatedAt.") {
			if patcher.UpdatedAt == nil {
				patchee.UpdatedAt = nil
				continue
			}
			if patchee.UpdatedAt == nil {
				patchee.UpdatedAt = &timestamppb.Timestamp{}
			}
			childMask := &field_mask.FieldMask{}
			for j := i; j < len(updateMask.Paths); j++ {
				if trimPath := strings.TrimPrefix(updateMask.Paths[j], prefix+"UpdatedAt."); trimPath != updateMask.Paths[j] {
					childMask.Paths = append(childMask.Paths, trimPath)
				}
			}
			if err := gorm1.MergeWithMask(patcher.UpdatedAt, patchee.UpdatedAt, childMask); err != nil {
				return nil, nil
			}
		}
		if f == prefix+"UpdatedAt" {
			updatedUpdatedAt = true
			patchee.UpdatedAt = patcher.UpdatedAt
			continue
		}
	}
	if err != nil {
		return nil, err
	}
	return patchee, nil
}

// DefaultListUserPermissions executes a gorm list call
func DefaultListUserPermissions(ctx context.Context, db *gorm.DB) ([]*UserPermissions, error) {
	in := UserPermissions{}
	ormObj, err := in.ToORM(ctx)
	if err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&ormObj).(UserPermissionsORMWithBeforeListApplyQuery); ok {
		if db, err = hook.BeforeListApplyQuery(ctx, db); err != nil {
			return nil, err
		}
	}
	db, err = gorm1.ApplyCollectionOperators(ctx, db, &UserPermissionsORM{}, &UserPermissions{}, nil, nil, nil, nil)
	if err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&ormObj).(UserPermissionsORMWithBeforeListFind); ok {
		if db, err = hook.BeforeListFind(ctx, db); err != nil {
			return nil, err
		}
	}
	db = db.Where(&ormObj)
	ormResponse := []UserPermissionsORM{}
	if err := db.Find(&ormResponse).Error; err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&ormObj).(UserPermissionsORMWithAfterListFind); ok {
		if err = hook.AfterListFind(ctx, db, &ormResponse); err != nil {
			return nil, err
		}
	}
	pbResponse := []*UserPermissions{}
	for _, responseEntry := range ormResponse {
		temp, err := responseEntry.ToPB(ctx)
		if err != nil {
			return nil, err
		}
		pbResponse = append(pbResponse, &temp)
	}
	return pbResponse, nil
}

type UserPermissionsORMWithBeforeListApplyQuery interface {
	BeforeListApplyQuery(context.Context, *gorm.DB) (*gorm.DB, error)
}
type UserPermissionsORMWithBeforeListFind interface {
	BeforeListFind(context.Context, *gorm.DB) (*gorm.DB, error)
}
type UserPermissionsORMWithAfterListFind interface {
	AfterListFind(context.Context, *gorm.DB, *[]UserPermissionsORM) error
}

// DefaultCreateUser executes a basic gorm create call
func DefaultCreateUser(ctx context.Context, in *User, db *gorm.DB) (*User, error) {
	if in == nil {
		return nil, errors.NilArgumentError
	}
	ormObj, err := in.ToORM(ctx)
	if err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&ormObj).(UserORMWithBeforeCreate_); ok {
		if db, err = hook.BeforeCreate_(ctx, db); err != nil {
			return nil, err
		}
	}
	if err = db.Create(&ormObj).Error; err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&ormObj).(UserORMWithAfterCreate_); ok {
		if err = hook.AfterCreate_(ctx, db); err != nil {
			return nil, err
		}
	}
	pbResponse, err := ormObj.ToPB(ctx)
	return &pbResponse, err
}

type UserORMWithBeforeCreate_ interface {
	BeforeCreate_(context.Context, *gorm.DB) (*gorm.DB, error)
}
type UserORMWithAfterCreate_ interface {
	AfterCreate_(context.Context, *gorm.DB) error
}

func DefaultReadUser(ctx context.Context, in *User, db *gorm.DB) (*User, error) {
	if in == nil {
		return nil, errors.NilArgumentError
	}
	ormObj, err := in.ToORM(ctx)
	if err != nil {
		return nil, err
	}
	if ormObj.Id == "" {
		return nil, errors.EmptyIdError
	}
	if hook, ok := interface{}(&ormObj).(UserORMWithBeforeReadApplyQuery); ok {
		if db, err = hook.BeforeReadApplyQuery(ctx, db); err != nil {
			return nil, err
		}
	}
	if db, err = gorm1.ApplyFieldSelection(ctx, db, nil, &UserORM{}); err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&ormObj).(UserORMWithBeforeReadFind); ok {
		if db, err = hook.BeforeReadFind(ctx, db); err != nil {
			return nil, err
		}
	}
	ormResponse := UserORM{}
	if err = db.Where(&ormObj).First(&ormResponse).Error; err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&ormResponse).(UserORMWithAfterReadFind); ok {
		if err = hook.AfterReadFind(ctx, db); err != nil {
			return nil, err
		}
	}
	pbResponse, err := ormResponse.ToPB(ctx)
	return &pbResponse, err
}

type UserORMWithBeforeReadApplyQuery interface {
	BeforeReadApplyQuery(context.Context, *gorm.DB) (*gorm.DB, error)
}
type UserORMWithBeforeReadFind interface {
	BeforeReadFind(context.Context, *gorm.DB) (*gorm.DB, error)
}
type UserORMWithAfterReadFind interface {
	AfterReadFind(context.Context, *gorm.DB) error
}

func DefaultDeleteUser(ctx context.Context, in *User, db *gorm.DB) error {
	if in == nil {
		return errors.NilArgumentError
	}
	ormObj, err := in.ToORM(ctx)
	if err != nil {
		return err
	}
	if ormObj.Id == "" {
		return errors.EmptyIdError
	}
	if hook, ok := interface{}(&ormObj).(UserORMWithBeforeDelete_); ok {
		if db, err = hook.BeforeDelete_(ctx, db); err != nil {
			return err
		}
	}
	err = db.Where(&ormObj).Delete(&UserORM{}).Error
	if err != nil {
		return err
	}
	if hook, ok := interface{}(&ormObj).(UserORMWithAfterDelete_); ok {
		err = hook.AfterDelete_(ctx, db)
	}
	return err
}

type UserORMWithBeforeDelete_ interface {
	BeforeDelete_(context.Context, *gorm.DB) (*gorm.DB, error)
}
type UserORMWithAfterDelete_ interface {
	AfterDelete_(context.Context, *gorm.DB) error
}

func DefaultDeleteUserSet(ctx context.Context, in []*User, db *gorm.DB) error {
	if in == nil {
		return errors.NilArgumentError
	}
	var err error
	keys := []string{}
	for _, obj := range in {
		ormObj, err := obj.ToORM(ctx)
		if err != nil {
			return err
		}
		if ormObj.Id == "" {
			return errors.EmptyIdError
		}
		keys = append(keys, ormObj.Id)
	}
	if hook, ok := (interface{}(&UserORM{})).(UserORMWithBeforeDeleteSet); ok {
		if db, err = hook.BeforeDeleteSet(ctx, in, db); err != nil {
			return err
		}
	}
	err = db.Where("id in (?)", keys).Delete(&UserORM{}).Error
	if err != nil {
		return err
	}
	if hook, ok := (interface{}(&UserORM{})).(UserORMWithAfterDeleteSet); ok {
		err = hook.AfterDeleteSet(ctx, in, db)
	}
	return err
}

type UserORMWithBeforeDeleteSet interface {
	BeforeDeleteSet(context.Context, []*User, *gorm.DB) (*gorm.DB, error)
}
type UserORMWithAfterDeleteSet interface {
	AfterDeleteSet(context.Context, []*User, *gorm.DB) error
}

// DefaultStrictUpdateUser clears / replaces / appends first level 1:many children and then executes a gorm update call
func DefaultStrictUpdateUser(ctx context.Context, in *User, db *gorm.DB) (*User, error) {
	if in == nil {
		return nil, fmt.Errorf("Nil argument to DefaultStrictUpdateUser")
	}
	ormObj, err := in.ToORM(ctx)
	if err != nil {
		return nil, err
	}
	lockedRow := &UserORM{}
	db.Model(&ormObj).Set("gorm:query_option", "FOR UPDATE").Where("id=?", ormObj.Id).First(lockedRow)
	if hook, ok := interface{}(&ormObj).(UserORMWithBeforeStrictUpdateCleanup); ok {
		if db, err = hook.BeforeStrictUpdateCleanup(ctx, db); err != nil {
			return nil, err
		}
	}
	if hook, ok := interface{}(&ormObj).(UserORMWithBeforeStrictUpdateSave); ok {
		if db, err = hook.BeforeStrictUpdateSave(ctx, db); err != nil {
			return nil, err
		}
	}
	if err = db.Save(&ormObj).Error; err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&ormObj).(UserORMWithAfterStrictUpdateSave); ok {
		if err = hook.AfterStrictUpdateSave(ctx, db); err != nil {
			return nil, err
		}
	}
	pbResponse, err := ormObj.ToPB(ctx)
	if err != nil {
		return nil, err
	}
	return &pbResponse, err
}

type UserORMWithBeforeStrictUpdateCleanup interface {
	BeforeStrictUpdateCleanup(context.Context, *gorm.DB) (*gorm.DB, error)
}
type UserORMWithBeforeStrictUpdateSave interface {
	BeforeStrictUpdateSave(context.Context, *gorm.DB) (*gorm.DB, error)
}
type UserORMWithAfterStrictUpdateSave interface {
	AfterStrictUpdateSave(context.Context, *gorm.DB) error
}

// DefaultPatchUser executes a basic gorm update call with patch behavior
func DefaultPatchUser(ctx context.Context, in *User, updateMask *field_mask.FieldMask, db *gorm.DB) (*User, error) {
	if in == nil {
		return nil, errors.NilArgumentError
	}
	var pbObj User
	var err error
	if hook, ok := interface{}(&pbObj).(UserWithBeforePatchRead); ok {
		if db, err = hook.BeforePatchRead(ctx, in, updateMask, db); err != nil {
			return nil, err
		}
	}
	pbReadRes, err := DefaultReadUser(ctx, &User{Id: in.GetId()}, db)
	if err != nil {
		return nil, err
	}
	pbObj = *pbReadRes
	if hook, ok := interface{}(&pbObj).(UserWithBeforePatchApplyFieldMask); ok {
		if db, err = hook.BeforePatchApplyFieldMask(ctx, in, updateMask, db); err != nil {
			return nil, err
		}
	}
	if _, err := DefaultApplyFieldMaskUser(ctx, &pbObj, in, updateMask, "", db); err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&pbObj).(UserWithBeforePatchSave); ok {
		if db, err = hook.BeforePatchSave(ctx, in, updateMask, db); err != nil {
			return nil, err
		}
	}
	pbResponse, err := DefaultStrictUpdateUser(ctx, &pbObj, db)
	if err != nil {
		return nil, err
	}
	if hook, ok := interface{}(pbResponse).(UserWithAfterPatchSave); ok {
		if err = hook.AfterPatchSave(ctx, in, updateMask, db); err != nil {
			return nil, err
		}
	}
	return pbResponse, nil
}

type UserWithBeforePatchRead interface {
	BeforePatchRead(context.Context, *User, *field_mask.FieldMask, *gorm.DB) (*gorm.DB, error)
}
type UserWithBeforePatchApplyFieldMask interface {
	BeforePatchApplyFieldMask(context.Context, *User, *field_mask.FieldMask, *gorm.DB) (*gorm.DB, error)
}
type UserWithBeforePatchSave interface {
	BeforePatchSave(context.Context, *User, *field_mask.FieldMask, *gorm.DB) (*gorm.DB, error)
}
type UserWithAfterPatchSave interface {
	AfterPatchSave(context.Context, *User, *field_mask.FieldMask, *gorm.DB) error
}

// DefaultPatchSetUser executes a bulk gorm update call with patch behavior
func DefaultPatchSetUser(ctx context.Context, objects []*User, updateMasks []*field_mask.FieldMask, db *gorm.DB) ([]*User, error) {
	if len(objects) != len(updateMasks) {
		return nil, fmt.Errorf(errors.BadRepeatedFieldMaskTpl, len(updateMasks), len(objects))
	}

	results := make([]*User, 0, len(objects))
	for i, patcher := range objects {
		pbResponse, err := DefaultPatchUser(ctx, patcher, updateMasks[i], db)
		if err != nil {
			return nil, err
		}

		results = append(results, pbResponse)
	}

	return results, nil
}

// DefaultApplyFieldMaskUser patches an pbObject with patcher according to a field mask.
func DefaultApplyFieldMaskUser(ctx context.Context, patchee *User, patcher *User, updateMask *field_mask.FieldMask, prefix string, db *gorm.DB) (*User, error) {
	if patcher == nil {
		return nil, nil
	} else if patchee == nil {
		return nil, errors.NilArgumentError
	}
	var err error
	var updatedCreatedAt bool
	var updatedUpdatedAt bool
	for i, f := range updateMask.Paths {
		if f == prefix+"Id" {
			patchee.Id = patcher.Id
			continue
		}
		if f == prefix+"Email" {
			patchee.Email = patcher.Email
			continue
		}
		if f == prefix+"Username" {
			patchee.Username = patcher.Username
			continue
		}
		if f == prefix+"Password" {
			patchee.Password = patcher.Password
			continue
		}
		if f == prefix+"Firstname" {
			patchee.Firstname = patcher.Firstname
			continue
		}
		if f == prefix+"Lastname" {
			patchee.Lastname = patcher.Lastname
			continue
		}
		if f == prefix+"Telephone" {
			patchee.Telephone = patcher.Telephone
			continue
		}
		if f == prefix+"Token" {
			patchee.Token = patcher.Token
			continue
		}
		if f == prefix+"Role" {
			patchee.Role = patcher.Role
			continue
		}
		if f == prefix+"ImageUrl" {
			patchee.ImageUrl = patcher.ImageUrl
			continue
		}
		if !updatedCreatedAt && strings.HasPrefix(f, prefix+"CreatedAt.") {
			if patcher.CreatedAt == nil {
				patchee.CreatedAt = nil
				continue
			}
			if patchee.CreatedAt == nil {
				patchee.CreatedAt = &timestamppb.Timestamp{}
			}
			childMask := &field_mask.FieldMask{}
			for j := i; j < len(updateMask.Paths); j++ {
				if trimPath := strings.TrimPrefix(updateMask.Paths[j], prefix+"CreatedAt."); trimPath != updateMask.Paths[j] {
					childMask.Paths = append(childMask.Paths, trimPath)
				}
			}
			if err := gorm1.MergeWithMask(patcher.CreatedAt, patchee.CreatedAt, childMask); err != nil {
				return nil, nil
			}
		}
		if f == prefix+"CreatedAt" {
			updatedCreatedAt = true
			patchee.CreatedAt = patcher.CreatedAt
			continue
		}
		if !updatedUpdatedAt && strings.HasPrefix(f, prefix+"UpdatedAt.") {
			if patcher.UpdatedAt == nil {
				patchee.UpdatedAt = nil
				continue
			}
			if patchee.UpdatedAt == nil {
				patchee.UpdatedAt = &timestamppb.Timestamp{}
			}
			childMask := &field_mask.FieldMask{}
			for j := i; j < len(updateMask.Paths); j++ {
				if trimPath := strings.TrimPrefix(updateMask.Paths[j], prefix+"UpdatedAt."); trimPath != updateMask.Paths[j] {
					childMask.Paths = append(childMask.Paths, trimPath)
				}
			}
			if err := gorm1.MergeWithMask(patcher.UpdatedAt, patchee.UpdatedAt, childMask); err != nil {
				return nil, nil
			}
		}
		if f == prefix+"UpdatedAt" {
			updatedUpdatedAt = true
			patchee.UpdatedAt = patcher.UpdatedAt
			continue
		}
	}
	if err != nil {
		return nil, err
	}
	return patchee, nil
}

// DefaultListUser executes a gorm list call
func DefaultListUser(ctx context.Context, db *gorm.DB) ([]*User, error) {
	in := User{}
	ormObj, err := in.ToORM(ctx)
	if err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&ormObj).(UserORMWithBeforeListApplyQuery); ok {
		if db, err = hook.BeforeListApplyQuery(ctx, db); err != nil {
			return nil, err
		}
	}
	db, err = gorm1.ApplyCollectionOperators(ctx, db, &UserORM{}, &User{}, nil, nil, nil, nil)
	if err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&ormObj).(UserORMWithBeforeListFind); ok {
		if db, err = hook.BeforeListFind(ctx, db); err != nil {
			return nil, err
		}
	}
	db = db.Where(&ormObj)
	db = db.Order("id")
	ormResponse := []UserORM{}
	if err := db.Find(&ormResponse).Error; err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&ormObj).(UserORMWithAfterListFind); ok {
		if err = hook.AfterListFind(ctx, db, &ormResponse); err != nil {
			return nil, err
		}
	}
	pbResponse := []*User{}
	for _, responseEntry := range ormResponse {
		temp, err := responseEntry.ToPB(ctx)
		if err != nil {
			return nil, err
		}
		pbResponse = append(pbResponse, &temp)
	}
	return pbResponse, nil
}

type UserORMWithBeforeListApplyQuery interface {
	BeforeListApplyQuery(context.Context, *gorm.DB) (*gorm.DB, error)
}
type UserORMWithBeforeListFind interface {
	BeforeListFind(context.Context, *gorm.DB) (*gorm.DB, error)
}
type UserORMWithAfterListFind interface {
	AfterListFind(context.Context, *gorm.DB, *[]UserORM) error
}
