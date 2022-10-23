package blogs

type blogUsecase struct {
	blogRepository Repository
}

func NewBlogUsecase(br Repository) Usecase {
	return &blogUsecase{
		blogRepository: br,
	}
}

func (bu *blogUsecase) GetAll(name string) []Domain {
	return bu.blogRepository.GetAll(name)
}

func (bu *blogUsecase) GetByID(id string) Domain {
	return bu.blogRepository.GetByID(id)
}

func (bu *blogUsecase) GetByCategoryID(categoryId string) []Domain {
	return bu.blogRepository.GetByCategoryID(categoryId)
}

func (bu *blogUsecase) Create(blogDomain *Domain) Domain {
	return bu.blogRepository.Create(blogDomain)
}

func (bu *blogUsecase) Update(id string, blogDomain *Domain) Domain {
	return bu.blogRepository.Update(id, blogDomain)
}

func (bu *blogUsecase) Delete(id string) bool {
	return bu.blogRepository.Delete(id)
}