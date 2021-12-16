package packet

var HEX_TABLE = map[rune][]byte{
    '0': {0, 0, 0, 0},
    '1': {0, 0, 0, 1},
    '2': {0, 0, 1, 0},
    '3': {0, 0, 1, 1},
    '4': {0, 1, 0, 0},
    '5': {0, 1, 0, 1},
    '6': {0, 1, 1, 0},
    '7': {0, 1, 1, 1},
    '8': {1, 0, 0, 0},
    '9': {1, 0, 0, 1},
    'A': {1, 0, 1, 0},
    'B': {1, 0, 1, 1},
    'C': {1, 1, 0, 0},
    'D': {1, 1, 0, 1},
    'E': {1, 1, 1, 0},
    'F': {1, 1, 1, 1},
}

const (
    SUM_TYPE_ID int = iota
    PRODUCT_TYPE_ID
    MINIMUM_TYPE_ID
    MAXIMUM_TYPE_ID
    LITERAL_VALUE_TYPE_ID
    GREATER_THAN_TYPE_ID
    LESS_THAN_TYPE_ID
    EQUAL_TO_TYPE_ID
)

const (
    TOTAL_LENGTH_TYPE_ID int = iota
    AMOUNT_LENGTH_TYPE_ID
)
