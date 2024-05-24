package pkg

type Pagination interface {
	GetOffset() int
	GetLimit() int
	GetPage() int
	GetSort() string
}

