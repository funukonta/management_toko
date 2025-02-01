package constants

const (
	PRODUCT_TYPE_CREAM    = "CREAM"
	PRODUCT_TYPE_SERUM    = "SERUM"
	PRODUCT_TYPE_FACEWASH = "FACEWASH"

	PRODUCT_TYPE_CREAM_LABEL    = "Cream"
	PRODUCT_TYPE_SERUM_LABEL    = "Serum"
	PRODUCT_TYPE_FACEWASH_LABEL = "Facewash"
)

func MapToLabelProductType() map[string]string {
	return map[string]string{
		PRODUCT_TYPE_CREAM:    PRODUCT_TYPE_CREAM_LABEL,
		PRODUCT_TYPE_SERUM:    PRODUCT_TYPE_SERUM_LABEL,
		PRODUCT_TYPE_FACEWASH: PRODUCT_TYPE_FACEWASH_LABEL,
	}
}
