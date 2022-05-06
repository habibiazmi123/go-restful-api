package book

type Service interface{
	FindAll() ([]Book, error)
	FindById(ID int) (Book, error)
	Create(book BookRequest) (Book, error)
	Update(ID int, book BookRequest) (Book, error)
	Delete(ID int) (Book, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) FindAll() ([]Book, error){
	books, err := s.repository.FindAll()
	return books, err
}

func (s *service) FindById(ID int) (Book, error){
	book, err := s.repository.FindById(ID)
	return book, err
}

func (s *service) Create(bookrequest BookRequest) (Book, error) {
	price, _ := bookrequest.Price.Int64()
	rating, _ := bookrequest.Rating.Int64()
	discount, _ := bookrequest.Discount.Int64()

	book := Book{
		Title: bookrequest.Title,
		Price: int(price),
		Description: bookrequest.Description,
		Rating: int(rating),
		Discount: int(discount),
	}

	newBook, err := s.repository.Create(book)
	return newBook, err
}

func (s *service) Update(ID int, bookrequest BookRequest) (Book, error) {
	book, err := s.repository.FindById(ID)

	price, _ := bookrequest.Price.Int64()
	rating, _ := bookrequest.Rating.Int64()
	discount, _ := bookrequest.Discount.Int64()

	book.Title = bookrequest.Title
	book.Price = int(price)
	book.Description = bookrequest.Description
	book.Rating = int(rating)
	book.Discount = int(discount)

	newBook, err := s.repository.Update(book)

	return newBook, err
}

func (s *service) Delete(ID int) (Book, error){
	book, err := s.repository.FindById(ID)
	newBook, err := s.repository.Delete(book)
	return newBook, err
}