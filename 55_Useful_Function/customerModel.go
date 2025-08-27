package main

// Customer represents a customer table in the database.
// GORM uses struct fields as columns and struct tags to define behavior.
type Customer struct {
	ID    int    `gorm:"primaryKey" json:"id"`           // Primary key column
	Name  string `gorm:"size:100;not null" json:"name"` // String column with max length 100, cannot be null
	Email string `gorm:"unique;size:100" json:"email,omitempty"` // Unique email, optional in JSON output
}

// Product represents a product table.
type Product struct {
	ID   int    `gorm:"primaryKey" json:"id"`
	Name string `gorm:"size:100;not null" json:"name"`
}

// Order represents an example table with a foreign key reference.
// CustomerID references Customer(ID)
type Order struct {
	ID         int      `gorm:"primaryKey" json:"id"`
	CustomerID int      `gorm:"not null" json:"customer_id"`     // Foreign key column
	Customer   Customer `gorm:"foreignKey:CustomerID" json:"customer"` // GORM will auto-preload customer info if configured
	Total      float64  `gorm:"type:numeric" json:"total"`
}

// Person is a simple struct, but without GORM tags, it won't auto-map to a table.
type Person struct {
	Name string
	Age  int
}

// Person1 is another simple struct. Acceptable in GORM, but without gorm tags it will use defaults:
// - Table name: pluralized struct name (person1s)
// - Column names: same as struct field names (Name, etc.)
type Person1 struct {
	Name string
}

/*
Notes on GORM struct modeling:

1. Field types:
   - int, uint, float64, string, bool, time.Time, slices for JSON columns
2. Tags:
   - `gorm:"primaryKey"` → marks a primary key
   - `gorm:"unique"` → unique constraint
   - `gorm:"not null"` → NOT NULL constraint
   - `gorm:"size:100"` → maximum string length
   - `gorm:"type:jsonb"` → for PostgreSQL JSONB fields
   - `json:"fieldname"` → JSON serialization
3. Relations:
   - Has One: use struct field and `foreignKey` tag
   - Has Many: use slice of structs and `foreignKey`
   - Belongs To: define a foreign key column and struct reference
   - Many-to-Many: define slice of structs with `many2many` tag
*/
