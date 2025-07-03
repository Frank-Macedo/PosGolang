package entity

import (
	"database/sql/driver"
	"fmt"

	"github.com/google/uuid"
)

type ID uuid.UUID

func NewID() ID {
	return ID(uuid.New())
}

func (id ID) String() string {
	return uuid.UUID(id).String()
}

func ParseID(s string) (ID, error) {
	id, err := uuid.Parse(s)

	return ID(id), err
}

func (id ID) Value() (driver.Value, error) {
	return uuid.UUID(id).String(), nil
}
func (id ID) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, id.String())), nil
}

func (id *ID) Scan(value interface{}) error {
	switch v := value.(type) {
	case string:
		uid, err := uuid.Parse(v)
		if err != nil {
			return err
		}
		*id = ID(uid)
		return nil
	case []byte:
		uid, err := uuid.ParseBytes(v)
		if err != nil {
			return err
		}
		*id = ID(uid)
		return nil
	default:
		return fmt.Errorf("cannot scan type %T into ID", value)
	}
}
