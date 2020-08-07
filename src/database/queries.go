package database

const (
	AddSuper = `
		INSERT INTO TB_SUPER (
			SUPER_ID,
			SUPER_UUID,
			SUPER_NAME,
			SUPER_FULLNAME,
			SUPER_INTELLIGENCE,
			SUPER_POWER,
			SUPER_OCCUPATION,
			SUPER_IMAGE,
		)

		VALUES (
			default,
			$1,
			$2,
			$3,
			$4,
			$5,
			$6,
			$7,
			$8,
		) RETURNING SUPER_ID
	`

	RemoveSuper = `
		DELETE FROM TB_SUPER WHERE SUPER_UUID = $1
	`

	FilterNameSuper = `
		SELECT * FROM TB_SUPER WHERE SUPER_NAME = $1
	`

	FilterUUIDSuper = `
		SELECT * FROM TB_SUPER WHERE SUPER_UUID = $1
	`

	SelectSupers = `
		SELECT * FROM TB_SUPER
	`
)